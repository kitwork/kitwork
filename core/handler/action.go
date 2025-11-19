package handler

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Task đại diện cho một action/workflow
type Task struct {
	Name    string
	Type    string
	Config  map[string]interface{}
	Actions []*Task
	Success []*Task
	Error   []*Task
	Switch  map[string]*Task
	Result  *TaskResult
}

type TaskResult struct {
	Data  interface{}
	Error error
	Case  string
}

type TaskContext struct {
}

// ParseAction parse map[string]interface{} thành Task đệ quy
func ParseAction(data map[string]interface{}) *Task {
	task := &Task{
		Config: make(map[string]interface{}),
		Switch: make(map[string]*Task),
	}

	for key, val := range data {

		// Nếu value là map -> đây là một action có type
		if vMap, ok := val.(map[string]interface{}); ok {
			task.Type = key // type chính

			for k, v := range vMap {
				switch k {

				case "name":
					if name, ok := v.(string); ok {
						task.Name = name
					}

				case "actions":
					if actions, ok := v.([]interface{}); ok {
						for _, a := range actions {
							if aMap, ok := a.(map[string]interface{}); ok {
								task.Actions = append(task.Actions, ParseAction(aMap))
							}
						}
					}

				case "success":
					if items, ok := v.([]interface{}); ok {
						for _, s := range items {
							if sMap, ok := s.(map[string]interface{}); ok {
								task.Success = append(task.Success, ParseAction(sMap))
							}
						}
					}

				case "error":
					if items, ok := v.([]interface{}); ok {
						for _, e := range items {
							if eMap, ok := e.(map[string]interface{}); ok {
								task.Error = append(task.Error, ParseAction(eMap))
							}
						}
					}

				case "switch":
					if sw, ok := v.(map[string]interface{}); ok {
						for sk, sv := range sw {
							if svMap, ok := sv.(map[string]interface{}); ok {
								task.Switch[sk] = ParseAction(svMap)
							}
						}
					}

				default:
					task.Config[k] = v
				}
			}

		} else {
			// Nếu không phải map => config bình thường
			task.Config[key] = val
		}
	}

	return task
}

// ---------------- EXECUTOR ENGINE ----------------

// Run toàn bộ workflow
func Run() {
	workflow, err := Readfile("./tasks/example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	root := ParseAction(workflow)

	fmt.Println("== Workflow Loaded ==")
	fmt.Printf("Root Type: %s | Name: %s\n", root.Type, root.Name)
	var ctx = new(TaskContext)
	// Thông thường type đầu tiên sẽ là "schedule"
	ExecuteTask(root, ctx)
}

// ExecuteTask xử lý từng action theo Type
func ExecuteTask(t *Task, context *TaskContext) {
	if t == nil {
		return
	}

	fmt.Printf("→ Running Task: [%s] %s\n", t.Type, t.Name)

	// Dispatch theo t.Type
	switch t.Type {

	case "schedule":
		fmt.Println("Schedule detected — chạy các actions...")
		for _, action := range t.Actions {
			ExecuteTask(action, context)
		}

	case "script":
		fmt.Println("Thực thi script...")
		// TODO: chạy file hoặc inline script
		RunSuccess(t, context)

	case "fetch":
		fmt.Println("Gọi API fetch...")
		// TODO: HTTP GET/POST
		RunSuccess(t, context)

	case "foreach":
		fmt.Println("Chạy vòng lặp foreach...")
		// TODO: lặp qua range
		for _, child := range t.Actions {
			ExecuteTask(child, context)
		}
		RunSuccess(t, context)

	default:
		fmt.Printf("⚠ Unknown task type: %s\n", t.Type)
		RunError(t, context)
	}
}

// helper: chạy success
func RunSuccess(t *Task, context *TaskContext) {
	for _, s := range t.Success {
		ExecuteTask(s, context)
	}
}

// helper: chạy error
func RunError(t *Task, context *TaskContext) {
	for _, e := range t.Error {
		ExecuteTask(e, context)
	}
}

func Readfile(path string) (map[string]interface{}, error) {
	// 1. Đọc file YAML
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	// 2. Parse YAML vào map[string]interface{}
	var workflow map[string]interface{}
	err = yaml.Unmarshal(data, &workflow)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}

	return workflow, nil
}
