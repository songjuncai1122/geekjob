package slice

func Delete(slice []int, ele int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == ele {
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice)-1]
		}
	}

	return slice
}

func Delete2(slice []int, ele int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == ele {
			slice[i] = slice[len(slice)-1]
			slice = slice[:len(slice)-1]
		}
	}
	return slice
}
