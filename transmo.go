package transmo

import (
	"errors"
	"reflect"
)

const transmo = "transmo"
const transmoIgnore = "ignore"

func Transform(input, output interface{}) (err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("output must pointer")
		}
	}()

	inputValue := reflect.ValueOf(input)
	outputValue := reflect.ValueOf(output).Elem()

	inputType := inputValue.Type()
	for i := 0; i < inputType.NumField(); i++ {
		name := inputType.Field(i).Name
		tag := inputType.Field(i).Tag
		if tag.Get(transmo) == transmoIgnore {
			continue
		}

		outputField := outputValue.FieldByName(name)
		if outputField.IsValid() {
			outputField.Set(inputValue.Field(i))
		}
	}
	return err
}
