# Estrutura do Projeto DevBook

Este documento descreve a arquitetura e organização do projeto DevBook, uma API REST desenvolvida em Go seguindo padrões de Clean Architecture.

## Visão Geral da Arquitetura

O projeto segue uma arquitetura em camadas (layered architecture) com separação clara de responsabilidades:

```
Requisição HTTP → Router → Middlewares → Controllers → Repositórios → Banco de Dados
                                                              ↓
                                                         Modelos
```

## Estrutura de Pastas e Responsabilidades

### Raiz do Projeto

**`main.go`** (Arquivo Principal)
- **Responsabilidade**: Ponto de entrada da aplicação
- **Funções**: 
  - Carrega configurações via `config.Carregar()`
  - Gera router via `router.Gerar()`
  - Inicia servidor HTTP na porta configurada
- **Dependências**: `config`, `router`
- **Por que**: Centraliza a inicialização da aplicação, seguindo o princípio de responsabilidade única

**`docker-compose.yml`**
- **Responsabilidade**: Configuração do ambiente de desenvolvimento
- **Funções**: Define serviço MySQL 8 com variáveis de ambiente
- **Por que**: Facilita setup do ambiente sem necessidade de instalação manual do MySQL

**`.env`**
- **Responsabilidade**: Variáveis de ambiente (não versionado)
- **Contém**: Configurações de banco, porta da API e secret key para JWT
- **Por que**: Separa configurações sensíveis do código fonte

### `src/` - Código Fonte Principal

#### `src/config/` - Configuração da Aplicação

**`config.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de configurações da aplicação
- **Variáveis exportadas**:
  - `StringConexaoBanco`: String de conexão com MySQL
  - `Porta`: Porta onde a API roda
  - `SecretKey`: Chave secreta para assinatura de tokens JWT
- **Função principal**: `Carregar()` - Carrega variáveis de ambiente do arquivo `.env`
- **Dependências**: `github.com/joho/godotenv`
- **Usado por**: `main.go`, `banco/banco.go`, `autenticacao/token.go`
- **Por que**: Centraliza todas as configurações em um único lugar, facilitando manutenção e testes

#### `src/banco/` - Conexão com Banco de Dados

**`banco.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de conexão com MySQL
- **Função principal**: `Conectar()` - Abre e valida conexão com banco de dados
- **Dependências**: `config` (para StringConexaoBanco), `github.com/go-sql-driver/mysql`
- **Usado por**: Todos os controllers que precisam acessar o banco
- **Retorna**: `*sql.DB` para uso em repositórios
- **Por que**: Encapsula lógica de conexão, permitindo reutilização e facilitando testes com mocks

#### `src/modelos/` - Modelos de Dados (Domain Layer)

**`Usuario.go`** (Arquivo Principal)
- **Responsabilidade**: Representa a entidade Usuario
- **Campos**: ID, Nome, Nick, Email, Senha, CriadoEm
- **Métodos principais**:
  - `Preparar(etapa string)`: Orquestra validação e formatação
  - `validar(etapa string)`: Valida campos obrigatórios e formato de email
  - `formatar(etapa string)`: Remove espaços e aplica hash na senha (se cadastro)
- **Dependências**: `seguranca` (para hash de senha), `github.com/badoux/checkmail` (validação de email)
- **Usado por**: Controllers, Repositórios
- **Por que**: Centraliza regras de negócio da entidade Usuario, seguindo Domain-Driven Design

**`Publicacao.go`** (Arquivo Principal)
- **Responsabilidade**: Representa a entidade Publicacao
- **Campos**: ID, Titulo, Conteudo, AutorID, AutorNick, Curtidas, CriadaEm
- **Métodos principais**:
  - `Preparar()`: Valida e formata publicação
  - `validar()`: Verifica titulo e conteudo obrigatórios
  - `formatar()`: Remove espaços em branco
- **Usado por**: Controllers, Repositórios
- **Por que**: Encapsula regras de validação e formatação de publicações

**`Senha.go`** (Arquivo Principal)
- **Responsabilidade**: Modelo para atualização de senha
- **Campos**: Nova, Atual
- **Usado por**: Controller de atualização de senha
- **Por que**: DTO específico para operação de atualização de senha, separando do modelo completo Usuario

#### `src/repositorios/` - Camada de Acesso a Dados

