package day10

type CircularList struct {
	List []byte
}

func (l CircularList) Get(i int) byte {
	return l.List[i % len(l.List)]
}

func (l CircularList) Set(i int, value byte) {
	l.List[i % len(l.List)] = value
}

func (l CircularList) subslice(start int, length int) []byte {
	res := make([]byte, length, length)
	for i := 0; i<length; i++ {
		res[i] = l.Get(start + i)
	}
	return res
}


func ComputeOneRound(list CircularList, lengths []byte, skip int, pos int) (CircularList, int, int) {
	res := list
	for _, length := range lengths {
		res = ComputeNewList(res, length, pos)
		pos = pos + int(length) + skip
		skip++
	}
	return res, skip, pos
}

func ComputeNewList(list CircularList, length byte, pos int) CircularList {
	newList := CircularList{make([]byte, len(list.List),len(list.List))}
	subSlice := list.subslice(pos, int(length))
	reverse(subSlice)
	for i := 0; i<len(list.List); i++ {
		newList.Set(i, list.Get(i))
	}
	for i := 0; i<int(length); i++ {
		newList.Set(pos + i, subSlice[i])
	}
	return newList
}

func reverse(arr []byte) {
	for i:= 0; i<len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}
