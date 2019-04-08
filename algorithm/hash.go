package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

func main() {
	fmt.Println(hash2("sfs"))
}

func hash2(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}

func hash1() {
	teststring := "welcome to beijing"

	//MD5
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(teststring))
	result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)

	//SHA1
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(teststring))
	result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)

	// file md5 and sha1
	testfile := "123.txt"
	infile, err := os.Open(testfile)
	if err != nil {
		//md5
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%x %s\n", md5h.Sum([]byte("")), testfile)

		//sha1
		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x %s\n", sha1.Sum([]byte("")), testfile)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
