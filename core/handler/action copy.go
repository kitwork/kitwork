package handler

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// )

// // Task đại diện cho một action/workflow
// type Task struct {
// 	Name    string
// 	Type    string
// 	Config  map[string]interface{}
// 	Actions []*Task
// 	Success []*Task
// 	Error   []*Task
// 	Switch  map[string]*Task
// 	Result  *TaskResult
// }

// func (t *Task) Fetch(context *TaskContext) (result interface{}, err error) {
// 	if t.Type != "fetch" {
// 		err = errors.New("Type not as fetch")
// 		return
// 	}
// 	return
// }

// func (t *Task) Script(context *TaskContext) (result interface{}, err error) {
// 	if t.Type != "fetch" {
// 		err = errors.New("Type not as fetch")
// 		return
// 	}
// 	return
// }

// func (t *Task) Cron(context *TaskContext) (result interface{}, err error) {
// 	if t.Type != "fetch" {
// 		err = errors.New("Type not as fetch")
// 		return
// 	}
// 	return
// }

// func (t *Task) Run(context *TaskContext) (result interface{}, err error) {
// 	if t == nil {
// 		return
// 	}

// 	fmt.Printf("→ Running Task: [%s] %s\n", t.Type, t.Name)

// 	// Dispatch theo t.Type
// 	switch t.Type {

// 	// case "schedule":
// 	// 	fmt.Println("Schedule detected — chạy các actions...")
// 	// 	for _, action := range t.Actions {
// 	// 		ExecuteTask(action, context)
// 	// 	}

// 	case "script":
// 		fmt.Println("Thực thi script...")
// 		// TODO: chạy file hoặc inline script
// 		return t.Script(context)

// 	case "fetch":
// 		fmt.Println("Gọi API fetch...")
// 		// TODO: HTTP GET/POST
// 		return t.Fetch(context)

// 	case "foreach":
// 		fmt.Println("Chạy vòng lặp foreach...")
// 		// TODO: lặp qua range
// 		// for _, child := range t.Actions {
// 		// 	ExecuteTask(child, context)
// 		// }
// 		// RunSuccess(t, context)
// 	case "cron":

// 		return t.Cron(context)

// 	default:
// 		fmt.Printf("⚠ Unknown task type: %s\n", t.Type)
// 		// RunError(t, context)
// 	}

// 	// Xử lý success / error branches
// 	if err != nil && len(t.Error) > 0 {
// 		fmt.Printf("Error occurred → run error branch\n")
// 		for _, e := range t.Error {
// 			e.Run(context)
// 		}
// 	} else if len(t.Success) > 0 {
// 		fmt.Printf("Success → run success branch\n")
// 		for _, s := range t.Success {
// 			s.Run(context)
// 		}
// 	}

// 	return result, err

// }

// type TaskResult struct {
// 	Data  interface{}
// 	Error error
// 	Case  string
// }

// type TaskContext struct {
// }

// // ParseAction parse map[string]interface{} thành Task đệ quy
// func ParseAction(data map[string]interface{}) *Task {
// 	task := &Task{
// 		Config: make(map[string]interface{}),
// 		Switch: make(map[string]*Task),
// 	}

// 	for key, val := range data {

// 		// Nếu value là map -> đây là một action có type
// 		if vMap, ok := val.(map[string]interface{}); ok {
// 			task.Type = key // type chính

// 			for k, v := range vMap {
// 				switch k {

// 				case "name":
// 					if name, ok := v.(string); ok {
// 						task.Name = name
// 					}

// 				case "actions":
// 					if actions, ok := v.([]interface{}); ok {
// 						for _, a := range actions {
// 							if aMap, ok := a.(map[string]interface{}); ok {
// 								task.Actions = append(task.Actions, ParseAction(aMap))
// 							}
// 						}
// 					}

// 				case "success":
// 					if items, ok := v.([]interface{}); ok {
// 						for _, s := range items {
// 							if sMap, ok := s.(map[string]interface{}); ok {
// 								task.Success = append(task.Success, ParseAction(sMap))
// 							}
// 						}
// 					}

// 				case "error":
// 					if items, ok := v.([]interface{}); ok {
// 						for _, e := range items {
// 							if eMap, ok := e.(map[string]interface{}); ok {
// 								task.Error = append(task.Error, ParseAction(eMap))
// 							}
// 						}
// 					}

// 				case "switch":
// 					if sw, ok := v.(map[string]interface{}); ok {
// 						for sk, sv := range sw {
// 							if svMap, ok := sv.(map[string]interface{}); ok {
// 								task.Switch[sk] = ParseAction(svMap)
// 							}
// 						}
// 					}

// 				default:
// 					task.Config[k] = v
// 				}
// 			}

// 		} else {
// 			// Nếu không phải map => config bình thường
// 			task.Config[key] = val
// 		}
// 	}

// 	return task
// }

// // ---------------- EXECUTOR ENGINE ----------------

// // Run toàn bộ workflow
// func Run() {
// 	workflow, err := Readfile("./services/tasks/example.yaml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var ctx = new(TaskContext)
// 	root := ParseAction(workflow)

// 	root.Run(ctx)
// }
