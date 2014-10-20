// package pathreflect provides the ability to address an object graph using
// a path notation and then modify the addressed node.
package pathreflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Path []string

func Parse(pathString string) Path {
	return Path(strings.Split(pathString, "/"))
}

func (p Path) Set(to interface{}, val interface{}) error {
	if len(p) == 0 {
		return fmt.Errorf("Path must contain at least one element")
	}

	var parent reflect.Value
	current := reflect.ValueOf(to)
	nameOrIndex := ""
	var err error
	for i := 0; i < len(p); i++ {
		if i > 0 {
			parent = current
		}
		nameOrIndex = p[i]
		current, err = getChild(current, nameOrIndex)
		if err != nil {
			return fmt.Errorf("Error traversing beyond path %s", p.through(i))
		}
	}

	if parent.Kind() == reflect.Map {
		// For maps, set the value on the parent
		parent.SetMapIndex(reflect.ValueOf(nameOrIndex), reflect.ValueOf(val))
	} else {
		// For structs and slices, set the value using Set on the terminal field
		current.Set(reflect.ValueOf(val))
	}
	return nil
}

func (p Path) through(i int) string {
	return strings.Join(p[:i], "/")
}

func getChild(parent reflect.Value, nameOrIndex string) (val reflect.Value, err error) {
	if parent.Kind() == reflect.Ptr || parent.Kind() == reflect.Interface {
		if parent.IsNil() {
			err = fmt.Errorf("Empty parent value")
			return
		}
		parent = parent.Elem()
	}

	switch parent.Kind() {
	case reflect.Map:
		val = parent.MapIndex(reflect.ValueOf(nameOrIndex))
		return
	case reflect.Struct:
		val = parent.FieldByName(nameOrIndex)
		return
	case reflect.Array, reflect.Slice:
		i, err2 := strconv.Atoi(nameOrIndex)
		if err2 != nil {
			err = fmt.Errorf("%s is not a valid index for an array or slice", nameOrIndex)
			return
		}
		val = parent.Index(i)
		return
	default:
		err = fmt.Errorf("Unable to extract value from value of kind %s", parent.Kind())
		return
	}
}
