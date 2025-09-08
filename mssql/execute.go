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
		LastInserted *int64 `json:"last-inserted"`
	}
)

func (cmd ExecuteCmd) Run(db *sql.DB) error {
	arguments, err := Arguments(cmd.Args)
	if err != nil {
		return err
	}

	result, err := db.Exec(cmd.Sql, arguments...)
	if err != nil {
		return fmt.Errorf("while executing command db: %e", err)
	}

	var out ExecuteResult
	out.RowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("while getting affected rows: %e", err)
	}

	lastInserted, err := result.LastInsertId()
	if err == nil {
		out.LastInserted = &lastInserted
	} else {
		out.LastInserted = nil
	}

	Vjsonln(out)

	return nil
}
