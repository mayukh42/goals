package lang

import (
	"log"

	"github.com/mayukh42/goals/utils"
)

func TestLanguage() {
	// testDeferredStack()
	testArrayOutofBounds()
}

func TestLanguageAll() {
	testDeferredStack()
	testArrayOutofBounds()
}

func testDeferredStack() {
	DeferredStack(utils.List{1, 2, 3, 4, 5})
}

func testArrayOutofBounds() {
	xs := utils.List{1, 2, 3, 4, 5}
	printArray(xs)
	log.Printf("howdy")
}

func printArray(xs utils.List) {
	i := 0
	defer ArrayIndexOutofBounds()
	for {
		log.Printf("xs[%d] = %v", i, xs[i])
		i++

		if i >= 10 {
			// lazy prog that bound the array to 10 elems
			break
		}
	}

	// this won't be executed
	log.Printf("howdy from print")
}
