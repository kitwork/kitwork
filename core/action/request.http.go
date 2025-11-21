package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (r *Request) HTTP(ctx *Context) error {
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

	ctx.Result, err = r.parseResponse(resp)
	if err != nil {
		return err
	}
	return err
}

func (r *Request) parseResponse(resp *http.Response) (interface{}, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Lưu Response kèm body bytes để dùng sau
	r.Response = &http.Response{
		Status:        resp.Status,
		StatusCode:    resp.StatusCode,
		Header:        resp.Header,
		ContentLength: resp.ContentLength,
		Body:          io.NopCloser(bytes.NewReader(bodyBytes)),
	}

	contentType := resp.Header.Get("Content-Type")
	switch {
	case strings.Contains(contentType, "application/json"):
		var data interface{}
		if err := json.Unmarshal(bodyBytes, &data); err != nil {
			return nil, fmt.Errorf("failed to parse JSON response: %w", err)
		}
		return data, nil

	case strings.Contains(contentType, "text/") || strings.Contains(contentType, "application/xml"):
		return string(bodyBytes), nil

	default:
		return string(bodyBytes), nil
	}
}
