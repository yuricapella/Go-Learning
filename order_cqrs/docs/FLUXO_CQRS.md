# Mapeamento do Fluxo CQRS - Order CQRS

Este documento descreve o fluxo completo CQRS (Command Query Responsibility Segregation) implementado para o domínio `customer`, servindo como referência para implementação de novos domínios (ex: `product`, `order`).

## Arquitetura Geral

```
┌─────────────────────────────────────────────────────────────────┐
│                        CQRS ARCHITECTURE                         │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  COMMAND SIDE (Write)          │  QUERY SIDE (Read)            │
│  ─────────────────────         │  ────────────────              │
│  MySQL (Write DB)              │  MongoDB (Read DB)            │
│  ─────────────────────         │  ────────────────              │
│                                 │                                │
│  POST /customers                │  GET /customers/:id           │
│  ↓                              │  ↓                             │
│  Handler → Repository → MySQL   │  Handler → Repository → MongoDB│
│  ↓                              │                                │
│  Event → RabbitMQ               │                                │
│  ↓                              │                                │
│  Projection → MongoDB           │                                │
│                                 │                                │
└─────────────────────────────────────────────────────────────────┘
```

## Fluxo de Command (Criação de Customer)

### 1. Início da Aplicação (`cmd/api/main.go`)

```
main()
├── config.Load()                    # Carrega variáveis de ambiente (.env)
├── database.ConnectRabbitMQ()       # Conecta ao RabbitMQ
├── eventbus.SetChannel()            # Configura canal RabbitMQ
├── projections.StartAllConsumers()  # Inicia consumers de eventos
└── router.SetupRouter()            # Configura rotas HTTP
```

**Arquivos envolvidos:**
- `cmd/api/main.go` - Entry point
- `src/config/config.go` - Carregamento de configurações
- `src/database/rabbitmq.go` - Conexão RabbitMQ
- `src/eventbus/eventbus.go` - Abstração do event bus
- `src/projections/runner.go` - Inicialização de projections
- `src/router/setup.go` - Setup do router Gin

---

### 2. Roteamento (`src/router/setup.go` + `src/customer/router/routes.go`)

```
SetupRouter()
├── Configure()
    └── Para cada rota em customerRoutes.Routes:
        ├── Aplica middleware Logger
        └── Registra rota no Gin (GET/POST/PUT/DELETE)
```

**Arquivos envolvidos:**
- `src/router/setup.go` - Setup principal do router
- `src/customer/router/routes.go` - Definição de rotas do customer
- `src/router/routeschema/route.go` - Schema de rota
- `src/middlewares/middlewares.go` - Middleware Logger

**Rotas definidas:**
- `POST /customers` → `handlers.Create`
- `GET /customers/:id` → `handlers.GetByID`

---

### 3. Command Handler (`src/customer/handlers/commands.go`)

```
POST /customers
↓
handlers.Create(ginContext)
├── Bind JSON → commands.Create { Name, Email }
├── database.ConnectMySQL() → *sql.DB
├── repositories.NewMySQLWriteRepository(db)
├── mysqlWriteRepository.Create(command)
│   └── Retorna: (id int64, createdAt time.Time, error)
├── Cria events.Created { ID, Name, Email, CreatedAt }
└── eventbus.PublishEvent(queue, event)
    └── Publica evento no RabbitMQ
```

**Arquivos envolvidos:**
- `src/customer/handlers/commands.go` - Handler de comandos
- `src/customer/commands/create.go` - Struct do comando
- `src/database/mysql.go` - Conexão MySQL
- `src/customer/repositories/mysql_write.go` - Repository de escrita
- `src/customer/events/created.go` - Struct do evento
- `src/eventbus/eventbus.go` - Publicação de eventos
- `src/customer/utils/constants.go` - Constante da queue

**Fluxo detalhado:**

1. **Recebe requisição HTTP POST** com JSON `{ "name": "...", "email": "..." }`
2. **Valida e bind** para `commands.Create`
3. **Conecta ao MySQL** usando `database.ConnectMySQL()` → usa `config.MySQLConnectionString`
4. **Cria repository** `MySQLWriteRepository`
5. **Executa comando** `Create()` que:
   - Insere no MySQL: `INSERT INTO customers (name, email) VALUES (?, ?)`
   - Retorna `id` (LastInsertId) e `created_at` (SELECT após insert)
