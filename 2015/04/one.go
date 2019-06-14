package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var key string
	fmt.Scan(&key)
	var i int
	for i = 1; true; i++ {
		block := fmt.Sprintf("%s%d", key, i)
		md5 := md5.Sum([]byte(block))
		if md5[0] == 0 && md5[1] == 0 && md5[2] <= 0x0f {
			break
		}
	}
	fmt.Println(i)
}
