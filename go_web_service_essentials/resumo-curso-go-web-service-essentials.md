# Perguntas e Aprofundamentos do Curso Go Web Service Essentials

Este documento complementa as aulas do curso com perguntas que ficaram durante o estudo e respostas mais diretas para fixar os conceitos. O objetivo nao e repetir o conteudo das aulas, mas registrar os pontos que exigiram mais interpretacao.

## Aula 1: Structs, Zero Value e Literal Construction

### A aula sobre structs se resume a preferir `var` para zero value e evitar inicializacao pela metade?

Sim, esse e um dos principais pontos da aula. A ideia nao e dizer que `example{}` esta sempre errado, mas que `var e example` comunica melhor a intencao quando o objetivo e obter o zero value.

Em uma struct simples, `var e example` e `example{}` podem produzir o mesmo valor final: todos os campos ficam com seus zero values.

`var e example` comunica que voce quer o zero value daquele tipo. Ja `example{}` usa uma construcao literal vazia. Para structs, isso geralmente cai no mesmo resultado pratico, mas a ideia da aula e escrever codigo que deixe clara a intencao para quem esta lendo.

Esse cuidado fica mais importante porque, em outros tipos, uma construcao vazia nem sempre representa exatamente o mesmo conceito de zero value. Por exemplo, um slice declarado com `var` pode ser `nil`, enquanto um slice criado com literal vazio pode estar vazio, mas nao ser `nil`.

Sobre inicializacao pela metade: o problema e aumentar o risco de usar ou retornar um valor incompleto. Quando os dados ja sao conhecidos, a aula recomenda reunir os valores primeiro e construir a struct de uma vez.

## Aula 2: Funcoes Literais

### Em Go, funcoes tambem sao valores que podem ser guardados em variaveis?

Sim. Uma funcao pode ser criada anonimamente, guardada em uma variavel, passada como argumento ou executada imediatamente.

Quando uma funcao anonima e atribuida a uma variavel, ela so executa quando essa variavel e chamada como funcao. Isso reforca a ideia de que, em Go, funcoes podem circular pelo codigo como valores.

### Qual e a diferenca entre declarar uma funcao anonima e executar essa funcao na hora?

Porque a primeira parte apenas declara a funcao. Os parenteses finais fazem a chamada.

Sem os parenteses finais, voce apenas criou um valor do tipo funcao. Com os parenteses, voce executa esse valor.

### No `defer`, a funcao anonima usa o valor da variavel no momento da declaracao ou no final?

`defer` executa a chamada no final da funcao atual.

Se a funcao adiada acessa uma variavel externa diretamente, ela enxerga o valor que essa variavel tiver quando o `defer` realmente executar.

Se voce quiser preservar o valor do momento em que o `defer` foi declarado, passe esse valor como parametro para a funcao anonima. Assim, o valor e copiado no momento da chamada adiada.

Essa e a diferenca entre capturar a variavel externa e passar um valor para os parametros da funcao anonima.

## Aulas 3 e 4: Serializacao e JSON

### Serializar e transformar dados em bytes, e desserializar e transformar bytes de volta em valores Go?

Serializar e transformar um valor em um formato transmissivel ou armazenavel, normalmente bytes.

Desserializar e fazer o caminho inverso: pegar esses bytes e reconstruir um valor que a linguagem consiga usar.

Em APIs HTTP, esse formato costuma ser JSON. O cliente envia JSON, a aplicacao Go converte esse JSON para structs, processa a regra de negocio e depois converte a resposta de volta para JSON.

O ponto importante e que o frontend ou outro servico nao envia "codigo Go". Ele envia dados em um formato combinado. JSON e um desses formatos.

### Quando usar `json.NewDecoder(r.Body).Decode(&valor)`, `json.NewEncoder(w).Encode(resposta)`, `Marshal` e `Unmarshal`?

Use `Decoder` quando os dados estao vindo de um fluxo, como o body de uma request HTTP. Nesse caso, `r.Body` contem os bytes recebidos na entrada da API, e `Decode(&valor)` preenche uma struct Go.

Use `Encoder` quando voce quer escrever a resposta em um destino, como o `ResponseWriter`. Nesse caso, `Encode(resposta)` transforma um valor Go em JSON e escreve esse JSON na response.

