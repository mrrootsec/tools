package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Parse command line arguments
	inputPtr := flag.String("d", "", "input to domain name crt.sh for")
	outputPtr := flag.String("o", "", "output file to save results")
	flag.Parse()

	if *inputPtr == "" || *outputPtr == "" {
		fmt.Println("usage: go run main.go -d <input> -o <output>")
		os.Exit(1)
	}

	// Build the request URL
	url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", *inputPtr)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response:", err)
		os.Exit(1)
	}

	// Parse the JSON response
	var domains []struct {
		NameValue string `json:"name_value"`
	}
	err = json.Unmarshal(body, &domains)
	if err != nil {
		fmt.Println("error parsing JSON:", err)
		os.Exit(1)
	}

	// Filter out the quotes and newlines
	var filteredDomains []string
	for _, domain := range domains {
		filteredDomain := strings.Replace(domain.NameValue, "\"", "", -1)
		filteredDomain = strings.Replace(filteredDomain, "\\n", "\n", -1)
		filteredDomains = append(filteredDomains, filteredDomain)
	}

	// Sort and deduplicate the domains
	sortCmd := exec.Command("sort", "-u")
	sortCmd.Stdin = strings.NewReader(strings.Join(filteredDomains, "\n"))
	sortOut, err := sortCmd.Output()
	if err != nil {
		fmt.Println("error sorting domains:", err)
		os.Exit(1)
	}

	// Write the output to a file
	outputFile, err := os.Create(*outputPtr)
	if err != nil {
		fmt.Println("error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	_, err = writer.WriteString(string(sortOut))
	if err != nil {
		fmt.Println("error writing output:", err)
		os.Exit(1)
	}

	writer.Flush()
	fmt.Printf("Results saved to %s\n", *outputPtr)
}
