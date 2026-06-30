In this video, we discuss 3 places you can go to find information about security vulnerabilities, review the OWASP Top Ten, and overview a go application with known vulnerabilities. 

Golang security resources:
Go CVE list: https://www.cvedetails.com/vulnerability-list/vendor_id14185/golang.html?__cf_chl_f_tk=eCBVKFV.s8WSuoOA8TKmAG.zA9kLjCHKWd29gzuuWd0-1782842396-1.0.1.1-p0ZQpF1a.2ATMOVnrnNa3wbuX2caQWIJkwb5zBxu1qg
Synk Vulnerability DB: www.synk.io/vuln?type=golang
Golang-announce: https://groups.google.com/g/golang-announce?pli=1

OWASP Top Ten: https://owasp.org/www-project-top-ten/ 
Injection
Broken Authentication
Sensitive Data Exposure
XML External Entities
Broken Access Control
Security Misconfiguration
Cross-site scripting
Insecure Deserialization
Using Components with Known Vulnerabilities
Insufficient Logging and Monitoring 

GitHub Repo (code & slides): https://github.com/tebeka/talks/tree/master/secure-go-code


[INICIO_DO_MARKDOWN]

# Segurança em Go e vulnerabilidades comuns

## Foco da palestra em Go

Nesta palestra, o foco vai ser em Go, mas muitas dessas coisas também se aplicam a outras linguagens.

A maior parte das coisas que vamos falar pode ser implementada em outras linguagens também. Mas o que eu realmente gosto em Go é que ele tem segurança incorporada no time que constrói Go. Ele vem de uma empresa, o Google, que tem muito foco em segurança, e a própria linguagem é bem rápida em atualizar vulnerabilidades de segurança e outras coisas quando elas aparecem.

## Política de segurança do Go

Go, como eu disse, tem uma política de segurança.

Se eu clicar nela, é claro que o navegador vai abrir em uma tela que vocês não veem.

Então, essa é a mentalidade de segurança e essa é a política de segurança do Go: como reportar bugs de segurança, qual é o processo de divulgação, etc.

Tem até a chave PGP, então você pode assinar coisas e fazer isso de forma segura quando divulga questões de segurança. Nem todas as linguagens ou empresas têm isso, mas Go é realmente focado nisso. Então, se você encontrar alguma coisa, existe uma forma de reportar.

## Onde encontrar informações sobre vulnerabilidades

A pergunta é: onde você pode encontrar informações sobre vulnerabilidades de segurança?

Existem vários lugares.

Um deles é a lista de CVE para vulnerabilidades de segurança em Go. Há fontes de dados de vulnerabilidades de segurança e segurança diária na linguagem Go, e você pode procurar informações sobre isso.

Outra coisa é consultar dados sobre vulnerabilidades em pacotes que não estão na standard library, mas que são pacotes Go. Então você pode olhar o seu pacote favorito que está usando e ver se está usando algo que pode ter problemas de segurança.

A mailing list `golang-announce` — e sim, mailing lists ainda existem — ou o user group, tem anúncios sobre segurança e sobre correções de segurança futuras.

Então você pode, e deve, ficar atualizado sobre aquilo que você usa.

## OWASP Top 10

Vamos focar em algumas das vulnerabilidades mais conhecidas, que são conhecidas como OWASP Top 10.

O OWASP Top 10 é a lista das 10 principais vulnerabilidades encontradas. Para mim, é uma lista muito triste, porque as coisas que estão lá existem há muito tempo.

E toda vez que olhamos para a lista, dizemos: “Sério? Isso é 2021, por favor. Podemos parar com isso?”

## As dez principais vulnerabilidades

Quais são as Top 10?

A primeira é `injection`, como em SQL injection, que já existe há bastante tempo.

Também temos:

- `broken authentication`;
- `sensitive data exposure`;
- `XML external entities`;
- `broken access control`;
- `security misconfiguration`;
- `XSS`, que é `cross-site scripting`;
- `insecure deserialization`;
- uso de componentes com vulnerabilidades conhecidas;
- `insufficient logging and monitoring`.

Se você olhar para todas essas coisas, elas não são revolucionárias. Elas simplesmente estão ali. Mas quando você começa a olhar para o seu código, pode descobrir que na verdade está violando algumas delas.

Então essa é uma boa checklist para começar e ver se você está fazendo algo errado ou não.

## Agrupando as vulnerabilidades

Eu gosto de agrupá-las em quatro grupos.

O primeiro grupo fala sobre `input`. Aqui entram:

- `injection`;
- `XML external entities`;
- `insecure deserialization`.

Ou seja: como você lida com entrada de dados.

O segundo grupo é sobre `output`, que inclui:

- `cross-site scripting`;
- `sensitive data exposure`.

Depois temos autenticação, que fala sobre:

- `authentication`;
- `access control`, que também é conhecido como `authorization`.

E, finalmente, temos o que eu chamo de infraestrutura, porque “miscellaneous” é um nome ruim.

Nesse grupo entram:

- `security misconfiguration`;
- uso de componentes com vulnerabilidades conhecidas;
- `insufficient logging and monitoring`.

## Aplicação de exemplo

O que vamos fazer?

Vamos olhar para uma aplicação que eu escrevi.

E eu preciso dizer que foi realmente doloroso para mim escrever essa aplicação, porque eu tive que superar muitos hábitos internos de como escrever código para conseguir colocar todos os erros ruins aqui. Mas eu consegui.

Essa aplicação é basicamente como um journal, um web journal, com uma camada de banco de dados que trabalha com um banco SQLite para escrever dados no banco.

Ela também tem algo que define alguns tipos, existe até um teste aqui, e tem o servidor HTTP que serve o HTTP. Ele tem vários endpoints que vamos olhar.

O que vamos fazer é olhar para esses aspectos e ver o que podemos fazer.

[FIM_DO_MARKDOWN]