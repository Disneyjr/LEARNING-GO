# Semana 4 — Concorrência I

🎯 Objetivo: goroutines, channels e select.

## Exercícios

1. **Hello concorrente**
   - Goroutine imprime "Olá do Go!" enquanto a main imprime "Olá da main!".

2. **Soma concorrente**
   - Divida um slice em duas partes e some em goroutines diferentes.
   - Junte o resultado via channel.

3. **Select**
   - Dois canais enviam mensagens aleatórias.
   - Use `select` para ler ambos.

4. **Timer**
   - Channel envia mensagem após 3 segundos.