package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os/exec"
// 	"strings"
// 	"sync"

// 	"gopkg.in/yaml.v3"
// )

// // ---------------- Models ----------------

// type Task struct {
// 	Name    string
// 	Type    string
// 	Config  map[string]interface{}
// 	Actions []*Task
// 	Success []*Task
// 	Error   []*Task
// 	Switch  map[string]*Task
// 	Result  *Result
// }

// type Result struct {
// 	Data  interface{}
// 	Error error
// 	Case  string
// }

// // Context chia sẻ giữa các task (thread-safe)
// type Context struct {
// 	mu   sync.RWMutex
// 	Vars map[string]interface{}
// }

// func NewContext() *Context {
// 	return &Context{
// 		Vars: make(map[string]interface{}),
// 	}
// }

// func (c *Context) Get(key string) (interface{}, bool) {
// 	c.mu.RLock()
// 	defer c.mu.RUnlock()
// 	v, ok := c.Vars[key]
// 	return v, ok
// }

// func (c *Context) Set(key string, val interface{}) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.Vars[key] = val
// }

// func (c *Context) Clone() *Context {
// 	c.mu.RLock()
// 	defer c.mu.RUnlock()
// 	nc := NewContext()
// 	for k, v := range c.Vars {
// 		nc.Vars[k] = v
// 	}
// 	return nc
// }

// // ---------------- YAML Reader ----------------

// func Readfilex(path string) (map[string]interface{}, error) {
// 	b, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var m map[string]interface{}
// 	if err := yaml.Unmarshal(b, &m); err != nil {
// 		return nil, err
// 	}
// 	// sometimes YAML top-level has a key like "schedule:" around the object
// 	// If top-level has exactly one key and that key's value is a map, use that map as workflow
// 	if len(m) == 1 {
// 		for _, v := range m {
// 			if mm, ok := v.(map[string]interface{}); ok {
// 				return mm, nil
// 			}
// 		}
// 	}
// 	return m, nil
// }

// // ---------------- ParseAction (as you had) ----------------

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

// // ---------------- Helpers: template-lite and as binding ----------------

// // ResolveString: nếu s chứa template like "{{ key }}" — trả về ctx.Vars["key"]
// // nếu không phải template trả về s unchanged
// func ResolveString(s string, ctx *Context) string {
// 	s = strings.TrimSpace(s)
// 	// quick detection: contains "{{" and "}}"
// 	if !strings.Contains(s, "{{") || !strings.Contains(s, "}}") {
// 		return s
// 	}
// 	// handle simple case "{{ key }}" or "{{key}}"
// 	inside := s
// 	inside = strings.ReplaceAll(inside, "{{", "")
// 	inside = strings.ReplaceAll(inside, "}}", "")
// 	inside = strings.TrimSpace(inside)
// 	// support nested dot access like "product.id"
// 	parts := strings.Split(inside, ".")
// 	val, ok := ctx.Get(parts[0])
// 	if !ok {
// 		return "" // not found -> empty
// 	}
// 	// if single var and no dot, return as string (try json marshal)
// 	if len(parts) == 1 {
// 		// if value is string return it
// 		if str, ok := val.(string); ok {
// 			return str
// 		}
// 		// else marshal to json string
// 		b, _ := json.Marshal(val)
// 		return string(b)
// 	}
// 	// dot access: navigate maps
// 	cur := val
// 	for _, p := range parts[1:] {
// 		if m, ok := cur.(map[string]interface{}); ok {
// 			cur = m[p]
// 		} else {
// 			return ""
// 		}
// 	}
// 	// result to string
// 	if s2, ok := cur.(string); ok {
// 		return s2
// 	}
// 	b, _ := json.Marshal(cur)
// 	return string(b)
// }

// // BindResultToContext: nếu task.Config có "as" => lưu result.Data vào ctx
// func BindResultToContext(t *Task, ctx *Context) {
// 	if t.Result == nil {
// 		return
// 	}
// 	if asRaw, ok := t.Config["as"]; ok {
// 		if asName, ok := asRaw.(string); ok && asName != "" {
// 			ctx.Set(asName, t.Result.Data)
// 		}
// 	}
// }

// // ---------------- Execution ----------------

// func Run() {
// 	workflow, err := Readfile("./tasks/example.yaml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	root := ParseAction(workflow)
// 	fmt.Println("== Workflow Loaded ==")
// 	fmt.Printf("Root Type: %s | Name: %s\n", root.Type, root.Name)

// 	ctx := NewContext()

// 	ExecuteTask(ctx, root)
// }

