package main

import (
	"fmt"

	"github.com/s-kmmr/slice-remover/strslice"
)

func main() {
	// usage
	strarr := []string{"a", "b", "c", "d", "e"}
	// 削除したい要素番号を指定（この例だと1番）
	delarr, err := strslice.Remove(strarr, 1)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	// [a c d e]
	fmt.Printf("%+v", delarr)

}
