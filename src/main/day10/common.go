package day10

import (
	"fmt"
	"strconv"
)

type CircularList struct {
	List []byte
}

func (l CircularList) Get(i int) byte {
	return l.List[i%len(l.List)]
}

func (l CircularList) Set(i int, value byte) {
	l.List[i%len(l.List)] = value
}

func (l CircularList) subslice(start int, length int) []byte {
	res := make([]byte, length, length)
	for i := 0; i < length; i++ {
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
	newList := CircularList{make([]byte, len(list.List), len(list.List))}
	subSlice := list.subslice(pos, int(length))
	reverse(subSlice)
	for i := 0; i < len(list.List); i++ {
		newList.Set(i, list.Get(i))
	}
	for i := 0; i < int(length); i++ {
		newList.Set(pos+i, subSlice[i])
	}
	return newList
}

func reverse(arr []byte) {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func ComputeHash(input string) string {
	lengths := make([]byte, 0)
	for _, c := range input {
		lengths = append(lengths, byte(c))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	skip := 0
	pos := 0
	l := 256
	list := CircularList{List: make([]byte, l, l)}
	for i := 0; i < l; i++ {
		list.Set(i, byte(i))
	}

	for i := 0; i < 64; i++ {
		list, skip, pos = ComputeOneRound(list, lengths, skip, pos)
	}

	hashval := make([]byte, 16, 16)

	for i := 0; i < 16; i++ {
		hashval[i] = list.Get(i * 16)
		for j := 1; j < 16; j++ {
			hashval[i] = hashval[i] ^ list.Get(i*16+j)
		}
	}

	res := ""
	for i := 0; i < 16; i++ {
		res = res + fmt.Sprintf("%02x", hashval[i])
	}
	return res
}

func ToBinary(hash string) string {
	res := ""
	for _, c := range hash {
		i, _ := strconv.ParseInt(string(c), 16, 5)
		digit := strconv.FormatInt(i, 2)
		for len(digit) < 4 {
			digit = "0" + digit
		}
		res = res + digit
	}
	return res
}
