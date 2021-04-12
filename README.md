# slice-remover
Removes the specified idx value from the slice

# usage

```
go get -u github.com/s-kmmr/slice-remover
```

```
import (
	"fmt"

	"github.com/s-kmmr/slice-remover/strslice"
)
func main() {
	
	strarr := []string{"a", "b", "c", "d", "e"}

	delarr, err := strslice.Remove(strarr, 1)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	
    // [a c d e]
	fmt.Printf("%+v", delarr)

}
```