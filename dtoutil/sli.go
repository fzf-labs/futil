package dto

import (
	"fmt"
	"reflect"
)

const maxDepth = 100

// NilSliceToEmptySlice 将 nil 切片、映射和指针转换为空等价物。
func NilSliceToEmptySlice(inter any) (any, error) {
	if inter == nil {
		return nil, nil
	}
	return nilSliceToEmptySliceWithDepth(inter, 0)
}

// nilSliceToEmptySliceWithDepth 是一个递归辅助函数，用于处理输入并将 nil 切片、映射和指针转换为空等价物，同时适当处理其他类型。
func nilSliceToEmptySliceWithDepth(inter any, depth int) (any, error) {
	// 防止递归深度过大
	if depth > maxDepth {
		return nil, fmt.Errorf("max recursion depth exceeded")
	}
	// 如果输入为 nil，返回 nil
	if inter == nil {
		return nil, nil
	}
	// 获取输入值的反射值
	val := reflect.ValueOf(inter)
	// 处理接口类型
	if val.Kind() == reflect.Interface {
		if val.IsNil() {
			return nil, nil
		}
		// 处理接口内部的实际值
		processed, err := nilSliceToEmptySliceWithDepth(val.Elem().Interface(), depth+1)
		if err != nil {
			return nil, fmt.Errorf("failed to process interface value: %w", err)
		}
		return processed, nil
	}
	// 根据反射值的类型进行处理
	switch val.Kind() {
	case reflect.Slice:
		if val.IsNil() {
			// 对于 nil 切片，返回空切片
			return reflect.MakeSlice(val.Type(), 0, 0).Interface(), nil
		}
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Cap())
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			processed, err := nilSliceToEmptySliceWithDepth(item.Interface(), depth+1)
			if err != nil {
				return nil, fmt.Errorf("failed to process slice element at index %d: %w", i, err)
			}

			if processed != nil {
				processedVal := reflect.ValueOf(processed)
				if !canAssign(processedVal.Type(), val.Type().Elem()) {
					return nil, fmt.Errorf("type mismatch: cannot assign %v to slice element of type %v",
						processedVal.Type(), val.Type().Elem())
				}
				newSlice = reflect.Append(newSlice, processedVal)
			}
		}
		return newSlice.Interface(), nil
	case reflect.Struct:
		// 创建新的结构体，并复制原始值
		newStruct := reflect.New(val.Type()).Elem()
		newStruct.Set(val) // 先复制所有字段，包括未导出字段
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			if !field.IsExported() {
				continue // 保持未导出字段的原始值
			}

			oldField := val.Field(i)
			newField := newStruct.Field(i)

			processed, err := nilSliceToEmptySliceWithDepth(oldField.Interface(), depth+1)
			if err != nil {
				return nil, fmt.Errorf("failed to process struct field %s: %w", field.Name, err)
			}

			if processed != nil {
				processedVal := reflect.ValueOf(processed)
				if !canAssign(processedVal.Type(), newField.Type()) {
					return nil, fmt.Errorf("type mismatch: cannot assign %v to struct field %s of type %v",
						processedVal.Type(), field.Name, newField.Type())
				}
				newField.Set(processedVal)
			}
		}
		return newStruct.Interface(), nil
	case reflect.Map:
		if val.IsNil() {
			// 对于 nil map，返回空 map
			return reflect.MakeMap(val.Type()).Interface(), nil
		}
		newMap := reflect.MakeMap(val.Type())
		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()

			processed, err := nilSliceToEmptySliceWithDepth(value.Interface(), depth+1)
			if err != nil {
				return nil, fmt.Errorf("failed to process map value for key %v: %w",
					key.Interface(), err)
			}

			if processed != nil {
				processedVal := reflect.ValueOf(processed)
				if !canAssign(processedVal.Type(), val.Type().Elem()) {
					return nil, fmt.Errorf("type mismatch: cannot assign %v to map value of type %v",
						processedVal.Type(), val.Type().Elem())
				}
				newMap.SetMapIndex(key, processedVal)
			}
		}
		return newMap.Interface(), nil
	case reflect.Ptr:
		if val.IsNil() {
			return nil, nil
		}
		processed, err := nilSliceToEmptySliceWithDepth(val.Elem().Interface(), depth+1)
		if err != nil {
			return nil, fmt.Errorf("failed to process pointer value: %w", err)
		}
		if processed == nil {
			return nil, nil
		}
		// 创建新的指针并设置处理后的值
		processedVal := reflect.ValueOf(processed)
		newPtr := reflect.New(processedVal.Type())
		newPtr.Elem().Set(processedVal)
		return newPtr.Interface(), nil

	case reflect.Array:
		// 处理数组类型
		newArray := reflect.New(val.Type()).Elem()
		for i := 0; i < val.Len(); i++ {
			processed, err := nilSliceToEmptySliceWithDepth(val.Index(i).Interface(), depth+1)
			if err != nil {
				return nil, fmt.Errorf("failed to process array element at index %d: %w", i, err)
			}
			if processed != nil {
				processedVal := reflect.ValueOf(processed)
				if !canAssign(processedVal.Type(), val.Type().Elem()) {
					return nil, fmt.Errorf("type mismatch: cannot assign %v to array element of type %v",
						processedVal.Type(), val.Type().Elem())
				}
				newArray.Index(i).Set(processedVal)
			}
		}
		return newArray.Interface(), nil

	default:
		// 对于其他类型，直接返回原值
		return inter, nil
	}
}

// canAssign 检查类型 'from' 的值是否可以赋值给类型 'to'。
func canAssign(from, to reflect.Type) bool {
	// 处理 nil 值
	if from == nil {
		return false
	}
	// 如果目标类型是接口，检查源类型是否实现了该接口
	if to.Kind() == reflect.Interface {
		return from.Implements(to)
	}
	// 处理指针类型
	if from.Kind() == reflect.Ptr && to.Kind() != reflect.Ptr {
		from = from.Elem()
	}
	if to.Kind() == reflect.Ptr && from.Kind() != reflect.Ptr {
		to = to.Elem()
	}
	// 检查类型是否直接可赋值
	if from.AssignableTo(to) {
		return true
	}
	// 检查基础类型是否兼容（例如：*int 到 int）
	return from.ConvertibleTo(to)
}
