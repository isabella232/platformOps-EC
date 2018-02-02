package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type EC_Manifest struct {
	Title    string `json:"title"`
	Command  string `json:"command"`
	Baseline string `json:"baseline"`
}

func (p EC_Manifest) ToString() string {
	return ToJson(p)
}

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}