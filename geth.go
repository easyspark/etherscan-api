package etherscan

import (
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"strconv"
)

func (c *Client) GetBlockByNumber(blockNum int64) (block Block, err error) {
	param := M{
		"boolean": "true",
		"tag":     strconv.FormatInt(blockNum, 16),
	}

	err = c.call("proxy", "eth_getBlockByNumber", param, &block)
	return
}

func (c *Client) GasPrice() (price *big.Int, err error) {
	param := M{}
	result := ""
	err = c.call("proxy", "eth_gasPrice", param, &result)
	price, _ = math.ParseBig256(result)
	return
}

func (c *Client) EstimateGas(from, to string, value, gasPrice *big.Int) (limit uint64, err error) {
	param := M{
		"value":    value.Text(16),
		"to":       to,
		"form":     from,
		"gasPrice": gasPrice.Text(16),
	}
	err = c.call("proxy", "eth_estimateGas", param, &limit)
	return
}

func (c *Client) SendRawTransaction(hex string) (hash string, err error) {
	param := M{
		"hex": hex,
	}
	err = c.call("proxy", "eth_sendRawTransaction", param, &hash)
	return
}
