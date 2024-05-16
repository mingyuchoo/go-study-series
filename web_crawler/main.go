package main

import (
    "fmt"
    "log"
    "net/http"
    "golang.org/x/net/html"
    "strings"
)

// Fetches the HTML content of the given URL.
func fetch(url string) (*html.Node, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching URL %s: %v", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error fetching URL %s: %v", url, resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error parsing HTML: %v", err)
    }

    return doc, nil
}

// Extracts links from the HTML node.
func extractLinks(n *html.Node, baseURL string) []string {
    var links []string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    href := attr.Val
                    if strings.HasPrefix(href, "/") {
                        href = baseURL + href
                    }
                    links = append(links, href)
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(n)
    return links
}

func main() {
    url := "http://jsonplaceholder.typicode.com"
    doc, err := fetch(url)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    links := extractLinks(doc, url)
    for _, link := range links {
        fmt.Println(link)
    }
}

