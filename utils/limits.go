package utils

const (
	MIN_INT = 0
	// use math.*
	MAX_UINT  = ^uint(0)
	MIN_UINT  = 0
	MAX_INT   = int(MAX_UINT >> 1)
	MIN_INT_2 = -MAX_INT - 1
)
