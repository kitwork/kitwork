package action

import (
	"errors"
	"fmt"
	"time"
)

// Action đại diện cho 1 action/workflow node
type Action struct {
	Name    string
	Type    Type
	Kind    Kind // short, full, value, list, switch
	Config  map[string]interface{}
	Actions []*Action // chain mặc định
	Success []*Action // nếu OK
	Error   []*Action // nếu lỗi
	Timeout time.Duration
}

func (t *Action) Script(ctx *Context) error {
	if t.Type != TypeScript {
		return errors.New("type is not script")
	}

	fmt.Println("→ [script] chạy script ...")
	return nil
}

func (t *Action) Cron(ctx *Context) error {
	if t.Type != TypeCron {
		return errors.New("type is not cron")
	}

	fmt.Println("→ [cron] chạy lịch ...")
	return nil
}

// ========================
//  ACTION RUNNER
// ========================

func (t *Action) Run(ctx *Context) (err error) {

	fmt.Printf("\n→ Running Action: [%s] %s\n", t.Type, t.Name)

	// 1. chạy action chính
	switch t.Type {
	case TypeScript:
		err = t.Script(ctx)

	case TypeFetch, TypeHTTP, TypeClient, TypeRequest:
		err = t.Request(ctx)

	case TypeCron:
		err = t.Cron(ctx)

	case TypeLog:
		err = t.Log(ctx)

	case TypeParser:

		err = t.Parse(ctx)

	case TypeForeach:
		fmt.Println("→ foreach chưa implement")

	default:
		err = fmt.Errorf("unknown action type: %s", t.Type)
	}

	// 2. Nếu action chính OK → chạy chuỗi Actions
	if err == nil && len(t.Actions) > 0 {
		fmt.Println("→ run actions chain...")
		for _, a := range t.Actions {
			if err = a.Run(ctx); err != nil {
				break
			}
		}
	}

	// 3. Nếu lỗi → chạy Error branch
	if err != nil && len(t.Error) > 0 {
		fmt.Println("→ Error xảy ra → chạy error branch")
		for _, e := range t.Error {
			e.Run(ctx)
		}
	}

	// 4. Nếu thành công → chạy Success branch
	if err == nil && len(t.Success) > 0 {
		fmt.Println("→ Success → chạy success branch")
		for _, s := range t.Success {
			s.Run(ctx)
		}
	}

	return
}
