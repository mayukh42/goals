package lang

import (
	"log"

	"github.com/mayukh42/goals/utils"
)

/** just for fun
 * elements will print in reverse order
 */
func DeferredStack(xs []utils.Any) error {
	for _, x := range xs {
		defer log.Print(utils.AnyToString(x))
	}
	return nil
}

func ArrayIndexOutofBounds() {
	if r := recover(); r != nil {
		log.Printf("array index out of bounds: %v", r)
	}
}
