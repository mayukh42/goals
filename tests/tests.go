package tests

import (
	"fmt"

	a "github.com/mayukh42/goals/algo"
	c "github.com/mayukh42/goals/collections"
	"github.com/mayukh42/goals/dp"
	g "github.com/mayukh42/goals/graphs"
	l "github.com/mayukh42/goals/list"
	o "github.com/mayukh42/goals/osal"
	t "github.com/mayukh42/goals/tree"
)

func Tests() {
	fmt.Printf("Test Goals!\n")

	// TestLists()
	// TestTrees()
	// TestDP()
	// TestCollections()
	// TestAlgo()
	// TestOSAlgorithms()
	// TestClassics()
	TestGraphs()
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

func TestOSAlgorithms() {
	// o.LRUTest()
	// o.TestBitStr()
	// o.TestByteStream()
	o.TestLRURefMatrixSet()
}

func TestGraphs() {
	// g.TestUFDS1()
	g.TestUFDS2()
}