6. **Cria evento** `events.Created` com dados do comando + dados do MySQL
7. **Publica evento** no RabbitMQ usando `eventbus.PublishEvent()` com queue `customer.QueueCustomerCreated`

---

### 4. Write Repository (`src/customer/repositories/mysql_write.go`)

```
mysqlWriteRepository.Create(command)
├── Prepare: INSERT INTO customers (name, email) VALUES (?, ?)
├── Exec(command.Name, command.Email)
├── LastInsertId() → id
└── QueryRow: SELECT created_at FROM customers WHERE id = ?
    └── Retorna: (id, createdAt, nil)
```

**Arquivos envolvidos:**
- `src/customer/repositories/mysql_write.go` - Repository de escrita MySQL

**Responsabilidades:**
- Persistir dados no MySQL (write database)
- Retornar dados gerados pelo banco (id, created_at)

---

### 5. Event Bus (`src/eventbus/eventbus.go`)

```
eventbus.PublishEvent(queue, event)
├── json.Marshal(event) → []byte
├── QueueDeclare(queue, durable=true)
└── Publish("", queue, message)
```

**Arquivos envolvidos:**
- `src/eventbus/eventbus.go` - Abstração do RabbitMQ

**Responsabilidades:**
- Serializar evento para JSON
- Declarar queue no RabbitMQ
- Publicar mensagem na queue

---

### 6. Projection Consumer (`src/projections/runner.go` + `src/projections/customer.go`)

```
projections.StartAllConsumers(ctx)
├── database.ConnectMongoDB() → mongoDatabase
├── mongoDatabase.Collection(customerUtils.CollectionCustomers)
├── repositories.NewMongoDBReadRepository(collection)
└── ConsumeCustomerCreatedEvent(ctx, repo)
    └── eventbus.ConsumeEvent(queue, handler)
        └── Loop infinito:
            ├── Recebe mensagem do RabbitMQ
            ├── json.Unmarshal → events.Created
            └── mongoDBReadRepository.InsertCreatedEvent(ctx, event)
```

**Arquivos envolvidos:**
- `src/projections/runner.go` - Inicialização de consumers
- `src/projections/customer.go` - Consumer específico de customer
- `src/database/mongodb.go` - Conexão MongoDB
- `src/customer/repositories/mongodb_read.go` - Repository de leitura
- `src/customer/utils/constants.go` - Constantes (queue e collection)

**Fluxo detalhado:**

1. **Inicialização** (`runner.go`):
   - Conecta ao MongoDB usando `database.ConnectMongoDB()` → usa `config.MongoDBDatabaseName`
   - Obtém collection usando `customerUtils.CollectionCustomers`
   - Cria `MongoDBReadRepository`
   - Inicia consumer em goroutine

2. **Consumer** (`customer.go`):
   - Consome mensagens da queue `customerUtils.QueueCustomerCreated`
   - Deserializa JSON para `events.Created`
   - Insere no MongoDB usando `InsertCreatedEvent()`

---

### 7. Read Repository - Insert (`src/customer/repositories/mongodb_read.go`)

```
mongoDBReadRepository.InsertCreatedEvent(ctx, event)
└── collection.InsertOne(ctx, event)
```

**Arquivos envolvidos:**
- `src/customer/repositories/mongodb_read.go` - Repository de leitura MongoDB

**Responsabilidades:**
- Inserir eventos no MongoDB (read database)
- Usar tags BSON para mapeamento correto dos campos

---

## Fluxo de Query (Leitura de Customer)

### 1. Query Handler (`src/customer/handlers/queries.go`)

```
GET /customers/:id
↓
handlers.GetByID(ginContext)
├── Bind URI → queries.GetByID { ID }
├── database.ConnectMongoDB() → mongoDatabase
├── mongoDatabase.Collection(customerUtils.CollectionCustomers)
├── repositories.NewMongoDBReadRepository(collection)
├── mongoDBReadRepository.GetByID(ctx, query.ID)
└── responses.JSON(ginContext, http.StatusOK, customer)
```