`Marshal` e `Unmarshal` sao usados quando voce ja tem os dados em memoria, normalmente como `[]byte`.

Na pratica, em handlers HTTP, `Decoder` costuma aparecer ao ler o body da request, e `Encoder` costuma aparecer ao escrever a resposta. `Marshal` e `Unmarshal` aparecem mais quando voce quer transformar valores em bytes, ou bytes em valores, sem escrever diretamente em um stream HTTP.

### Big endian e little endian importam no dia a dia de uma API JSON?

Na maioria dos casos de API JSON, nao. Mas a intuicao de que isso tem a ver com "ordem" esta correta.

Big endian e little endian falam sobre a ordem dos bytes ao representar numeros em formato binario. Isso importa em protocolos binarios, arquivos binarios e comunicacao de baixo nivel.

Quando a API usa JSON, esses detalhes ficam abstraidos pelo formato. Entao, para APIs web comuns, o mais importante e entender serializacao e desserializacao, nao decorar endianess.

## Aula 5: Hello World, Pacotes e Executaveis

### `go run .` procura `package main` e `func main()` para executar o programa?

Um programa Go executavel precisa de `package main` e de uma funcao `main`.

O pacote `main` indica que aquele codigo pode gerar um executavel. A funcao `main` e o ponto de entrada chamado quando o programa inicia.

### Em Go, letra maiuscula e minuscula funcionam como `public`, `private` ou `protected` do Java?

Teoricamente, sim, da para fazer uma comparacao parcial com Java, mas nao e uma equivalencia perfeita.

Em Go, a visibilidade e definida pela primeira letra do nome:

- nomes com letra maiuscula sao exportados e funcionam de forma parecida com `public`;
- nomes com letra minuscula ficam restritos ao proprio pacote e se parecem mais com o acesso package-private/default do Java.

O que nao existe em Go e um equivalente direto a `protected`, nem um `private` limitado a uma classe ou arquivo. Se algo comeca com letra minuscula, qualquer arquivo dentro do mesmo pacote pode acessar.

Entao a comparacao mais justa seria:

- Go maiusculo: parecido com `public`;
- Go minusculo: parecido com package-private/default;
- Go nao tem equivalente direto a `protected`;
- Go nao tem `private` por arquivo ou por struct.

A menor unidade de encapsulamento em Go e o pacote.

### `gofmt` e tipo um lint?

Da para comparar superficialmente porque ambos sao ferramentas de qualidade de codigo, mas eles resolvem problemas diferentes.

`gofmt` e um formatador automatico. Ele padroniza a aparencia do codigo Go: indentacao, quebras de linha, espacos e estilo geral.

Lint tem outro objetivo: apontar possiveis problemas, bugs, inconsistencias ou mas praticas. Entao, em resumo: `gofmt` cuida do formato; linters cuidam mais da qualidade e dos riscos do codigo.

## Aula 6: HTTP Handler, Request e ResponseWriter

### Por que um handler precisa receber `http.ResponseWriter` e `*http.Request`?

Porque um handler fica no meio de uma comunicacao HTTP.

O `Request` representa o que chegou do cliente para o servidor. Ele contem metodo HTTP, URL, headers, body, query params, cookies e outras informacoes da chamada.

O `ResponseWriter` representa a resposta que o servidor vai montar e devolver ao cliente. Por ele, o handler define headers, status code e body da resposta.

Pensando no fluxo:

cliente envia request, handler processa, servidor devolve response.

Portanto, o `ResponseWriter` nao monta a URL. A URL pertence a request. O `ResponseWriter` representa o canal de escrita da resposta.

### A aula recomenda router externo porque criar rotas e middlewares do zero seria mais verboso?

Sim. O Go ja possui um roteador simples na standard library, suficiente para muitos casos.

Routers externos passam a fazer mais sentido quando a API precisa de recursos como parametros de rota, agrupamento de rotas, middlewares mais elaborados, validacao por metodo HTTP ou uma organizacao mais sofisticada.

## Aula 7: Vulnerabilidades

### A aula de vulnerabilidades e uma checklist para revisar uma API antes de entrega-la?

A aula apresentou uma forma de pensar em seguranca durante a revisao de uma aplicacao.

