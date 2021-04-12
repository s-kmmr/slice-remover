package strslice

import (
	"fmt"

	"github.com/s-kmmr/slice-remover/code"
)

// Remove Removes the specified idx value from the slice(string)
func Remove(slice []string, idx uint) ([]string, error) {
	if (len(slice) - 1) < int(idx) {
		return slice, code.NewSliceErrorWithMsg(code.CapOver, fmt.Sprintf("slice bounds out of range [:%v] with capacity %v", idx, len(slice)))
	}
	return append(slice[:idx], slice[idx+1:]...), nil
}
