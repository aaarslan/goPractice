package main

import (
	"fmt"
	"sort"
)

/*  Use https://godoc.org/sort to sort the following:

(1)
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}

(2)
s := []string{"Zeno", "John", "Al", "Jenny"}

(3)
n := []int{7,4,8,2,9,19,12,32,3}

Also sort the above in reverse order
*/

func main() {
	s := []string{"Meiko", "Tharja", "Nanoka", "Matsumoto", "Camilla"}
	sort.Strings(s)
	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	type people []string
	studyGroup := people{"Nanoka", "Miu", "Rumia", "Mito"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

}
