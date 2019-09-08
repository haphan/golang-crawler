package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jackdanger/collectlinks"
	"github.com/kr/pretty"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Must enter address")
		os.Exit(1)
	}
	client := http.Client{}
	resp, err := client.Get(args[0])
	defer resp.Body.Close()

	if nil != err {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	links := collectlinks.All(resp.Body)
	pretty.Print(links)
}
