package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func ap(slice []int, e int) []int {
	slice = append(slice, e)
	return slice
}

func TestSomeSlice(t *testing.T) {
	// slice := make([]int, 0, 5)
	slice := []int{}
	slice = append(slice, 1, 2, 3, 5, 6)
	fmt.Printf("%d, %d, %v\n", len(slice), cap(slice), slice)
	newSlice := ap(slice, 4)
	fmt.Printf("%d, %d, %v\n", len(newSlice), cap(newSlice), newSlice)

	fmt.Println(&slice[0] == &newSlice[0])
}

func TestSomeSlice1(t *testing.T) {
	slice := []int{1, 2, 3, 5, 6, 7, 7}
	// slice = append(slice, 1, 2, 3, 5, 6)
	fmt.Printf("%d, %d, %v\n", len(slice), cap(slice), slice)
	newSlice := ap(slice, 4)
	fmt.Printf("%d, %d, %v\n", len(newSlice), cap(newSlice), newSlice)

	fmt.Println(&slice[0] == &newSlice[0])
}

func TestHt(t *testing.T) {
	s := []int{}
	v := reflect.ValueOf(&s).Elem()
	va := reflect.Append(v, reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3), reflect.ValueOf(4), reflect.ValueOf(5), reflect.ValueOf(6))
	fmt.Printf("%d %d, %v\n", va.Len(), va.Cap(), va)
}

func TestSomeSLiceAppend(t *testing.T) {
	s := make([]int, 5, 5)
	s[0] = 1
	s[1] = 3
	App(s...)
	fmt.Printf("%d %d %v\n", len(s), cap(s), s)
}

func App(slice ...int) {
	fmt.Printf("%d %d %v\n", len(slice), cap(slice), slice)
	slice = append(slice, 3)
	fmt.Printf("%d %d %v\n", len(slice), cap(slice), slice)
}

func TestDK(t *testing.T) {
	data := [10]int{}
	fmt.Println(cap(data))
	slice := data[5:8]
	fmt.Printf("%d %d %v\n", len(slice), cap(slice), slice)
	slice = append(slice, 9) // slice=? data=?
	fmt.Printf("%d %d %v\n", len(slice), cap(slice), slice)
	slice = append(slice, 10, 11, 12) // slice=? data=?
	fmt.Printf("%d %d %v\n", len(slice), cap(slice), slice)
}

func TestSDs(t *testing.T) {
	// tt := []int{0, 8, 16, 24, 32, 48, 64, 80}
	ttt := []uint8{0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8}
	fmt.Println(ttt[5])
	fmt.Println(1 << 13)
	fmt.Println((40 + 1<<13 - 1) &^ (1 << 13))
}

func TestJH(t *testing.T) {
	slice := make([]string, 0, 3)
	slice = append(slice, "0", "1", "3")
	fmt.Printf("%d, %d, %v\n", len(slice), cap(slice), slice)
	slice = append(slice, "566")
	fmt.Printf("%d, %d, %v\n", len(slice), cap(slice), slice)
}

func TestIsMatch(t *testing.T) {
	s := "ab"
	p := ".*"

	match := isMatch(s, p)
	fmt.Println(match)
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		} else if p[j-1] == '.' {
			return true
		} else {
			return s[i-1] == p[j-1]
		}
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if match(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if match(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	for now := range ticker.C {
		fmt.Println(now, now.Second(), now.Day(), now.Hour())
	}

}

func TestKmp(t *testing.T) {
	h, s := "sadbutsad", "sad"
	str := strStr(h, s)
	fmt.Println(str)
}

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}

		next[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == m {
			return i - j + 1
		}
	}
	return -1
}

func TestGetNext(t *testing.T) {
	getNext("ababaca")
}

func getNext(needle string) {
	m := len(needle)
	next := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}

		next[i] = j
	}

	fmt.Printf("%v\n", next)
}

type S struct {
	name string
}

func TestMap(t *testing.T) {
	m := make(map[int32]*S, 10)
	m[1] = nil
	for k, v := range m {
		fmt.Println(k, v)
	}

	v, ok := m[1]
	if ok {
		fmt.Println("111")
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}

func TestStructMap(t *testing.T) {
	ff := make(map[int32]struct{}, 5)
	ff[1] = struct{}{}
	ff[2] = struct{}{}
	for k, v := range ff {
		fmt.Println(k, v)
	}

	_, ok := ff[1]
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("un ok")
	}
	ff[2] = struct{}{}
	_, ok = ff[2]
	if ok {
		fmt.Println(ok)
	} else {
		fmt.Println(!ok)
	}

	for k, v := range ff {
		fmt.Println(k, v)
	}
}
