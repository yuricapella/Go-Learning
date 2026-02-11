# Estrutura do Projeto Order CQRS

Este documento descreve a arquitetura e organização do projeto Order CQRS, uma API desenvolvida em Go seguindo padrões CQRS (Command Query Responsibility Segregation) com separação de escrita (MySQL) e leitura (MongoDB).

## Visão Geral da Arquitetura

O projeto segue arquitetura CQRS com separação clara entre comandos (escrita) e queries (leitura):

```
Requisição HTTP → Router → Middlewares → Handlers
                                              ↓
                                    Commands (Write) → MySQL → Event → RabbitMQ → Projection → MongoDB
                                    Queries (Read) → MongoDB
```

## Estrutura de Pastas e Responsabilidades

### Raiz do Projeto

**`cmd/api/main.go`** (Arquivo Principal)
- **Responsabilidade**: Ponto de entrada da aplicação
- **Funções**: 
  - Carrega configurações via `config.Load()`
  - Conecta ao RabbitMQ e configura eventbus
  - Gera router via `router.SetupRouter()`
  - Inicia servidor HTTP na porta configurada
- **Dependências**: `config`, `database`, `eventbus`, `router`
- **Por que**: Centraliza a inicialização da aplicação e configuração de serviços externos

**`docker-compose.yml`**
- **Responsabilidade**: Configuração do ambiente de desenvolvimento
- **Funções**: Define serviços MySQL, MongoDB e RabbitMQ com variáveis de ambiente
- **Por que**: Facilita setup do ambiente sem necessidade de instalação manual dos serviços

**`.env`**
- **Responsabilidade**: Variáveis de ambiente (não versionado)
- **Contém**: Configurações de MySQL, MongoDB, RabbitMQ e porta da API
- **Por que**: Separa configurações sensíveis do código fonte

**`sql/sql.sql`**
- **Responsabilidade**: Schema do banco de dados MySQL (write database)
- **Contém**: Definição da tabela customers
- **Por que**: Versiona estrutura do banco, facilitando setup e migrações

### `src/` - Código Fonte Principal

#### `src/config/` - Configuração da Aplicação

