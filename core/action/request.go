package action

import (
	"fmt"
)

type Request struct {
	URL    string            `action:"url,required,fallback"`       // URL endpoint
	Method string            `action:"method" default:"GET"`        // GET / POST / PUT...
	Header map[string]string `action:"header"`                      // request headers
	Agent  Type              `action:"agent,ignore" default:"http"` // http | client

	Parse    string `action:"parse,ignore" default:"auto"`
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
