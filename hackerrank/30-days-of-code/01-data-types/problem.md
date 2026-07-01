Objective
Today, we're discussing data types. Check out the Tutorial tab for learning materials and an instructional video!
Tutorial de datatypes porém em java: https://www.youtube.com/watch?v=XLCka0noTY4


[INICIO_DO_MARKDOWN]

# Desafio: Data Types

## Tarefa

Complete o código no editor abaixo.

As variáveis `i`, `d` e `s` já estão declaradas e inicializadas para você.

Você deve:

- Declarar 3 variáveis:
  - uma do tipo `int`;
  - uma do tipo `double`;
  - uma do tipo `String`.

- Ler 3 linhas de entrada a partir de `stdin`, de acordo com a sequência dada na seção **Formato de Entrada**, e inicializar suas variáveis.

- Usar o operador `+` para realizar as seguintes operações:
  - imprimir a soma de `i` com sua variável `int` em uma nova linha;
  - imprimir a soma de `d` com sua variável `double`, com uma casa decimal, em uma nova linha;
  - concatenar `s` com a string lida como entrada e imprimir o resultado em uma nova linha.

## Observação

Se você estiver usando uma linguagem que não suporta o uso de `+` para concatenação de strings, como C, você pode simplesmente imprimir uma variável imediatamente após a outra na mesma linha.

A string fornecida no editor deve ser impressa primeiro, imediatamente seguida pela string que você leu como entrada.

## Formato de Entrada

A primeira linha contém um inteiro que você deve somar com `i`.

A segunda linha contém um `double` que você deve somar com `d`.

A terceira linha contém uma string que você deve concatenar com `s`.

## Formato de Saída

Imprima:

- a soma dos dois inteiros na primeira linha;
- a soma dos dois `double`, com uma casa decimal, na segunda linha;
- as duas strings concatenadas na terceira linha.

## Exemplo de Entrada

    12
    4.0
    is the best place to learn and practice coding!

## Exemplo de Saída

    16
    8.0
    HackerRank is the best place to learn and practice coding!

## Explicação

Quando somamos os inteiros `4` e `12`, obtemos o inteiro `16`.

Quando somamos os números de ponto flutuante `4.0` e `4.0`, obtemos `8.0`.

Quando concatenamos `HackerRank` com `is the best place to learn and practice coding!`, obtemos:

    HackerRank is the best place to learn and practice coding!

## Aviso importante

Você não passará neste desafio se tentar atribuir os valores do exemplo diretamente às suas variáveis, em vez de seguir as instruções acima e ler a entrada a partir de `stdin`.

[FIM_DO_MARKDOWN]