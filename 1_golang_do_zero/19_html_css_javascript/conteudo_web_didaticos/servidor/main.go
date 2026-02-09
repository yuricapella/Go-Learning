package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Obter diretório atual (onde está o servidor)
	diretorioAtual, erro := filepath.Abs(".")
	if erro != nil {
		log.Fatalf("Erro ao obter diretório atual: %v", erro)
	}

	// Caminho para a pasta static (um nível acima do servidor)
	diretorioStatic := filepath.Join(diretorioAtual, "..", "static")

	fmt.Println("=== SERVIDOR DE ARQUIVOS ESTÁTICOS ===")
	fmt.Printf("Servindo arquivos de: %s\n", diretorioStatic)
	fmt.Println("Servidor rodando em: http://localhost:8080")
	fmt.Println("Acesse: http://localhost:8080/static/index.html")
	fmt.Println("Pressione Ctrl+C para parar o servidor\n")

	// Criar FileServer para servir arquivos estáticos
	arquivoServidor := http.FileServer(http.Dir(diretorioStatic))

	// Configurar rota /static/ para servir arquivos
	http.Handle("/static/", http.StripPrefix("/static/", arquivoServidor))

	// Rota raiz serve index.html diretamente
	http.HandleFunc("/", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		if requisicao.URL.Path == "/" {
			http.ServeFile(escritorResposta, requisicao, filepath.Join(diretorioStatic, "index.html"))
			return
		}
		// Para outras rotas, tentar servir do diretório static
		arquivoServidor.ServeHTTP(escritorResposta, requisicao)
	})

	// Iniciar servidor na porta 8080
	erroServidor := http.ListenAndServe(":8080", nil)
	if erroServidor != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", erroServidor)
	}
}
