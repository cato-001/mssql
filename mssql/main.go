package main

import (
	"os"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		DatabaseHost string `arg:"--host,required,env:MSSQL_DATABASE_HOST" help:"database host to access"`
		DatabaseUser string `arg:"--user,required,env:MSSQL_DATABASE_USER" help:"user for the database access"`
		DatabaseToken string `arg:"--token,required,env:MSSQL_DATABASE_TOKEN" help:"token for the database access"`

		Query *QueryCmd `arg:"subcommand"`
		Execute *ExecuteCmd `arg:"subcommand"`
	}

	parser, err := arg.NewParser(arg.Config{
		Program: "mssql",
	}, &args)
	Ejsonln(err)
	parser.MustParse(os.Args[1:])

	db, err := OpenDB(args.DatabaseHost, args.DatabaseUser, args.DatabaseToken)
	Ejsonln(err)
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

