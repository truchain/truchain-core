package storage

import "github.com/dgraph-io/badger/v3"

func OpenStateDb() error {
	db, err := badger.Open(badger.DefaultOptions("/state"))
	if err != nil {
		return err
	}
	_ = db.Close()

	return nil
}
