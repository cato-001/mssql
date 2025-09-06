package main

import (
	"database/sql"
	"fmt"
)

type (
	ExecuteCmd struct {
		Sql string `arg:"positional,required"`
		Args string `arg:"--args"`
	}

	ExecuteResult struct {
		RowsAffected int64 `json:"rows-affected"`
		LastInserted int64 `json:"last-inserted"`
	}
)

func (cmd ExecuteCmd) Run(db *sql.DB) error {
	arguments := Arguments(cmd.Args)

	result, err := db.Exec(cmd.Sql, arguments...)
	if err != nil {
		return fmt.Errorf("while executing command db: %e", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("while getting affected rows: %e", err)
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("while getting last id: %e", err)
	}

	Vjsonln(ExecuteResult {
		RowsAffected: rowsAffected,
		LastInserted: lastInserted,
	})

	return nil
}
