package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    // You can get an API key at https://pdfshift.io
    apiKey := "sk_3c03755c6e2e9b6f453b3151a445006e5bae4f80"

    params := map[string]interface{}{
        "source":            
    `
        <!DOCTYPE html>
         <html lang="en">
          <head>
            <meta charset="UTF-8" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
            <title>Document</title>
          </head>
          <body>
            <h1>send me your iphone :)</h1>
          </body>
        </html>
    `,
    }

    jsonParams, err := json.Marshal(params)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    client := &http.Client{}

    req, err := http.NewRequest("POST", "https://api.pdfshift.io/v3/convert/pdf", bytes.NewBuffer(jsonParams))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    req.Header.Set("Content-Type", "application/json")

    auth := "api:" + apiKey
    req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error performing request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    if resp.StatusCode >= 400 {
        fmt.Printf("Request failed with status code %d: %s\n", resp.StatusCode, string(body))
        return
    }

    err = ioutil.WriteFile("main.pdf", body, 0644)
    if err != nil {
        fmt.Println("Error saving PDF document:", err)
        return
    }

    fmt.Println("Malades!")
}