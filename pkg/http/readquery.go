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
		tag := string(outStruct.Field(i).Tag)
		tags, err := structtag.Parse(tag)
		if err != nil {
			return err
		}

		for _, item := range tags.Tags() {
			if item.Key == "query" {
				query := requestQuery.Get(item.Name)

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
				}
			}
		}
	}

	return nil
}
