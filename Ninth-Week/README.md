# Semana 9 — Arquitetura Go

🎯 Objetivo: estruturar projeto em camadas.

## Exercícios

1. **Camadas**
   - Reestruturar API `/soma` em:
     - `handler` (HTTP)
     - `service` (regra de negócio)
     - `repository` (dados em memória)

2. **Interfaces**
   - Criar `Repository` com 2 implementações: memória e fake.

3. **Injeção de dependência**
   - Fazer injeção manual em `main.go`.
