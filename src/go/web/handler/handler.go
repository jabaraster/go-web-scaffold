package handler

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func IndexHtmlHandler(w http.ResponseWriter, r *http.Request) {
	to := tagOutput{"./assets"}
	funcMap := template.FuncMap{
		"jsTag":  to.getJsTag,
		"cssTag": to.getCssTag,
	}
	data, readErr := ioutil.ReadFile("./assets/index.html")
	if readErr != nil {
		panic(readErr)
	}
	tmpl := template.Must(template.New("base").Funcs(funcMap).Parse(string(data)))
	tmplErr := tmpl.ExecuteTemplate(w, "base", nil)
	if tmplErr != nil {
		panic(tmplErr)
	}
}

type ContentType struct {
	Value      string
	AssetsRoot string
}

type fileInfo struct {
	directory            string
	name                 string
	extension            string
	nameWithoutExtension string
}

func (e *ContentType) StaticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", e.Value)

	fileInfo := parse(r.URL.Path)
	li := strings.LastIndex(fileInfo.name, "___")
	if li < 0 {
		writeFile(e.AssetsRoot+"/"+fileInfo.directory+"/"+fileInfo.name, w, r)
		return
	}
	plainFileName := fileInfo.name[:li]
	writeFile(e.AssetsRoot+"/"+fileInfo.directory+"/"+plainFileName, w, r)
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
		panic(err1)
	}
	targetAbs, err2 := filepath.Abs(target)
	if err2 != nil {
		panic(err2)
	}
	return strings.Index(targetAbs, rootAbs) == 0
}

func (t *tagOutput) getTag(path string, fn func(fileInfo, string) string) (template.HTML, error) {
	fileInfo := parse(path)
	dir := t.root + fileInfo.directory
	if !checkDir(dir, t.root) {
		panic("invalid path -> " + path)
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

func (t *tagOutput) getCssTag(path string) template.HTML {
	tag, err := t.getTag(path, func(fileInfo fileInfo, hash string) string {
		return fmt.Sprintf("<link href=\"%s/%s___%s\" type=\"text/css\" rel=\"stylesheet\"></link>", fileInfo.directory, fileInfo.name, hash)
	})
	if err != nil {
		panic(err)
	}
	return tag
}

func (t *tagOutput) getJsTag(path string) template.HTML {
	tag, err := t.getTag(path, func(fileInfo fileInfo, hash string) string {
		return fmt.Sprintf("<script src=\"%s/%s___%s\" type=\"text/javascript\"></script>", fileInfo.directory, fileInfo.name, hash)
	})
	if err != nil {
		panic(err)
	}
	return tag
}

func writeFile(file string, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Write(data)
}
