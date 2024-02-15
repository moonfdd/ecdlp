package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//比自己小
	if false {
		strList := make([]string, 0)
		for ii := 1; ii < 100; ii += 1 {

			iList := []int{4*ii - 1}
			// i := 2*ii + 1 //1个1
			// i := 4*ii - 1
			// i := 8*ii - 1 //3个1
			// i := 16*ii - 1 //4哥1
			// i := 32*ii - 1 //5个1

			for j := 0; j < len(iList); j++ {
				k := iList[j]
				itemp := k
				str := ""
				for {
					if itemp%2 == 1 {
						itemp = itemp*3 + 1
						itemp /= 2
						str = "1" + str
						// str += "1"
					} else {
						itemp /= 2
						str = "0" + str
						// str += "0"
					}
					if itemp < k {
						break
					}
					// fmt.Println(itemp)
					if itemp == 1 {
						break
					}
				}

				str = strings.TrimLeft(str, "0")
				str = strings.Trim(str, "0")
				strList = append(strList, str)
				fmt.Println(k, "_", len(str), ":", str)
			}
		}
		sort.StringSlice(strList).Sort()
		for i := 0; i < len(strList); i++ {
			// fmt.Println(i, ":", strList[i])
		}
		return
	}
	//到1
	if true {
		strList := make([]string, 0)
		for ii := 1; ii < 100; ii += 1 {

			iList := []int{2*ii + 1}
			// i := 2*ii + 1 //1个1
			// i := 4*ii - 1
			// i := 8*ii - 1 //3个1
			// i := 16*ii - 1 //4哥1
			// i := 32*ii - 1 //5个1
			// 0703 _ 78 : 101000111110000111101110010100111011101111101111000100111010110111111010111111
			// 1055 _ 77 : 10100011111000011110111001010011101110111110111100010011101011011111101011111

			for j := 0; j < len(iList); j++ {
				k := iList[j]
				itemp := k
				str := ""
				for {
					if itemp%2 == 1 {
						itemp = itemp*3 + 1
						itemp /= 2
						str = "1" + str
						// str += "1"
					} else {
						itemp /= 2
						// str = "0" + str
						// str += "0"
					}
					// if itemp < k {
					// 	break
					// }
					// fmt.Println(itemp)
					if itemp == 1 {
						break
					}
				}

				str = strings.TrimLeft(str, "0")
				str = strings.Trim(str, "0")
				strList = append(strList, str)
				fmt.Println(k, "_", len(str), ":", str)
			}
		}
		sort.StringSlice(strList).Sort()
		for i := 0; i < len(strList); i++ {
			// fmt.Println(i, ":", strList[i])
		}
		return
	}
	//根据2到n，求n+1次数
	if true {
		a := make(map[int]struct{})
		for ii := 1; ii <= 243; ii++ {
			i := ii
			itemp := i
			count := 0
			for {
				if _, ok := a[itemp]; ok {
					break
				}
				if itemp == 1 {
					break
				}
				a[itemp] = struct{}{}
				if itemp%2 == 1 {
					itemp = itemp*3 + 1
					itemp /= 2
				} else {
					itemp /= 2
				}
				count++
			}
			a[i] = struct{}{}
			fmt.Println(i, count)
		}
		return
	}
	if true {
		a := make(map[int]struct{})
		for i := 1; i <= 128; i++ {
			a[i] = struct{}{}
			itemp := i
			for {
				if itemp == 1 {
					break
				}
				if itemp%2 == 1 {
					itemp = itemp*3 + 1
					itemp /= 2
				} else {
					itemp /= 2
				}
				a[itemp] = struct{}{}
			}
		}
		fmt.Println(len(a))
		return
	}
	fmt.Println("Modular Inverse")
}
