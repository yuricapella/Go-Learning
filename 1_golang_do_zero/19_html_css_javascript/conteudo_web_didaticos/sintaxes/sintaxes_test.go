package sintaxeWeb_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sintaxeWeb "github.com/yuricapella/Go-Learning/1_golang_do_zero/19_html_css_javascript/conteudo_web_didaticos/sintaxes"
)

// TestSintaxeHTMLBasico testa geração de HTML básico
func TestSintaxeHTMLBasico(t *testing.T) {
	html := sintaxeWeb.SintaxeHTMLBasico()

	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Error("HTML deve conter DOCTYPE")
	}

	if !strings.Contains(html, "<html") {
		t.Error("HTML deve conter tag html")
	}

	if !strings.Contains(html, "<head>") {
		t.Error("HTML deve conter tag head")
	}

	if !strings.Contains(html, "<body>") {
		t.Error("HTML deve conter tag body")
	}
}

// TestSintaxeCSSBasico testa geração de CSS básico
func TestSintaxeCSSBasico(t *testing.T) {
	css := sintaxeWeb.SintaxeCSSBasico()

	if !strings.Contains(css, "h1") {
		t.Error("CSS deve conter seletor h1")
	}

	if !strings.Contains(css, "color:") {
		t.Error("CSS deve conter propriedade color")
	}

	if !strings.Contains(css, ".classe") {
		t.Error("CSS deve conter seletor de classe")
	}

	if !strings.Contains(css, "#identificador") {
		t.Error("CSS deve conter seletor de ID")
	}
}

// TestSintaxeJavaScriptBasico testa geração de JavaScript básico
func TestSintaxeJavaScriptBasico(t *testing.T) {
	javascript := sintaxeWeb.SintaxeJavaScriptBasico()

	if !strings.Contains(javascript, "let") {
		t.Error("JavaScript deve conter declaração let")
	}

	if !strings.Contains(javascript, "function") {
		t.Error("JavaScript deve conter função")
	}

	if !strings.Contains(javascript, "addEventListener") {
		t.Error("JavaScript deve conter manipulação de eventos")
	}
}

// TestSintaxeFileServer testa FileServer
func TestSintaxeFileServer(t *testing.T) {
	diretorioTemporario := t.TempDir()
	arquivoTeste := filepath.Join(diretorioTemporario, "teste.html")
	os.WriteFile(arquivoTeste, []byte("<html><body>Teste</body></html>"), 0644)

	fileServer := sintaxeWeb.SintaxeFileServer(diretorioTemporario)
	requisicao := httptest.NewRequest(http.MethodGet, "/teste.html", nil)
	gravadorResposta := httptest.NewRecorder()

	fileServer.ServeHTTP(gravadorResposta, requisicao)

	if gravadorResposta.Code != http.StatusOK {
		t.Errorf("Status code recebido: %d, Esperado: %d", gravadorResposta.Code, http.StatusOK)
	}

	if !strings.Contains(gravadorResposta.Body.String(), "Teste") {
		t.Error("Resposta deve conter conteúdo do arquivo")
	}
}

// TestSintaxeVincularCSSAoHTML testa vinculação de CSS ao HTML
func TestSintaxeVincularCSSAoHTML(t *testing.T) {
	html := sintaxeWeb.SintaxeVincularCSSAoHTML()

	if !strings.Contains(html, "<link rel=\"stylesheet\"") {
		t.Error("HTML deve conter link para CSS")
	}

	if !strings.Contains(html, "estilo.css") {
		t.Error("HTML deve referenciar arquivo CSS")
	}
}

// TestSintaxeVincularJSAoHTML testa vinculação de JavaScript ao HTML
func TestSintaxeVincularJSAoHTML(t *testing.T) {
	html := sintaxeWeb.SintaxeVincularJSAoHTML()

	if !strings.Contains(html, "<script src=") {
		t.Error("HTML deve conter script tag")
	}

	if !strings.Contains(html, "script.js") {
		t.Error("HTML deve referenciar arquivo JavaScript")
	}
}

// TestSintaxeEstruturaHTMLCompleta testa estrutura HTML completa
func TestSintaxeEstruturaHTMLCompleta(t *testing.T) {
	html := sintaxeWeb.SintaxeEstruturaHTMLCompleta()

	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Error("HTML deve conter DOCTYPE")
	}

	if !strings.Contains(html, "<link rel=\"stylesheet\"") {
		t.Error("HTML deve conter link para CSS")
	}

	if !strings.Contains(html, "<script src=") {
		t.Error("HTML deve conter script tag")
	}
}

// TestSintaxeContentType testa definição de Content-Type
func TestSintaxeContentType(t *testing.T) {
	gravadorResposta := httptest.NewRecorder()

	sintaxeWeb.SintaxeContentType(gravadorResposta, "html")
	if gravadorResposta.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Errorf("Content-Type HTML incorreto: %s", gravadorResposta.Header().Get("Content-Type"))
	}

	gravadorResposta2 := httptest.NewRecorder()
	sintaxeWeb.SintaxeContentType(gravadorResposta2, "css")
	if gravadorResposta2.Header().Get("Content-Type") != "text/css; charset=utf-8" {
		t.Errorf("Content-Type CSS incorreto: %s", gravadorResposta2.Header().Get("Content-Type"))
	}

	gravadorResposta3 := httptest.NewRecorder()
	sintaxeWeb.SintaxeContentType(gravadorResposta3, "js")
	if gravadorResposta3.Header().Get("Content-Type") != "application/javascript; charset=utf-8" {
		t.Errorf("Content-Type JS incorreto: %s", gravadorResposta3.Header().Get("Content-Type"))
	}
}
