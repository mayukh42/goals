package list

// split into 2 lists
func SplitIntoLists(l *Node) (*Node, *Node) {
	var firstH, secondH, firstT, secondT *Node
	if l == nil {
		return firstH, secondH
	}

	one := true
	curr := l
	for curr != nil {
		node := NewNode(curr.Value)
		if one {
			if firstH == nil {
				firstH = node
				firstT = firstH
			} else {
				// silently append and change tail
				firstT = Link(firstT, node)
			}
		} else {
			if secondH == nil {
				secondH = node
				secondT = secondH
			} else {
				secondT = Link(secondT, node)
			}
		}
		curr = curr.Next
		// DFF?
		one = !one
	}

	return firstH, secondH
}

func MergeLists(first, second *Node) *Node {
	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	var head, tail *Node
	currF, currS := first, second
	for currF != nil && currS != nil {
		nodeF := NewNode(currF.Value)
		if head == nil {
			head = nodeF
			tail = head
		} else {
			tail = Link(tail, nodeF)
		}
		currF = currF.Next

		nodeS := NewNode(currS.Value)
		if head == nil {
			head = nodeS
			tail = head
		} else {
			tail = Link(tail, nodeS)
		}
		currS = currS.Next
	}

	for currF != nil {
		nodeF := NewNode(currF.Value)
		if head == nil {
			head = nodeF
			tail = head
		} else {
			tail = Link(tail, nodeF)
		}
		currF = currF.Next
	}

	for currS != nil {
		nodeS := NewNode(currS.Value)
		if head == nil {
			head = nodeS
			tail = head
		} else {
			tail = Link(tail, nodeS)
		}
		currS = currS.Next
	}

	return head
}

/**
 * 1 -> 2 -> 3
 * 1 <-	c
 *		2 -> 3
 *
 * invariant
 * 	prev, curr, head (new)
 * 	h = reversed list up to that node in iteration
 * 	p = next to tag to h
 * 	c = iterator
 */
func Reverse(l *Node) *Node {
	var p, c, h *Node
	c = l
	p = nil

	for c != nil {
		h = c
		c = c.Next
		h.Next = p
		p = h
	}

	return h
}
