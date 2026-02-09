package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/15_go_routines/goroutines"
)

func main() {
	fmt.Println("=== GOROUTINES E CONCORRÊNCIA EM GO ===\n")

	goroutines.DemonstrarGoroutinesBasicas()
	goroutines.DemonstrarWaitGroup()
	goroutines.DemonstrarCanais()
	goroutines.DemonstrarCanaisComBuffer()
	goroutines.DemonstrarSelect()
	goroutines.DemonstrarWorkerPools()
	goroutines.DemonstrarGenerator()
	goroutines.DemonstrarMultiplexador()
}
