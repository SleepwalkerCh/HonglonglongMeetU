package tools

func InList(element interface{}, list interface{}) (res bool) {
	switch element.(type) {
	case int:
		return InIntList(element.(int), list)
	case int32:
		return InInt32List(element.(int32), list)
	case string:
		return InStringList(element.(string), list)
	}
	return false
}

func InIntList(element int, list interface{}) (res bool) {
	for _, singleEle := range list.([]int) {
		singleEleInt := singleEle
		if element == singleEleInt {
			return true
		}
	}
	return false
}

func InInt32List(element int32, list interface{}) (res bool) {
	for _, singleEle := range list.([]int32) {
		singleEleInt := singleEle
		if element == singleEleInt {
			return true
		}
	}
	return false
}

func InStringList(element string, list interface{}) (res bool) {
	for _, singleEle := range list.([]string) {
		singleEleInt := singleEle
		if element == singleEleInt {
			return true
		}
	}
	return false
}
