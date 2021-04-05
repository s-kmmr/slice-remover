package strslice

import (
	"fmt"

	"github.com/s-kmmr/slice-remover/code"
)

// Remove 指定したidxの値をslice(string)から削除する。(順序は保持される)
func Remove(slice []string, idx int) ([]string, error) {
	if (len(slice) - 1) < idx {
		return slice, code.NewSliceErrorWithMsg(code.CapOver, fmt.Sprintf("slice bounds out of range [:%v] with capacity %v", idx, len(slice)))
	}
	return append(slice[:idx], slice[idx+1:]...), nil
}
