package localdb

import (
	"github.com/dgraph-io/badger"
	"log"
	"os"
	"regexp"
)

const (
	dbLocation string = "/tmp/data"
)

func init() {
	info, err := os.Stat(dbLocation)
	if err != nil {
		if isNotPresent := os.IsNotExist(err); isNotPresent {
			log.Fatalf("The directory %s need to exist", dbLocation)
		}
	}

	isMatch, err := regexp.MatchString("drwx.*", info.Mode().String())
	if !isMatch {
		log.Fatalf("The directory %s need to have read write execute permission", dbLocation)
	}
}

var db = func() *badger.DB {
	badgerOpt := badger.DefaultOptions
	badgerOpt.Dir = dbLocation
	badgerOpt.ValueDir = dbLocation

	dbPtr, err := badger.Open(badgerOpt)
	if err == nil {
		return dbPtr
	}
	log.Fatalf("Not able to initialise DB: %v", err)
	return nil
}()

func Write(key string, data []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

func MultiWrite(multiData map[string][]byte) error {
	return db.Update(func(txn *badger.Txn) error {
		for key, val := range multiData {
			err := txn.Set([]byte(key), val)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func Read(key string) ([]byte, error) {
	var data = make([]byte, 0)
	err := db.View(func(txn *badger.Txn) error {
		item, readError := txn.Get([]byte(key))
		if readError == nil {
			dst, copyError := item.ValueCopy(data)
			if copyError == nil {
				data = dst
				return nil
			}
			return copyError
		}
		return readError
	})
	return data, err
}

func ListKeys() []string {
	keys := make([]string, 0)
	_ = db.View(func(txn *badger.Txn) error {

		options := badger.DefaultIteratorOptions
		options.PrefetchValues = false
		iterator := txn.NewIterator(options)

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			item := iterator.Item()
			keys = append(keys, string(item.Key()))
		}
		return nil
	})
	return keys
}

func Remove(key string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func Close() {
	_ = db.Close()
}
