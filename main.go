package main

import (
	"rental-mobil/cmd"
	"rental-mobil/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupMySQL()

	// run http
	cmd.ServeHTTP()
}
