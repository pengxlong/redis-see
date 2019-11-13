package app

import (
	"net/http"
	"text/template"
    "fmt"
)

type App struct {
	Name string
}

func NewApp(name string) *App {
	app := &App{Name: name}
	return app
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var filter string
    var key string
    if len(r.Form["filter"]) > 0 {
        filter = r.Form["filter"][0]
    }
    if len(r.Form["key"]) > 0 {
        key = r.Form["key"][0]
    }
    host := r.Host
    tls := r.TLS
    scheme := "http://"
    if tls != nil {
        scheme = "https://"
    }
    // 解析指定文件生成模板对象
    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        fmt.Println("create template failed, err:", err)
        return
    }
    // 搜索keys
    var keys []string
    if filter == "" {
        keys = RedisKeys()
    } else {
        keys = RedisKeysFilter(filter)
    }
    var value string
    // 查询value,目前只支持string
    if key != "" {
        value = RedisGet(key)
    }

    data := make(map[string]interface{})
    
    data["host"] = scheme + host
    data["keys"] = keys
    data["value"] = value
    // 利用给定数据渲染模板，并将结果写入w
    tmpl.Execute(w, data)
}

func (app *App) Run() {
	http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Println("HTTP server failed,err:", err)
        return
    }
}