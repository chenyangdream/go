package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	//"path/filepath"
	"runtime/debug"
)

const (
	UPLOAD_DIR      = "./uploads"
	TEMPLATE_DIR    = "./views"
	UPLOAD_TEMPLATE = "upload.html"
	LIST_TEMPLATE   = "list.html"
)

var templates = make(map[string]*template.Template)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, UPLOAD_TEMPLATE, nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)

		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	locals := make(map[string]interface{})
	images := []string{}

	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images
	fmt.Println(locals)
	renderHtml(w, LIST_TEMPLATE, locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//check(err)
	//fmt.Println("Current working dir is ", dir)
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	fmt.Println(fileInfoArr)

	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		fmt.Println("templateName: = ", templateName)
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		fmt.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
	fmt.Println(templates)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				log.Println("WARN: panic in %v - %v", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()

		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
