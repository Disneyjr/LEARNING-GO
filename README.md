# LEARNING-GO

> 📚 Repositório de estudos para dominar o básico da **linguagem Go** usando apenas a **stdlib**.  
>
> Organizado em 10 semanas, cada uma com metas e exercícios práticos.

---

## 📅 Roadmap de Estudos

### ✅ Semana 1 [31/08 - 06/09] — Fundamentos
- Sintaxe básica, funções, erros simples  
- Exercícios: calculadora CLI, troca de valores  

### 🚀 Semana 2 [07/09 - 13/09] — Estruturas de Dados
- Arrays, slices, maps, structs e interfaces  
- Exercícios: agenda telefônica, stack/queue  

### ⏳ Semana 3 [14/09 - 20/09] — Ponteiros, Composição e Generics
- Ponteiros, embedding, generics básicos  
- Exercícios: set genérico, função `Max`  

### ⏳ Semana 4 [21/09 - 27/09] — Concorrência I
- Goroutines, channels, `select`  
- Exercícios: soma concorrente, crawler  

### ⏳ Semana 5 [28/09 - 04/10] — Concorrência II
- WaitGroup, Mutex, Context, Worker Pool  
- Exercícios: fatorial com workers  

### ⏳ Semana 6 [05/10 - 11/10] — Stdlib Essencial
- Strings, arquivos, CSV/JSON, random  
- Exercícios: parser de CSV → JSON  

### ⏳ Semana 7 [12/10 - 18/10] — HTTP e Rede
- Servidor/cliente HTTP, middlewares  
- Exercícios: API REST básica  

### ⏳ Semana 8 [19/10 - 25/10] — Testes e Qualidade
- Unit tests, table-driven tests, benchmarks  
- Exercícios: testes para funções e API  

### ⏳ Semana 9 [26/10 - 01/11] — Arquitetura Go
- Organização de pacotes, interfaces, DI manual  
- Exercícios: refatorar API em camadas  

### ⏳ Semana 10 [02/11 - 08/11] — Projeto Final
- **Mini-sistema de pedidos** com API REST, workers, logs, env config e testes  

---

## 🎯 Objetivo
Chegar ao final das 10 semanas dominando:
- Fundamentos e idiomáticos da linguagem Go  
- Estruturas de dados e concorrência (core da linguagem)  
- Arquitetura limpa em Go sem bibliotecas externas  

---

## 📅 Aprendizado Extra.
- Na segunda semana resolvi introduzir os testes unitarios e testes de Benchmark para ir me adaptando a esses conceitos na linguagem.
> `go test -v # run the unit test #`
> `go test -bench . -benchmem -benchtime=100000x # run the benchmark test X times #`