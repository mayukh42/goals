package algo

import (
	"log"

	"github.com/mayukh42/goals/utils"
)

func TestInequality() {
	res := PointsInCircle(27)
	log.Printf("res: %v", res)
}

func TestPower() {
	a, b := 3, 10
	res := Power(a, b)
	log.Printf("%d ** %d = %d", a, b, res)
}

func TestIntDivide() {
	inputs := []utils.Tuple{
		*utils.NewTuple(24, 3),
		*utils.NewTuple(25, 3),
		*utils.NewTuple(28, 3),
		*utils.NewTuple(42, 4),
		*utils.NewTuple(45, 9),
		*utils.NewTuple(63, 10),
		*utils.NewTuple(63, 36),
		*utils.NewTuple(3, 7),
		*utils.NewTuple(7, 7),
	}

	for _, t := range inputs {
		a := t.First.(int)
		b := t.Second.(int)
		q, r := Divide(a, b)
		log.Printf("%d/ %d = %d, %d %% %d = %d", a, b, q, a, b, r)
	}
}

func TestNewtonSqrt() {
	inputs := []int{4, 7, 49, 101, 144, 161, 299, 1024, 2047, 2147483648}
	for _, i := range inputs {
		s := NewtonSqrt(i)
		log.Printf("int sqrt(%d) = %d", i, s)
	}
}
