package pointers

func swap(px, py *int) {
	var temp int
	temp = *px
	*px = *py
	*py = temp
}
