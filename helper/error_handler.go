package helper

import "database/sql"

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RollbackOrCommit(tx *sql.Tx) {
	errRecover := recover()
	if errRecover != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(errRecover)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}

}
