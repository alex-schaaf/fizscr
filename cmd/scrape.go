package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

const url = "https://my.sport.uni-goettingen.de/fiz/"

func main() {
	html := getHTML()
	value := getValue(html)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}
	filename := os.Args[1]

	appendToFile(filename, value)
}

func appendToFile(filename string, value float64) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	entry := fmt.Sprintf("%s\t%f\n", timestamp, value)

	if _, err := f.WriteString(entry); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func getValue(html []byte) float64 {
	pattern := `gauge\.set\((\d+(\.\d+)?)\);`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(string(html))

	if len(matches) > 1 {
		fmt.Printf("Extracted number: %s\n", matches[1])
	} else {
		fmt.Println("No match found")
	}
	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		log.Fatalf("Failed to convert the extracted number to float64: %v", err)
	}
	return value
}

func getHTML() []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to load the URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}
	return body
}