**`config.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de configurações da aplicação
- **Variáveis exportadas**:
  - `APIPort`: Porta onde a API roda
  - `MySQLConnectionString`: String de conexão com MySQL (write database)
  - `MongoDBConnectionString`: String de conexão com MongoDB (read database)
  - `MongoDBDatabaseName`: Nome do banco MongoDB
  - `RabbitMQHost`, `RabbitMQPort`, `RabbitMQUser`, `RabbitMQPassword`: Parâmetros de conexão RabbitMQ
- **Função principal**: `Load()` - Carrega variáveis de ambiente do arquivo `.env`
- **Dependências**: `github.com/joho/godotenv`
- **Usado por**: `main.go`, `database/*.go`
- **Por que**: Centraliza todas as configurações em um único lugar, facilitando manutenção e testes

#### `src/database/` - Conexões com Bancos de Dados

**`mysql.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de conexão com MySQL (write database)
- **Função principal**: `ConnectMySQL()` - Abre e valida conexão com MySQL
- **Dependências**: `config` (para MySQLConnectionString), `github.com/go-sql-driver/mysql`
- **Usado por**: Repositórios de escrita (write repositories)
- **Retorna**: `*sql.DB` para uso em repositórios
- **Por que**: Encapsula lógica de conexão com MySQL, permitindo reutilização

**`mongodb.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de conexão com MongoDB (read database)
- **Função principal**: `ConnectMongoDB()` - Abre e valida conexão com MongoDB
- **Dependências**: `config` (para MongoDBConnectionString), `go.mongodb.org/mongo-driver/v2`
- **Usado por**: Repositórios de leitura (read repositories) e projections
- **Retorna**: `*mongo.Client` e `*mongo.Database`
- **Por que**: Encapsula lógica de conexão com MongoDB para queries

**`rabbitmq.go`** (Arquivo Principal)
- **Responsabilidade**: Gerenciamento de conexão com RabbitMQ
- **Função principal**: `ConnectRabbitMQ()` - Abre conexão e canal com RabbitMQ
- **Dependências**: `config` (para parâmetros RabbitMQ), `github.com/rabbitmq/amqp091-go`
- **Usado por**: `eventbus`, projections (futuro)
- **Retorna**: `*amqp.Connection` e `*amqp.Channel`
- **Por que**: Encapsula lógica de conexão com RabbitMQ para publicação e consumo de eventos

#### `src/router/` - Configuração de Rotas

**`routeschema/route.go`** (Arquivo Principal)
- **Responsabilidade**: Define tipo compartilhado para configuração de rotas
- **Struct**: `Route` com campos `Path`, `Method`, `HandlerFunction`
- **Dependências**: `github.com/gin-gonic/gin`
- **Usado por**: `setup.go`, rotas de domínios (customer, products, orders)
- **Por que**: Centraliza definição de estrutura de rota, evitando duplicação

**`setup.go`** (Arquivo Principal)
- **Responsabilidade**: Configuração do router principal
- **Funções principais**:
  - `SetupRouter()`: Cria instância do Gin router
  - `Configure()`: Aplica todas as rotas dos domínios ao router com middlewares
- **Dependências**: `github.com/gin-gonic/gin`, `middlewares`, rotas de domínios
- **Usado por**: `main.go`
- **Por que**: Centraliza configuração de rotas, facilitando manutenção e visualização de todas as rotas

#### `src/middlewares/` - Interceptadores HTTP

**`middlewares.go`** (Arquivo Principal)
- **Responsabilidade**: Middlewares para interceptação de requisições
- **Middlewares**:
  - `Logger()`: Registra método, URI e host de cada requisição
- **Dependências**: `github.com/gin-gonic/gin`
- **Usado por**: `router/setup.go` (aplicado a todas as rotas)
- **Por que**: Centraliza lógica transversal (logging) que se aplica a múltiplas rotas

#### `src/responses/` - Padronização de Respostas HTTP

**`responses.go`** (Arquivo Principal)
- **Responsabilidade**: Padronização de respostas JSON da API
- **Funções**:
  - `JSON()`: Retorna resposta JSON com status code e dados
  - `Error()`: Retorna erro padronizado em formato JSON
  - `Success()`: Retorna mensagem de sucesso com dados opcionais
- **Dependências**: `github.com/gin-gonic/gin`
- **Usado por**: Todos os handlers para retornar respostas consistentes
- **Por que**: Garante formato consistente de respostas e erros em toda a API

#### `src/eventbus/` - Publicação de Eventos

**`eventbus.go`** (Arquivo Principal)
- **Responsabilidade**: Wrapper para publicação de eventos no RabbitMQ
- **Funções principais**:
  - `SetChannel()`: Configura canal RabbitMQ (chamado no main.go)
  - `PublishEvent()`: Publica evento em uma fila do RabbitMQ
- **Dependências**: `github.com/rabbitmq/amqp091-go`
- **Usado por**: Handlers de comandos para publicar eventos após escrita no MySQL
- **Por que**: Encapsula lógica de publicação de eventos, facilitando testes e manutenção

### Domínios - Estrutura CQRS

#### `src/customer/` - Domínio Customer (Implementado)

**`commands/create.go`** (Arquivo Principal)
- **Responsabilidade**: Define command para criação de customer
- **Struct**: `Create` com campos `Name` e `Email`
- **Usado por**: `handlers/commands.go`
- **Por que**: Representa intenção de criar customer, separando dados de entrada da lógica de negócio

**`handlers/commands.go`** (Arquivo Principal)
- **Responsabilidade**: Handler HTTP para comandos de customer
- **Handlers principais**:
  - `Create()`: Processa command Create, cria no MySQL via repository, cria event e publica no RabbitMQ
- **Fluxo**:
  1. Recebe JSON e faz bind para `commands.Create`
  2. Conecta ao MySQL
  3. Chama `MySQLWriteRepository.Create()` que retorna `id` e `created_at`
  4. Cria `events.Created` com dados do command + dados gerados pelo banco
  5. Publica event no RabbitMQ via `eventbus.PublishEvent()`
  6. Retorna resposta JSON com dados criados
- **Dependências**: `commands`, `repositories`, `events`, `eventbus`, `database`, `responses`
- **Usado por**: Rota definida em `routes.go`
- **Por que**: Orquestra fluxo completo de comando: validação, escrita, criação de event e publicação

**`repositories/mysql_write.go`** (Arquivo Principal)
- **Responsabilidade**: Operações de escrita no MySQL para customer
- **Struct**: `MySQLWriteRepository` com campo `db *sql.DB`
- **Função construtora**: `NewMySQLWriteRepository(db *sql.DB)` - Cria instância do repositório
- **Métodos principais**:
  - `Create()`: Insere customer no MySQL e retorna `id` e `created_at` gerados pelo banco
- **Dependências**: `commands` (para tipo Create), `database/sql`
- **Usado por**: `handlers/commands.go`
- **Por que**: Encapsula lógica de escrita no MySQL, retornando dados gerados pelo banco para criação do event

**`events/created.go`** (Arquivo Principal)
- **Responsabilidade**: Define event CustomerCreated
- **Struct**: `Created` com campos `ID`, `Name`, `Email`, `CreatedAt`
- **Usado por**: `handlers/commands.go` para criar event após escrita no MySQL
- **Por que**: Representa fato ocorrido (customer criado) que será consumido por projections

**`routes.go`** (Arquivo Principal)
- **Responsabilidade**: Define rotas HTTP para domínio customer
- **Variável**: `Routes []routeschema.Route` - Slice com todas as rotas de customer
- **Rotas definidas**: 
  - POST `/customers` → `handlers.Create`
- **Usado por**: `router/setup.go` (importado e agregado)
- **Por que**: Separa rotas por domínio, mantendo código organizado e modular

#### `src/products/` - Domínio Products (Esboço - Não Implementado)

**Estrutura planejada seguindo padrão CQRS:**

```
/products/
  commands/
    create.go        # struct CreateProductCommand (name, price, description)
  queries/
    get_by_id.go     # struct GetProductByIDQuery
    get_all.go       # struct ListProductsQuery
  handlers/
    commands.go      # Handler para POST/PUT/DELETE, usa write repo, publish event
    queries.go       # Handler para GET, usa read repo (MongoDB)
  repositories/
    mysql_write.go   # Repo com métodos para MySQL (escrita)
    mongodb_read.go  # Repo com métodos para MongoDB (leitura)
  events/
    created.go       # struct ProductCreatedEvent
    updated.go       # struct ProductUpdatedEvent
  projections/
    projection.go    # Projection: consumer RabbitMQ, atualiza MongoDB
  routes.go          # Registra rotas POST/GET/PUT/DELETE
```

**Nota**: Esta é uma estrutura planejada. A implementação seguirá o mesmo padrão do domínio customer.

#### `src/orders/` - Domínio Orders (Esboço - Não Implementado)

**Estrutura planejada seguindo padrão CQRS:**

```
/orders/
  commands/
    create.go        # struct CreateOrderCommand (customer_id, product_ids, total)
  queries/
    get_by_id.go     # struct GetOrderByIDQuery
    get_by_customer.go # struct GetOrdersByCustomerQuery
  handlers/
    commands.go      # Handler para POST/PUT/DELETE, usa write repo, publish event
    queries.go       # Handler para GET, usa read repo (MongoDB)
  repositories/
    mysql_write.go   # Repo com métodos para MySQL (escrita)
    mongodb_read.go  # Repo com métodos para MongoDB (leitura)
  events/
    created.go       # struct OrderCreatedEvent
    updated.go       # struct OrderUpdatedEvent
  projections/
    projection.go    # Projection: consumer RabbitMQ, atualiza MongoDB
  routes.go          # Registra rotas POST/GET/PUT/DELETE
```

**Nota**: Esta é uma estrutura planejada. A implementação seguirá o mesmo padrão do domínio customer.

## Fluxo de Requisição - Command (Write)

### Exemplo: POST /customers (Criar Customer)

```
1. Requisição HTTP chega ao servidor
   POST /customers
   Body: { "name": "Ana", "email": "ana@email.com" }
   ↓
2. main.go → router.SetupRouter() cria router
   ↓
3. router/setup.go → Configure() aplica rotas
   ↓
4. Middleware Logger() → Registra requisição
   ↓
5. customer/routes.go → Rota POST /customers aponta para handlers.Create
   ↓
6. customer/handlers/commands.go → Create()
   6.1 Faz bind do JSON para commands.Create
   6.2 Conecta ao MySQL via database.ConnectMySQL()
   6.3 Cria MySQLWriteRepository e chama Create(command)
   ↓
7. customer/repositories/mysql_write.go → Create()
   7.1 Insere customer no MySQL
   7.2 Retorna id e created_at gerados pelo banco
   ↓
8. customer/handlers/commands.go → Create() (continuação)
   8.1 Cria events.Created com dados do command + dados do banco
   8.2 Publica event via eventbus.PublishEvent("customer_created", event)
   ↓
9. eventbus/eventbus.go → PublishEvent()
   9.1 Serializa event para JSON
   9.2 Declara fila "customer_created" no RabbitMQ
   9.3 Publica mensagem na fila
   ↓
10. Resposta HTTP enviada ao cliente
    Status: 201 Created
    Body: { "id": 1, "name": "Ana", "email": "ana@email.com" }
```

## Fluxo de Requisição - Query (Read) - Planejado

### Exemplo: GET /customers/{id} (Buscar Customer)

```
1. Requisição HTTP chega ao servidor
   GET /customers/10
   ↓
2. main.go → router.SetupRouter() cria router
   ↓
3. router/setup.go → Configure() aplica rotas
   ↓
4. Middleware Logger() → Registra requisição
   ↓
5. customer/routes.go → Rota GET /customers/{id} aponta para handlers.GetByID
   ↓
6. customer/handlers/queries.go → GetByID()
   6.1 Extrai id da URL
   6.2 Conecta ao MongoDB via database.ConnectMongoDB()
   6.3 Cria MongoDBReadRepository e chama GetByID(id)
   ↓
7. customer/repositories/mongodb_read.go → GetByID()
   7.1 Consulta MongoDB pelo id
   7.2 Retorna customer (já atualizado pela projection)
   ↓
8. Resposta HTTP enviada ao cliente
    Status: 200 OK
    Body: { "id": 10, "name": "Ana", "email": "ana@email.com", "created_at": "2026-02-10T..." }
```

**Nota**: Este fluxo está planejado e será implementado na parte 2 do CQRS.

## Fluxo de Projection - Planejado

### Exemplo: Customer Projection (Worker separado)

```
1. Projection worker inicia e conecta ao RabbitMQ
   ↓
2. customer/projections/projection.go → Consume()
   2.1 Escuta fila "customer_created" do RabbitMQ
   ↓
3. Para cada evento CustomerCreatedEvent recebido:
   3.1 Deserializa JSON para events.Created
   3.2 Conecta ao MongoDB via database.ConnectMongoDB()
   3.3 Cria documento de leitura otimizado para queries
   3.4 Salva (ou atualiza) documento no MongoDB
   ↓
4. MongoDB agora contém dados atualizados para leitura rápida
```

**Nota**: Este fluxo está planejado e será implementado na parte 2 do CQRS.

## Relacionamentos e Dependências

### Hierarquia de Dependências

```
main.go
├── config/ (carrega primeiro)
│   └── config.go
├── database/
│   ├── mysql.go
│   ├── mongodb.go
│   └── rabbitmq.go
├── eventbus/
│   └── eventbus.go
└── router/
    ├── routeschema/
    │   └── route.go
    └── setup.go
    │
handlers/
├── dependem de: commands, repositories, events, eventbus, database, responses
└── customer/handlers/commands.go
│
repositories/
├── dependem de: commands, database (via *sql.DB ou *mongo.Database)
└── customer/repositories/mysql_write.go
│
middlewares/
├── dependem de: gin
└── middlewares.go
│
responses/
└── responses.go (sem dependências internas)
│
events/
└── customer/events/created.go (sem dependências internas)
│
commands/
└── customer/commands/create.go (sem dependências internas)
```

## Padrões Arquiteturais Utilizados

- **CQRS Pattern**: Separação entre comandos (escrita) e queries (leitura)
- **Event Sourcing**: Eventos representam fatos ocorridos no sistema
- **Repository Pattern**: Repositórios abstraem acesso a dados
- **Dependency Injection**: Repositórios recebem conexões via construtor
- **Middleware Pattern**: Middlewares interceptam requisições antes dos handlers
- **Factory Pattern**: Funções `New*Repository` criam instâncias de repositórios

## Por que esta estrutura?

1. **Separação de Responsabilidades**: Cada camada tem responsabilidade única e bem definida
2. **Escalabilidade**: Write e Read podem escalar independentemente
3. **Manutenibilidade**: Mudanças em uma camada não afetam outras diretamente
4. **Testabilidade**: Camadas podem ser testadas independentemente usando mocks
5. **Organização**: Rotas e handlers separados por domínio facilitam localização e manutenção
6. **Reutilização**: Módulos como `responses`, `eventbus` e `database` são reutilizados em múltiplos lugares
