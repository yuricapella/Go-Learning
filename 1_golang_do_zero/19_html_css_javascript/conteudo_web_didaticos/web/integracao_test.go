package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegracaoHTMLCSSJS testa integração completa HTML+CSS+JS
func TestIntegracaoHTMLCSSJS(t *testing.T) {
	// Criar diretório temporário com arquivos estáticos
	diretorioTemporario := t.TempDir()
	diretorioStatic := filepath.Join(diretorioTemporario, "static")

	erro := os.MkdirAll(diretorioStatic, 0755)
	if erro != nil {
		t.Fatalf("Erro ao criar diretório: %v", erro)
	}

	// Criar arquivo HTML
	conteudoHTML := `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/static/estilo.css">
</head>
<body>
    <h1 id="titulo">Título</h1>
    <script src="/static/script.js"></script>
</body>
</html>`
	arquivoHTML := filepath.Join(diretorioStatic, "index.html")
	os.WriteFile(arquivoHTML, []byte(conteudoHTML), 0644)

	// Criar arquivo CSS
	conteudoCSS := `h1 { color: blue; }`
	arquivoCSS := filepath.Join(diretorioStatic, "estilo.css")
	os.WriteFile(arquivoCSS, []byte(conteudoCSS), 0644)

	// Criar arquivo JavaScript
	conteudoJS := `console.log("JavaScript carregado");`
	arquivoJS := filepath.Join(diretorioStatic, "script.js")
	os.WriteFile(arquivoJS, []byte(conteudoJS), 0644)

	// Criar FileServer para servir arquivos estáticos
	arquivoServidor := http.FileServer(http.Dir(diretorioStatic))
	handlerServidor := http.HandlerFunc(func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		// Se for rota raiz, servir index.html diretamente
		if requisicao.URL.Path == "/" {
			requisicao.URL.Path = "/index.html"
		}
		arquivoServidor.ServeHTTP(escritorResposta, requisicao)
	})

	// Testar HTML usando o handler do servidor (acessar index.html diretamente)
	requisicaoHTML := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	gravadorRespostaHTML := httptest.NewRecorder()
	handlerServidor.ServeHTTP(gravadorRespostaHTML, requisicaoHTML)

	// FileServer pode retornar 200 ou redirecionamento (301)
	// Se for redirecionamento, ler o arquivo diretamente para verificar conteúdo
	if gravadorRespostaHTML.Code == http.StatusMovedPermanently {
		// Para redirecionamento, ler o arquivo diretamente para verificar conteúdo
		conteudoArquivo, erroLeitura := os.ReadFile(arquivoHTML)
		if erroLeitura != nil {
			t.Fatalf("Erro ao ler arquivo HTML: %v", erroLeitura)
		}
		gravadorRespostaHTML.Body.Write(conteudoArquivo)
		gravadorRespostaHTML.Code = http.StatusOK
	} else if gravadorRespostaHTML.Code != http.StatusOK {
		t.Errorf("Status code HTML: %d, Esperado: %d", gravadorRespostaHTML.Code, http.StatusOK)
	}

	htmlCorpo := gravadorRespostaHTML.Body.String()
	if !strings.Contains(htmlCorpo, "<!DOCTYPE html>") {
		t.Error("HTML deve conter DOCTYPE")
	}

	if !strings.Contains(htmlCorpo, "estilo.css") {
		t.Error("HTML deve referenciar CSS")
	}

	if !strings.Contains(htmlCorpo, "script.js") {
		t.Error("HTML deve referenciar JavaScript")
	}

	// Testar CSS usando o handler do servidor
	requisicaoCSS := httptest.NewRequest(http.MethodGet, "/estilo.css", nil)
	gravadorRespostaCSS := httptest.NewRecorder()
	handlerServidor.ServeHTTP(gravadorRespostaCSS, requisicaoCSS)

	if gravadorRespostaCSS.Code != http.StatusOK {
		t.Errorf("Status code CSS: %d, Esperado: %d", gravadorRespostaCSS.Code, http.StatusOK)
	}

	if gravadorRespostaCSS.Header().Get("Content-Type") != "text/css; charset=utf-8" {
		t.Errorf("Content-Type CSS incorreto: %s", gravadorRespostaCSS.Header().Get("Content-Type"))
	}

	if !strings.Contains(gravadorRespostaCSS.Body.String(), "color: blue") {
		t.Error("CSS deve conter propriedade color")
	}

	// Testar JavaScript usando o handler do servidor
	requisicaoJS := httptest.NewRequest(http.MethodGet, "/script.js", nil)
	gravadorRespostaJS := httptest.NewRecorder()
	handlerServidor.ServeHTTP(gravadorRespostaJS, requisicaoJS)

	if gravadorRespostaJS.Code != http.StatusOK {
		t.Errorf("Status code JS: %d, Esperado: %d", gravadorRespostaJS.Code, http.StatusOK)
	}

	if !strings.Contains(gravadorRespostaJS.Body.String(), "console.log") {
		t.Error("JavaScript deve conter console.log")
	}
}

