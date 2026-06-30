In our previous video we wrote our http handler. Now, we want to test it.

More information: https://pkg.go.dev/net/http/httptest

[INICIO_DO_MARKDOWN]

Testando handlers HTTP em Go
Criando os objetos necessários para testar o handler
Quando você quer testar um handler, precisa criar alguns objetos.

Lembre que um handler é apenas uma função que recebe um w e um request.

Se formos olhar para ele, ele recebe:

w;
request.
E é isso.

O w é uma interface. ResponseWriter é uma interface. Já request é uma estrutura específica, mas você precisa fornecer esses dois valores.

Então você consegue testar handlers sem realmente rodar o servidor.

Usando httptest.NewRecorder e httptest.NewRequest
Você pode fazer:

w = httptest.NewRecorder();
r = httptest.NewRequest(...).
No seu request, você informa:

o método;
o path;
e pode passar um body como algum tipo de io.Reader.
No nosso caso, nós realmente precisamos do body.

Criando o body com strings.NewReader
O texto usado no exemplo é:

“Have you seen Who's on First?”
Alguém viu “Who's on First”?

É engraçado, certo? É antigo e engraçado.

Exatamente, é sobre Abbott and Costello.

Então fazemos um strings.NewReader em cima do texto. Assim obtemos um body.

Chamando o handler diretamente
Agora chamamos o tokenizeHandler com w e r.

Ou seja:

preparamos o w e o r;
depois rodamos o handler com esses valores.
Obtendo a resposta
Depois disso, pegamos a resposta de volta a partir de w.

Usamos algo como:

w.Result().
Agora temos o resultado e podemos começar os testes.

Verificando o status code
Podemos fazer uma verificação com:

require.Equal(http.StatusOK, res.StatusCode, ...).
Por exemplo.

Também podemos adicionar uma mensagem dizendo que isso é o status code, para sabermos o que falhou.

Precisamos importar também net/http.

Eu prefiro usar as constantes, e não o 200. Acho mais legível.

Decodificando o JSON esperado
Depois podemos criar um valor expected.

Também precisamos de encoding/json.

Usamos um novo decoder para decodificar a resposta.

Essa é a parte usual que você faz:

cria uma request com w e r;
chama o handler;
pega o resultado de volta;
começa a verificar o que está saindo.
[FIM_DO_MARKDOWN]