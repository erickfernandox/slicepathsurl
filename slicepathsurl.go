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
	return strings.Contains(word, "%")
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
	levelPtr := flag.Int("l", 2, "Level [1, 2 or ... N]")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	var urlsList []string

	for scanner.Scan() {
		urlStr := scanner.Text()
		if urlStr == "" {
			continue
		}
		if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
			continue
		}
		if containsPercent(urlStr) {
			continue
		}

		urlStr = strings.Split(urlStr, "?")[0]

		urlParts, err := url.Parse(urlStr)
		if err != nil {
			continue
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
					suffixes := []string{".php", ".php3", ".php4", ".aspx", ".jsf", ".asp", ".html", ".jsonp", ".json", ".jsp", ".axd", ".htm", ".esp", ".cgi", ".do", ".jsx", ".xhtml", ".jhtm"}
					hasValidSuffix := false
					for _, suffix := range suffixes {
						if strings.HasSuffix(path, suffix) {
							hasValidSuffix = true
							break
						}
					}
					if hasValidSuffix {
						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path
						urlsList = append(urlsList, cleanDoubleSlashes(resultantUrl))
					}
				} else {
					values := strings.Split(path, "/")
					hasNumber := false
					for _, value := range values {
						if _, err := strconv.Atoi(value); err == nil {
							hasNumber = true
							break
						}
					}
					if hasNumber {
						continue
					} else {
						resultantUrl := urlParts.Scheme + "://" + urlParts.Host + "/" + path
						urlsList = append(urlsList, cleanDoubleSlashes(resultantUrl))
					}
				}
			}
		}
	}

	// Remover duplicados
	urlsMap := make(map[string]bool)
	for _, url := range urlsList {
		urlsMap[url] = true
	}

	for url := range urlsMap {
		fmt.Println(url)
	}
}

func cleanDoubleSlashes(url string) string {
	// Remove "//" no caminho sem afetar o schema (https://)
	parts := strings.SplitN(url, "//", 3)
	if len(parts) == 3 {
		return parts[0] + "//" + parts[1] + "/" + strings.ReplaceAll(parts[2], "//", "/")
	}
	return url
}
