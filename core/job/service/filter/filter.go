package filter

type filter struct {
	bits  uint64
	value map[Operand]interface{}
}

func NewFilter(opts ...FilterOpt) *filter { _ = "STUB: not implemented"; return nil }

func (f *filter) GetStringValue(operand Operand) string { _ = "STUB: not implemented"; return "" }

func (f *filter) GetStringArrayValue(operand Operand) []string {
	_ = "STUB: not implemented"
	return nil
}

// Contains provide conditional check for the filter if all operands satisfied by the filter.
func (f *filter) Contains(operands ...Operand) bool { _ = "STUB: not implemented"; return false }
