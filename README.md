# Go CLI Tools

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

A collection of command-line tools and web utilities built in Go.

## Tools

### `cli-todo`  
A persistent to-do list manager with JSON storage  
Supports add/list/complete/delete tasks  
`todo [add|list|done|delete] [task]`

### `go-calc`
Simple calculator with basic operations  
Handles add/sub via command arguments  
`calc [add|sub] [num1] [num2]`

### `web-app`  -- not complete
Note-taking web app with Gorilla Mux  
CRUD operations with HTML templates  
Run and visit `localhost:8000`

### `rss` 
RSS feed reader and web server  
Parses and displays formatted articles with elegant UI 
Run and visit `localhost:8000`

## Usage
```bash
cd [tool-directory]
go run main.go
