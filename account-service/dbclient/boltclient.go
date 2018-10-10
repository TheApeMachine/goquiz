package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/theapemachine/goquiz/account-service/model"
)

// IBoltClient : Sets up an interface for the bolt embedded key/value store.
type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
	Check() bool
}

// BoltClient : Implements the bolt interface.
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDb : Opens a connection to our embedded key/value store.
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))

		if err != nil {
			return fmt.Errorf("Create bucket failed: %s", err)
		}

		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100

	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)

		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		jsonBytes, _ := json.Marshal(acc)

		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}

	fmt.Printf("Seeded %v fake accounts...\n", total)
}

// Seed : Seed the DB with some data.
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

// QueryAccount : Implements a way to retrieve an Account by ID.
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))

		accountBytes := b.Get([]byte(accountID))

		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}

		json.Unmarshal(accountBytes, &account)

		return nil
	})

	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

// Check : Naive healthcheck.
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}
