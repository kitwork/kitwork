package action

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

type Context struct {
	Return interface{} // dữ liệu cuối cùng trả về khi workflow kết thúc
	Result interface{} // dữ liệu hiện tại của workflow, được các action đọc/ghi
	templ  *template.Template
}

func NewContext() *Context {
	return &Context{}
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

	tmpl, err := template.New("action").Funcs(defaultFuncs()).Parse(val)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, c); err != nil {
		return "", err
	}

	return buf.String(), nil
}
