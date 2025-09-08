# Semana 10 — Projeto Final

🎯 Objetivo: consolidar tudo em um sistema.

## Projeto Final

**Mini-sistema de pedidos**:

- Endpoints:
  - `POST /pedido`
  - `GET /pedido/:id`
  - `GET /pedido`

- Pedido contém:
  - `ID`
  - `Cliente`
  - `Valor`

- Worker pool:
  - Processa pedidos simulando geração de nota fiscal.

- Configuração:
  - Variáveis de ambiente (`os.Getenv`).
  - Flags (`flag`).

- Qualidade:
  - Testes unitários no service.
  - Logs básicos (`log`).
  - Estrutura em pacotes (`cmd/`, `internal/`, `pkg/`).
