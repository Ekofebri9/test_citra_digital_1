package main

import (
	"fmt"
	"sort"
	"strings"
)

func sorter(inputan string) string {
	arr := strings.Split(inputan, "")
	vocal := "aAiIuUeEoO"
	hurufHidup := []string{}
	hurufMati := []string{}
	huruf := []string{}
	for i := 0; i < len(arr); i++ {
		if strings.Contains(vocal, arr[i]) {
			hurufHidup = append(hurufHidup, arr[i])
		} else {
			hurufMati = append(hurufMati, arr[i])
		}
	}
	sort.Strings(hurufHidup)
	sort.Strings(hurufMati)
	huruf = append(hurufHidup, hurufMati...)
	return strings.Join(huruf, "")
}

func main() {
	fmt.Println(sorter("omama"))
	fmt.Println(sorter("osama"))
}
