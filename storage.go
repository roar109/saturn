package saturn

import (
	"log"

	"github.com/tidwall/buntdb"
)

var (
	db      *buntdb.DB
	storage = &FileStorage{}
)

const storageName string = "data.db"

type Storage interface {
	save(sJob SJob)
	read(messageId string)
}

type FileStorage struct{}

func init() {
	log.Println("Configuring DB...")
	// Open the data.db file. It will be created if it doesn't exist.
	db_, err := buntdb.Open(storageName)
	if err != nil {
		log.Fatal(err)
	}
	db = db_
}
func CloseDB() {
	log.Println("Closing DB")
	db.Close()
}

func (strg *FileStorage) read(messageId string) string {
	var value string
	log.Println("Reading from " + messageId)
	err := db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(messageId)
		if err != nil {
			return err
		}
		value = val
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (strg *FileStorage) save(sJob SJob) {
	log.Println(sJob.toString())
	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(sJob.Key, sJob.toString(), nil)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}
