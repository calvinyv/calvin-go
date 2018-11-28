package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	re, err := ioutil.ReadFile("/Users/calvin/abc1234")
	if err != nil {
		s := []byte("1")
		ioutil.WriteFile("/Users/calvin/abc1234", s, os.ModePerm)
	}
	re, _ = ioutil.ReadFile("/Users/calvin/abc1234")

	fmt.Println("calvin---1" + string(re[:]))
	asInit, _ := strconv.Atoi(string(re[:]))
	asInit = asInit +1
	fmt.Println(strconv.Itoa(asInit))
	ioutil.WriteFile("/Users/calvin/abc1234", []byte(strconv.Itoa(asInit)), os.ModePerm)
	re, _ = ioutil.ReadFile("/Users/calvin/abc1234")
	fmt.Println("calvin---2" + string(re[:]))
	os.Remove("/Users/calvin/abcd1234")
}
