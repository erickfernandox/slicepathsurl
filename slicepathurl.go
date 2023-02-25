package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func containsPercent(word string) bool {
	for _, char := range word {
		if char == '%' {
			return true
		}
	}

	return false
}

func main() {
	// Define os argumentos do script
	//listaPtr := flag.String("l", "", "Caminho da lista de URLs")
	nivelPtr := flag.Int("n", 2, "Nível [1, 2 ou 3]")
	flag.Parse()

	// Abre o arquivo com a lista de URLs
	//	arquivo, err := os.Open(*listaPtr)
	//	if err != nil {
	//	panic(err)
	//}
	//defer arquivo.Close()

	//scanner := bufio.NewScanner(arquivo)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	//url := strings.TrimSpace(scanner.Text())

	var listaUrls []string

	// Percorre a lista de URLs e cria a lista resultante com base no nível informado
	for scanner.Scan() {

		urlStr := scanner.Text()
		if containsPercent(urlStr) {
			continue
		}

		partesUrl, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}

		partesPath := strings.Split(partesUrl.Path, "/")

		for i := 0; i < *nivelPtr; i++ {
			if i == 0 {
				urlBase := partesUrl.Scheme + "://" + partesUrl.Host
				listaUrls = append(listaUrls, urlBase)
			} else {
				if len(partesPath) < i+1 {
					break
				}
				path := strings.Join(partesPath[:i+1], "/")

				if strings.ContainsRune(strings.TrimSuffix(path, "/"), ';') || strings.ContainsRune(strings.TrimSuffix(path, "/"), '.') || strings.ContainsRune(strings.TrimSuffix(path, "/"), '-') {
					continue
				} else {
					urlResultante := partesUrl.Scheme + "://" + partesUrl.Host + "/" + path

					index := strings.Index(urlResultante, "//") // Encontra o índice da primeira ocorrência de '//'

					if index != -1 { // Se encontrar a primeira ocorrência de '//'
						// Procura o índice da próxima ocorrência de '//', a partir do índice da primeira ocorrência encontrada
						secondIndex := strings.Index(urlResultante[index+2:], "//")

						if secondIndex != -1 { // Se encontrar a segunda ocorrência de '//'
							// Substitui a segunda ocorrência de '//' por '/'
							urlResultante = urlResultante[:index+2+secondIndex] + "/" + urlResultante[index+2+secondIndex+2:]
						}
					}

					listaUrls = append(listaUrls, urlResultante)
				}
			}
		}
	}

	// Remove URLs repetidas da lista resultante
	mapaUrls := make(map[string]bool)
	for _, url := range listaUrls {
		mapaUrls[url] = true
	}

	var listaUrlsResultantes []string
	for url := range mapaUrls {
		listaUrlsResultantes = append(listaUrlsResultantes, url)
	}

	// Imprime a lista resultante
	for _, url := range listaUrlsResultantes {

		fmt.Println(url)
	}
}
