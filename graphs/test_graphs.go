package graphs

import (
	u "github.com/mayukh42/goals/utils"
)

func CreateGraphAL() *VGraph {
	vs := map[u.Any]u.List{
		6: {4},
		4: {3, 5, 6},
		3: {2, 4},
		2: {1, 3, 5},
		5: {1, 2, 4},
		1: {2, 5},
	}

	return NewVGraph(6, vs)
}
