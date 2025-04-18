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
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>{{.Title}}</h1>
	<ul>
		{{range .Items}}
			<li>
				<a href="{{.Link}}">{{.Title}}</a><br>
				<p>{{.Description}}</p>
			</li>
		{{end}}
	</ul>
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