**`usuarios.go`** (Arquivo Principal)
- **Responsabilidade**: Operações CRUD e relacionadas a usuários no banco de dados
- **Struct**: `usuarios` com campo `db *sql.DB`
- **Função construtora**: `NovoRepositorioDeUsuarios(db *sql.DB)` - Cria instância do repositório
- **Métodos principais**:
  - `Criar()`: Insere novo usuário
  - `Buscar()`: Busca usuários por nome ou nick
  - `BuscarPorID()`: Busca usuário específico
  - `BuscarPorEmail()`: Busca para autenticação
  - `Atualizar()`: Atualiza dados do usuário
  - `Deletar()`: Remove usuário
  - `Seguir()`, `PararDeSeguir()`: Gerencia relacionamentos de seguimento
  - `BuscarSeguidores()`, `BuscarSeguidos()`: Lista relacionamentos
  - `BuscarSenha()`, `AtualizarSenha()`: Gerencia senhas
- **Dependências**: `modelos` (para tipos), `database/sql` (para queries)
- **Usado por**: Controllers
- **Por que**: Separa lógica de acesso a dados da lógica de negócio, facilitando testes e manutenção

**`publicacoes.go`** (Arquivo Principal)
- **Responsabilidade**: Operações CRUD e relacionadas a publicações no banco de dados
- **Struct**: `publicacoes` com campo `db *sql.DB`
- **Função construtora**: `NovoRepositorioDePublicacoes(db *sql.DB)`
- **Métodos principais**:
  - `Criar()`: Insere nova publicação
  - `Buscar()`: Busca feed (publicações do usuário e seguidos)
  - `BuscarPorID()`: Busca publicação específica com join para nick do autor
  - `BuscarPorUsuario()`: Lista publicações de um usuário
  - `Atualizar()`: Atualiza titulo e conteudo
  - `Deletar()`: Remove publicação
  - `Curtir()`, `Descurtir()`: Gerencia curtidas
- **Dependências**: `modelos` (para tipos), `database/sql`
- **Usado por**: Controllers de publicações
- **Por que**: Encapsula queries SQL complexas, mantendo controllers limpos

#### `src/controllers/` - Camada de Controle (Handlers HTTP)

**`usuarios.go`** (Arquivo Principal)
- **Responsabilidade**: Handlers HTTP para operações de usuários
- **Função privada**: `verificarUsuarioNaRequisicao()` - Valida se usuário pode realizar ação
- **Handlers principais**:
  - `CriarUsuario()`: Cria novo usuário (não autenticado)
  - `BuscarUsuarios()`: Lista usuários (autenticado)
  - `BuscarUsuarioPorID()`: Busca usuário específico (autenticado)
  - `AtualizarUsuario()`: Atualiza usuário (autenticado + validação de ownership)
  - `DeletarUsuario()`: Remove usuário (autenticado + validação de ownership)
  - `SeguirUsuario()`, `PararDeSeguirUsuario()`: Gerencia seguimento
  - `BuscarSeguidores()`, `BuscarSeguidos()`: Lista relacionamentos
  - `AtualizarSenha()`: Atualiza senha (autenticado + validação de ownership)
- **Dependências**: `repositorios`, `modelos`, `respostas`, `autenticacao`, `seguranca`, `banco`
- **Usado por**: Rotas definidas em `router/rotas/usuarios.go`
- **Por que**: Separa lógica HTTP da lógica de negócio, facilitando testes e manutenção

**`publicacoes.go`** (Arquivo Principal)
- **Responsabilidade**: Handlers HTTP para operações de publicações
- **Handlers principais**:
  - `CriarPublicacao()`: Cria publicação (autenticado, extrai autorID do token)
  - `BuscarPublicacoes()`: Busca feed do usuário (autenticado)
  - `BuscarPublicacaoPorID()`: Busca publicação específica (autenticado)
  - `AtualizarPublicacao()`: Atualiza publicação (autenticado + validação de autor)
  - `DeletarPublicacao()`: Remove publicação (autenticado + validação de autor)
  - `BuscarPublicacoesPorUsuario()`: Lista publicações de um usuário (autenticado)
  - `CurtirPublicacao()`, `DescurtirPublicacao()`: Gerencia curtidas (autenticado)
- **Dependências**: `repositorios`, `modelos`, `respostas`, `autenticacao`, `banco`
- **Usado por**: Rotas definidas em `router/rotas/publicacoes.go`
- **Por que**: Centraliza lógica HTTP relacionada a publicações

**`login.go`** (Arquivo Principal)
- **Responsabilidade**: Handler de autenticação
- **Handler**: `Login()` - Autentica usuário e retorna token JWT
- **Fluxo**: Recebe email/senha → Busca usuário → Valida senha → Gera token → Retorna token
- **Dependências**: `repositorios`, `modelos`, `autenticacao`, `seguranca`, `banco`
- **Usado por**: Rota definida em `router/rotas/login.go`
- **Por que**: Separa lógica de autenticação em controller específico

#### `src/router/` - Configuração de Rotas

