Project Titls : Design an api to fetch user transactions from etherum node hosted on infura's Kovan testnet.

Step1: Configure following variables in constants 

Ethclient = [kovan testnet api end point]

Cassandrahost = [cassandra host ip address]

Keyspace = [keyspace name]

TotalBlocks = [no of recent blocks have to insert into the database]

Step 2: Run main.go file - go run main.io

    	start local server with 8000 port
	

Step 3: To insert user transaction details into the database(getting from kovan testnet api)

call api end point POST: http://127.0.0.1:8000/addTransactions/

Step 4: To get user transactions

call api end point GET: http://127.0.0.1:8000/getTransactions/{address}

Ex : http://127.0.0.1:8000/getTransactions/0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb

Result : 

[{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x4Ab3687cf35F6a76792A2EB47cF2d53A02f836e0","tx_id":"0x2041599693f07b63da47f08804e7431ae7c45fb67e77347fc6ca70dc49ebf08f","block_no":12953715},{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x4Ab3687cf35F6a76792A2EB47cF2d53A02f836e0","tx_id":"0xbe8d174158d0649f1a107b81ab4192651352f440e02c460f3fa7e11f2158e91f","block_no":12953712},{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x4Ab3687cf35F6a76792A2EB47cF2d53A02f836e0","tx_id":"0xd8205dc114773d90064dc1c98423619ba1ae754eb7c14908749a05b5ebea5c95","block_no":12953718},{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x8ccEe4Aa55219B6E0346d77e8a84aF1451bC6b89","tx_id":"0x96533d96ba06d241dbdc2a1cd658f2873309642ccf1d5b32a5181171dd42007b","block_no":12951000},{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x8ccEe4Aa55219B6E0346d77e8a84aF1451bC6b89","tx_id":"0x9a82d6ac35d9b1cf1bab8e93c89daa11d46fed38fda265d80630eecf3cfe4c0c","block_no":12950997},{"user_add":"0x01aD0cb736dbBf3c3124Bef0d7050E8D85D32BAb","to_add":"0x8ccEe4Aa55219B6E0346d77e8a84aF1451bC6b89","tx_id":"0xdd9782e66f80c13777b5036224a93df31ce7751274534786dbe791179b34b646","block_no":12951003}]


-----------------------------------------------------------------------------------------------------------------------------------------------


Db Design :

CREATE KEYSPACE blockchain WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE blockchain.transactions (
    user_add text,
    to_add text,
    tx_id text,
    block_diff varint,
    block_hash text,
    block_no varint,
    block_time varint,
    tx_gas varint,
    tx_gas_price varint,
    tx_val varint,
    PRIMARY KEY (user_add, to_add, tx_id)
)


user_add --> sender address
to_add --> receiver address
tx_id --> transaction id

PRIMARY KEY (user_add, to_add, tx_id): The partition key is user_add, the composite clustering key is (to_add, tx_id)

create secondary index on to_add -- > to get receiver address transactions.


3rd party library : 
gorilla/mux --> implements a request router and dispatcher for matching incoming requests to their respective handler.The name mux stands for "HTTP request multiplexer".

 





