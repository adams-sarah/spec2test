// generated by stringer -type=Error; DO NOT EDIT

package spec2test

import "fmt"

const _Error_name = "ErrorStructTypeRequiredErrorResourceAlreadyAdded"

var _Error_index = [...]uint8{0, 23, 48}

func (i Error) String() string {
	if i < 0 || i >= Error(len(_Error_index)-1) {
		return fmt.Sprintf("Error(%d)", i)
	}
	return _Error_name[_Error_index[i]:_Error_index[i+1]]
}
