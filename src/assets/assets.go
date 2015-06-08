package assets
import (
    "os"
    "time"
    "fmt"
    "strings"
    "strconv"

    "html/template"
    "net/http"

    "../env"
    "../bindata"
    "../bindata/debug"
)

const (
    assetsFileDelimiter = "___"
)

func BasicLayoutHtmlHandler(htmlPath string) http.Handler {
    return &basicLayoutHtmlHandler{htmlPath}
}

type basicLayoutHtmlHandler struct {
    htmlPath string
}

func (h *basicLayoutHtmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    baseData := mustGetData("html/common/basic-layout.html")
    contentsData := mustGetData(h.htmlPath)

    funcMap := template.FuncMap{
        "cssTag": cssTagOutputer,
        "jsTag": scriptTagOutputer,
    }

    tmpl := template.Must(template.New("base").Funcs(funcMap).Parse(string(baseData)))
    tmpl = template.Must(tmpl.New("contents").Funcs(funcMap).Parse(string(contentsData)))

    err := tmpl.ExecuteTemplate(w, "base", nil)
    if err != nil {
        panic(err)
    }
}

func cssTagOutputer(text string) template.HTML {
    cssInfo := mustGetInfo(text[1:])
    descriptor := newDescriptor(cssInfo)
    tag := fmt.Sprintf("<link rel=\"stylesheet\" href=\"%s%s%s\" type=\"text/css\">", text, assetsFileDelimiter, descriptor)
    return template.HTML(tag)
}

func scriptTagOutputer(text string) template.HTML {
    jsInfo := mustGetInfo(text[1:])
    descriptor := newDescriptor(jsInfo)
    tag := fmt.Sprintf("<script src=\"%s%s%s\" type=\"text/javascript\"></script>", text, assetsFileDelimiter, descriptor)
    return template.HTML(tag)
}

func newDescriptor(fileInfo os.FileInfo) string {
    base := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
    descriptor := fileInfo.ModTime().Sub(base).Nanoseconds()
    return strconv.FormatInt(descriptor, 10)
}

func ParseAsset(path string) (*template.Template, error) {
    src, err := getData(path)
    if err != nil {
        return nil, err
    }
    return template.New(path).Parse(string(src))
}

func ContentTypeHandler(contentType string) http.Handler {
    return &contentTypeHandler{contentType}
}

type contentTypeHandler struct {
     contentType string
}

func extractFilePath(r *http.Request) *assetFile {
    originalPath := r.URL.Path[1:];
    tokens := strings.Split(originalPath, assetsFileDelimiter)

    switch (len(tokens)) {
        case 1:
            return &assetFile{tokens[0], false, "0"}
        case 2:
            return &assetFile{tokens[0], true, tokens[1]}
    }
    panic(originalPath)
}

type assetFile struct {
    filePath string
    cacheable bool
    delimiter string
}

func (h *contentTypeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    asset := extractFilePath(r)

    data, err := getData(asset.filePath)
    if err != nil {
        if env.IsDebugMode() {
            panic(err)
        } else {
            http.NotFound(w, r)
        }
        return
    }
    w.Header().Add("content-type", h.contentType)
    if asset.cacheable || (asset.filePath == "img/unset.png") {
        SetCacheToHeader(w)
    }

    _, err2 := w.Write(data)
    if err2 != nil {
        if env.IsDebugMode() {
            panic(err)
        } else {
            http.NotFound(w, r)
        }
        return
    }
}

func GetData(path string) ([]byte, error) {
    return getData(path)
}

func getData(path string) ([]byte, error) {
    if env.IsProductionMode() {
        return bindata.Asset("assets/" + path)
    }
    return debug.Asset("assets/" + path)
}

func mustGetData(path string) []byte {
    d, e := getData(path)
    if e != nil {
        panic(e)
    }
    return d
}

func mustGetInfo(path string) os.FileInfo {
    f,e := bindata.AssetInfo("assets/" + path)
    if e != nil {
        panic(e)
    }
    return f
}

func SetCacheToHeader(w http.ResponseWriter) {
    s := time.Hour * 24 * 365
    w.Header().Set("Cache-Control", fmt.Sprintf("pulibc, max-age=%d", s)) //1年間キャッシュを有効にする
    w.Header().Del("Pragma")
    w.Header().Del("Expires")
}
