package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Display files in the web browser.
const tpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Video Files</title>
</head>
<body>
    <h1>Select a Video File</h1>
    <ul>
        {{range .}}
            <li><a href="/videos/{{.}}">{{.}}</a></li>
        {{end}}
    </ul>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("/output"))))
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("/output")
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, filepath.Base(file.Name()))
		}
	}

	tmpl, err := template.New("page").Parse(tpl)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, fileNames)
}
