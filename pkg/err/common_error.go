package err

func HandleError(e error) bool {
	if e == nil {
		return false
	} else {
		return true
	}
}
