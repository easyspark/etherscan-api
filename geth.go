package etherscan

import (
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
	err = c.call("proxy", "eth_gasPrice", param, price)
	return
}
