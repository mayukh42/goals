package main

import (
	"fmt"

	a "github.com/mayukh42/goals/algo"
	c "github.com/mayukh42/goals/collections"
	"github.com/mayukh42/goals/dp"
	l "github.com/mayukh42/goals/list"
	t "github.com/mayukh42/goals/tree"
)

func main() {
	fmt.Printf("Hello, Go!\n")

	// TestLists()
	// TestTrees()
	// TestDP()
	// TestCollections()
	TestAlgo()
	// TestClassics()
}

// tests - the old school way
func TestLists() {
	// l.CreateAndPrintIntList()
	// l.CreateAndPrintStringList()
	// l.DeleteIntList()
	// l.SplitMergeIntList()
	// l.TestReverse()
	// l.TestRotate()
	l.TestDListSuffixed()
	l.TestDListPrefixed()
}

func TestTrees() {
	// t.CreateAndInfix()
	// t.CreateAndPrefix()
	// t.CreateAndPostfix()
	// t.HeightBalanced()
	t.TestLCA()
}

func TestDP() {
	// dp.TestLevenshtein()
	// dp.TestLCSeq()
	dp.TestMaximalSum()
}

func TestCollections() {
	// c.EnqueueTest()
	// c.DequeueTest()
	// c.PushTest()
	c.PopTest()
}

func TestAlgo() {
	// a.TestTreeBFS()
	// a.TestTreeDFS()
	// a.TestInequality()
	// a.TestPower()
	// a.TestIntDivide()
	// a.TestNewtonSqrt()
	a.TestGraphBFS()
}

func TestClassics() {
	a.TestBinarySearch()
}
