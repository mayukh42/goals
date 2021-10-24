package list

import (
	"testing"
)

func TestCreateAndPrintList(t *testing.T) {
	xs := List{1, 2, 3, 4, 5}
	ll := CreateFromArray(xs)
	t.Logf("list: %v\n", ll.String())
}
