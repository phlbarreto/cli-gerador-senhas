package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower   = "abcdefghijklmnopqrstuvwxyz"
	numbers = "0123456789"
	symbols = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func passGenerate(length int, charset string) string {
	var senha strings.Builder
	for i := 0; i < length; i++ {
		c := randomChar(charset)
		senha.WriteByte(c)
	}

	return senha.String()
}

func randomChar(charset string) byte {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	errorHandler(err)
	return charset[n.Int64()]
}

func main() {
	length := flag.Int("l", 16, "Define o tamanho da senha.")
	count := flag.Int("c", 1, "Define a quantidade de senhas a serem geradas.")
	haveSymbols := flag.Bool("s", false, "Inclui simbolos (!@#$%...)")
	noUpper := flag.Bool("nu", false, "Exclui letras maiusculas")
	noLower := flag.Bool("nl", false, "Exclui letras minusculas")
	noNumber := flag.Bool("nn", false, "Exclui numeros")
	flag.Parse()

	var builder strings.Builder
	maiusculas := "não"
	if !*noUpper {
		builder.WriteString(upper)
		maiusculas = "sim"
	}

	minusculas := "não"
	if !*noLower {
		builder.WriteString(lower)
		minusculas = "sim"
	}

	numeros := "não"
	if !*noNumber {
		builder.WriteString(numbers)
		numeros = "sim"
	}

	simbolos := "não"
	if *haveSymbols {
		builder.WriteString(symbols)
		simbolos = "sim"
	}

	charset := builder.String()

	if len(charset) == 0 {
		fmt.Println("Erro: pelo menos um conjunto de caracteres deve ser selecionado.")
		return
	}
	if *length <= 0 {
		fmt.Println("Erro: o tamanho da senha deve ser maior que zero.")
		return
	}

	var senhasGeradas []string
	for c := 0; c < *count; c++ {
		senhasGeradas = append(senhasGeradas, passGenerate(*length, charset))
	}

	palavra := "senha"
	conjugacao := "gerada"
	if *count > 1 {
		palavra = "senhas"
		conjugacao = "geradas"
	}

	fmt.Printf("🔐 %s %s:\n\n", palavra, conjugacao)
	for _, s := range senhasGeradas {
		fmt.Printf("- %s\n", s)
	}
	fmt.Print("\n")

	fmt.Printf("✓ %d %s de %d caracteres %s.\n", *count, palavra, *length, conjugacao)

	fmt.Println("Letras maiúsculas:", maiusculas)
	fmt.Println("Letras minúsculas:", minusculas)
	fmt.Println("Números:", numeros)
	fmt.Println("Símbolos:", simbolos)
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
