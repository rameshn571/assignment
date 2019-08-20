package api

import (
	"../db"
	"../utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

// insert transaction date into the database
func InsertTransactionDetails(w http.ResponseWriter, r *http.Request) {
	insertIntoDatabase()
	json.NewEncoder(w).Encode("Data inserted")
}

//get transactions using user address
func GetTransactionDetails(w http.ResponseWriter, r *http.Request) {

	var userList []utils.User
	m := map[string]interface{}{}

	vars := mux.Vars(r)
	address := vars["address"]
	get_query := fmt.Sprintf("select user_add,to_add,tx_id,block_no FROM %s.transactions where user_add = ?", utils.Keyspace)
	iterable := db.Session.Query(get_query, address).Iter()
	for iterable.MapScan(m) {
		userList = append(userList, utils.User{
			UserAddress: m["user_add"].(string),
			ToAddress:   m["to_add"].(string),
			TxnId:       m["tx_id"].(string),
			BlockNo:     m["block_no"].(*big.Int),
		})
		m = map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(userList)
}

func insertIntoDatabase() {
	client, err := ethclient.Dial(utils.Ethclient)
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	latest, err := strconv.ParseInt(header.Number.String(), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := int64(1); i < int64(utils.Totalblocks); i++ {
		latestBN := latest - i
		block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(latestBN))
		if err != nil {
			log.Fatal(err)
		}
		blockNumber := block.Number().Uint64()
		blockTime := block.Time()
		blockDiff := block.Difficulty().Uint64()
		blockHash := block.Hash().Hex()
		fromAddress := ""

		for _, tx := range block.Transactions() {
			// some times tx.To() getting nil (invalid memory address or nil pointer dereference)
			toAdd := tx.To()
			if toAdd != nil {
				txId := tx.Hash().Hex()
				txVal := tx.Value().String()
				txGas := tx.Gas()
				txGasPrice := tx.GasPrice().Uint64()
				toAddress := toAdd.Hex()
				fmt.Println("***toAddress***  ", toAddress)
				chainID, err := client.NetworkID(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
					fromAddress = msg.From().Hex()
				}
				fmt.Println("***fromAddress***  ", fromAddress)
				db.Insert(fromAddress, toAddress, txId, blockDiff, blockHash, blockNumber, blockTime, txGas, txGasPrice, txVal)
			}
		}
	}
}

