package cfgloader

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfigFromEnv[T any]() (*T, error) {
	var result T

	envFile := ".env"
	env := os.Getenv("ENV")
	if env != "" {
		envFile = fmt.Sprintf(".env.%s", env)
	}
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("load config err")
		return nil, err
	}
	err = loadFromEnvironment(reflect.ValueOf(&result).Elem())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func loadFromEnvironment(v reflect.Value) error {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.Struct {
			err := loadFromEnvironment(field)
			if err != nil {
				return err
			}
			continue
		}
		envKey := fieldType.Tag.Get("env")
		if envKey == "" {
			continue
		}

		value := os.Getenv(envKey)
		if value == "" {
			return fmt.Errorf("missing environment variable: %s", envKey)
		}

		err := setFieldValue(field, value)
		if err != nil {
			return err
		}
	}
	return nil

}

func setFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	case reflect.Slice:
		arrayValue := strings.Split(value, ",")
		slice := reflect.MakeSlice(field.Type(), len(arrayValue), len(arrayValue))
		for i, str := range arrayValue {
			elem := slice.Index(i)
			if err := setFieldValue(elem, str); err != nil {
				return err
			}
		}
		field.Set(slice)
	default:
		return fmt.Errorf("unsupported type: %s", field.Kind())
	}

	return nil
}
