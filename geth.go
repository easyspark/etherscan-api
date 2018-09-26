package etherscan

import (
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
