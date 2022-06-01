package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/reginaldosousa/lm-direct-go/messages"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lm_parser MESSAGE")
		fmt.Println("Ex: lm_parser 83050102030405010102010001000000363067")
		os.Exit(1)
	}
	h := os.Args[1]
	decodedBytes, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(decodedBytes)

	message, err := messages.Parse(buf)
	if err != nil {
		log.Fatalf("fail to parse message: %s", err)
	}
	formatted, err := json.MarshalIndent(&message, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", formatted)
}
