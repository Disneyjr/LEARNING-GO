# Semana 1 — Fundamentos

🎯 Objetivo: dominar sintaxe básica e funções.

## Exercícios

1. **Olá Go**
   - Crie um programa que receba seu nome via `os.Args` e imprima:
     ```
     Olá, <nome>! Bem-vindo ao Go.
     ```

2. **Calculadora CLI**
   - Programa que receba 3 argumentos:
     ```
     go run main.go soma 10 20
     go run main.go sub 50 30
     ```
   - Saída esperada:
     ```
     Resultado: 20
     ```

3. **Troca de valores**
   - Função `troca(a, b int) (int, int)` que inverte os números.

4. **Divisão com erro**
   - Função `divisao(a, b int) (int, error)` que retorna erro se `b == 0`.