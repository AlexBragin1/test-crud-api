package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CommitOrRollback(tx *sqlx.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		log.Printf("Cancel transaction: %s", errRollback)
		panic(err)
	} else {
		if errCommit := tx.Commit(); errCommit == nil {
			log.Println("OK transaction")
		} else {
			log.Println("missing transaction")
		}
	}
}
