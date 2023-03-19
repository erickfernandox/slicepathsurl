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

func containsSpecialChars(word string, chars string) bool {
	for _, char := range chars {
		if strings.ContainsRune(word, char) {
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
				chars := ":*;(){}[]@\".&"

				if containsSpecialChars(path, chars) || counter > 1 {
					continue
				} else {

					values := strings.Split(path, "/")
					verify_numeric_values := 0

					for _, value := range values {
						if _, err := strconv.Atoi(value); err == nil {
							verify_numeric_values = verify_numeric_values + 1
						} else {
							continue
						}
					}
					if verify_numeric_values > 0 {
						continue
					} else {
						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path
						index := strings.Index(resultantUrl, "//") 

						if index != -1 {
							secondIndex := strings.Index(resultantUrl[index+2:], "//")
							if secondIndex != -1 { // If it finds the second occurrence of '//'
								resultantUrl = resultantUrl[:index+2+secondIndex] + "/" + resultantUrl[index+2+secondIndex+2:]
							}
						}

						urlsList = append(urlsList, resultantUrl)
					}
				}
			}
		}
	}

	// Removes URls duplicadas da lista
	urlsMap := make(map[string]bool)
	for _, url := range urlsList {
		urlsMap[url] = true
	}

	var resultantUrlsList []string
	for url := range urlsMap {
		resultantUrlsList = append(resultantUrlsList, url)
	}

	// Printa o resultado da lista
	for _, url := range resultantUrlsList {
		fmt.Println(url)
	}
}
