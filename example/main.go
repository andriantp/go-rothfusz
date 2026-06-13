package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/andriantp/go-rothfusz/rothfusz"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: go run . <temp> <rh>")
		return
	}

	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Failed to Parse temp [%s]:%v", os.Args[1], err)
	}

	rh, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("Failed to Parse rh [%s]:%v", os.Args[2], err)
	}

	repo := rothfusz.NewRothfusz(
		26.7, // min valid temp
		85,   // humid RH threshold
	)
	result := repo.CalculateHeatIndex(temp, rh)
	js, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatalf("Failed to construct result:%v", err)
	}
	log.Printf("result:\n%s", string(js))
}
