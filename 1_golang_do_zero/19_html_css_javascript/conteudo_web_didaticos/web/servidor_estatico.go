package web

import (
	"fmt"
)

// DemonstrarServidorEstatico demonstra como servir arquivos estáticos com Go
func DemonstrarServidorEstatico() {
	fmt.Println("--- SERVIDOR DE ARQUIVOS ESTÁTICOS ---")
	fmt.Println("Servidores HTTP podem servir arquivos estáticos (HTML, CSS, JS, imagens)")
	fmt.Println("usando o FileServer do pacote net/http.")
	fmt.Println()

	fmt.Println("Por que servir arquivos estáticos?")
	fmt.Println("  - Páginas web precisam de HTML, CSS e JavaScript")
	fmt.Println("  - Arquivos estáticos são servidos diretamente sem processamento")
	fmt.Println("  - Navegadores fazem requisições HTTP para obter esses arquivos")
	fmt.Println("  - FileServer facilita servir diretórios inteiros")
	fmt.Println("  - Essencial para aplicações web")
	fmt.Println()

	fmt.Println("Exemplo 1: Servidor básico servindo arquivos estáticos")
	fmt.Println("Usando FileServer para servir um diretório:")
	fmt.Println()

	fmt.Println("  package main")
	fmt.Println("  import (")
	fmt.Println("      \"net/http\"")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("  func main() {")
	fmt.Println("      diretorioArquivos := \"./static\"")
	fmt.Println("      arquivoServidor := http.FileServer(http.Dir(diretorioArquivos))")
	fmt.Println("      http.Handle(\"/\", arquivoServidor)")
	fmt.Println("      http.ListenAndServe(\":8080\", nil)")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("Exemplo 2: Servir arquivos em rota específica")
	fmt.Println("Servindo arquivos estáticos em /static/:")
	fmt.Println()

	fmt.Println("  diretorioArquivos := \"./static\"")
	fmt.Println("  arquivoServidor := http.FileServer(http.Dir(diretorioArquivos))")
	fmt.Println("  http.Handle(\"/static/\", http.StripPrefix(\"/static/\", arquivoServidor))")
	fmt.Println("  // Arquivo static/index.html será acessível em /static/index.html")
	fmt.Println()

	fmt.Println("Exemplo 3: Content-Type automático")
	fmt.Println("FileServer define Content-Type automaticamente baseado na extensão:")
	fmt.Println()

	fmt.Println("  - .html → text/html")
	fmt.Println("  - .css → text/css")
	fmt.Println("  - .js → application/javascript")
	fmt.Println("  - .json → application/json")
	fmt.Println("  - .png → image/png")
	fmt.Println("  - .jpg → image/jpeg")
	fmt.Println()

	fmt.Println("Exemplo 4: Servir arquivo específico")
	fmt.Println("Servindo um arquivo específico em uma rota:")
	fmt.Println()

	fmt.Println("  http.HandleFunc(\"/\", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {")
	fmt.Println("      http.ServeFile(escritorResposta, requisicao, \"./static/index.html\")")
	fmt.Println("  })")
	fmt.Println()

	fmt.Println("Exemplo 5: Servidor com múltiplos diretórios")
	fmt.Println("Servindo diferentes tipos de arquivos:")
	fmt.Println()

	fmt.Println("  // Arquivos CSS")
	fmt.Println("  cssServidor := http.FileServer(http.Dir(\"./css\"))")
	fmt.Println("  http.Handle(\"/css/\", http.StripPrefix(\"/css/\", cssServidor))")
	fmt.Println()
	fmt.Println("  // Arquivos JavaScript")
	fmt.Println("  jsServidor := http.FileServer(http.Dir(\"./js\"))")
	fmt.Println("  http.Handle(\"/js/\", http.StripPrefix(\"/js/\", jsServidor))")
	fmt.Println()
	fmt.Println("  // Imagens")
	fmt.Println("  imgServidor := http.FileServer(http.Dir(\"./images\"))")
	fmt.Println("  http.Handle(\"/images/\", http.StripPrefix(\"/images/\", imgServidor))")
	fmt.Println()

	fmt.Println("Exemplo 6: Página padrão (index.html)")
	fmt.Println("FileServer serve index.html automaticamente quando acessa diretório:")
	fmt.Println()

	fmt.Println("  // Acessar / serve automaticamente /index.html")
	fmt.Println("  // Acessar /pasta/ serve automaticamente /pasta/index.html")
	fmt.Println()

	fmt.Println("Vantagens do FileServer:")
	fmt.Println("  ✓ Serve arquivos automaticamente")
	fmt.Println("  ✓ Define Content-Type correto")
	fmt.Println("  ✓ Suporta index.html automático")
	fmt.Println("  ✓ Trata diretórios e arquivos")
	fmt.Println("  ✓ Simples de configurar")
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Use caminhos relativos seguros (evite ../)")
	fmt.Println("  - FileServer serve todo o diretório (incluindo subdiretórios)")
	fmt.Println("  - Use StripPrefix para remover prefixo da URL")
	fmt.Println("  - Content-Type é definido automaticamente")
	fmt.Println("  - index.html é servido automaticamente para diretórios")
	fmt.Println("  - Caminhos são case-sensitive")
	fmt.Println("  - Use http.Dir() para criar FileSystem")
	fmt.Println()
}
