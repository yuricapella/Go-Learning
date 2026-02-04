package http

import (
	"fmt"
)

// DemonstrarServidorGin demonstra como criar servidor HTTP usando Gin framework
func DemonstrarServidorGin() {
	fmt.Println("--- SERVIDOR HTTP COM GIN ---")
	fmt.Println("Gin é um framework web rápido e minimalista para Go.")
	fmt.Println("Ele simplifica a criação de APIs RESTful e aplicações web.\n")

	fmt.Println("Por que usar Gin?")
	fmt.Println("  - Sintaxe mais limpa e expressiva")
	fmt.Println("  - Roteamento mais poderoso e flexível")
	fmt.Println("  - Middleware integrado")
	fmt.Println("  - JSON binding automático")
	fmt.Println("  - Validação de dados simplificada")
	fmt.Println("  - Performance excelente")
	fmt.Println("  - Grande comunidade e documentação\n")

	fmt.Println("Instalação:")
	fmt.Println("  go get -u github.com/gin-gonic/gin\n")

	fmt.Println("Exemplo 1: Servidor Gin básico")
	fmt.Println("Criando um servidor simples com Gin:\n")

	fmt.Println("  package main")
	fmt.Println("  import (")
	fmt.Println("      \"github.com/gin-gonic/gin\"")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("  func main() {")
	fmt.Println("      roteador := gin.Default()")
	fmt.Println("      roteador.GET(\"/\", func(contexto *gin.Context) {")
	fmt.Println("          contexto.String(200, \"Olá, mundo!\")")
	fmt.Println("      })")
	fmt.Println("      roteador.Run(\":8080\")")
	fmt.Println("  }\n")

	fmt.Println("Exemplo 2: Múltiplas rotas com Gin")
	fmt.Println("Definindo várias rotas de forma simples:\n")

	fmt.Println("  roteador := gin.Default()")
	fmt.Println("  roteador.GET(\"/\", handlerInicial)")
	fmt.Println("  roteador.GET(\"/sobre\", handlerSobre)")
	fmt.Println("  roteador.GET(\"/contato\", handlerContato)")
	fmt.Println("  roteador.Run(\":8080\")\n")

	fmt.Println("Exemplo 3: Métodos HTTP com Gin")
	fmt.Println("Gin facilita definir métodos HTTP específicos:\n")

	fmt.Println("  roteador := gin.Default()")
	fmt.Println("  roteador.GET(\"/usuarios\", listarUsuarios)")
	fmt.Println("  roteador.POST(\"/usuarios\", criarUsuario)")
	fmt.Println("  roteador.PUT(\"/usuarios/:id\", atualizarUsuario)")
	fmt.Println("  roteador.DELETE(\"/usuarios/:id\", deletarUsuario)")
	fmt.Println("  roteador.Run(\":8080\")\n")

	fmt.Println("Exemplo 4: Parâmetros de rota")
	fmt.Println("Gin facilita capturar parâmetros da URL:\n")

	fmt.Println("  roteador.GET(\"/usuarios/:id\", func(contexto *gin.Context) {")
	fmt.Println("      idUsuario := contexto.Param(\"id\")")
	fmt.Println("      contexto.String(200, \"ID do usuário: %s\", idUsuario)")
	fmt.Println("  })")
	fmt.Println("  // Uso: /usuarios/123\n")

	fmt.Println("Exemplo 5: Query parameters com Gin")
	fmt.Println("Lendo query parameters de forma simples:\n")

	fmt.Println("  roteador.GET(\"/buscar\", func(contexto *gin.Context) {")
	fmt.Println("      nome := contexto.Query(\"nome\")")
	fmt.Println("      idade := contexto.DefaultQuery(\"idade\", \"0\")")
	fmt.Println("      contexto.JSON(200, gin.H{\"nome\": nome, \"idade\": idade})")
	fmt.Println("  })")
	fmt.Println("  // Uso: /buscar?nome=João&idade=30\n")

	fmt.Println("Exemplo 6: Body JSON com Gin")
	fmt.Println("Gin facilita binding de JSON:\n")

	fmt.Println("  type Usuario struct {")
	fmt.Println("      Nome  string `json:\"nome\"`")
	fmt.Println("      Idade int    `json:\"idade\"`")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  roteador.POST(\"/usuarios\", func(contexto *gin.Context) {")
	fmt.Println("      var usuario Usuario")
	fmt.Println("      if erro := contexto.ShouldBindJSON(&usuario); erro != nil {")
	fmt.Println("          contexto.JSON(400, gin.H{\"erro\": erro.Error()})")
	fmt.Println("          return")
	fmt.Println("      }")
	fmt.Println("      contexto.JSON(201, gin.H{\"status\": \"criado\", \"usuario\": usuario})")
	fmt.Println("  })\n")

	fmt.Println("Exemplo 7: Retornando JSON")
	fmt.Println("Gin facilita retornar JSON:\n")

	fmt.Println("  roteador.GET(\"/dados\", func(contexto *gin.Context) {")
	fmt.Println("      dados := gin.H{")
	fmt.Println("          \"nome\": \"João\",")
	fmt.Println("          \"idade\": 30,")
	fmt.Println("      }")
	fmt.Println("      contexto.JSON(200, dados)")
	fmt.Println("  })\n")

	fmt.Println("Exemplo 8: Grupos de rotas")
	fmt.Println("Organizando rotas em grupos:\n")

	fmt.Println("  api := roteador.Group(\"/api\")")
	fmt.Println("  {")
	fmt.Println("      api.GET(\"/usuarios\", listarUsuarios)")
	fmt.Println("      api.POST(\"/usuarios\", criarUsuario)")
	fmt.Println("  }")
	fmt.Println("  // Rotas: /api/usuarios\n")

	fmt.Println("Exemplo 9: Middleware básico")
	fmt.Println("Adicionando middleware simples:\n")

	fmt.Println("  roteador.Use(func(contexto *gin.Context) {")
	fmt.Println("      fmt.Println(\"Middleware executado\")")
	fmt.Println("      contexto.Next()")
	fmt.Println("  })\n")

	fmt.Println("Comparação: net/http padrão vs Gin")
	fmt.Println("  net/http padrão:")
	fmt.Println("    ✓ Sem dependências externas")
	fmt.Println("    ✓ Controle total")
	fmt.Println("    ✓ Ideal para aprender fundamentos")
	fmt.Println("    ✗ Mais verboso")
	fmt.Println("    ✗ Menos recursos prontos\n")

	fmt.Println("  Gin:")
	fmt.Println("    ✓ Sintaxe mais limpa")
	fmt.Println("    ✓ Roteamento mais poderoso")
	fmt.Println("    ✓ Middleware integrado")
	fmt.Println("    ✓ JSON binding automático")
	fmt.Println("    ✗ Dependência externa")
	fmt.Println("    ✗ Menos controle sobre detalhes\n")

	fmt.Println("Quando usar cada abordagem:")
	fmt.Println("  Use net/http quando:")
	fmt.Println("    - Aprendendo fundamentos HTTP")
	fmt.Println("    - API muito simples")
	fmt.Println("    - Quer evitar dependências")
	fmt.Println("    - Precisa controle total\n")

	fmt.Println("  Use Gin quando:")
	fmt.Println("    - API RESTful mais complexa")
	fmt.Println("    - Precisa de roteamento avançado")
	fmt.Println("    - Quer middleware e validação")
	fmt.Println("    - Prioriza velocidade de desenvolvimento\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Gin precisa ser instalado: go get github.com/gin-gonic/gin")
	fmt.Println("  - Use gin.Default() para incluir middleware padrão")
	fmt.Println("  - Use gin.New() para servidor sem middleware padrão")
	fmt.Println("  - Parâmetros de rota são capturados com :nome")
	fmt.Println("  - Query params usam Query() ou DefaultQuery()")
	fmt.Println("  - Use ShouldBindJSON() para binding de JSON")
	fmt.Println("  - Contexto.Next() continua para próximo handler/middleware")
	fmt.Println("  - Gin.H é um alias para map[string]interface{}")
	fmt.Println()
}
