package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Lat, Long float64
}

func (l location) UnmarshalJSON([]byte) error {
	return nil
}

func main() {
	var curiosity location
	var bytes []byte = []byte("{\"Lat\":-4.5895,\"Long\":137.4417}")
	json.Unmarshal(bytes, &curiosity)
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
