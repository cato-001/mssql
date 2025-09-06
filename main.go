package main

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		DatabaseHost string `arg:"--host,required,env:MSSQL_DATABASE_HOST" help:"the database host to connect to"`
		DatabaseUser string `arg:"--user,required,env:MSSQL_DATABASE_USER" help:"the username for the connection"`
		DatabaseToken string `arg:"--user,required,env:MSSQL_DATABASE_TOKEN" help:"the username for the connection"`

		Query *QueryCmd `arg:"subcommand"`
		Execute *ExecuteCmd `arg:"subcommand"`
	}

	parser, err := arg.NewParser(arg.Config{
		Program: "mssql",
	}, &args)
	if err != nil {
		Ejsonln(err)
		return
	}
	parser.MustParse(os.Args[1:])

	db, err := OpenDB(args.DatabaseHost, args.DatabaseUser, args.DatabaseToken)
	if err != nil {
		Ejsonln(err)
		return
	}
	defer db.Close()

	if args.Query != nil {
		err := args.Query.Run(db)
		Ejsonln(err)
	}
	if args.Execute != nil {
		err := args.Execute.Run(db)
		Ejsonln(err)
	}
}

func Arguments(args string) []any {
	var (
		arguments = make([]any, 0)
		jsonArgs any
	)
	err := json.Unmarshal([]byte(args), &jsonArgs)
	if err != nil {
		panic(err)
	}

	switch argsValue := jsonArgs.(type) {
	case []any:
		arguments = argsValue
	case map[string]any:
		for key, value := range argsValue {
			arguments = append(arguments, sql.Named(key, value))
		}
	}
	return arguments
}

