package graphs

import (
	"fmt"
	"log"
	"strings"

	u "github.com/mayukh42/goals/utils"
)

/** union find (partition) data structure
 * linked list
 * tree*
 *
 * requires a ref map for nodes/ indices
 * 	maintained on the caller side
 */

type UFDSRef struct {
	Id    u.Any
	Addrs map[u.Any]int
	Set   *UFDS
	Size  int
}

func NewUFDSRef(id u.Any) *UFDSRef {
	return &UFDSRef{
		Id:    id,
		Addrs: make(map[u.Any]int),
		Set:   NewUFDS(fmt.Sprintf("%v.ufds", u.AnyToString(id))),
		Size:  0,
	}
}

func (ur *UFDSRef) MakeSet(v u.Any) *UFDSRef {
	_, pos := ur.Set.MakeSet(v)
	ur.Addrs[v] = pos
	ur.Size++
	return ur
}

func (ur *UFDSRef) MakeSetWithNode(n *UFNode) *UFDSRef {
	_, pos := ur.Set.MakeSetWithNode(n)
	ur.Addrs[n.Value] = pos
	ur.Size++
	return ur
}

func (ur *UFDSRef) GetSingleton(v u.Any) (*UFNode, int) {
	var n *UFNode = nil
	if p, ok := ur.Addrs[v]; ok {
		return ur.Set.Nodes[p], p
	}
	return n, -1
}

func (ur *UFDSRef) UnionSingleton(u2 *UFNode) *UFNode {
	return ur.Set.Union(ur.Set.Root, u2)
}

func (ur *UFDSRef) Union(vr *UFDSRef) *UFDSRef {
	if vr == nil {
		return ur
	}
	if vr.Size == 0 {
		return ur
	}

	// union the two sets in ur and vr
	if ur.Size >= vr.Size {
		ur.mergeMap(vr)
		ur.union(ur.Set, vr.Set)
	} else {
		// TODO: refactor to cleaner method sigs
		vr.mergeMap(ur)
		ur.Addrs = vr.Addrs
		ur.union(vr.Set, ur.Set)
	}

	return ur
}

func (ur *UFDSRef) union(u1, u2 *UFDS) {
	// assume u1.Size >= u2.Size
	u1.Root = u1.Union(u1.Root, u2.Root)
	u1.Nodes = append(u1.Nodes, u2.Nodes...)
	u1.Size = len(u1.Nodes)
	ur.Set = u1
	ur.Size = u1.Size
}

func (ur *UFDSRef) mergeMap(vr *UFDSRef) {
	for k, v := range vr.Addrs {
		ur.Addrs[k] = v + ur.Size
	}
}

func (ur *UFDSRef) Find(e *UFNode) *UFNode {
	return ur.Set.Find(e)
}

func (ur *UFDSRef) Consistent() bool {
	return ur.Set.Consistent()
}

func (ur *UFDSRef) Identity(v u.Any) u.Any {
	vs, _ := ur.GetSingleton(v)
	if vs == nil {
		log.Printf("%v not found in any set", v)
		return nil
	}

	root := ur.Set.Find(vs)
	if root == nil {
		log.Printf("root not found")
		return nil
	}
	return root.Value
}

func (ur *UFDSRef) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("ufdsref[%d] \n\t", ur.Size))
	sb.WriteString(fmt.Sprintf("addr map: %v", ur.Addrs))
	sb.WriteString(ur.Set.String())
	return sb.String()
}
