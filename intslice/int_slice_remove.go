package intslice

import (
	"fmt"

	"github.com/s-kmmr/slice-remover/code"
)

// Remove 指定したidxの値をslice(int)から削除する。(順序は保持される)
func Remove(slice []int, idx uint) ([]int, error) {
	if (len(slice) - 1) < int(idx) {
		return slice, code.NewSliceErrorWithMsg(code.CapOver, fmt.Sprintf("slice bounds out of range [:%v] with capacity %v", idx, len(slice)))
	}
	return append(slice[:idx], slice[idx+1:]...), nil
}