**`router.go`** (Arquivo Principal)
- **Responsabilidade**: Geração do router principal
- **Função**: `Gerar()` - Cria instância do Gorilla Mux e configura rotas
- **Dependências**: `github.com/gorilla/mux`, `router/rotas`
- **Usado por**: `main.go`
- **Por que**: Encapsula criação do router, facilitando testes e manutenção

**`router/rotas/rotas.go`** (Arquivo Principal de Rotas)
- **Responsabilidade**: Configuração centralizada de todas as rotas
- **Struct**: `Rota` - Define estrutura de uma rota (URI, Método, Função, RequerAutenticacao)
- **Função**: `Configurar(router *mux.Router)` - Aplica todas as rotas ao router com middlewares apropriados
- **Lógica**: 
  - Agrega rotas de `rotasUsuarios`, `rotaLogin` e `rotasPublicacoes`
  - Aplica middlewares `Logger` e `Autenticar` conforme configuração
- **Dependências**: `middlewares`, `github.com/gorilla/mux`
- **Usado por**: `router/router.go`
- **Por que**: Centraliza configuração de rotas, facilitando manutenção e visualização de todas as rotas

**`router/rotas/usuarios.go`** (Subarquivo de Rotas)
- **Responsabilidade**: Define rotas relacionadas a usuários
- **Variável**: `rotasUsuarios []Rota` - Slice com todas as rotas de usuários
- **Rotas definidas**: 10 rotas (POST, GET, PUT, DELETE para CRUD + seguir, seguidores, atualizar-senha)
- **Usado por**: `rotas.go` (importado e agregado)
- **Por que**: Separa rotas por domínio, mantendo código organizado e modular

**`router/rotas/publicacoes.go`** (Subarquivo de Rotas)
- **Responsabilidade**: Define rotas relacionadas a publicações
- **Variável**: `rotasPublicacoes []Rota` - Slice com todas as rotas de publicações
- **Rotas definidas**: 8 rotas (CRUD + buscar por usuário + curtir/descurtir)
- **Usado por**: `rotas.go` (importado e agregado)
- **Por que**: Mantém rotas organizadas por funcionalidade

**`router/rotas/login.go`** (Subarquivo de Rotas)
- **Responsabilidade**: Define rota de autenticação
- **Variável**: `rotaLogin Rota` - Rota única de login
- **Rota**: POST `/login` sem autenticação
- **Usado por**: `rotas.go` (importado e agregado)
- **Por que**: Separa rota de autenticação das demais rotas

#### `src/middlewares/` - Interceptadores HTTP

**`middlewares.go`** (Arquivo Principal)
- **Responsabilidade**: Middlewares para interceptação de requisições
- **Middlewares**:
  - `Logger()`: Registra método, URI e host de cada requisição
  - `Autenticar()`: Valida token JWT antes de permitir acesso à rota
- **Dependências**: `autenticacao` (para ValidarToken), `respostas` (para retornar erros)
- **Usado por**: `router/rotas/rotas.go` (aplicado condicionalmente conforme RequerAutenticacao)
- **Por que**: Centraliza lógica transversal (logging, autenticação) que se aplica a múltiplas rotas

#### `src/autenticacao/` - Autenticação JWT

**`token.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento completo de tokens JWT
- **Funções principais**:
  - `CriarToken(usuarioID uint64)`: Gera token JWT com expiração de 6 horas
  - `ValidarToken(request *http.Request)`: Valida se token é válido
  - `ExtrairUsuarioID(request *http.Request)`: Extrai ID do usuário do token
  - `extrairToken()`: Função privada que extrai token do header Authorization
  - `retornarChaveDeVerificacao()`: Função privada para validação de assinatura
- **Dependências**: `config` (para SecretKey), `github.com/dgrijalva/jwt-go`
- **Usado por**: `middlewares`, `controllers` (para extrair usuarioID)
- **Por que**: Encapsula toda lógica relacionada a JWT em um único módulo

#### `src/seguranca/` - Segurança e Criptografia

**`seguranca.go`** (Arquivo Principal)
- **Responsabilidade**: Operações de segurança relacionadas a senhas
- **Funções**:
  - `Hash(senha string)`: Gera hash bcrypt da senha
  - `VerificarSenha(senhaString, senhaHash string)`: Compara senha com hash
- **Dependências**: `golang.org/x/crypto/bcrypt`
- **Usado por**: `modelos/Usuario.go` (no método formatar), `controllers/usuarios.go` (na atualização de senha)
- **Por que**: Centraliza operações de segurança, facilitando manutenção e testes

#### `src/respostas/` - Padronização de Respostas HTTP

**`respostas.go`** (Arquivo Principal)
- **Responsabilidade**: Padronização de respostas JSON da API
- **Funções**:
  - `JSON()`: Retorna resposta JSON com status code e dados
  - `Erro()`: Retorna erro padronizado em formato JSON
- **Usado por**: Todos os controllers para retornar respostas consistentes
- **Por que**: Garante formato consistente de respostas e erros em toda a API, facilitando consumo pelo frontend

### `sql/` - Scripts de Banco de Dados

**`sql.sql`** (Arquivo Principal)
- **Responsabilidade**: Schema do banco de dados
- **Conteúdo**: 
  - Criação do banco `devbook`
  - Tabela `usuarios` com campos id, nome, nick, email, senha, criadoEm
  - Tabela `seguidores` com relacionamento many-to-many entre usuários
  - Tabela `publicacoes` com relacionamento com usuarios
- **Usado por**: Script de inicialização do banco
- **Por que**: Versiona estrutura do banco, facilitando setup e migrações

**`dados.sql`** (Subarquivo)
- **Responsabilidade**: Dados de teste para desenvolvimento
- **Conteúdo**: Inserts de usuários e relacionamentos de teste
- **Usado por**: População inicial do banco para testes
- **Por que**: Facilita desenvolvimento e testes sem necessidade de criar dados manualmente

## Fluxo de Requisição

### Requisição Autenticada (exemplo: GET /usuarios)

```
1. Requisição HTTP chega ao servidor
   ↓
