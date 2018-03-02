package main

import (
	"fmt"
	"io"
	"os"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	input := os.Args[1]
	
	mdh := md5.New()
	io.WriteString(mdh, input)
	mdhByte := mdh.Sum(nil) //this return []Byte
	mdhString := hex.EncodeToString(mdhByte) //convert the Byte to string
	
	fmt.Println("MD5 of input string is: ", mdhString)
}
