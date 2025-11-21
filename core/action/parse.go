package action

import (
	"errors"
	"fmt"
)

type Parse struct {
	Value interface{} `action:"value,ignore"`
	To    string      `action:"to,required,fallback"`
}

func (t *Action) Parse(ctx *Context) error {
	if t.Type != TypeParser {
		return errors.New("type is not parse")
	}

	cfg := Parse{Value: ctx.Result} // mặc định lấy dữ liệu từ ctx.Data

	if err := t.classify(ctx, &cfg); err != nil {
		return err
	}
	fmt.Print(cfg)
	return cfg.Handle(ctx)
}

func (p *Parse) Handle(ctx *Context) error {
	if p.Value == nil {
		return errors.New("parse: no input data in ctx.Result")
	}

	switch p.To {
	case "json":

	case "string":

	case "int":

	default:

	}

	return nil
}
