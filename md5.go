package main

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(e []byte) {
	m := md5.Sum(e)
	fmt.Printf("%x\n", m) //将[]byte转成16进制
}
