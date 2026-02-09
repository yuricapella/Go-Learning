# Guia de Pull Requests - Padrão do Projeto

Padrão baseado em boas práticas de Pull Requests com mensagens em português, organizado de forma similar ao padrão de commits do projeto.

## Formato do Título

```
[Feature | Refactor | Hotfix]: [objetivo/área/módulo principal]
```

O título deve ser autoexplicativo e refletir o objetivo principal da PR. Analise todos os commits da branch para entender o contexto completo e gerar um título que represente o assunto principal e o motivo da implementação.

## Tipos de PR

- `Feature`: Nova funcionalidade, módulo ou projeto completo
- `Refactor`: Refatoração de código sem mudança de comportamento
- `Hotfix`: Correção urgente de bug em produção
- `Docs`: Documentação
- `Chore`: Tarefas de manutenção, configuração ou organização

## Estrutura do Corpo

O corpo do PR deve ser organizado em seções claras que facilitem a revisão e compreensão das mudanças. Use a mesma lógica de organização por arquivo/módulo do padrão de commits.

### Estrutura Recomendada:

```
## Resumo

Breve descrição (2-3 linhas) do que foi implementado e por que foi necessário.

## Principais Alterações

[Seção por módulo/arquivo]:
- Mudança específica 1
- Mudança específica 2

[Outra seção]:
- Mudança específica 1
- Mudança específica 2

## Dependências

- Adicionada dependência <nome> versão <versão>
- Removida dependência <nome>

## Breaking Changes

[**OMITIR SEÇÃO COMPLETA se não houver breaking changes** - Apenas inclua se realmente existirem mudanças que quebram compatibilidade]
- Mudança que quebra compatibilidade e como migrar

## Testes

[**OMITIR SEÇÃO COMPLETA se não houver testes automatizados** - Apenas inclua se realmente existirem testes implementados]
- Testes automatizados implementados
- Cobertura de testes (se relevante)
- Padrões de teste utilizados (table-driven, subtestes, etc.)
- Módulos que possuem arquivos de teste além do módulo de testes dedicado

## Melhorias Implementadas

[**OMITIR SEÇÃO COMPLETA se não houver melhorias** - Apenas inclua se realmente existirem melhorias identificadas durante desenvolvimento]
- Refatorações realizadas após identificar oportunidades de melhoria
- Otimizações de estrutura ou organização
- Decisões técnicas que melhoram a qualidade do código
```

## Regras Essenciais

1. **Resumo conciso**: Comece com um resumo de 2-3 linhas explicando o que foi feito e por que
2. **Organização por seção**: Agrupe mudanças relacionadas por módulo, arquivo ou funcionalidade
3. **Evite redundâncias**: Não repita informações já presentes nos commits
4. **Seja específico**: Liste arquivos, funções, endpoints e mudanças concretas
5. **Destaque impactos**: Mencione mudanças que afetam padrões, fluxos ou configurações globais
6. **Dependências**: Sempre liste dependências adicionadas ou removidas
7. **Omita seções vazias**: **REGRA CRÍTICA** - Não inclua nenhuma seção que não possua conteúdo relevante. Se uma seção não tem pelo menos um bullet point com informação útil, ela não deve aparecer no PR. Isso se aplica a TODAS as seções: Breaking Changes, Testes, Melhorias Implementadas, ou qualquer outra seção opcional
8. **Foco no que foi feito**: Evite mencionar detalhes de commits específicos que não agregam ao entendimento geral do PR
9. **Separe contexto**: Diferencie estrutura/organização original de melhorias implementadas ao longo do desenvolvimento
10. **Evite commits individuais**: Não mencione refatorações ou melhorias específicas de commits isolados que não representam o escopo principal
11. **Verifique arquivos de teste**: Ao documentar testes, verifique todos os módulos do projeto, não apenas o módulo de testes dedicado. Muitos módulos podem possuir arquivos _test.go próprios
12. **Análise profunda**: AntES de criar o PR, analise profundamente todos os commits, arquivos modificados e estrutura do projeto para não esquecer detalhes importantes. Verifique padrões, arquivos de teste em todos os módulos, dependências e melhorias implementadas

## Organização por Seções

Para PRs com múltiplas alterações, organize o corpo em seções que agrupem mudanças relacionadas.

### Seções Comuns:

