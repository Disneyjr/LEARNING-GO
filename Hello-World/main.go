package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	PrintMessageDeclaredBasicVariable()
	PrintLMessageDeclaredEnxutedVariable()
	PrintFAllPlaceholdersComplete()
}

func PrintMessageDeclaredBasicVariable() {
	var nome string = "Disney Junior"
	fmt.Print("Olá, " + nome + "! Bem-vindo ao Go. \n")
}
func PrintLMessageDeclaredEnxutedVariable() {
	nome := "Disney Junior"
	fmt.Println("Olá, " + nome + "! Bem-vindo ao Go.")
}
func PrintFAllPlaceholdersComplete() {
	str := "GoLang"
	i := 255
	f := 3.14159
	b := true
	ch := 'A'
	s := struct{ Nome string }{"Disney"}
	ptr := &i

	fmt.Println("===== Strings =====")
	fmt.Printf("%%s: %s\n", str)
	fmt.Printf("%%q: %q\n", str)
	fmt.Printf("%%x (hex bytes): %x\n", str)
	fmt.Printf("%%X (hex bytes maiúsculo): %X\n", str)
	fmt.Printf("%%-10s (alinhado à esquerda, largura 10): %-10s|\n", str)
	fmt.Printf("%%10s (alinhado à direita, largura 10): %10s|\n", str)

	fmt.Println("\n===== Inteiros =====")
	fmt.Printf("%%d: %d\n", i)
	fmt.Printf("%%b (binário): %b\n", i)
	fmt.Printf("%%o (octal): %o\n", i)
	fmt.Printf("%%x (hex minúsculo): %x\n", i)
	fmt.Printf("%%X (hex maiúsculo): %X\n", i)
	fmt.Printf("%%c (caractere Unicode): %c\n", ch)
	fmt.Printf("%%5d (largura 5): |%5d|\n", i)
	fmt.Printf("%%-5d (largura 5, esquerda): |%-5d|\n", i)

	fmt.Println("\n===== Float =====")
	fmt.Printf("%%f: %.2f\n", f)
	fmt.Printf("%%8.2f (largura 8, 2 casas decimais): |%8.2f|\n", f)
	fmt.Printf("%%-8.2f (alinhado esquerda): |%-8.2f|\n", f)
	fmt.Printf("%%e: %e\n", f)
	fmt.Printf("%%E: %E\n", f)
	fmt.Printf("%%g: %g\n", f)
	fmt.Printf("%%G: %G\n", f)

	fmt.Println("\n===== Boolean =====")
	fmt.Printf("%%t: %t\n", b)

	fmt.Println("\n===== Ponteiros =====")
	fmt.Printf("%%p: %p\n", ptr)

	fmt.Println("\n===== Genérico =====")
	fmt.Printf("%%v: %v\n", s)
	fmt.Printf("%%+v (com campos): %+v\n", s)
	fmt.Printf("%%#v (literal Go): %#v\n", s)
	fmt.Printf("%%T (tipo): %T\n", s)

	fmt.Println("\n===== Combinações avançadas =====")
	fmt.Printf("String: |%10s|\n", str)
	fmt.Printf("String: |%-10s|\n", str)
	fmt.Printf("Inteiro: |%05d|\n", i)
	fmt.Printf("Float: |%08.3f|\n", f)
	fmt.Printf("Float científica: |%10.2e|\n", f)
}
