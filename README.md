# WebScraper

This is a simple web scraper written in Go that extracts data from uri using BFS & golang.org/x/net/html.

## Technologies Used
- GoLang: Backend development
- ReGex: For Valid URL
- HTML Library: Html Parser (golang.org/x/net/html). 

## Getting Started

### Prerequisites
- GoLang installed on your system. You can download it from [here](https://golang.org/dl/).

### Working
The web scraper uses Go's built-in HTTP client to fetch the HTML content of the specified URL. It then utilizes the golang.org/x/net/html library, which provides a convenient way to extract and manipulate data from HTML documents using Go's syntax similar to Js.

Here's a basic overview of the scraper's workflow:

- HTTP Request: The scraper makes an HTTP GET request to the specified URL and retrieves the HTML content of the page.

- HTML Parsing: Using golang.org/x/net/html library, the HTML content is parsed into a Document object, which allows easy traversal and selection of elements.

- Data Extraction: The scraper specifies rules for extracting data from the HTML document. These rules typically involve selecting specific HTML elements based on their class names, IDs, or other attributes, and extracting relevant data such as text, links, or attributes.

- Output: The extracted data is then saved to a file, depending on the requirements of the application. There will be some hero links that remain in log only.

## Features
- Extracts specified data from Valid Url
- Customizable scraping rules using DFS or BFS or anyother way
- Lightweight and efficient

### Installation
1. Clone the repository:
    ```bash
    git clone github.com/iamgak/go-scraper
    ```

2. Navigate to the project directory:
    ```bash
    cd go-scraper
    ```

3. Install dependencies:
This project uses the following third-party libraries:
`golang.org/x/net/html` - A library for querying HTML documents in Go
Install the dependencies using:

    ```bash
    go mod tidy
    ```

4. Run the script:
    ```bash
    go run .
    ```
    optional if you want to change port number and url info
    ```bash

    go run . -timeout=10 -filename = "csvFile.csv"
    ```

5. Output:

See, the latest file
    ```bash
    cd csv; ls 
    ```

OR
    Open your filesystem and navigate to filename like `links20240602.csv` latest one.See, the list of all files.



## Contributing
Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request with your changes.

