package sintaxeWeb

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// SintaxeHTMLBasico demonstra estrutura HTML básica
func SintaxeHTMLBasico() string {
	html := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>Título da Página</title>
</head>
<body>
    <h1>Título Principal</h1>
    <p>Parágrafo de texto</p>
    <div>Divisão de conteúdo</div>
</body>
</html>`
	return html
}

// SintaxeCSSBasico demonstra CSS básico
func SintaxeCSSBasico() string {
	css := `/* Seletor de elemento */
h1 {
    color: blue;
    font-size: 24px;
}

/* Seletor de classe */
.classe {
    margin: 10px;
    padding: 5px;
}

/* Seletor de ID */
#identificador {
    background-color: yellow;
}`
	return css
}

// SintaxeJavaScriptBasico demonstra JavaScript básico
func SintaxeJavaScriptBasico() string {
	javascript := `// Variáveis
let nome = "João";
const idade = 30;

// Função
function saudar() {
    console.log("Olá, " + nome);
}

// Manipulação do DOM
document.getElementById("botao").addEventListener("click", function() {
    alert("Botão clicado!");
});`
	return javascript
}

// SintaxeServidorEstatico demonstra servidor HTTP servindo arquivos estáticos
func SintaxeServidorEstatico(diretorioArquivos string, porta string) error {
	arquivoServidor := http.FileServer(http.Dir(diretorioArquivos))
	http.Handle("/", arquivoServidor)
	fmt.Printf("Servidor iniciado em http://localhost:%s\n", porta)
	return http.ListenAndServe(":"+porta, nil)
}

// SintaxeVincularCSSAoHTML demonstra como vincular CSS ao HTML
func SintaxeVincularCSSAoHTML() string {
	html := `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="estilo.css">
</head>
<body>
    <h1>Título</h1>
</body>
</html>`
	return html
}

// SintaxeVincularJSAoHTML demonstra como vincular JavaScript ao HTML
func SintaxeVincularJSAoHTML() string {
	html := `<!DOCTYPE html>
<html>
<head>
    <script src="script.js"></script>
</head>
<body>
    <button id="botao">Clique</button>
</body>
</html>`
	return html
}

// SintaxeEstruturaHTMLCompleta demonstra estrutura HTML completa com CSS e JS
func SintaxeEstruturaHTMLCompleta() string {
	html := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Página Completa</title>
    <link rel="stylesheet" href="estilo.css">
</head>
<body>
    <h1>Título</h1>
    <p>Conteúdo</p>
    <script src="script.js"></script>
</body>
</html>`
	return html
}

// SintaxeContentType demonstra como definir Content-Type correto
func SintaxeContentType(escritorResposta http.ResponseWriter, tipoArquivo string) {
	switch tipoArquivo {
	case "html":
		escritorResposta.Header().Set("Content-Type", "text/html; charset=utf-8")
	case "css":
		escritorResposta.Header().Set("Content-Type", "text/css; charset=utf-8")
	case "js":
		escritorResposta.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	}
}

// SintaxeServirArquivoEstatico demonstra como servir um arquivo estático específico
func SintaxeServirArquivoEstatico(caminhoArquivo string) ([]byte, error) {
	conteudoArquivo, erro := os.ReadFile(caminhoArquivo)
	if erro != nil {
		return nil, erro
	}
	return conteudoArquivo, nil
}

// SintaxeFileServer demonstra uso de FileServer
func SintaxeFileServer(diretorio string) http.Handler {
	return http.FileServer(http.Dir(diretorio))
}

// SintaxeCaminhoAbsoluto demonstra como obter caminho absoluto
func SintaxeCaminhoAbsoluto(caminhoRelativo string) (string, error) {
	return filepath.Abs(caminhoRelativo)
}