- **Módulos Didáticos**: Para conteúdo educacional ou exemplos
- **Modelos**: Novos modelos, estruturas de dados, tipos
- **Repositórios**: Funções de acesso a dados, queries, operações de banco
- **Controllers**: Handlers HTTP, lógica de negócio, validações
- **Rotas**: Definições de endpoints, configuração de rotas, middlewares
- **Serviços**: Camada de serviços, lógica de aplicação
- **Segurança**: Autenticação, autorização, criptografia
- **Banco de Dados**: Schemas, tabelas, migrations, scripts SQL
- **Configuração**: Arquivos de config, variáveis de ambiente, docker-compose
- **Refatoração**: Mudanças estruturais, renomeações, melhorias de código
- **Testes**: Novos testes, correções de testes, cobertura
- **Documentação**: Atualizações em README, comentários, docs

### Regras para Seções:

1. **Organize por contexto**: Cada seção representa um módulo, funcionalidade ou tipo de mudança
2. **Evite redundâncias**: Não repita o nome da seção dentro dos bullet points
3. **Seja específico**: Liste apenas o que foi feito, não o que já existia
4. **Agrupe quando fizer sentido**: Módulos muito relacionados podem estar na mesma seção
5. **Mantenha ordem lógica**: Organize por ordem de dependência (ex: modelos antes de controllers)
6. **Use nomes descritivos**: Seções devem ter nomes claros que identifiquem o grupo de mudanças
7. **Seções opcionais**: Para PRs simples, algumas seções podem ser omitidas

## Verificação de Alterações (OBRIGATÓRIO)

**ANTES de criar qualquer Pull Request, SEMPRE analise todos os commits da branch para entender o contexto completo.**

### Processo Obrigatório:

1. Execute `git log <branch-base>..<sua-branch>` para ver todos os commits da PR
2. Analise o histórico completo para entender o objetivo principal
3. Identifique os módulos/arquivos principais afetados
4. Liste dependências adicionadas ou removidas
5. Verifique se há breaking changes (se não houver, não crie a seção)
6. Verifique arquivos de teste em TODOS os módulos usando `find . -name "*_test.go"` ou similar
7. Identifique melhorias implementadas durante desenvolvimento que não fazem parte do escopo original
8. Considere o impacto das mudanças no projeto
9. **Validação final**: Revise cada seção do PR e remova qualquer seção que não tenha conteúdo relevante

### Modo de Entrega:

- **SEMPRE** gere a descrição do Pull Request em uma caixa de texto para o usuário copiar
- **NUNCA** crie o PR automaticamente
- **NUNCA** tente fazer merge ou operações de Git durante a análise
- O usuário sempre usa o modo "ask" do Cursor e não há permissão de escrita
- A descrição completa do PR deve ser apresentada em formato de texto dentro de uma caixa delimitada para fácil cópia

## Seções Especiais

### Melhorias Implementadas

Quando houver refatorações, melhorias de estrutura ou otimizações feitas ao longo do desenvolvimento (não parte do escopo original), crie uma seção específica:

```
## Melhorias Implementadas

- Refatoração X: descrição da melhoria e motivo
- Otimização Y: descrição e benefício
```

Esta seção deve conter apenas melhorias que foram identificadas e implementadas durante o desenvolvimento, não fazem parte do escopo original mas agregam valor ao código.

### Testes

A seção de testes deve focar apenas em testes automatizados implementados. Não inclua testes manuais ou validações feitas durante desenvolvimento manual, a menos que sejam parte essencial da funcionalidade (ex: scripts de teste automatizados). **Se não houver testes automatizados implementados, não crie esta seção.**

### Validação de Conteúdo

Antes de finalizar o PR, valide que:
- Cada seção possui pelo menos um bullet point com informação útil
- Não há seções vazias ou com mensagens genéricas como "não houve X"
- Todos os módulos foram verificados para arquivos de teste
- Melhorias implementadas foram identificadas e documentadas (se existirem)
- Breaking changes foram verificados e documentados (se existirem)
- O PR reflete fielmente o trabalho realizado sem omitir detalhes importantes

## Dicas para Boas PRs

1. **Contexto é importante**: Explique o "porquê" além do "o que"
2. **Seja conciso mas completo**: Resuma sem perder informações essenciais
3. **Facilite a revisão**: Organize de forma que o revisor encontre rapidamente o que precisa
4. **Documente decisões**: Se houver decisões técnicas importantes, mencione-as
5. **Considere o reviewer**: Pense no que o revisor precisa saber para aprovar a PR
6. **Foque no que foi feito**: Evite mencionar detalhes de commits específicos que não agregam ao entendimento geral
7. **Separe contexto**: Diferencie estrutura/organização original de melhorias implementadas ao longo do desenvolvimento
8. **Separe aulas de projetos**: Quando houver conteúdo educacional e projetos práticos, organize em seções distintas
9. **Testes automatizados apenas**: Na seção de testes, inclua apenas testes automatizados implementados, não testes manuais
