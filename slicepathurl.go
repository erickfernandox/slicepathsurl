package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
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

	levelPtr := flag.Int("l", 2, "Level [1, 2 or 3]")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var urlsList []string

	for scanner.Scan() {

		urlStr := scanner.Text()
		if containsPercent(urlStr) {
			continue
		}

		urlParts, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}

		pathParts := strings.Split(urlParts.Path, "/")

		for i := 0; i < *levelPtr; i++ {
			if i == 0 {
				baseUrl := urlParts.Scheme + "://" + urlParts.Host
				urlsList = append(urlsList, baseUrl)
			} else {
				if len(pathParts) < i+1 {
					break
				}
				path := strings.Join(pathParts[:i+1], "/")
				counter := strings.Count(path, "-")

				if strings.ContainsRune(strings.TrimSuffix(path, "/"), ';') || strings.ContainsRune(strings.TrimSuffix(path, "/"), '.') || counter != 1 {
					continue
				} else {
					pathVerify := path[1:]
					if _, err := strconv.Atoi(pathVerify); err == nil {
						continue
					} else {

						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path

						index := strings.Index(resultantUrl, "//") // Finds the index of the first occurrence of '//'

						if index != -1 { 
							
							secondIndex := strings.Index(resultantUrl[index+2:], "//")

							if secondIndex != -1 { 
								resultantUrl = resultantUrl[:index+2+secondIndex] + "/" + resultantUrl[index+2+secondIndex+2:]
							}
						}

						urlsList = append(urlsList, resultantUrl)
					}
				}
			}
		}
	}

	// Remove URLs duplicados dentro da Lista
	urlsMap := make(map[string]bool)
	for _, url := range urlsList {
		urlsMap[url] = true
	}

	var resultantUrlsList []string
	for url := range urlsMap {
		resultantUrlsList = append(resultantUrlsList, url)
	}

	// Prints the resulting list
	for _, url := range resultantUrlsList {

		fmt.Println(url)
	}

}
