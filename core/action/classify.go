package action

import (
	"fmt"
	"reflect"
	"strings"
)

// classify gán cấu hình t.Config vào struct out (Request, Parse, Log...)
func (t *Action) classify(ctx *Context, out interface{}) error {
	val := reflect.ValueOf(out)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("out must be a non-nil pointer")
	}
	elem := val.Elem()
	elemType := elem.Type()

	if t.Kind == KindValue || t.Kind == KindShort {
		var uniqueField reflect.Value
		var fallbackField reflect.Value

		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			f := elem.Field(i)

			// chỉ xét field có thể set và là string
			if f.CanSet() && f.Kind() == reflect.String {
				if fallbackField.Kind() == 0 {
					fallbackField = f // ghi nhớ field string đầu tiên làm fallback
				}
			}

			tag := field.Tag.Get("action")
			parts := strings.Split(tag, ",")
			for _, p := range parts {
				if p == "fallback" {
					uniqueField = f
					break
				}
			}
			if uniqueField.IsValid() {
				break
			}
		}

		// Chọn field để gán: uniqueField > fallbackField
		targetField := uniqueField
		if !targetField.IsValid() {
			targetField = fallbackField
		}

		if targetField.IsValid() {
			s, _ := t.Config["value"].(string)
			rendered, err := ctx.render(s)
			if err != nil {
				return err
			}
			targetField.SetString(rendered)
		}
		return nil
	}

	// 3. Full / List kind → quét tất cả field
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tag := field.Tag.Get("action")
		parts := strings.Split(tag, ",")
		key := strings.ToLower(field.Name)
		ignore := false
		required := false
		defaultVal := ""

		if len(parts) > 0 && parts[0] != "" {
			key = parts[0]
		}

		for _, p := range parts[1:] {
			switch p {
			case "ignore":
				ignore = true
			case "required":
				required = true
			default:
				if strings.HasPrefix(p, "default:") {
					defaultVal = strings.TrimPrefix(p, "default:")
				}
			}
		}

		v, ok := t.Config[key]
		f := elem.Field(i)
		if !f.CanSet() {
			continue
		}

		// Bỏ qua nếu ignore và không có config
		if ignore && !ok {
			continue
		}

		switch f.Kind() {
		case reflect.String:
			s := defaultVal
			if ok {
				if ss, ok2 := v.(string); ok2 {
					s = ss
				}
			}
			rendered, _ := ctx.render(s)
			f.SetString(rendered)
		case reflect.Map:
			m := reflect.MakeMap(f.Type())
			if ok {
				if mm, ok2 := v.(map[string]interface{}); ok2 {
					for mk, mv := range mm {
						ms := fmt.Sprint(mv)
						rendered, _ := ctx.render(ms)
						m.SetMapIndex(reflect.ValueOf(mk), reflect.ValueOf(rendered))
					}
				}
			}
			f.Set(m)
		case reflect.Int, reflect.Int64:
			n := int64(0)
			if ok {
				switch vv := v.(type) {
				case int:
					n = int64(vv)
				case int64:
					n = vv
				}
			}
			f.SetInt(n)
		}

		if required && (f.Kind() == reflect.String && f.String() == "") {
			return fmt.Errorf("required field '%s' is empty", key)
		}
	}

	return nil
}
