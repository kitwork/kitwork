package action

import (
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	URL    string
	Method string
	Header map[string]string
	Agent  Type // "http" | "client"
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

	switch t.Kind {
	case KindValue, KindShort:
		if s, ok := t.Config["value"].(string); ok {
			cfg.URL = s
		} else {
			return errors.New("invalid value for fetch")
		}

	case KindFull, KindList:
		for k, v := range t.Config {
			switch k {
			case "url":
				if s, ok := v.(string); ok {
					cfg.URL = s
				} else {
					return fmt.Errorf("invalid type for url: %T", v)
				}
			case "method":
				if s, ok := v.(string); ok {
					cfg.Method = strings.ToUpper(s)
				} else {
					return fmt.Errorf("invalid type for method: %T", v)
				}
			case "header":
				if m, ok := v.(map[string]interface{}); ok {
					cfg.Header = make(map[string]string)
					for hk, hv := range m {
						if hs, ok := hv.(string); ok {
							cfg.Header[hk] = hs
						} else {
							fmt.Printf("⚠️ Warning: header value %v is not string\n", hv)
						}
					}
				}

			default:
				fmt.Printf("⚠️ Unknown config key: %s\n", k)
			}
		}
	case KindUnknown:
		fmt.Println("⚠️ Unknown fetch config:", t.Config)
	}

	if cfg.URL == "" {
		return errors.New("fetch action missing URL")
	}

	if cfg.Method == "" {
		cfg.Method = "GET"
	}

	return cfg.Handle(ctx)
}

func (r *Request) Handle(ctx *Context) error {
	switch r.Agent {
	case TypeHTTP:
		return r.HTTP(ctx)
	case TypeClient:
		return r.Client(ctx)
	default:
		return fmt.Errorf("unknown request agent: %s", r.Agent)
	}
}
