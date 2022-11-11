package main

import (
	"log"
	"os"

	"github.com/deroproject/derohe/walletapi"
)

func main() {
	count := 1
	tableSize := 1 << 19

	walletapi.Initialize_LookupTable(count, tableSize)
	data, err := walletapi.Balance_lookup_table.Serialize()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./lookuptable", data, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
}
