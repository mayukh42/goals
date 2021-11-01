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
	Addrs map[u.Any]int
	Set   *UFDS
	Size  int
}

func NewUFDSRef(id u.Any) *UFDSRef {
	return &UFDSRef{
		Addrs: make(map[u.Any]int),
		Set:   NewUFDS(id),
		Size:  0,
	}
}

func (ur *UFDSRef) Add(v u.Any) *UFDSRef {
	_, pos := ur.Set.MakeSet(v)
	ur.Addrs[v] = pos
	ur.Size++
	return ur
}

func (ur *UFDSRef) AddNode(n *UFNode) *UFDSRef {
	_, pos := ur.Set.MakeSetWithNode(n)
	ur.Addrs[n.Value] = pos
	ur.Size++
	return ur
}

func (ur *UFDSRef) GetSingleton(v u.Any) *UFNode {
	var n *UFNode = nil
	if p, ok := ur.Addrs[v]; ok {
		return ur.Set.Nodes[p]
	}
	return n
}

func (ur *UFDSRef) Identity(v u.Any) u.Any {
	vs := ur.GetSingleton(v)
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
