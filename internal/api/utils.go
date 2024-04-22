package api

import "reflect"

func includes[T any](arr []T, el T) bool {
	for _, e := range arr {
		if reflect.DeepEqual(e, el) {
			return true
		}
	}

	return false
}
