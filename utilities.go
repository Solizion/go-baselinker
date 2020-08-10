package baselinker

func (list intArray) has(number int) bool {
	for _, item := range list {
		if item == number {
			return true
		}
	}

	return false
}
