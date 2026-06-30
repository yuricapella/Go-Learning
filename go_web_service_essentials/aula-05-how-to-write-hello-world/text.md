[INICIO_DO_MARKDOWN]

# Introdução a um programa simples em Go

## Pedido inicial e objetivo da aula

Em seguida, ele explica que a aula vai passar linha por linha pelo código e explicar o que está acontecendo. Depois, será mostrado como executar o código. A intenção é garantir que todo mundo consiga rodar esse código simples.

## Declaração de pacote em Go

O código começa com a declaração de pacote.

O professor comenta que Go foi construído para times grandes. Cada pedaço de código fica dentro de um pacote. Então, qualquer arquivo Go que você começa normalmente começa com a palavra-chave `package`.

Depois disso, você define o nome do pacote. É possível definir muitos nomes de pacotes, mas o pacote `main` tem um significado especial.

O pacote `main` é o pacote que o runtime de Go vai procurar quando estiver executando o programa.

Então:

- `package main` define o pacote.
- O pacote `main` tem um significado especial para execução.
- É nele que o runtime de Go procura o ponto de entrada do programa.

## Imports e uso de outros pacotes

Depois da declaração do pacote, normalmente vêm os imports.

`import` é a forma de pegar código de outro pacote. 

Quando você olha para a linguagem Go e começa um programa em Go, ela tem muito pouco embutido. Existem poucas funções predefinidas e built-in que você pode usar. Fora isso, a funcionalidade está em outros pacotes, e você usa uma declaração `import` para acessá-la.

O professor comenta que é possível fazer o `import` de formas diferentes. Se você está importando apenas um pacote, pode fazer em uma linha só, e isso também está correto.

Mas ele diz que pegou o hábito de fazer tudo dentro dos parênteses, porque a maioria dos programas importa mais do que apenas um pacote.

## Definição de função com `func`

Depois dos imports, há a definição de uma função.

A definição de uma função começa com `func`.

O professor compara com outras linguagens:

- Algumas linguagens usam `def`.
- Rust usa `fn`.
- JavaScript usa `function`.
- Cada linguagem escolhe o que quer.

Em Go, usa-se `func`.

Depois vem o nome da função, que neste caso é `main`.

Essa função não tem argumentos. Então não há argumentos dentro dos parênteses, mas os parênteses ainda são obrigatórios. É preciso escrevê-los.

Depois vem a chave de abertura, o código dentro da função e a chave de fechamento.

## A função `main`

A função `main`, dentro do pacote `main`, é uma função especial.

Essa é a função que o runtime de Go executa quando o programa começa. Ele procura por `main` e vai executá-la.

Dentro do corpo de `main`, neste exemplo, há apenas uma instrução. O programa está chamando `fmt.Println`.

## Chamada de `fmt.Println`

O professor explica a chamada `fmt.Println`.

`fmt` é o nome do pacote. Às vezes ele comenta que é pronunciado como “F”.

Em Go, é necessário escrever o nome do pacote antes do nome da função. Em algumas linguagens, você pode usar apenas o nome da função, como `print` ou `len`, porque você importou. Mas em Go, por design, o nome do pacote sempre aparece antes da função ou do tipo que você usa.

Então a forma é:

- nome do pacote;
- ponto;
- nome da função.

No caso:

- `fmt` é o pacote;
- `.` é o ponto;
- `Println` é a função.

## Letras maiúsculas e símbolos exportados

O professor chama atenção para o `P` maiúsculo em `Println`. Isso também é algo exigido pela linguagem.

Em Go, não existem `private`, `protected` e várias outras categorias como em algumas linguagens. Em vez disso, há símbolos exportados e não exportados.

Símbolos exportados podem ser usados de fora do pacote. Isso significa que podemos usar `Println` a partir do nosso próprio pacote.

Símbolos não exportados só podem ser usados dentro do mesmo pacote. Então ninguém pode chamar `main`, a não ser outro código dentro desse mesmo pacote `main`.

A distinção entre símbolos exportados e não exportados é simples:

- Se começam com letra maiúscula, são exportados.
- Caso contrário, não são exportados.

E é isso. Não há mais nada para lembrar. Não há palavra-chave especial. Não há algo extra para aprender.

Quando você vê uma letra maiúscula, sabe que é exportado. Quando vê letra minúscula, é interno ao pacote.

## Funções de impressão no pacote `fmt`

O pacote `fmt` tem várias funções de impressão. É um pacote usado para imprimir coisas.

O professor comenta que existem várias funções de impressão:

- As que começam com `P` imprimem na saída padrão.
- Algumas começam com `F` e imprimem para algo que se parece com um arquivo. Isso é chamado de `IO writer`, e será discutido depois.
- Algumas começam com `S` e criam uma nova string. É uma forma de formatar uma nova string.

No caso de `Println`, o `ln` no final significa que será adicionada uma nova linha ao que está sendo impresso.

Também existe `Print`, sem a nova linha, que imprime sem adicionar essa quebra de linha.

## Parâmetros da função e strings

A chamada de função usa parênteses de abertura e fechamento. Dentro deles ficam os parâmetros da função.

Neste caso, há um único parâmetro para a função: uma string.

A string começa com aspas duplas. O professor observa que essa não é a única forma de criar strings em Go, mas é uma das duas formas.

Ele também comenta que, como dá para ver pelo símbolo de coração no final, strings em Go têm suporte a Unicode.

Go tem suporte a Unicode. O professor reforça que é uma linguagem moderna, que está por aí há algum tempo, e que o Google lida muito com Unicode. Por isso, é possível ver essa influência e por que isso é importante.

## Formatação do código com `gofmt`

Depois de explicar o código, o professor fala sobre a formatação.

Você pode pensar que quer mudar algumas coisas, talvez adicionar espaços ou mudar a posição de algo. Mas, quando ele salva o código, tudo volta.

Existe uma ferramenta chamada `gofmt`, como o pacote `fmt`, que formata o código.

Quase todo código Go por aí, praticamente sem exceção, passa por essa ferramenta de formatação.

O editor ou IDE que você escolher — seja o que o professor está usando, Visual Studio Code ou GoLand — vai formatar o código do mesmo jeito.

Há um ditado na comunidade Go:

- todos amam o `gofmt`;
- todos odeiam o `gofmt`.

A parte boa é que você não precisa pensar. Esse é o jeito como código Go é escrito.

O professor menciona algumas características:

- usa tabs;
- as chaves ficam no topo;
- normalmente não se usam ponto e vírgula, mesmo que você possa colocar.

Ele comenta que, se você colocar um ponto e vírgula, isso ainda é Go válido. Mas, quando salva o arquivo, o `gofmt` remove o ponto e vírgula, porque ele não é necessário.

Então o código fica indentado de forma uniforme. Isso é bom porque as pessoas se acostumam a ver uma única maneira de escrever código.

## Opiniões sobre indentação e a referência a Silicon Valley

O professor diz que todo mundo também odeia isso porque todo mundo tem sua própria opinião sobre como o código deveria ser indentado.

Ele compara com um episódio da série `Silicon Valley`, sobre a startup com o Piper.

No episódio, um personagem termina com a namorada porque ele usa tabs e ela usa espaços.

O professor brinca que, se eles estivessem usando Go, provavelmente ainda estariam juntos.

[FIM_DO_MARKDOWN]