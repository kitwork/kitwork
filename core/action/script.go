package action

import (
	"errors"
	"fmt"
)

type Script struct {
	Run   string `action:"url,required,fallback"`       // URL endpoint
	Agent Type   `action:"agent,ignore" default:"http"` // http | client

}

func (t *Action) Script(ctx *Context) error {
	if t.Type != TypeScript {
		return errors.New("type is not script")
	}

	fmt.Println("→ [script] chạy script ...")
	return nil
}
