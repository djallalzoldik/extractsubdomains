package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func extractSubdomain(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		rawURL := strings.TrimSpace(scanner.Text())
		if rawURL != "" {
			subdomain, err := extractSubdomain(rawURL)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error extracting subdomain: %v\n", err)
			} else {
				fmt.Println(subdomain)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
