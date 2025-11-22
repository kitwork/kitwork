package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/template"
)

type Key struct {
	Telegram string
}

type Context struct {
	Return interface{} // dữ liệu cuối cùng trả về khi workflow kết thúc
	Result interface{} // dữ liệu hiện tại của workflow, được các action đọc/ghi

	Debug bool

	Version string

	Key Key

	templ *template.Template
	pipes *template.FuncMap
}

func NewContext() *Context {
	defFuncs := defaultFuncs()
	return &Context{Key: Key{Telegram: "apikey"}, pipes: &defFuncs}
}

// addPipe thêm một func mới vào pipes
func (c *Context) addPipe(name string, function interface{}) error {
	if name == "" {
		return fmt.Errorf("pipe name cannot be empty")
	}

	// check function có phải là func
	if function == nil || reflect.TypeOf(function).Kind() != reflect.Func {
		return fmt.Errorf("value for pipe '%s' is not a function", name)
	}

	// khởi tạo pipes nếu nil
	if c.pipes == nil {
		m := make(template.FuncMap)
		c.pipes = &m
	}

	// check key đã tồn tại chưa
	if _, exists := (*c.pipes)[name]; exists {
		return fmt.Errorf("pipe '%s' already exists", name)
	}

	// thêm vào pipes
	(*c.pipes)[name] = function
	return nil
}

func (c *Context) pipeFuncs() template.FuncMap {
	return *c.pipes
}

// Tạo func map cho pipeline
func defaultFuncs() template.FuncMap {
	return template.FuncMap{
		"json": func(v interface{}) string {
			b, _ := json.Marshal(v)
			return string(b)
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
		"trim": func(s string) string {
			return strings.TrimSpace(s)
		},
	}
}

func (c *Context) render(val string) (string, error) {
	if val == "" {
		return "", nil
	}

	tmpl, err := template.New("action").Funcs(c.pipeFuncs()).Parse(val)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return "", err
	}

	return buf.String(), nil
}
