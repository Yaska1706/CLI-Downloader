package internal

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func handledownload(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return resp
	}
	defer resp.Body.Close()

	return resp
}

func createfile(filepath string) *os.File {
	output, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		return output
	}

	defer output.Close()

	return output
}

func cleanURl(url string) string {
	return strings.TrimSpace(url)
}
