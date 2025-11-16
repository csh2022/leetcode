package lib

import (
	"reflect"
)

func IsTwoIntSlicesEqual(a []int, b []int) bool {
	return reflect.DeepEqual(a, b)
}
