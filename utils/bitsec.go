package utils

type Comp func(left interface{}, wright interface{}) int

// BisectLeft all(val >= x for val in a[i:hi]) for the right side.
// TODO need to optimize the implementation
func BisectLeft(items []interface{}, value interface{}, cmp Comp) int {
	for index, item := range items {
		if cmp(item, value) >= 0 {
			return index
		}
	}
	return -1
}

//all(val > x for val in a[i:hi])
func BisectRight(items []interface{}, value interface{}, cmp Comp) int {
	for index, item := range items {
		if cmp(item, value) > 0 {
			return index
		}
	}
	return -1
}

// how to defined common method lib
//This is equivalent to a.insert(bisect.bisect_left(a, x, lo, hi), x)
/*func InsortLeft(items []interface{}, value interface{}, cmp Comp) []interface{} {
	for index, item := range items {
		if cmp(item, value) > 0 {
			return index
		}
	}
	return -1
}

func InsortRight(items []interface{}, value interface{}, cmp Comp) int {
	for index, item := range items {
		if cmp(item, value) > 0 {
			return index
		}
	}
	return -1
}*/
