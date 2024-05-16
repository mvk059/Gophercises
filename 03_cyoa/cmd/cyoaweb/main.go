package main

import (
	_3_cyoa "Gophercises/03_cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func readArguments() (string, int) {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	port := flag.Int("port", 3000, "the port to start the CYOA web application")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *filename)
	return *filename, *port
}

func openFile(filename string) os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return *file
}

func main() {
	filename, port := readArguments()
	file := openFile(filename)
	story, err := _3_cyoa.JsonStory(&file)
	if err != nil {
		panic(err)
	}

	// Create our custom CYOA story handler
	//tpl := template.Must(template.New("").Parse(storyTmpl))
	h := _3_cyoa.NewHandler(
		story,
		//_3_cyoa.WithTemplate(tpl),
		//_3_cyoa.WithPathFunc(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/", h) //story/
	fmt.Printf("Starting the server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))

	//handler := _3_cyoa.NewHandler(story)
	//fmt.Printf("Starting the server on port %d\n", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

// Slightly altered template to show how this feature works
var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
