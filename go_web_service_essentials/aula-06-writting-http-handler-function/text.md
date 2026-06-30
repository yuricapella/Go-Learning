[INICIO_DO_MARKDOWN]

# HTTP handler function e routing em Go

## Funções handler HTTP

`HTTP handler function`.

Nós escrevemos nossos handlers como uma função, e essa função normalmente recebe dois parâmetros.

O primeiro é `w`, que é um `http.ResponseWriter`. Esse é o objeto no qual você vai escrever. Nele, você pode definir os `HTTP headers`, o `status code`, etc.

O segundo parâmetro é `r`, que é o `http.Request`. Nele, você também tem os headers, pode ter o body da request, tem o path, com as coisas que entraram no caminho que ele está tentando acessar. Você pode consultar informações sobre basic authentication e qualquer coisa que se pareça com uma HTTP request.

## Handler de health check

Esse é o handler que o professor está fazendo. Ele comenta que está adicionando bastante ao código.

É um handler que verifica a saúde do servidor, o `health` do servidor.

Existem muitos sistemas de monitoramento, como `Prometheus` e outros. O que eles fazem é: você informa “este é o servidor, esta é a URL”, e de tempos em tempos eles vão lá e chamam essa URL, esperando que ela diga `OK`.

No caso da aula, o handler simplesmente vai responder `OK`.

Mas, em um caso real, você poderia, por exemplo:

- rodar uma dummy query no banco de dados;
- rodar uma query simples no banco de dados para garantir que a conexão com o banco está funcionando;
- tentar chamar a authentication API para verificar se há conexão com a API de autenticação;
- e assim por diante.

Então esse é apenas um `health handler` básico. Ele não faz métricas nem outras coisas.

## Routing

No `main`, o que será feito é informar ao pacote `http` como fazer o que é conhecido como `routing`.

Ou seja: se alguém acessar `/health` no web server, chame esse handler.

Isso é o routing.

O professor comenta que qualquer web server que você escolher terá algum tipo de noção de routing.

## Routing simples no servidor HTTP de Go

O servidor HTTP em Go tem um mecanismo de routing muito simples.

Você pode fornecer algo como um caminho, e ele será um match exato. Ou você pode terminar com uma `/`, e então tudo abaixo daquele caminho será correspondido.

Se você vem de frameworks que suportam regular expressions nos routers, esse router não tem suporte para regular expressions.

Ele também não tem suporte para dizer algo como:

- para uma request `GET`, use este handler;
- para uma request `POST`, use aquele handler.

Ele não tem isso. É um router muito, muito simples.

Mas, se você está escrevendo APIs pequenas, APIs internas, normalmente isso é suficiente.

## Quando usar routers externos

Se você quiser algo além disso, existem muitas opções por aí.

O professor menciona um router popular chamado `Gorilla Mux`.

Esse router suporta:

- regular expressions;
- métodos específicos;
- variáveis dentro das rotas;
- outros recursos.

Você pode adicionar variáveis dentro das routes e definir algum tipo de regular expression.

Se você vem mais de outras linguagens, provavelmente está acostumado com esse tipo de coisa.

## Por que o router padrão não tem mais recursos

O professor diz para não tentar pedir ao time de Go para adicionar mais coisas ao router padrão, porque eles não vão fazer isso.

Uma das razões é que, especialmente com regular expressions, é muito difícil saber quanto CPU uma regular expression vai consumir.

Ele pede desculpas à Cloudflare por usar esse exemplo, mas diz que é o caso que ele lembra, embora não sejam os únicos.

A Cloudflare teve uma queda que, segundo ele, também durou algumas horas. A causa foi alguém escrever uma regular expression que esgotou a CPU usada para o servidor HTTP e HTTPS.

Então regular expressions podem ser realmente perigosas. É preciso saber como trabalhar com elas.

O time de Go diz algo como: “não vamos dar essa arma para você atirar no próprio pé”.

Se você quiser fazer isso, pode pegar um pacote de terceiros e ficar à vontade para usar. Mas o router padrão não vai incluir isso.

## Experiência prática com o router padrão

O professor comenta, por experiência, que isso não é um problema tão grande.

Se você não tem um routing muito sofisticado na sua aplicação, o router padrão costuma ser suficiente.

No pior caso, você faz uma verificação:

- se é uma request `POST`, faça isso;
- caso contrário, faça aquilo.

E pronto.

## Iniciando o servidor

Depois do routing, é necessário iniciar o servidor.

Existem funções como:

- `ListenAndServe`;
- `ListenAndServeTLS`, que é para HTTPS.

No caso de `ListenAndServeTLS`, é necessário fornecer um certificate, uma key, etc.

Na aula, será usado o caso básico.

Você informa ao servidor em que endereço ele deve escutar.

O que vem antes dos dois-pontos é a interface. Se estiver vazio, significa que ele vai escutar em todas as interfaces.

Há também um parâmetro, um global handler, que você pode passar. Na maior parte do tempo, ele será `nil`.

Então esse é um servidor HTTP básico.

[FIM_DO_MARKDOWN]