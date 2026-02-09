# Guia de Commits - Padrão do Projeto

Padrão baseado em **Conventional Commits** com mensagens em português.


## Conventional Commits

Este projeto segue o padrão [Conventional Commits](https://www.conventionalcommits.org/), uma especificação para mensagens de commit que facilita a leitura e automação.

## Formato

```
<tipo>(<escopo>): <mensagem principal>

- Criado/Criada <pasta/arquivo> com <descrição>
- Modificado <arquivo> <descrição específica da mudança>
- Adicionado <campo/funcionalidade> em <arquivo>
- Atualizado <arquivo> <descrição da atualização>
- Atualizados go.mod e go.sum com dependência <nome>
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

## Organização por Seções

Para commits com múltiplas alterações, organize o corpo do commit em seções por arquivo ou grupo de arquivos relacionados.

### Estrutura Recomendada:

Use seções com o nome do arquivo ou caminho relativo seguido de dois pontos (`:`). Dentro de cada seção, liste apenas as mudanças específicas sem repetir o nome do arquivo:

```
<tipo>(<escopo>): <mensagem principal>

<arquivo1>:
- Mudança específica 1
- Mudança específica 2

<arquivo2>:
- Mudança específica 1
- Mudança específica 2
```

### Regras para Seções:

1. **Organize por arquivo**: Cada seção representa um arquivo ou grupo de arquivos relacionados
2. **Evite redundâncias**: Não repita o nome do arquivo dentro dos bullet points
3. **Seja específico**: Liste apenas o que foi feito, não o que já existia
4. **Agrupe quando fizer sentido**: Arquivos muito relacionados podem estar na mesma seção
5. **Mantenha ordem lógica**: Organize por ordem de dependência (ex: modelos antes de controllers)
6. **Use caminhos relativos**: Use o caminho do arquivo relativo ao projeto
7. **Seções opcionais**: Para commits simples com poucos arquivos, seções podem ser omitidas

## Verificação de Alterações (OBRIGATÓRIO)

**ANTES de criar qualquer mensagem de commit, SEMPRE verifique o histórico do Git para distinguir arquivos criados de arquivos modificados.**

### Processo Obrigatório:

1. Execute `git status` e `git diff --name-status HEAD -- <caminho>` para ver arquivos modificados
2. Execute `git ls-files --others --exclude-standard <caminho>` para ver arquivos novos
3. Execute `git diff HEAD -- <arquivo>` para ver mudanças específicas em arquivos modificados
4. Use "Criado/Criada" apenas para arquivos novos (não rastreados)
5. Use "Modificado" para arquivos existentes que foram alterados
6. Seja específico sobre o que mudou em cada arquivo modificado

### Modo de Entrega:

- **SEMPRE** gere a mensagem de commit em uma caixa de texto para o usuário copiar
- **NUNCA** execute `git commit` automaticamente
- **NUNCA** tente criar ou modificar arquivos durante a análise
- O usuário sempre usa o modo "ask" do Cursor e não há permissão de escrita
- A mensagem completa do commit deve ser apresentada em formato de texto dentro de uma caixa delimitada para fácil cópia

## Validação de Conteúdo

Antes de finalizar o commit, valide que:
- Cada seção possui pelo menos um bullet point com informação útil
- Não há seções vazias ou com mensagens genéricas
- Todos os arquivos modificados foram analisados em detalhes
- Mudanças específicas foram identificadas e documentadas
- O commit reflete fielmente o trabalho realizado sem omitir detalhes importantes
- Análise profunda foi realizada para não esquecer nenhum detalhe relevante
