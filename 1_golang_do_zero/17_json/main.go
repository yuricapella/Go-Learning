package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/17_json/json"
)

func main() {
	fmt.Println("=== JSON EM GO ===")
	fmt.Println()

	json.DemonstrarMarshal()
	json.DemonstrarUnmarshal()
	json.DemonstrarTagsJSON()
	json.DemonstrarMarshalIndent()
}
