package helpers

import (
	"os"
	"strings"
)

func EnforceHTTPS(url string) string {
	if url[:4] != "http" {
		return "https://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newUrl := strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(newUrl, "https://", "", 1)
	newUrl = strings.Replace(newUrl, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]

	return newUrl != os.Getenv("DOMAIN")
}