O foco foi mostrar fontes para acompanhar vulnerabilidades e usar listas como OWASP Top 10 como checklist mental para avaliar riscos comuns em aplicacoes web.

### Quais pontos de seguranca devo observar antes de entregar uma API?

Alguns pontos importantes sao:

- validacao de entradas externas;
- protecao contra injection;
- autenticacao;
- autorizacao;
- exposicao de dados sensiveis;
- dependencias vulneraveis;
- configuracao insegura;
- logs e monitoramento.

A aula reforca que seguranca nao e apenas uma etapa final. Ela precisa fazer parte da forma de revisar e construir codigo.

## Aula 8: Build Tags e pprof

### Build tags sao parecidas com profiles de ambiente, como `application.properties`, ou fazem outra coisa?

Elas lembram profiles de ambiente apenas no sentido de permitir variacoes por contexto. Mas tecnicamente fazem outra coisa.

Build tags controlam quais arquivos entram na compilacao.

Elas permitem gerar versoes diferentes do mesmo programa a partir do mesmo codigo-fonte. Por exemplo, uma versao com profiling habilitado e outra sem profiling.

A diferenca para profiles de configuracao e que build tags atuam antes do programa rodar. Elas mudam o que entra no binario. Ja configuracoes de ambiente normalmente mudam valores em runtime.

### `pprof` e um teste, um log ou uma ferramenta para medir desempenho?

`pprof` e uma ferramenta de profiling do Go.

Ele pode lembrar logs no sentido de ajudar a observar o que acontece na aplicacao, mas nao e apenas registro textual de eventos. Tambem pode ser usado durante investigacoes parecidas com testes de desempenho, mas nao e um teste automatizado por si so.

Ele ajuda a investigar como a aplicacao usa recursos como CPU, memoria e goroutines.

O objetivo do `pprof` e diagnosticar desempenho e comportamento da aplicacao em execucao. Ele pode ajudar a encontrar gargalos, consumo excessivo de memoria ou funcoes que usam CPU demais.

### `pprof` conseguiria apontar problemas como teste mal isolado, race condition ou dependencia entre testes?

Indiretamente, talvez ele mostre sintomas, mas nao e a ferramenta certa para afirmar esse tipo de problema.

`pprof` pode mostrar sintomas de desempenho, como excesso de goroutines, CPU alta ou memoria crescendo. Mas ele nao e a ferramenta principal para validar isolamento de testes ou detectar race condition.

Para race condition, use `go test -race`. Para testes que dependem de ordem, estado compartilhado ou falta de reset, o ideal e melhorar os proprios testes e isolar dados, mocks e setup.

## Aula 9: Testando HTTP Handler

### Por que testar um handler com `httptest` sem subir servidor real?

Porque o objetivo do teste e validar o comportamento do handler, nao a rede.

Com `httptest`, o teste fica mais rapido, mais simples e mais isolado. Ele consegue verificar se o handler leu a entrada corretamente, processou a regra esperada e montou a resposta correta.

## Conclusao

O curso abordou os principais fundamentos para criar serviços web em Go.

Na aula 1, foram vistos `struct`, zero value, literal construction e boas práticas para inicializar valores compostos.

Na aula 2, foram estudadas funções literais, funções anônimas, atribuição de funções a variáveis e uso de `defer`.

Nas aulas 3 e 4, o foco foi serialização e desserialização, incluindo transformação de dados em bytes, uso de JSON, `json.NewDecoder`, `json.NewEncoder`, `Marshal` e `Unmarshal`.

Na aula 5, foram apresentados os elementos básicos de um programa executável em Go: `package main`, `func main`, imports, `fmt.Println`, nomes exportados e `gofmt`.

Na aula 6, foram estudados handlers HTTP, `http.ResponseWriter`, `*http.Request`, health check, roteamento com a standard library e uso de routers externos.

Na aula 7, o conteúdo tratou de segurança, fontes de vulnerabilidades, OWASP Top 10 e pontos de revisão antes de entregar uma API.

Na aula 8, foram abordadas build tags, criação de builds diferentes a partir do mesmo código e uso de `pprof` para profiling.

Na aula 9, foi estudado como testar handlers HTTP com `httptest.NewRequest`, `httptest.NewRecorder`, validação de status code, headers e body da resposta.