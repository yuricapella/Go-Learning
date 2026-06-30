By using build tags, you can create different versions of your Go application from the same source code and toggle between them with ease.

In this example MikiTebeka shows us how to use both explicit and implicit build tags when building Go executables. 

View the code used in this video: https://github.com/353words/building-go

Golang 1.17 Release Notes: https://go.dev/doc/go1.17
Package Build Documentation: https://pkg.go.dev/go/build
Pprof Library: https://pkg.go.dev/net/http/pprof

go run . nao funciona
mas go run -tags profile . funciona
porque?

ele mostra um arquivo prof.go com os codigos:

//go:build profile
é um comentário...

package main

// Will export /debug/pprof/
import _ "net/http/pprof"

ai no terminal ele usa
go env GOOS GOARCH
e retorna
darwin
amd64
entao ele explica que dos 3 arquivos
by_darwin.go
by_windwos.go
by_linux.go
apenas o by_darwin.go vai executar porque é o tipo de maquina
codigo dele:
package main

var Founder = "Steve"

ai ele usa o comando no terminal novamente
go run .
e em outro terminal usa
curl localhost:8080/by 
e no terminal apos rodar curl aparece Steve

e depois recomenda o site: github.com/etcd-io/bbolt
key value database, many files, go operation systems e archteture mas não compreendi o intuito do video em si, nao ficou claro e nao faz sentido esses códigos e comandos kkkkk


[INICIO_DO_MARKDOWN da transcrição]

# Build Tags em executáveis Go

## Introdução à série

Bem-vindos à nossa série sobre construção de executáveis Go.

Desta vez, vamos falar sobre `build tags`.

## Exemplo com servidor web e profiling

Digamos que você tenha um web server e queira adicionar profiling usando a biblioteca `http/pprof`.

Isso vai adicionar um endpoint `/debug/pprof` ao seu web server.

Mas, como isso pode ser uma questão de segurança, você não quer construir isso por padrão.

Uma solução é adicionar uma `build tag`.

Você pode adicionar um comentário como:

- `//go:build profile`

Aqui, `profile` é o nome da tag.

O formato `//go:build` é novo desde o Go 1.17, mas o formato antigo com `// +build` também é suportado por compatibilidade com versões anteriores.

## Testando sem a build tag

Depois que temos isso, se rodarmos nosso servidor normalmente:

- `go run server`

E fizermos uma requisição com `curl` para:

- `http://localhost:8080/debug/pprof`

Vamos ver que a página não foi encontrada.

## Testando com a build tag

No entanto, se fizermos:

- `go run -tags profile`

E rodarmos o servidor, agora, ao acessar essa página, ela existe.

As `build tags` são suportadas nos subcomandos:

- `go run`;
- `go test`;
- e, é claro, `go build`.

## Tags explícitas

Essas são tags explícitas. Ou seja, você as define explicitamente pelo nome.

Mas também existem tags implícitas.

## Tags implícitas

Essas tags implícitas são principalmente para sistema operacional e arquitetura.

Então, se você fizer algo como consultar `GOOS` e `GOARCH`, verá que estou rodando em:

- `darwin`;
- `amd64`.

## Build por sistema operacional

Aqui eu tenho três arquivos:

- `by_darwin`;
- `by_linux`;
- `by_windows`.

Isso significa que, quando eu fizer build nesta máquina, somente o arquivo `by_darwin` será construído.

Então, se eu rodar meu servidor com:

- `go run .`

E depois fizer uma requisição com `curl` para:

- `http://localhost:8080/by`

Vou ver `Steve`.

## Uso de tags implícitas para máquinas diferentes

Essas `build tags` implícitas são principalmente para construir em máquinas diferentes.

Então, se você tem configurações diferentes para máquinas diferentes, pode usar isso.

Você pode olhar, por exemplo, para o Bolt. Ele é um banco de dados key-value. Você verá que ele tem muitos arquivos que terminam com o sistema operacional Go ou com a arquitetura Go.

## Encerramento

É isso sobre `build tags`.

Muito obrigado e nos vemos na próxima vez.

[FIM_DO_MARKDOWN]