// func ExecuteTask(ctx *Context, t *Task) {
// 	if t == nil {
// 		return
// 	}
// 	fmt.Printf("→ Run: [%s] %s\n", t.Type, t.Name)

// 	// default empty result
// 	t.Result = &Result{}

// 	switch t.Type {
// 	case "schedule":
// 		// run children actions sequentially
// 		for _, c := range t.Actions {
// 			ExecuteTask(ctx, c)
// 		}
// 		// schedule node itself typically has no result; treat as success
// 		RunSuccess(ctx, t)

// 	case "script":
// 		runScript(ctx, t)
// 		postExecute(ctx, t)

// 	case "fetch":
// 		runFetch(ctx, t)
// 		postExecute(ctx, t)

// 	case "foreach":
// 		runForeach(ctx, t)
// 		postExecute(ctx, t)

// 	default:
// 		// unknown type -> mark error
// 		t.Result.Error = fmt.Errorf("unknown type %s", t.Type)
// 		postExecute(ctx, t)
// 	}
// }

// // postExecute handles binding & routing (error/switch/success)
// func postExecute(ctx *Context, t *Task) {
// 	// bind result to context if 'as' provided
// 	BindResultToContext(t, ctx)

// 	// if error -> run error handlers
// 	if t.Result != nil && t.Result.Error != nil {
// 		fmt.Printf("  ✖ Task error: %v\n", t.Result.Error)
// 		for _, e := range t.Error {
// 			ExecuteTask(ctx, e)
// 		}
// 		return
// 	}

// 	// if case -> run switch branch
// 	if t.Result != nil && t.Result.Case != "" {
// 		if next, ok := t.Switch[t.Result.Case]; ok {
// 			ExecuteTask(ctx, next)
// 			return
// 		}
// 		// unknown case => just warn
// 		fmt.Printf("  ⚠ switch case '%s' not found\n", t.Result.Case)
// 	}

// 	// default: run success actions
// 	for _, s := range t.Success {
// 		ExecuteTask(ctx, s)
// 	}
// }

// // ---------------- Implementations of action types ----------------

// func runScript(ctx *Context, t *Task) {
// 	// script expects config "run" (string) which can be a shell command or path
// 	runRaw, _ := t.Config["run"].(string)
// 	cmdStr := ResolveString(runRaw, ctx)

// 	if cmdStr == "" {
// 		t.Result = &Result{Data: nil, Error: nil}
// 		return
// 	}

// 	// execute via sh -c
// 	cmd := exec.Command("sh", "-c", cmdStr)
// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		t.Result = &Result{Data: nil, Error: fmt.Errorf("script error: %v, %s", err, stderr.String())}
// 		return
// 	}

// 	// try parse stdout as json, otherwise string
// 	outStr := strings.TrimSpace(out.String())
// 	var parsed interface{}
// 	if json.Unmarshal([]byte(outStr), &parsed) == nil {
// 		t.Result = &Result{Data: parsed, Error: nil}
// 	} else {
// 		t.Result = &Result{Data: outStr, Error: nil}
// 	}
// }

// func runFetch(ctx *Context, t *Task) {
// 	// fetch expects config "url", optional "method", optional "body"
// 	urlRaw, _ := t.Config["url"].(string)
// 	url := ResolveString(urlRaw, ctx)
// 	if url == "" {
// 		t.Result = &Result{Data: nil, Error: fmt.Errorf("fetch: empty url")}
// 		return
// 	}

// 	method := "GET"
// 	if m, ok := t.Config["method"].(string); ok && m != "" {
// 		method = strings.ToUpper(m)
// 	}

// 	var bodyReader io.Reader
// 	if b, ok := t.Config["body"].(string); ok && b != "" {
// 		bodyStr := ResolveString(b, ctx)
// 		bodyReader = strings.NewReader(bodyStr)
// 	}

// 	req, err := http.NewRequest(method, url, bodyReader)
// 	if err != nil {
// 		t.Result = &Result{Data: nil, Error: err}
// 		return
// 	}

// 	// optional headers
// 	if headers, ok := t.Config["headers"].(map[string]interface{}); ok {
// 		for hk, hv := range headers {
// 			if hs, ok := hv.(string); ok {
// 				req.Header.Set(hk, hs)
// 			}
// 		}
// 	}

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Result = &Result{Data: nil, Error: err}
// 		return
// 	}
// 	defer resp.Body.Close()
// 	b, _ := ioutil.ReadAll(resp.Body)

