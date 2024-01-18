package http

import (
	"errors"
	"gotemplate/pkg/structtag"
	"net/http"
	"reflect"
	"strconv"
)

var (
	ErrObjectIsNotStruct = errors.New("object is not struct")
	ErrUnknownFieldType  = errors.New("unknown field type")
)

func ReadQuery(r *http.Request, out interface{}) error {
	outStruct := reflect.TypeOf(out)
	fields := outStruct.NumField()
	valueStruct := reflect.ValueOf(out).Elem()
	if valueStruct.Kind() != reflect.Struct {
		return ErrObjectIsNotStruct
	}

	requestQuery := r.URL.Query()

	for i := 0; i < fields; i++ {
		tagForParse := string(outStruct.Field(i).Tag)
		tags, err := structtag.Parse(tagForParse)
		if err != nil {
			return err
		}

		tag, err := tags.Get("query")
		if err != nil {
			return err
		}

		query := requestQuery.Get(tag.Name)

		field := valueStruct.Field(i)
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		default:
			return ErrUnknownFieldType
		case reflect.Int:
			intElement, err := strconv.ParseInt(query, 10, 0)
			if err != nil {
				return err
			}

			field.SetInt(intElement)
			continue
		case reflect.String:
			field.SetString(query)
			continue
		case reflect.Bool:
			boolElement, err := strconv.ParseBool(query)
			if err != nil {
				return err
			}

			field.SetBool(boolElement)
			continue
		case reflect.Float64, reflect.Float32:
			floatElement, err := strconv.ParseFloat(query, 64)
			if err != nil {
				return err
			}

			field.SetFloat(floatElement)
			continue
		}
	}

	return nil
}
