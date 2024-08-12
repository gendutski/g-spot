package dotenvgenerator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var errType error = errors.New("invalid type")

// generate .env file from struct envconfig tags
func GenerateDotEnv(cfg interface{}, build, cancelBuildInError bool, p Prompter) (rows []string, err error) {
	// set prompter
	if p == nil {
		p = newPrompter(bufio.NewReader(os.Stdin))
	}

	// validate type
	switch reflect.TypeOf(cfg).Kind() {
	case reflect.Struct:
		rows = append(rows, readStruct(reflect.TypeOf(cfg), p)...)
	case reflect.Pointer:
		val := reflect.ValueOf(cfg).Elem()
		if val.Kind() != reflect.Struct {
			return nil, errType
		}
		rows = append(rows, readStruct(val.Type(), p)...)
	case reflect.Slice:
		val := reflect.ValueOf(cfg)
		for i := 0; i < val.Len(); i++ {
			elm, _err := getStructFromElement(val.Index(i), i)
			if _err != nil {
				if err != nil {
					err = fmt.Errorf("%v\n%v", err, _err)
				} else {
					err = _err
				}
				continue
			}
			// add comment
			if i > 0 {
				rows = append(rows, "")
			}
			rows = append(rows, fmt.Sprintf("#%s", elm.Type().Name()))
			rows = append(rows, readStruct(elm.Type(), p)...)
		}
	default:
		return nil, errType
	}
	if len(rows) == 0 {
		return
	}

	text := strings.Join(rows, "\n")
	if build && (err == nil || !cancelBuildInError) {
		_err := os.WriteFile(".env", []byte(text), 0644)
		if _err != nil {
			if err != nil {
				err = fmt.Errorf("%v\n%v", err, _err)
			} else {
				err = _err
			}
		}
	}
	return
}

func readStruct(elm reflect.Type, p Prompter) []string {
	var result []string
	numFields := elm.NumField()
	for i := 0; i < numFields; i++ {
		// set field, tag, kind
		field := elm.Field(i)
		tag := field.Tag
		kind := field.Type.Kind()

		// env variable
		envconfig := tag.Get("envconfig")
		if envconfig == "" || envconfig == "-" {
			continue
		}

		// init prompt
		prompt := tag.Get("prompt")
		_default := tag.Get("default")
		if _default != "" {
			prompt += fmt.Sprintf(" (default:%s)", _default)
		}
		if kind.String() == "bool" {
			prompt += " (true or false)"
		}
		prompt += ": "

		// is secret prompt
		secret, _ := strconv.ParseBool(tag.Get("secret"))

		// scan
		var scan string
		if secret {
			scan = p.PromptPassword(prompt)
			fmt.Println()
		} else {
			scan = p.PromptString(prompt)
		}
		if scan == "" && _default != "" {
			scan = _default
		}
		result = append(result, fmt.Sprintf("%s=\"%s\"", envconfig, strings.ReplaceAll(scan, `"`, "")))
	}
	return result
}

func getStructFromElement(itm reflect.Value, n int) (reflect.Value, error) {
	switch itm.Type().Kind() {
	case reflect.Interface:
		return getStructFromElement(itm.Elem(), n)
	case reflect.Pointer:
		return getStructFromElement(itm.Elem(), n)
	case reflect.Struct:
		return itm, nil
	}
	if n >= 0 {
		return itm, fmt.Errorf("index #%d invalid type", n)
	}
	return itm, errType
}