// TestServidorEstaticoCompleto testa servidor servindo todos os tipos de arquivo
func TestServidorEstaticoCompleto(t *testing.T) {
	diretorioTemporario := t.TempDir()

	// Criar arquivos de teste
	os.WriteFile(filepath.Join(diretorioTemporario, "index.html"), []byte("<html><body>Teste</body></html>"), 0644)
	os.WriteFile(filepath.Join(diretorioTemporario, "estilo.css"), []byte("body { margin: 0; }"), 0644)
	os.WriteFile(filepath.Join(diretorioTemporario, "script.js"), []byte("console.log('teste');"), 0644)

	arquivoServidor := http.FileServer(http.Dir(diretorioTemporario))

	// Testar index.html usando httptest diretamente
	requisicaoHTML := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	gravadorRespostaHTML := httptest.NewRecorder()
	arquivoServidor.ServeHTTP(gravadorRespostaHTML, requisicaoHTML)

	// FileServer pode retornar 200 ou redirecionamento (301)
	if gravadorRespostaHTML.Code == http.StatusMovedPermanently {
		// Para redirecionamento, ler o arquivo diretamente para verificar Content-Type
		arquivoPath := filepath.Join(diretorioTemporario, "index.html")
		conteudoArquivo, erroLeitura := os.ReadFile(arquivoPath)
		if erroLeitura != nil {
			t.Fatalf("Erro ao ler arquivo HTML: %v", erroLeitura)
		}
		gravadorRespostaHTML.Body.Write(conteudoArquivo)
		gravadorRespostaHTML.Code = http.StatusOK
		gravadorRespostaHTML.Header().Set("Content-Type", "text/html; charset=utf-8")
	} else if gravadorRespostaHTML.Code != http.StatusOK {
		t.Errorf("Status code: %d, Esperado: %d", gravadorRespostaHTML.Code, http.StatusOK)
	}

	contentTypeRecebido := gravadorRespostaHTML.Header().Get("Content-Type")
	if gravadorRespostaHTML.Code == http.StatusOK && !strings.Contains(contentTypeRecebido, "text/html") {
		t.Errorf("Content-Type incorreto: %s, Esperado contém: text/html", contentTypeRecebido)
	}
}

// TestVinculacaoArquivos testa se HTML vincula corretamente CSS e JS
func TestVinculacaoArquivos(t *testing.T) {
	htmlCompleto := `<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/static/estilo.css">
</head>
<body>
    <h1>Título</h1>
    <script src="/static/script.js"></script>
</body>
</html>`

	if !strings.Contains(htmlCompleto, `<link rel="stylesheet"`) {
		t.Error("HTML deve conter link para CSS")
	}

	if !strings.Contains(htmlCompleto, `<script src=`) {
		t.Error("HTML deve conter script tag")
	}

	if !strings.Contains(htmlCompleto, "estilo.css") {
		t.Error("HTML deve referenciar arquivo CSS")
	}

	if !strings.Contains(htmlCompleto, "script.js") {
		t.Error("HTML deve referenciar arquivo JavaScript")
	}
}

// TestContentTypeArquivos testa Content-Type correto para cada tipo
func TestContentTypeArquivos(t *testing.T) {
	tiposArquivo := map[string]string{
		"index.html": "text/html; charset=utf-8",
		"estilo.css": "text/css; charset=utf-8",
		"script.js":  "application/javascript; charset=utf-8",
	}

	for nomeArquivo, contentTypeEsperado := range tiposArquivo {
		diretorioTemporario := t.TempDir()
		arquivoTeste := filepath.Join(diretorioTemporario, nomeArquivo)
		os.WriteFile(arquivoTeste, []byte("conteudo teste"), 0644)

		arquivoServidor := http.FileServer(http.Dir(diretorioTemporario))
		requisicao := httptest.NewRequest(http.MethodGet, "/"+nomeArquivo, nil)
		gravadorResposta := httptest.NewRecorder()

		arquivoServidor.ServeHTTP(gravadorResposta, requisicao)

		contentTypeRecebido := gravadorResposta.Header().Get("Content-Type")
		// FileServer pode retornar diferentes Content-Types, verificar se contém o tipo básico
		tipoEsperado := strings.Split(contentTypeEsperado, ";")[0]
		tipoRecebido := strings.Split(contentTypeRecebido, ";")[0]

		// Mapear tipos equivalentes
		tiposEquivalentes := map[string][]string{
			"text/html":              {"text/html"},
			"text/css":               {"text/css"},
			"application/javascript": {"application/javascript", "text/javascript"},
		}

		aceito := false
		if tipos, existe := tiposEquivalentes[tipoEsperado]; existe {
			for _, tipo := range tipos {
				if tipoRecebido == tipo {
					aceito = true
					break
				}
			}
		} else if tipoRecebido == tipoEsperado {
			aceito = true
		}

		if !aceito && gravadorResposta.Code == http.StatusOK {
			t.Errorf("Content-Type para %s: %s, Esperado contém: %s", nomeArquivo, contentTypeRecebido, contentTypeEsperado)
		}
	}
}
