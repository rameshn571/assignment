package utils

import (
	"math/big"
)

const Ethclient = "https://kovan.infura.io/v3/6c6f87a10e12438f8fbb7fc7c762b37c"
const Cassandrahost = "192.168.14.18"
const Keyspace = "blockchain1"
const Totalblocks = 10000

type User struct {
	UserAddress string   `json:"user_add"`
	ToAddress   string   `json:"to_add"`
	TxnId       string   `json:"tx_id"`
	BlockNo     *big.Int `json:"block_no"`
}

