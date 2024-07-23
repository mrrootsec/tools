package main

import (
    "bufio"
    "fmt"
    "os"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

//converting payloads from txt file to yaml comfortable syntax

type Payload struct {
    Payloads []string `yaml:"payloads"`
}

func main() {
    // Open the file
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Slice to hold the lines
    var lines []string

    // Read each line and append to the slice
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Create the payload structure
    payload := Payload{
        Payloads: lines,
    }

    // Marshal the struct into YAML
    yamlData, err := yaml.Marshal(&payload)
    if err != nil {
        fmt.Println("Error marshaling YAML:", err)
        return
    }

    // Write the YAML data to a file
    err = ioutil.WriteFile("output.yaml", yamlData, 0644)
    if err != nil {
        fmt.Println("Error writing YAML file:", err)
        return
    }

    fmt.Println("YAML file created successfully!")
}
