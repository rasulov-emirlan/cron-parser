package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rasulov-emirlan/cron-parser/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("cron-parser: need exactly one argument")
	}
	// parse string from cmd into Command struct
	cmdParser := parser.NewParser()
	if err := cmdParser.ParseAll(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmdParser)
}
