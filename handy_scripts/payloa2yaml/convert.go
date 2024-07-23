package main

import (
    "bufio"
    "fmt"
    "os"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

// Payload struct to hold the array of payloads
type Payload struct {
    Payloads []string `yaml:"payloads"`
}

func main() {
    // Check if the correct number of arguments is provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    // Get the filename from the command line arguments
    filename := os.Args[1]

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    // Ensure the file is closed at the end of the function
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Slice to hold the lines from the file
    var lines []string

    // Read each line and append to the slice
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Check for errors during the scan
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Create the payload structure with the lines read
    payload := Payload{
        Payloads: lines,
    }

    // Marshal the struct into YAML format
    yamlData, err := yaml.Marshal(&payload)
    if err != nil {
        fmt.Println("Error marshaling YAML:", err)
        return
    }

    // Write the YAML data to a file named "output.yaml"
    err = ioutil.WriteFile("output.yaml", yamlData, 0644)
    if err != nil {
        fmt.Println("Error writing YAML file:", err)
        return
    }

    // Print success message
    fmt.Println("YAML file created successfully!")
}
