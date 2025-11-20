package handler

import (
	"errors"
	"fmt"
	"log"
)

// Task đại diện cho một action/workflow
type Task struct {
	Name    string
	Type    string
	Config  map[string]interface{}
	Actions []*Task
	Success []*Task
	Error   []*Task
}

func (t *Task) Fetch(ctx *TaskContext) error {

	if t.Type != "fetch" {

		return errors.New("type is not fetch")
	}

	return nil
}

func (t *Task) Script(ctx *TaskContext) error {

	if t.Type != "script" {

		return errors.New("type is not script")
	}
	return nil
}

func (t *Task) Cron(ctx *TaskContext) error {

	if t.Type != "cron" {
		return errors.New("type is not cron")

	}

	fmt.Println("→ [cron] chạy lịch...")

	return nil
}

func (t *Task) Run(ctx *TaskContext) (err error) {

	fmt.Printf("→ Running Task: [%s] %s\n", t.Type, t.Name)

	switch t.Type {

	case "script":
		err = t.Script(ctx)

	case "fetch":
		err = t.Fetch(ctx)

	case "cron":
		err = t.Cron(ctx)

	case "foreach":
		fmt.Println("→ foreach chưa implement")
		// TODO

	default:
		err = fmt.Errorf("unknown task type: %s", t.Type)

	}

	if err == nil && len(t.Actions) > 0 {
		fmt.Println("→ run actions")
		for _, a := range t.Actions {
			if err = a.Run(ctx); err != nil {
				break
			}
		}
	}
	// Branch control
	if err != nil && len(t.Error) > 0 {
		fmt.Println("→ Error xảy ra → chạy error branch")
		for _, e := range t.Error {
			e.Run(ctx)
		}
	}

	if err == nil && len(t.Success) > 0 {
		fmt.Println("→ Success → chạy success branch")
		for _, s := range t.Success {
			s.Run(ctx)
		}
	}

	return
}

type TaskContext struct {
	Data   interface{}
	Result interface{}
}

// ParseAction parse map[string]interface{} thành Task đệ quy
func ParseAction(data map[string]interface{}) *Task {
	task := &Task{Config: make(map[string]interface{})}

	for key, val := range data {
		if vMap, ok := val.(map[string]interface{}); ok {
			task.Type = key

			for k, v := range vMap {
				switch k {

				case "name":
					if name, ok := v.(string); ok {
						task.Name = name
					}

				case "actions":
					if list, ok := v.([]interface{}); ok {
						for _, a := range list {
							if aMap, ok := a.(map[string]interface{}); ok {
								task.Actions = append(task.Actions, ParseAction(aMap))
							}
						}
					}

				case "success":
					if list, ok := v.([]interface{}); ok {
						for _, s := range list {
							if sMap, ok := s.(map[string]interface{}); ok {
								task.Success = append(task.Success, ParseAction(sMap))
							}
						}
					}

				case "error":
					if list, ok := v.([]interface{}); ok {
						for _, e := range list {
							if eMap, ok := e.(map[string]interface{}); ok {
								task.Error = append(task.Error, ParseAction(eMap))
							}
						}
					}

				default:
					task.Config[k] = v
				}
			}

		} else {
			task.Config[key] = val
		}
	}

	return task
}

func Run() {
	workflow, err := Readfile("./services/tasks/example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	ctx := new(TaskContext)
	root := ParseAction(workflow)

	result := root.Run(ctx)
	fmt.Println(result)
}