// 	// try parse JSON
// 	var parsed interface{}
// 	if json.Unmarshal(b, &parsed) == nil {
// 		t.Result = &Result{Data: parsed, Error: nil}
// 	} else {
// 		// fallback to string
// 		t.Result = &Result{Data: string(b), Error: nil}
// 	}

// 	// optional: set case based on status code or body
// 	if statusCaseMap, ok := t.Config["case_by_status"].(map[string]interface{}); ok {
// 		codeStr := fmt.Sprintf("%d", resp.StatusCode)
// 		if _, ok := statusCaseMap[codeStr]; ok {
// 			t.Result.Case = statusCaseMap[codeStr].(string)
// 		}
// 	}
// }

// func runForeach(ctx *Context, t *Task) {
// 	// expect config "range": can be "{{ products }}" or a path in config
// 	rangeRaw, _ := t.Config["range"].(string)
// 	rangeResolved := ResolveString(rangeRaw, ctx)
// 	if rangeResolved == "" {
// 		// try direct config value
// 		if v, ok := t.Config["range_value"]; ok {
// 			rangeResolved = "" // will handle below
// 			_ = v
// 		} else {
// 			t.Result = &Result{Data: nil, Error: fmt.Errorf("foreach: empty range")}
// 			return
// 		}
// 	}

// 	// Attempt to resolve the range variable from context (prefer structured)
// 	// If the template pointed to a context variable, we should fetch the original object
// 	// Try to get inside braces key
// 	key := ""
// 	if strings.Contains(rangeRaw, "{{") && strings.Contains(rangeRaw, "}}") {
// 		inside := strings.ReplaceAll(rangeRaw, "{{", "")
// 		inside = strings.ReplaceAll(inside, "}}", "")
// 		inside = strings.TrimSpace(inside)
// 		parts := strings.Split(inside, ".")
// 		key = parts[0]
// 	}

// 	var list []interface{}
// 	if key != "" {
// 		if v, ok := ctx.Get(key); ok {
// 			// try to assert []interface{} or []string etc
// 			switch vv := v.(type) {
// 			case []interface{}:
// 				list = vv
// 			case []string:
// 				for _, s := range vv {
// 					list = append(list, s)
// 				}
// 			default:
// 				// If single value, wrap
// 				list = []interface{}{vv}
// 			}
// 		}
// 	}

// 	// if no list from ctx, attempt to parse rangeResolved as json array
// 	if len(list) == 0 && rangeResolved != "" {
// 		var parsed interface{}
// 		if json.Unmarshal([]byte(rangeResolved), &parsed) == nil {
// 			if arr, ok := parsed.([]interface{}); ok {
// 				list = arr
// 			}
// 		}
// 	}

// 	if len(list) == 0 {
// 		// nothing to iterate
// 		t.Result = &Result{Data: nil, Error: fmt.Errorf("foreach: empty list")}
// 		return
// 	}

// 	async := false
// 	if a, ok := t.Config["async"].(bool); ok {
// 		async = a
// 	}

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	results := make([]interface{}, 0, len(list))
// 	errorsFound := false

// 	for idx, item := range list {
// 		idx := idx
// 		item := item

// 		execOne := func() {
// 			defer wg.Done()
// 			// clone ctx so changes per item don't leak unless user writes to a shared var intentionally
// 			local := ctx.Clone()
// 			// set 'as' variable if provided
// 			if asName, ok := t.Config["as"].(string); ok && asName != "" {
// 				local.Set(asName, item)
// 			}
// 			// execute child actions sequentially
// 			for _, child := range t.Actions {
// 				ExecuteTask(local, child)
// 				// if child had error -> stop this item
// 				if child.Result != nil && child.Result.Error != nil {
// 					mu.Lock()
// 					errorsFound = true
// 					mu.Unlock()
// 					return
// 				}
// 			}
// 			// collect result from last action or the item itself
// 			var last interface{}
// 			if len(t.Actions) > 0 {
// 				lastT := t.Actions[len(t.Actions)-1]
// 				if lastT.Result != nil {
// 					last = lastT.Result.Data
// 				}
// 			} else {
// 				last = item
// 			}
// 			mu.Lock()
// 			results = append(results, last)
// 			mu.Unlock()
// 			fmt.Printf("  foreach item %d done\n", idx)
// 		}

// 		wg.Add(1)
// 		if async {
// 			go execOne()
// 		} else {
// 			execOne()
// 		}
// 	}

// 	wg.Wait()
// 	if errorsFound {
// 		t.Result = &Result{Data: results, Error: fmt.Errorf("foreach: some items failed")}
// 	} else {
// 		t.Result = &Result{Data: results, Error: nil}
// 	}
// }
