package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/18_http/conteudo_http_didaticos/http"
)

func main() {
	fmt.Println("=== HTTP EM GO ===\n")

	http.DemonstrarServidorNativo()
	http.DemonstrarMetodosHTTP()
	http.DemonstrarRotasHandlers()
	http.DemonstrarServidorGin()
}
