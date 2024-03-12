package sliutil

import (
	"reflect"
)

// NilSliceToEmptySlice 递归地查找给定接口中的所有nil切片，并用空切片替换它们。
// 下面列出了此功能的缺点，以避免可能的误用。
// 1.此函数将返回给定结构的指针，以防指针a传入，将返回一个指针，而不是双指针
// 2.在此过程中，结构的所有私有字段都将被删除。
func NilSliceToEmptySlice(any interface{}) interface{} {
	switch any.(type) {
	case string, []byte:
		return any
	}
	val := reflect.ValueOf(any)
	switch val.Kind() {
	case reflect.Ptr:
		// if, in some case, a pointer is passed in, We dereference it and do the normal stuff
		if val.IsNil() {
			if reflect.TypeOf(any).Elem().Kind() == reflect.Slice {
				res := reflect.New(reflect.TypeOf(any).Elem())
				resp := reflect.MakeSlice(reflect.TypeOf(any).Elem(), 0, 1)
				res.Elem().Set(resp)
				return res.Interface()
			}
			return any
		}
		return NilSliceToEmptySlice(val.Elem().Interface())
	case reflect.Slice:
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Len())
		// if this is not empty, copy it
		if !val.IsZero() {
			// iterate over each element in slice
			for i := 0; i < val.Len(); i++ {
				item := val.Index(i)
				var newItem reflect.Value
				switch item.Kind() {
				case reflect.Struct, reflect.Slice, reflect.Map:
					// recursively handle nested struct
					newItem = reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(item.Interface())))
				case reflect.Ptr:
					if item.IsNil() {
						if item.Type().Elem().Kind() == reflect.Slice {
							newSliceNotNil := reflect.MakeSlice(item.Type().Elem(), 0, 1)
							newItem = newSliceNotNil
							break
						}
						newItem = item
						break
					}
					if item.Elem().Kind() == reflect.Struct || item.Elem().Kind() == reflect.Slice {
						newItem = reflect.ValueOf(NilSliceToEmptySlice(item.Elem().Interface()))
						break
					}
					fallthrough
				default:
					newItem = item
				}
				newSlice = reflect.Append(newSlice, newItem)
			}
		}
		return newSlice.Interface()
	case reflect.Struct:
		// new struct that will be returned
		newStruct := reflect.New(reflect.TypeOf(any))
		newVal := newStruct.Elem()
		// iterate over input's fields
		for i := 0; i < val.NumField(); i++ {
			newValField := newVal.Field(i)
			if !newValField.CanSet() {
				continue
			}
			valField := val.Field(i)
			updatedField := reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))
			if valField.Kind() == reflect.Ptr {
				if updatedField.IsValid() {
					newValField.Set(updatedField)
				}
			} else {
				if updatedField.IsValid() {
					newValField.Set(reflect.Indirect(updatedField))
				}
			}
		}
		return newStruct.Interface()
	case reflect.Map:
		// new map to be returned
		newMap := reflect.MakeMap(reflect.TypeOf(any))
		// iterate over every key value pair in input map
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			// recursively handle nested value
			newV := reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(v.Interface())))
			newMap.SetMapIndex(k, newV)
		}
		return newMap.Interface()
	default:
		return any
	}
}
