package db

import (
	"database/sql"

	"github.com/geooooo/itk-go-test/internal/config"
	_ "github.com/lib/pq"
)

type Db struct {
	connStr string
}

func NewDb(config *config.Config) (*Db, error) {
	db := &Db{
		connStr: config.ConnStr(),
	}

	if !config.DbReset {
		return db, nil
	}

	if err := db.init(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Db) init() error {
	sqlDb, err := sql.Open("postgres", db.connStr)
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	if _, err := sqlDb.Exec(createTableSQL); err != nil {
		return err
	}

	if _, err := sqlDb.Exec(fillDataSQL); err != nil {
		return err
	}

	return nil
}

func (db *Db) UpdateWalletBalance(uuid string, amount uint, operation OperationType) error {
	sqlDb, err := sql.Open("postgres", db.connStr)
	if err != nil {
		return err
	}
	defer sqlDb.Close()

	sqlTx, err := sqlDb.Begin()
	if err != nil {
		return err
	}

	sqlRow := sqlDb.QueryRow(getWalletSumSQL, uuid)
	sum := uint(0)
	if err := sqlRow.Scan(&sum); err != nil {
		sqlTx.Rollback()
		return err
	}

	switch operation {
	case DepositOperation:
		sum += amount
	case WithdrawOperation:
		if amount > sum {
			sum = 0
		} else {
			sum -= amount
		}
	}

	if _, err := sqlDb.Exec(upateWalletSumSQL, sum, uuid); err != nil {
		sqlTx.Rollback()
		return err
	}

	sqlTx.Commit()

	return nil
}

func (db *Db) GetWalletBalance(uuid string) (uint, error) {
	sqlDb, err := sql.Open("postgres", db.connStr)
	if err != nil {
		return 0, err
	}
	defer sqlDb.Close()

	sqlRow := sqlDb.QueryRow(getWalletSumSQL, uuid)

	sum := uint(0)
	if err := sqlRow.Scan(&sum); err != nil {
		return 0, err
	}

	return sum, nil
}
