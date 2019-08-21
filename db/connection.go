package db

import (
	"../utils"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

// cassandra session
func init() {
	var err error
	cluster := gocql.NewCluster(utils.Cassandrahost)
	cluster.Consistency = gocql.Quorum
        cluster.ConnectTimeout = time.Second * 10
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}

	ks_query := fmt.Sprintf("CREATE KEYSPACE if NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}", utils.Keyspace)
	err = Session.Query(ks_query).Exec()
	if err != nil {
		log.Println(err)
	}

	t_query := fmt.Sprintf("CREATE TABLE if NOT EXISTS %s.transactions (user_add text,to_add text,tx_id text,block_diff varint,block_hash text,block_no varint,block_time varint,tx_gas varint,tx_gas_price varint,tx_val varint,PRIMARY KEY (user_add, to_add, tx_id))", utils.Keyspace)
	err = Session.Query(t_query).Exec()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
	ind_query := fmt.Sprintf("create INDEX if NOT EXISTS to_add_index ON %s.transactions(to_add)", utils.Keyspace)
	err = Session.Query(ind_query).Exec()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("cassandra init done")
}

//Insert data
func Insert(fromAddress string, toAddress string, txId string, blockDiff uint64, blockHash string, blockNumber uint64, blockTime uint64, txGas uint64, txGasPrice uint64, txVal string) {
	var err error
	i_query := fmt.Sprintf("INSERT INTO %s.transactions (user_add, to_add,tx_id,block_diff,block_hash,block_no,block_time,tx_gas,tx_gas_price,tx_val) VALUES (?,?,?,?,?,?,?,?,?,?)", utils.Keyspace)
	err = Session.Query(i_query,
		fromAddress, toAddress, txId, blockDiff, blockHash, blockNumber, blockTime, txGas, txGasPrice, txVal).Exec()

	if err != nil {
		log.Println(err)
		return
	}
}

func TruncateTable() {
	var err error
	trun_query := fmt.Sprintf("TRUNCATE %s.transactions", utils.Keyspace)
	err = Session.Query(trun_query).Exec()
	if err != nil {
		log.Println(err)
		return
	}
}

