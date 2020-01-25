package reflect

import "testing"

func TestCheckType(t *testing.T) {
	var (
		i int
		f float32
		d *float64
		s struct{}
	)
	CheckType(i)
	CheckType(f)
	CheckType(d)
	CheckType(s)
}
