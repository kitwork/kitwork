package action

import "fmt"

// Type represents an action type in KitWork workflow engine
type Type string

// --- Action Types ---
const (
	// --- Network / Fetch ---
	TypeFetch   Type = "fetch"
	TypeHTTP    Type = "http"
	TypeClient  Type = "client"
	TypeRequest Type = "request"

	// --- Script / Command ---
	TypeScript  Type = "script"
	TypeCmd     Type = "cmd"
	TypeCommand Type = "command"

	// --- Flow Control ---
	TypeForeach Type = "foreach"

	TypeSwitch   Type = "switch"
	TypeLoop     Type = "loop"
	TypeReturn   Type = "return"
	TypeWait     Type = "wait"
	TypeRoutines Type = "routines"
	// TypeIf      Type = "if"

	// --- Logging / Debug ---
	TypeLog Type = "log"

	// --- Cron / Schedule ---
	TypeCron Type = "cron"

	// --- IO / Mail / Storage ---
	TypeSendMail Type = "sendmail"
	TypeSave     Type = "save"
	TypeCheck    Type = "check"

	// --- Browser Automation ---
	TypeChrome   Type = "chrome"
	TypeChromedp Type = "chromedp"

	// --- Custom / Fallback ---
	TypeCustom Type = "custom"
	TypeUnknow Type = "unknow"

	// --- Parse /  ---
	TypeParser Type = "parse"
)

// TypeParse converts a string to a Type enum
// Returns error if the string is not a valid action type
func TypeParse(s string) (Type, error) {
	switch s {
	case "fetch":
		return TypeFetch, nil
	case "http":
		return TypeHTTP, nil
	case "client":
		return TypeHTTP, nil
	case "script":
		return TypeScript, nil
	case "cmd":
		return TypeCmd, nil
	case "command":
		return TypeCommand, nil
	case "foreach":
		return TypeForeach, nil
	case "routines":
		return TypeRoutines, nil
	// case "if":
	// 	return TypeIf, nil
	case "switch":
		return TypeSwitch, nil
	case "cron":
		return TypeCron, nil
	case "loop":
		return TypeLoop, nil
	case "return":
		return TypeReturn, nil
	case "wait":
		return TypeWait, nil
	case "log":
		return TypeLog, nil
	case "sendmail":
		return TypeSendMail, nil
	case "save":
		return TypeSave, nil
	case "check":
		return TypeCheck, nil
	case "chrome":
		return TypeChrome, nil
	case "chromedp":
		return TypeChromedp, nil
	case "custom":
		return TypeCustom, nil

	case "parse":
		return TypeParser, nil
	default:
		return "", fmt.Errorf("invalid action type: %s", s)
	}
}

// TypeParseSafe converts string to Type enum safely
// Logs a warning and returns TypeCustom if invalid
func TypeParseSafe(s string) Type {
	t, err := TypeParse(s)
	if err != nil {
		fmt.Printf("⚠️ Warning: %v, fallback to TypeCustom\n", err)
		return TypeCustom
	}
	return t
}
