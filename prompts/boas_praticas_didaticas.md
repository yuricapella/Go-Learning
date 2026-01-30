# Boas Práticas para Arquivos Didáticos em Go

## Princípios Fundamentais

### 1. Estrutura Hierárquica Clara

**Organização:**
- `main.go` na raiz: apenas orquestração, chama funções `Demonstrar*()` de subpacotes
- Subpastas por tópico: cada conceito em seu próprio arquivo/pacote
- Função principal `Demonstrar*()`: apresenta o conceito e chama exemplos individuais
- Funções `Exemplo*()`: cada exemplo é uma função separada, numerada e auto-contida

**Por quê:** Separação de responsabilidades facilita navegação e compreensão progressiva.

---

### 2. Explicação do "Por Quê" Antes do "Como"

**Sempre inclua:**
- O problema que o conceito resolve
- Por que usar essa abordagem (não apenas como usar)
- Comparação com alternativas quando relevante
- Contexto de quando aplicar

**Por quê:** Entender o problema motiva o aprendizado e ajuda na tomada de decisão prática.

---

### 3. Progressão Didática

**Ordem dos elementos:**
1. Título do conceito (`--- CONCEITO ---`)
2. Definição breve do que é
3. Seção "Por que usar?" com problema e solução
4. Exemplos numerados (`Exemplo 1:`, `Exemplo 2:`, etc.)
5. Seção `⚠️ IMPORTANTE:` com pontos-chave e armadilhas comuns

**Por quê:** Do conceito abstrato para a prática concreta, terminando com alertas importantes.

---

### 4. Exemplos Simples e Diretos

**Características:**
- Um exemplo = um conceito específico
- Código mínimo necessário para demonstrar
- Evite complexidade desnecessária
- Nomeie variáveis de forma descritiva
- Comente o que cada exemplo demonstra antes do código

**Por quê:** Exemplos simples são mais fáceis de entender e memorizar.

---

### 5. Evitar Redundâncias

**Regra:**
- Se dois exemplos demonstram o mesmo conceito, mantenha apenas o melhor
- Explique que o exemplo mantido é a "construção básica" ou "padrão recomendado"
- Remova variações que não agregam valor didático

**Por quê:** Redundância confunde e aumenta carga cognitiva desnecessariamente.

---

### 6. Arquivo `sintaxes.go` como Referência

**Propósito:**
- Apenas sintaxe básica, sem explicações extensas
- Sem complexidade adicional
- Serve como consulta rápida de "como construir"
- Funções nomeadas `sintaxe*()` para fácil identificação

**Por quê:** Separa referência rápida de explicação didática detalhada.

**⚠️ ATENÇÃO:** Arquivos que terminam com `_test.go` são **ignorados pelo compilador Go** quando não está em modo de teste. Use apenas para testes reais (`go test`). Para código didático/demonstrativo, use nomes normais (ex: `comandos.go`, não `comandos_test.go`).

---

### 7. Comparações Quando Relevante

**Quando usar:**
- Conceitos relacionados (ex: WaitGroup vs Canais)
- Abordagens alternativas para o mesmo problema
- Trade-offs entre opções

**Formato:**
- Lista de vantagens/desvantagens (✓/✗)
- Explicação clara de quando usar cada abordagem

**Por quê:** Ajuda a escolher a ferramenta certa para cada situação.

---

### 8. Seção "⚠️ IMPORTANTE"

**Conteúdo:**
- Armadilhas comuns
- Boas práticas essenciais
- Comportamentos não óbvios
- Regras de uso seguro

**Por quê:** Previne erros comuns e reforça conhecimento crítico.

---

### 9. Nomenclatura Consistente

**Padrões:**
- `Demonstrar*()`: função principal que apresenta o conceito
- `Exemplo*()`: exemplos individuais numerados
- `sintaxe*()`: funções de referência sintática
- Variáveis descritivas em português para clareza didática

**Por quê:** Consistência facilita navegação e compreensão do código.

---

### 10. Casos de Uso Práticos

**Quando apropriado, inclua:**
- Lista de casos de uso reais
- Vantagens do padrão/conceito
- Quando aplicar vs quando evitar

**Por quê:** Conecta teoria com prática real, facilitando aplicação futura.

---

### 11. Separação de Conteúdo Didático e Conteúdo das Aulas

**Regra fundamental:**
- Conteúdo didático deve estar em pasta separada (ex: `conteudo_*_didaticos/`)
- **NUNCA** duplicar arquivos das aulas originais na pasta didática
- Conteúdo das aulas permanece intacto em suas pastas originais
- Apenas arquivos didáticos (`main.go`, `sintaxes.go`, `testes/*.go`) na pasta didática

**Estrutura correta:**
```
pasta_assunto/
├── conteudo_assunto_didatico/    # Apenas conteúdo didático
│   ├── main.go
│   ├── sintaxes.go
│   └── testes/ ou conceitos/
├── introducao/                    # Conteúdo original (intacto)
└── avancado/                      # Conteúdo original (intacto)
```

**Por quê:** 
- Evita confusão entre conteúdo didático e exemplos práticos das aulas
- Mantém organização clara e separação de responsabilidades
- Facilita manutenção sem afetar conteúdo original
- Permite referenciar exemplos das aulas sem duplicá-los

---

## Checklist para Criar Arquivo Didático

- [ ] Função `Demonstrar*()` principal que orquestra o conteúdo
- [ ] Explicação do conceito no início
- [ ] Seção "Por que usar?" com problema e solução
- [ ] Exemplos separados em funções `Exemplo*()` numeradas
- [ ] Cada exemplo tem título descritivo antes do código
- [ ] Exemplos são simples e focados em um conceito
- [ ] Seção `⚠️ IMPORTANTE:` com pontos-chave
- [ ] Comparações com alternativas quando relevante
- [ ] Casos de uso práticos quando apropriado
- [ ] Sem redundâncias entre exemplos
- [ ] Arquivo `sintaxes.go` atualizado com sintaxe básica
- [ ] Conteúdo didático em pasta separada (não duplicar arquivos das aulas)

---

## Princípio Áureo

**"Do simples ao complexo, sempre explicando o porquê antes do como."**

Um arquivo didático eficaz:
1. **Contextualiza** (o que é e por que existe)
2. **Demonstra** (exemplos progressivos e simples)
3. **Alerta** (armadilhas e boas práticas)
4. **Aplica** (casos de uso reais)
