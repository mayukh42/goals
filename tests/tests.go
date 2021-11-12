package tests

import (
	"fmt"

	a "github.com/mayukh42/goals/algo"
	c "github.com/mayukh42/goals/collections"
	d "github.com/mayukh42/goals/dp"
	g "github.com/mayukh42/goals/graphs"
	"github.com/mayukh42/goals/lang"
	o "github.com/mayukh42/goals/osal"
	t "github.com/mayukh42/goals/tree"
)

func Tests() {
	fmt.Printf("Test Goals!\n")

	// TestLists()
	// TestTrees()
	// TestDP()
	// TestCollections()
	TestAlgo()
	// TestOSAlgorithms()
	// TestClassics()
	// TestGraphs()
	// TestLanguage()
}

// tests - the old school way
func TestLists() {
	// l.CreateAndPrintIntList()
	// l.CreateAndPrintStringList()
	// l.DeleteIntList()
	// l.SplitMergeIntList()
	// l.TestReverse()
	// l.TestRotate()
	// l.TestDListSuffixed()
	// l.TestDListPrefixed()
	// l.TestReverseArr()
}

func TestTrees() {
	// t.CreateAndInfix()
	// t.CreateAndPrefix()
	// t.CreateAndPostfix()
	// t.HeightBalanced()
	t.TestLCA()
}

func TestDP() {
	d.TestDP()
}

func TestAlgo() {
	a.TestAlgo()
}

func TestLanguage() {
	lang.TestLanguage()
}

func TestCollections() {
	// c.EnqueueTest()
	// c.DequeueTest()
	// c.PushTest()
	// c.PopTest()
	c.FSQTest()
}

func TestOSAlgorithms() {
	// o.LRUTest()
	// o.TestBitStr()
	// o.TestByteStream()
	o.TestLRURefMatrixSet()
}

func TestGraphs() {
	// g.TestUFDS1()
	// g.TestUFDS2()
	// g.TestUFDS3()
	// g.TestUFDSRef1()
	// g.TestUFDSRef2()
	g.MegaUFDSTest()
}