**Arquivos envolvidos:**
- `src/customer/handlers/queries.go` - Handler de queries
- `src/customer/queries/get_by_id.go` - Struct da query
- `src/database/mongodb.go` - Conexão MongoDB
- `src/customer/repositories/mongodb_read.go` - Repository de leitura
- `src/customer/utils/constants.go` - Constante da collection
- `src/responses/responses.go` - Helpers de resposta

**Fluxo detalhado:**

1. **Recebe requisição HTTP GET** com parâmetro `:id` na URL
2. **Valida e bind** para `queries.GetByID` usando `ShouldBindUri()`
3. **Conecta ao MongoDB** usando `database.ConnectMongoDB()` → usa `config.MongoDBDatabaseName`
4. **Obtém collection** usando `customerUtils.CollectionCustomers`
5. **Cria repository** `MongoDBReadRepository`
6. **Busca customer** usando `GetByID(id)`
7. **Retorna JSON** com dados do customer ou erro 404

---

### 2. Read Repository - Get (`src/customer/repositories/mongodb_read.go`)

```
mongoDBReadRepository.GetByID(ctx, id)
├── filter := bson.M{"id": id}
├── collection.FindOne(ctx, filter)
└── Decode(&customer) → events.Created
```

**Arquivos envolvidos:**
- `src/customer/repositories/mongodb_read.go` - Repository de leitura MongoDB

**Responsabilidades:**
- Buscar customer no MongoDB por ID
- Retornar `events.Created` ou erro se não encontrado

---

## Estrutura de Arquivos por Domínio (Customer)

```
src/customer/
├── commands/
│   └── create.go              # Struct do comando Create
├── events/
│   └── created.go             # Struct do evento Created (com tags BSON)
├── handlers/
│   ├── commands.go            # Handler Create (command side)
│   └── queries.go             # Handler GetByID (query side)
├── queries/
│   └── get_by_id.go           # Struct da query GetByID
├── repositories/
│   ├── mysql_write.go          # Repository de escrita (MySQL)
│   └── mongodb_read.go        # Repository de leitura (MongoDB)
├── router/
│   └── routes.go              # Definição de rotas HTTP
└── utils/
    └── constants.go           # Constantes (queues, collections)
```

---

## Configurações e Variáveis de Ambiente

### Arquivo `.env`

```env
# MongoDB Configuration (Read Database)
MONGODB_HOST=localhost
MONGODB_PORT=27017
MONGODB_DATABASE=ecommerce_read    # Nome do DATABASE MongoDB

# MySQL Configuration (Write Database)
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=ecommerce
MYSQL_PASSWORD=ecommerce123
MYSQL_DATABASE=ecommerce_write     # Nome do DATABASE MySQL

# RabbitMQ Configuration
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
```

### Constantes (`src/customer/utils/constants.go`)

```go
const (
    QueueCustomerCreated = "customer_created"  // Nome da queue RabbitMQ
    CollectionCustomers  = "customers"         // Nome da COLLECTION MongoDB
)
```

**Importante:**
- **DATABASE** = vem de `MONGODB_DATABASE` no `.env` (`ecommerce_read`)
- **COLLECTION** = vem da constante `CollectionCustomers` (`customers`)
- **TABELA MySQL** = `customers` (hardcoded no SQL, pode ser parametrizado futuramente)

---

## Fluxo Completo Visual

```
┌─────────────────────────────────────────────────────────────────────┐
│                    COMMAND FLOW (Write)                             │
└─────────────────────────────────────────────────────────────────────┘

HTTP POST /customers
    ↓
[router/routes.go] → handlers.Create
    ↓
[handlers/commands.go]
    ├── Bind: commands.Create { Name, Email }
    ├── Connect MySQL
    ├── mysql_write.Create() → INSERT → (id, createdAt)
    ├── Create: events.Created { ID, Name, Email, CreatedAt }
    └── PublishEvent → RabbitMQ queue "customer_created"
        ↓
[projections/customer.go] ConsumeEvent
    ├── Unmarshal JSON → events.Created
    └── mongodb_read.InsertCreatedEvent() → MongoDB collection "customers"
        ↓
MongoDB Database: "ecommerce_read"
Collection: "customers"
Document: { id, name, email, createdat }


┌─────────────────────────────────────────────────────────────────────┐
│                    QUERY FLOW (Read)                                 │
└─────────────────────────────────────────────────────────────────────┘

HTTP GET /customers/:id
    ↓
[router/routes.go] → handlers.GetByID
    ↓
[handlers/queries.go]
    ├── Bind URI: queries.GetByID { ID }
    ├── Connect MongoDB
    ├── Collection: "customers"
    └── mongodb_read.GetByID(id) → FindOne
        ↓
MongoDB Database: "ecommerce_read"
Collection: "customers"
Filter: { id: <id> }
    ↓
Return: events.Created { ID, Name, Email, CreatedAt }
```

