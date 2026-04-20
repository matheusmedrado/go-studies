package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	for _, searchedNum := range list{
		if searchedNum == num {
			return true
		}
	}
	return false
}
