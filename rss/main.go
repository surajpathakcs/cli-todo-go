package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}
type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

var tmpl = []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}} - RSS Feed</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f9f9f9;
        }
        header {
            text-align: center;
            margin-bottom: 30px;
            padding-bottom: 20px;
            border-bottom: 1px solid #eaeaea;
        }
        h1 {
            color: #2c3e50;
            margin-bottom: 10px;
        }
        .feed-container {
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.05);
            overflow: hidden;
        }
        .feed-item {
            padding: 20px;
            border-bottom: 1px solid #f0f0f0;
            transition: background-color 0.2s ease;
        }
        .feed-item:hover {
            background-color: #f8f8f8;
        }
        .feed-item:last-child {
            border-bottom: none;
        }
        .feed-item h2 {
            margin: 0 0 10px 0;
            color: #3498db;
            font-size: 1.2em;
        }
        .feed-item a {
            color: inherit;
            text-decoration: none;
        }
        .feed-item a:hover {
            text-decoration: underline;
        }
        .feed-description {
            color: #666;
            font-size: 0.95em;
        }
        .feed-link {
            display: inline-block;
            margin-top: 10px;
            color: #3498db;
            font-size: 0.9em;
            text-decoration: none;
        }
        .feed-link:hover {
            text-decoration: underline;
        }
        
    </style>
</head>
<body>
    <header>
        <h1>{{.Title}}</h1>
        <p>Latest updates from our RSS feed</p>
    </header>

    <div class="feed-container">
        {{range .Items}}
        <article class="feed-item">
            <h2><a href="{{.Link}}">{{.Title}}</a></h2>
            <div class="feed-description">
                {{.Description}}
            </div>
            <a href="{{.Link}}" class="feed-link">Read more →</a>
        </article>
        {{end}}
    </div>

</body>
</html>
`)

func fetchRSS(feedurl string) (*Channel, error) {
	resp, err := http.Get(feedurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss RSS
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		return nil, err
	}
	return &rss.Channel, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	feedURL := "https://thehimalayantimes.com/rssFeed/14"
	channel, err := fetchRSS(feedURL)

	if err != nil {
		http.Error(w, "Failed to fetch the Feed", http.StatusInternalServerError)
		fmt.Println("Error Fetching RSS", err)
		return
	}
	tmplParsed, err := template.New("rss").Parse(string(tmpl))
	if err != nil {
		http.Error(w, "Error on Parsing Template ", http.StatusInternalServerError)
		fmt.Println("Template parsing error", err)
		return
	}

	// Attempt to render the template with the provided channel data
	err = tmplParsed.Execute(w, channel)
	if err != nil {
		http.Error(w, "Error on Rendering Template with channel data", http.StatusInternalServerError)
		fmt.Println("Error Rendering Template")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening to Port 8000")

	http.ListenAndServe(":8000", nil)
}
