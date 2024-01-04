package tools

import (
	"reflect"
	"strings"
)

type Path struct {
	Key   string
	Value *string
}

func GetPaths[T any](v *T) []Path {
	if v == nil {
		return nil
	}

	var paths []Path
	view := reflect.ValueOf(*v)
	typeOfView := view.Type()

	for i := 0; i < view.NumField(); i++ {
		field := view.Field(i)
		fieldName := typeOfView.Field(i).Name
		// Check if the field name ends with "Path"
		if strings.HasSuffix(fieldName, "Path") && !field.IsNil() {
			pathValue := field.Elem().String()
			paths = append(paths, Path{
				Key:   fieldName,
				Value: &pathValue,
			})
		}
	}

	return paths
}

/* Pass in a function to get the inner object that contains the exapnded Paths
 * For example:
 * In the case of `v AppEntitlementWithUserBindings` you would pass in a getter that returns
 * v.AppEntitlementView which is a `*AppEntitlementView`, expects a pointer
 */

// func GetMappedJSONPaths[T, V any](list []T, getStructWithPaths func(T) *V) (map[string]int, error) {
// 	if len(list) == 0 {
// 		return make(map[string]int), nil
// 	}

// 	return MapJSONPaths[T](list,
// }

func MapJSONPaths[T any](list []T, getPaths func(T) []Path) (map[string]int, error) {
	res := make(map[string]int)
	for _, item := range list {
		for _, path := range getPaths(item) {
			index, err := GetJSONPathIndex(path.Value)
			if err != nil {
				return nil, err
			}
			if index != -1 {
				res[path.Key] = index
			}
		}
	}
	return res, nil
}
