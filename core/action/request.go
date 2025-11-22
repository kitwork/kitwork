package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	Name   string            `action:"name"`
	As     string            `action:"as"`
	URL    string            `action:"url,required,fallback"`       // URL endpoint
	Method string            `action:"method" default:"GET"`        // GET / POST / PUT...
	Header map[string]string `action:"header"`                      // request headers
	Agent  Type              `action:"agent,ignore" default:"http"` // http | client

	Body  interface{} `action:"body"`
	Parse string      `action:"parse,ignore" default:"auto"`

	Response interface{}
}

func (t *Action) Request(ctx *Context) error {

	cfg := Request{}

	switch t.Type {

	case TypeClient:
		cfg.Agent = TypeClient
	case TypeFetch, TypeHTTP, TypeRequest:
		cfg.Agent = TypeHTTP
	default:
		return fmt.Errorf("type is not a request/fetch type: %s", t.Type)
	}

	if err := t.classify(ctx, &cfg); err != nil {
		return err
	}

	return cfg.Handle(ctx)
}

func (r *Request) Handle(ctx *Context) error {

	switch r.Agent {
	case TypeHTTP:
		err := r.HTTP(ctx)
		return err
	case TypeClient:
		return r.Client(ctx)
	default:
		return fmt.Errorf("unknown request agent: %s", r.Agent)
	}
}

func (r *Request) body(ctx *Context) (io.Reader, error) {
	if r.Method == "" || r.Method == "GET" || r.Body == nil {
		return nil, nil
	}

	switch b := r.Body.(type) {
	case io.Reader:
		return b, nil
	case []byte:
		s := string(b)
		rendered, err := ctx.render(s)
		if err != nil {
			return bytes.NewReader(b), err
		}
		return bytes.NewReader([]byte(rendered)), nil
	case string:
		rendered, err := ctx.render(b)
		if err != nil {
			return strings.NewReader(b), err
		}
		return strings.NewReader(rendered), nil
	case map[string]interface{}, []interface{}:
		// deep render
		rendered, err := bodyDeepRender(b, ctx)
		if err != nil {
			return nil, err
		}
		data, err := json.Marshal(rendered)
		if err != nil {
			return nil, err
		}
		if r.Header == nil {
			r.Header = map[string]string{}
		}
		if _, ok := r.Header["Content-Type"]; !ok {
			r.Header["Content-Type"] = "application/json"
		}
		return bytes.NewReader(data), nil
	default:
		// struct, int, float… → marshal JSON
		data, err := json.Marshal(b)
		if err != nil {
			return nil, err
		}
		if r.Header == nil {
			r.Header = map[string]string{}
		}
		if _, ok := r.Header["Content-Type"]; !ok {
			r.Header["Content-Type"] = "application/json"
		}
		return bytes.NewReader(data), nil
	}
}

func bodyDeepRender(v interface{}, ctx *Context) (interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v2 := range val {
			rv, err := bodyDeepRender(v2, ctx)
			if err != nil {
				return nil, err
			}
			m[k] = rv
		}
		return m, nil
	case []interface{}:
		arr := make([]interface{}, len(val))
		for i, v2 := range val {
			rv, err := bodyDeepRender(v2, ctx)
			if err != nil {
				return nil, err
			}
			arr[i] = rv
		}
		return arr, nil
	case string:
		s, err := ctx.render(val)
		if err != nil {
			return nil, err
		}
		return s, nil
	default:
		return val, nil
	}
}
