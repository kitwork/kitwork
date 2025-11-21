package action

import (
	"fmt"
	"time"
)

// Parse converts YAML-decoded map into an Action struct (recursive).
func Parse(data map[string]interface{}) *Action {
	action := &Action{
		Config:  make(map[string]interface{}),
		Actions: []*Action{},
		Success: []*Action{},
		Error:   []*Action{},
		Kind:    KindUnknown,
	}

	for key, val := range data {

		switch v := val.(type) {

		// ================================
		// 1) FULL FORM
		//    log:
		//      name: hello
		//      actions: [...]
		// ================================
		case map[string]interface{}:
			action.Type = TypeParseSafe(key)
			action.Kind = KindFull

			for k, vv := range v {
				switch k {
				case "name":
					if s, ok := vv.(string); ok {
						action.Name = s
					}
				case "actions":
					action.Actions = parseList(vv)
				case "success":
					action.Success = parseList(vv)
				case "error":
					action.Error = parseList(vv)
				case "timeout":
					if tStr, ok := vv.(string); ok {
						d, err := time.ParseDuration(tStr)
						if err == nil {
							action.Timeout = d
						} else {
							fmt.Printf("⚠️  Warning: cannot parse timeout %v: %v\n", vv, err)
						}
					}
				default:
					action.Config[k] = vv
				}
			}

		// ================================
		// 2) LIST FORM
		//    actions:
		//      - log
		//      - fetch: {...}
		// ================================
		case []interface{}:
			action.Type = TypeParseSafe(key)
			action.Kind = KindList
			action.Actions = parseList(v)

		// ================================
		// 3) VALUE FORM
		//    log: "hello"
		// ================================
		case string, int, float64, bool:
			action.Type = TypeParseSafe(key)
			action.Kind = KindValue
			action.Config["value"] = v

		// ================================
		// 4) FALLBACK / UNKNOWN
		// ================================
		default:
			action.Type = TypeCustom
			action.Kind = KindUnknown
			action.Config["value"] = v
		}
	}

	return action
}

// parseList parses:
// - short string: "log"
// - full form: { fetch: {...} }
// - unknown primitives
func parseList(v interface{}) []*Action {
	result := []*Action{}

	list, ok := v.([]interface{})
	if !ok {
		return result
	}

	for _, item := range list {

		switch x := item.(type) {

		// -----------------------
		// SHORT FORM
		// - log
		// -----------------------
		case string:
			result = append(result, &Action{
				Type:   TypeParseSafe(x),
				Kind:   KindShort,
				Config: map[string]interface{}{},
			})

		// -----------------------
		// FULL FORM
		// - log: { msg: "hi" }
		// -----------------------
		case map[string]interface{}:
			result = append(result, Parse(x))

		// -----------------------
		// Unknown value in list
		// -----------------------
		default:
			result = append(result, &Action{
				Type:   TypeCustom,
				Kind:   KindUnknown,
				Config: map[string]interface{}{"value": x},
			})
		}
	}

	return result
}
