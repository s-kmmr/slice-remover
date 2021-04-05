package code

type SliceCode int

var (
	OK      SliceCode = 0
	CapOver SliceCode = 1
)

func (s SliceCode) String() string {
	switch s {
	case OK:
		return "ok"
	case CapOver:
		return "slice bounds out of range"
	}
	return ""
}
