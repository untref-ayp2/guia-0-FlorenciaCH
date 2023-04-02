package pointers

func swap1(px, py *int) (int, int) {
	*px = *py
	*py = *px
	return *py, *px
}
