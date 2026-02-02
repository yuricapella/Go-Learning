# Guia de Commits - Padrão do Projeto

Padrão baseado em **Conventional Commits** com mensagens em português.

## Conventional Commits

Este projeto segue o padrão [Conventional Commits](https://www.conventionalcommits.org/), uma especificação para mensagens de commit que facilita a leitura e automação.

### Formato

```
<tipo>(<escopo>): <mensagem principal>

- <detalhe 1>
- <detalhe 2>
```

## Tipos de Commit

- `feat`: Nova funcionalidade ou aula
- `refactor`: Refatoração de código sem mudança de comportamento
- `fix`: Correção de bug
- `docs`: Documentação
- `style`: Formatação, sem mudança de código
- `test`: Adição ou correção de testes
- `chore`: Tarefas de manutenção

## Regras Essenciais

1. **Escopo**: Nome da pasta/módulo com underscore (ex: `4_funcoes`, `14_interface`)
2. **Mensagem**: Verbo no infinitivo ou particípio passado, sem ponto final
3. **Corpo**: Bullet points diretos descrevendo o que foi feito
4. **Sem redundâncias**: Não mencione "em Go" ou "Go" (projeto já é sobre Go)
5. **Dependências**: Se adicionar, mencione go.mod e go.sum no final
6. **Foco no que foi feito**: Evite linhas genéricas. Foque apenas no que foi criado/modificado concretamente

## Exemplos Reais

### Nova aula
```
feat(14_interface): adiciona conteúdo de interfaces da aula e exemplos didáticos sobre interfaces

- Adicionado main.go como arquivo orquestrador que chama as funções Demonstrar* de cada tópico
- Criada pasta interfaces/ contendo cinco arquivos especializados
- Adicionado sintaxes.go exclusivamente com código puro e funções sintáticas
```

### Refatoração
```
refactor(4_funcoes): refatora estrutura do módulo distribuindo conteúdo didático em arquivos separados

- Refatorado main.go distribuindo cada tópico em arquivos separados na pasta funcoes/
- Criada pasta funcoes/ contendo seis arquivos especializados
- Transformado main.go em arquivo orquestrador
```

### Com dependências
```
feat(15_aplicacao_linha_de_comando): adiciona aula sobre aplicações CLI

- Criada pasta app_aula/ com aplicação CLI de exemplo
- Adicionado exemplos_didaticos/ com demonstrações
- Atualizados go.mod e go.sum com dependência github.com/urfave/cli
```

## Template

```
<tipo>(<escopo>): <mensagem principal>

- <o que foi feito>
- <o que foi feito>
- <o que foi feito>
```
