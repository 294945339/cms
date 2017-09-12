// 转结构体工具类
package utils

import (
	"reflect"
	"strconv"
	"errors"
)

func SliceToStruct(o interface{}, data []string, tag string) error {
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		return errors.New("value cannot be changed!")
	} else {
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		index, err := strconv.Atoi(t.Field(i).Tag.Get(tag))
		if err != nil {
			continue
		}
		f := v.Field(i)
		setValue(&f, data[index])
	}

	return nil
}

func MapToStruct(o interface{}, data map[string]string) error {
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		return errors.New("value cannot be changed!")
	} else {
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("mapKey")

		f := v.Field(i)
		setValue(&f, data[key])
	}

	return nil
}

func setValue(f *reflect.Value, data string) {
	switch f.Kind() {
	case reflect.String:
		f.SetString(data)
		break
	case reflect.Int, reflect.Int64:
		d, _ := strconv.ParseInt(data, 10, 64)
		f.SetInt(d)
		break
	case reflect.Float64, reflect.Float32:
		d, _ := strconv.ParseFloat(data, 64)
		f.SetFloat(d)
		break
	case reflect.Bool:
		d, _ := strconv.ParseBool(data)
		f.SetBool(d)
		break
	case reflect.UnsafePointer:
		f.SetPointer(nil)
	case reflect.Ptr:
		switch f.Type().String() {
		case "*string":
			if data != "" {
				f.Set(reflect.ValueOf(&data))
			}
			break
		case "*int":
			d, _ := strconv.Atoi(data)
			if data != "" {
				f.Set(reflect.ValueOf(&d))
			}
			break
		case "*int64":
			d, _ := strconv.ParseInt(data, 10, 64)
			if data != "" {
				f.Set(reflect.ValueOf(&d))
			}
			break
		case "*float32":
			d, _ := strconv.ParseFloat(data, 32)
			if data != "" {
				f.Set(reflect.ValueOf(&d))
			}
			break
		case "*float64":
			d, _ := strconv.ParseFloat(data, 64)
			if data != "" {
				f.Set(reflect.ValueOf(&d))
			}
			break
		case "*bool":
			d, _ := strconv.ParseBool(data)
			if data != "" {
				f.Set(reflect.ValueOf(&d))
			}
			break
		}
	}
}