2. main.go → router.Gerar() cria router
   ↓
3. router/rotas/rotas.go → Configurar() aplica rotas
   ↓
4. Middleware Logger() → Registra requisição
   ↓
5. Middleware Autenticar() → Valida token JWT
   ↓
6. Controller usuarios.go → BuscarUsuarios()
   ↓
7. Controller abre conexão → banco.Conectar()
   ↓
8. Controller cria repositório → repositorios.NovoRepositorioDeUsuarios(db)
   ↓
9. Repositório executa query → usuarios.Buscar()
   ↓
10. Controller formata resposta → respostas.JSON()
    ↓
11. Resposta HTTP enviada ao cliente
```

### Requisição de Login (POST /login)

```
1. Requisição HTTP chega ao servidor
   ↓
2. Middleware Logger() → Registra requisição
   ↓
3. Sem middleware Autenticar (rota pública)
   ↓
4. Controller login.go → Login()
   ↓
5. Controller busca usuário → repositorios.BuscarPorEmail()
   ↓
6. Controller valida senha → seguranca.VerificarSenha()
   ↓
7. Controller gera token → autenticacao.CriarToken()
   ↓
8. Token retornado como texto simples
```

## Relacionamentos e Dependências

### Hierarquia de Dependências

```
main.go
├── config/ (carrega primeiro)
│   └── config.go
├── router/
│   ├── router.go
│   └── rotas/
│       ├── rotas.go (orquestra todas as rotas)
│       ├── usuarios.go
│       ├── publicacoes.go
│       └── login.go
│
controllers/
├── dependem de: repositorios, modelos, respostas, autenticacao, seguranca, banco
└── usuarios.go, publicacoes.go, login.go
│
repositorios/
├── dependem de: modelos, banco (via *sql.DB)
└── usuarios.go, publicacoes.go
│
middlewares/
├── dependem de: autenticacao, respostas
└── middlewares.go
│
autenticacao/
├── dependem de: config
└── token.go
│
seguranca/
└── seguranca.go (sem dependências internas)
│
respostas/
└── respostas.go (sem dependências internas)
│
modelos/
├── dependem de: seguranca (Usuario.go)
└── Usuario.go, Publicacao.go, Senha.go
```

### Por que esta estrutura?

1. **Separação de Responsabilidades**: Cada camada tem uma responsabilidade única e bem definida
2. **Testabilidade**: Camadas podem ser testadas independentemente usando mocks
3. **Manutenibilidade**: Mudanças em uma camada não afetam outras diretamente
4. **Escalabilidade**: Fácil adicionar novas rotas, controllers ou repositórios seguindo o padrão
5. **Reutilização**: Módulos como `respostas`, `seguranca` e `autenticacao` são reutilizados em múltiplos lugares
6. **Organização**: Rotas separadas por domínio facilitam localização e manutenção

## Padrões Arquiteturais Utilizados

- **Repository Pattern**: Repositórios abstraem acesso a dados
- **Dependency Injection**: Repositórios recebem `*sql.DB` via construtor
- **Middleware Pattern**: Middlewares interceptam requisições antes dos handlers
- **DTO Pattern**: Modelos como `Senha` servem como Data Transfer Objects
- **Factory Pattern**: Funções `NovoRepositorioDe*` criam instâncias de repositórios
