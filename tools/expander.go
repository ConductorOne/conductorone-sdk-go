package tools

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const (
	atTypeApp             = "type.googleapis.com/c1.api.app.v1.App"
	atTypeAppResource     = "type.googleapis.com/c1.api.app.v1.AppResource"
	atTypeAppResourceType = "type.googleapis.com/c1.api.app.v1.AppResourceType"
)

type Path struct {
	Key   string
	Value *string
}
type marshallable interface {
	MarshalJSON() ([]byte, error)
}

type marshalJSON[T any] func(T) ([]byte, error)
type getStructWithPaths[T, V any] func(T) *V
type makeExpandedObject[T, V any] func(T, map[string]*any) *V

func ExpandResponse[T, K, I any, V marshallable](responseList []T, expandedList []V, structWithPaths getStructWithPaths[T, K], makeResult makeExpandedObject[T, I]) ([]*I, error) {
	expanded := make([]any, 0, len(expandedList))
	for _, x := range expandedList {
		x := x
		converted, err := GetMarshalledObject(x)
		if err != nil {
			return nil, err
		}
		expanded = append(expanded, converted)
	}
	result := make([]*I, 0, len(responseList))
	for _, response := range responseList {
		response := response
		expandedMap, err := GetMappedJSONPaths(response, structWithPaths)
		if err != nil {
			return nil, err
		}
		expandedObjects := PopulateExpandedMap(expandedMap, expanded)
		result = append(result, makeResult(response, expandedObjects))
	}

	return result, nil
}

// Populate the expanded map with references to the related objects.
func PopulateExpandedMap(expandMap map[string]int, expanded []any) map[string]*any {
	rv := make(map[string]*any)
	for k, v := range expandMap {
		rv[k] = &expanded[v]
	}
	return rv
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

/* Pass in the list you want to expand with a function to get the inner object that contains the exapnded Paths
 * For example:
 * In the case of `v AppEntitlementWithUserBindings` you would pass in a getter that returns
 * v.AppEntitlementView which is a `*AppEntitlementView`, expects a pointer
 */
func GetMappedJSONPaths[T, V any](item T, structWithPaths getStructWithPaths[T, V]) (map[string]int, error) {
	fn := func(t T) []Path {
		v := structWithPaths(t)
		paths := GetPaths(v)
		return paths
	}
	return mapJSONPaths[T](item, fn)
}

func mapJSONPaths[T any](item T, getPaths func(T) []Path) (map[string]int, error) {
	res := make(map[string]int)
	for _, path := range getPaths(item) {
		index, err := GetJSONPathIndex(path.Value)
		if err != nil {
			return nil, err
		}
		if index != -1 {
			res[path.Key] = index
		}
	}

	return res, nil
}

func GetAtTypeWithReflection[T any](input *T) *string {
	inputVal := reflect.ValueOf(input)
	if inputVal.Kind() != reflect.Ptr {
		return nil
	}

	asTypeMethod := inputVal.MethodByName("GetAtType")
	if !asTypeMethod.IsValid() {
		return nil
	}

	result := asTypeMethod.Call(nil)
	if len(result) != 1 {
		return nil
	}

	// Check if the result can be converted to *string
	asTypeValue, ok := result[0].Interface().(*string)
	if !ok {
		return nil
	}

	return asTypeValue
}

func GetMarshalledObject[T marshallable](input T) (any, error) {
	getAtType := GetAtTypeWithReflection[T]
	marshall := func(input T) ([]byte, error) {
		return input.MarshalJSON()
	}
	return AtTypeToObject(input, getAtType, marshall)
}

func As[T any, V any](input T, marshal marshalJSON[T]) (*V, error) {
	d, err := marshal(input)
	if err != nil {
		return nil, err
	}

	var rv V
	err = json.Unmarshal(d, &rv)
	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func AtTypeToObject[T any](input T, getAtType func(*T) *string, marshal marshalJSON[T]) (any, error) {
	inputType := getAtType(&input)
	if inputType == nil {
		return nil, errors.New("input type is nil")
	}

	switch *inputType {
	case atTypeApp:
		return As[T, shared.App](input, marshal)
	case atTypeAppResource:
		return As[T, shared.AppResource](input, marshal)
	case atTypeAppResourceType:
		return As[T, shared.AppResourceType](input, marshal)
	default:
		return nil, errors.New("unknown type")
	}
}
