package main

import (
	"fmt"
	"sort"
)

func main1() {
	// 个数 数字 是否存在
	set := make(map[int]map[int]struct{})
	for ii := 3; ii < 1048576; ii += 2 {
		i := ii
		count := 0
		for {
			if i&1 == 1 {
				i *= 3
				i++
				i >>= 1
				count++
			} else {
				i >>= 1
			}
			if i == 1 {
				break
			}
		}
		if _, ok := set[count]; !ok {
			set[count] = make(map[int]struct{})
		}
		set[count][ii] = struct{}{}
	}

	// set = JianHua(set)
	PrintMap(set)
}

func JianHua(set map[int]map[int]struct{}) map[int]map[int]struct{} {
	keys := make([]int, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	setout := make(map[int]map[int]struct{})
	for _, k := range keys {
		setout[k] = make(map[int]struct{})
		keys2 := make([]int, 0, len(set[k]))
		for k2 := range set[k] {
			keys2 = append(keys2, k2)
		}
		sort.Ints(keys2)
		setout[k][keys2[0]] = struct{}{}
		for i := len(keys2) - 1; i >= 1; i-- {
			b := false
			if !b { // 101变成1
				keys2iTemp := keys2[i]
				keys2iTemp--
				if keys2iTemp%4 == 0 {
					keys2iTemp /= 4
					if _, ok := set[k][keys2iTemp]; !ok {
						b = true
					}
				} else {
					b = true
				}
			}
			if b { // 1110001变成11
				keys2iTemp := keys2[i]
				if keys2iTemp&113 == 113 {
					// fmt.Println(keys2iTemp, "!=", 113)
					b = false
				}
			}
			if b { // 11100011变成11
				keys2iTemp := keys2[i]
				if keys2iTemp&227 == 227 {
					// fmt.Println(keys2iTemp, "!=", 227)
					b = false
				}
			}
			if b {
				setout[k][keys2[i]] = struct{}{}
			}
		}

	}
	return setout
}

func PrintMap(set map[int]map[int]struct{}) {
	keys := make([]int, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Print(k, ":")
		keys2 := make([]int, 0, len(set[k]))
		for k2 := range set[k] {
			keys2 = append(keys2, k2)
		}
		sort.Ints(keys2)
		fmt.Println(keys2)
	}
}

func main() {
	//比自己小
	if false {
		strList := make([]string, 0)
		for ii := 002; ii < 400; ii += 1 {

			// iList := []int{ii}
			iList := []int{2*ii - 1}
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

				// str = strings.TrimLeft(str, "0")
				// str = strings.Trim(str, "0")
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
	if false {
		strList := make([]string, 0)
		for ii := 3; ii < 343; ii += 2 {

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
						str = "0" + str
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

				// str = strings.TrimLeft(str, "0")
				// str = strings.Trim(str, "0")
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
	if false {
		a := make(map[int]struct{})
		for ii := 1; ii <= 2430; ii++ {
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
	//除了27很长，还有更长的
	//703 81
	if true {
		a := make(map[int]struct{})
		for ii := 1; ii <= 100; ii += 2 {
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
					count++
				} else {
					itemp /= 2
				}
				// count++
			}
			a[i] = struct{}{}
			// if ii != 401151 && count > 167 {
			// 	fmt.Println(ii, count)
			// }
			fmt.Println(i, count)
		}
		return
	}
	if false {
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
