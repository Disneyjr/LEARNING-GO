# Semana 5 — Concorrência II

🎯 Objetivo: WaitGroup, Mutex, Context e Worker Pool.

## Exercícios

1. **WaitGroup**
   - Espere 5 goroutines terminarem.

2. **Contador seguro**
   - Variável compartilhada entre goroutines protegida por `sync.Mutex`.

3. **Worker Pool**
   - Workers leem jobs de um channel e calculam fatorial.
   - Resultados voltam em outro channel.

4. **Cancelamento**
   - Adicione `context.WithCancel` ao worker pool.
