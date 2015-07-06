package handler

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zenazn/goji/web"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type html struct {
	htmlAssetsRoot string
	assetsRoot     string
}

func GetHtmlPathHandler(htmlPath string, assetsRoot string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		writeHtmlResponse(htmlPath, assetsRoot, w)
	}
}

func GetHtmlHandler(htmlAssetsRoot string, assetsRoot string) func(web.C, http.ResponseWriter, *http.Request) {
	html := html{htmlAssetsRoot: htmlAssetsRoot, assetsRoot: assetsRoot}
	return html.handler
}

func (e *html) handler(c web.C, w http.ResponseWriter, r *http.Request) {
	writeHtmlResponse(e.htmlAssetsRoot+"/"+getPage(c)+".html", e.assetsRoot, w)
}

func writeHtmlResponse(htmlPath string, assetsRoot string, w http.ResponseWriter) {
	data, readErr := ioutil.ReadFile(htmlPath)
	if readErr != nil {
		http.Error(w, "not found", http.StatusNotFound)
	}

	to := tagOutput{assetsRoot}
	funcMap := template.FuncMap{
		"jsTag":  to.getJsTag,
		"cssTag": to.getCssTag,
	}
	tmpl := template.Must(template.New("base").Funcs(funcMap).Parse(string(data)))
	tmplErr := tmpl.ExecuteTemplate(w, "base", nil)
	if tmplErr != nil {
		panic(tmplErr)
	}
}

func getPage(c web.C) string {
	ret := c.URLParams["page"]
	if len(ret) == 0 {
		return "index"
	}
	return ret
}

type staticFile struct {
	contentType string
	assetsRoot  string
}

type fileInfo struct {
	directory            string
	name                 string
	extension            string
	nameWithoutExtension string
}

func GetAssetsHandlerWithContentType(contentType string, assetsRoot string) func(http.ResponseWriter, *http.Request) {
	ct := staticFile{contentType: contentType, assetsRoot: assetsRoot}
	return ct.staticHandler
}

func (e *staticFile) staticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", e.contentType)

	fileInfo := parse(r.URL.Path)
	li := strings.LastIndex(fileInfo.name, "___")
	if li < 0 {
		writeFile(e.assetsRoot+"/"+fileInfo.directory+"/"+fileInfo.name, w, r)
		return
	}
	plainFileName := fileInfo.name[:li]
	writeFile(e.assetsRoot+"/"+fileInfo.directory+"/"+plainFileName, w, r)
}

func parse(path string) fileInfo {
	li := strings.LastIndex(path, "/")
	if li < 0 {
		nameWithoutExtension, extension := parseFileName(path)
		return fileInfo{directory: "",
			name:                 path,
			extension:            extension,
			nameWithoutExtension: nameWithoutExtension,
		}
	}
	dir := path[:li]
	file := path[li+1:]
	nameWithoutExtension, extension := parseFileName(file)
	return fileInfo{directory: dir,
		name:                 file,
		extension:            extension,
		nameWithoutExtension: nameWithoutExtension,
	}
}

func parseFileName(fileName string) (nameWithoutExtension, extension string) {
	li := strings.LastIndex(fileName, ".")
	if li < 0 {
		return fileName, ""
	}
	return fileName[:li], fileName[li+1:]
}

type tagOutput struct {
	root string
}

func checkDir(target string, root string) bool {
	rootAbs, err1 := filepath.Abs(root)
	if err1 != nil {
		return false
	}
	targetAbs, err2 := filepath.Abs(target)
	if err2 != nil {
		return false
	}
	return strings.Index(targetAbs, rootAbs) == 0
}

func (t *tagOutput) getTag(path string, fn func(fileInfo, string) string) (template.HTML, error) {
	fileInfo := parse(path)
	dir := t.root + fileInfo.directory
	if !checkDir(dir, t.root) {
		return "", errors.New("invalid path -> " + path)
	}
	data, readErr := ioutil.ReadFile(dir + "/" + fileInfo.name)
	if readErr != nil {
		return "", readErr
	}
	hash := md5.Sum(data)
	hashStr := fmt.Sprintf("%x", hash)
	tag := fn(fileInfo, hashStr)
	return template.HTML(tag), nil
}

func (t *tagOutput) getCssTag(path string) (template.HTML, error) {
	tag, err := t.getTag(path, func(fileInfo fileInfo, hash string) string {
		return fmt.Sprintf("<link href=\"%s/%s___%s\" type=\"text/css\" rel=\"stylesheet\"></link>", fileInfo.directory, fileInfo.name, hash)
	})
	if err != nil {
		return "", err
	}
	return tag, nil
}

func (t *tagOutput) getJsTag(path string) (template.HTML, error) {
	tag, err := t.getTag(path, func(fileInfo fileInfo, hash string) string {
		return fmt.Sprintf("<script src=\"%s/%s___%s\" type=\"text/javascript\"></script>", fileInfo.directory, fileInfo.name, hash)
	})
	if err != nil {
		return "", err
	}
	return tag, nil
}

func writeFile(file string, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Write(data)
}

func writeJsonResponse(data interface{}, w http.ResponseWriter) {
	b, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		panic(jsonErr)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, writeErr := w.Write(b)
	if writeErr != nil {
		panic(writeErr)
	}
}
