package main

import (
	"database/sql"
	"fmt"
)

type QueryCmd struct {
	Sql string `arg:"positional,required"`
	Args string `arg:"--args"`
}

func (cmd QueryCmd) Run(db *sql.DB) error {
	arguments, err := Arguments(cmd.Args)
	if err != nil {
		return err
	}

	rows, err := db.Query(cmd.Sql, arguments...)
	if err != nil {
		return fmt.Errorf("while querring db: %v", err)
	}

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("while reading columns: %v", err)
	}

	results := make([]any, len(columns))
	scanResults := make([]any, len(columns))
	for index := range results {
		scanResults[index] = &results[index]
	}

	for rows.Next() {
		rows.Scan(scanResults...)

		value := make(map[string]any, len(columns))
		for index, column := range columns {
			result := results[index]
			value[column] = result
		}

		Vjsonln(value)
	}

	return nil
}