---

## Padrão para Criar Novo Domínio (ex: Product)

Seguindo o mesmo padrão de `customer`, para criar `product`:

### 1. Estrutura de Pastas

```
src/product/
├── commands/
│   └── create.go              # Struct do comando
├── events/
│   └── created.go             # Struct do evento (com BSON tags)
├── handlers/
│   ├── commands.go            # Handler Create
│   └── queries.go             # Handler GetByID
├── queries/
│   └── get_by_id.go           # Struct da query
├── repositories/
│   ├── mysql_write.go         # Repository MySQL
│   └── mongodb_read.go        # Repository MongoDB
├── router/
│   └── routes.go               # Rotas HTTP
└── utils/
    └── constants.go           # Constantes (queue, collection)
```

### 2. Arquivos a Criar/Modificar

**Novos arquivos:**
- `src/product/commands/create.go` - Struct do comando
- `src/product/events/created.go` - Struct do evento
- `src/product/handlers/commands.go` - Handler Create
- `src/product/handlers/queries.go` - Handler GetByID
- `src/product/queries/get_by_id.go` - Struct da query
- `src/product/repositories/mysql_write.go` - Repository MySQL
- `src/product/repositories/mongodb_read.go` - Repository MongoDB
- `src/product/router/routes.go` - Rotas
- `src/product/utils/constants.go` - Constantes
- `src/projections/product.go` - Consumer de eventos

**Modificar:**
- `src/projections/runner.go` - Adicionar consumer de product
- `src/router/setup.go` - Adicionar rotas de product

### 3. Constantes Necessárias

```go
// src/product/utils/constants.go
const (
    QueueProductCreated = "product_created"
    CollectionProducts  = "products"
)
```

### 4. Fluxo Idêntico

- **Command**: POST /products → MySQL → Event → RabbitMQ → Projection → MongoDB
- **Query**: GET /products/:id → MongoDB → Return JSON

---

## Dependências e Conexões

### Conexões de Banco de Dados

- **MySQL**: `src/database/mysql.go` → `config.MySQLConnectionString`
- **MongoDB**: `src/database/mongodb.go` → `config.MongoDBConnectionString` + `config.MongoDBDatabaseName`
- **RabbitMQ**: `src/database/rabbitmq.go` → `config.RabbitMQHost/Port/User/Password`

### Event Bus

- **Publish**: `src/eventbus/eventbus.go` → `PublishEvent(queue, event)`
- **Consume**: `src/eventbus/eventbus.go` → `ConsumeEvent(queue, handler)`

### Responses

- **JSON**: `src/responses/responses.go` → `JSON(ginContext, status, data)`
- **Error**: `src/responses/responses.go` → `Error(ginContext, status, err)`

---

## Notas Importantes

1. **Separação de Responsabilidades:**
   - **Command Side**: MySQL (write) → RabbitMQ → MongoDB (projection)
   - **Query Side**: MongoDB (read) apenas

2. **Event Sourcing:**
   - Eventos são publicados após escrita no MySQL
   - Projections consomem eventos e atualizam MongoDB
   - MongoDB é sempre atualizado via eventos, nunca diretamente

3. **Configuração:**
   - Tudo parametrizado via `.env`
   - Constantes para queues e collections
   - Sem hardcode de nomes de database/collection

4. **Tags BSON:**
   - Eventos precisam de tags BSON para mapeamento correto no MongoDB
   - Exemplo: `CreatedAt time.Time `json:"created_at" bson:"createdat"``

5. **Inicialização:**
   - `main.go` inicializa conexões e consumers
   - Consumers rodam em goroutines (long-running)
   - Router inicia após tudo configurado
