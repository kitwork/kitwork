package action

import (
	"fmt"
	"net/http"
)

func (r *Request) Client(ctx *Context) error {
	fmt.Printf("→ [http] %s %s\n", r.Method, r.URL)

	req, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return err
	}

	for k, v := range r.Header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("→ Response status:", resp.Status)
	return nil
}
