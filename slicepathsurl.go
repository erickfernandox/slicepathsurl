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

		// Remove query parameters from URL
		urlStr = strings.Split(urlStr, "?")[0]

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
				chars := ":*;(){}[]@\"&'+"

				if containsSpecialChars(path, chars) || counter > 1 {
					continue
				} else if strings.Contains(path, ".") {
					suffixes := []string{".php", ".aspx", ".jsf", ".asp", ".html", ".jsonp", ".json", ".jsp"}

					// Check if the path has any of the valid suffixes
					hasValidSuffix := false
					for _, suffix := range suffixes {
						if strings.HasSuffix(path, suffix) {
							hasValidSuffix = true
							break
						}
					}

					// If a valid suffix was found, include the URL
					if hasValidSuffix {
						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path
						index := strings.Index(resultantUrl, "//")

						if index != -1 {
							secondIndex := strings.Index(resultantUrl[index+2:], "//")
							if secondIndex != -1 {
								resultantUrl = resultantUrl[:index+2+secondIndex] + "/" + resultantUrl[index+2+secondIndex+2:]
							}
						}

						urlsList = append(urlsList, resultantUrl)
					}
				} else {
					values := strings.Split(path, "/")
					verifyNumericValues := 0

					for _, value := range values {
						if _, err := strconv.Atoi(value); err == nil {
							verifyNumericValues = verifyNumericValues + 1
						} else {
							continue
						}
					}
					if verifyNumericValues > 0 {
						continue
					} else {
						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path
						index := strings.Index(resultantUrl, "//")

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

	// Removes duplicated URLs from the list
	urlsMap := make(map[string]bool)
	for _, url := range urlsList {
		urlsMap[url] = true
	}

	var resultantUrlsList []string
	for url := range urlsMap {
		resultantUrlsList = append(resultantUrlsList, url)
	}

	// Print the resulting list
	for _, url := range resultantUrlsList {
		fmt.Println(url)
	}
}
