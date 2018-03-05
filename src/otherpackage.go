package src

import "strconv"

func GetMessageFromOtherPackage(i int) (int, string) {
	i = i + 1
	return i, "Second Message " + strconv.Itoa(i)
}
