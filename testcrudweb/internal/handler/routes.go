package handler

import "net/http"

// Função para registrar rota GET
func Get(route string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			handlerFunc(writer, request)
		} else {
			http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}

// Função para registrar rota POST
func Post(route string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			handlerFunc(writer, request)
		} else {
			http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}

// Função para registrar rota PUT
func Put(route string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPut {
			handlerFunc(writer, request)
		} else {
			http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}

// Função para registrar rota DELETE
func Delete(route string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodDelete {
			handlerFunc(writer, request)
		} else {
			http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})
}
