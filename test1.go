//
//var (
//[]output string = [
//	"i-af0yxmfe",
//	"i-ijtunhbk",
//	"i-k2qyrnjz",
//	"i-362pv0q9",
//	"i-o5pmakm7",
//	"i-zvmlma11"
//]
//)
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Prints:
// <nil> [[0 1] [0 2] [0 3]]
// 0 0 0
// 0 1 1
// 1 0 0
// 1 1 2
// 2 0 0
// 2 1 3

func main() {
	jsonstring := "[\"i-af0yxmfe\",\"i-ijtunhbk\"]"
	var listoflists []string
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	err := dec.Decode(&listoflists)
	fmt.Println(err, listoflists)
	for i, list := range listoflists {
		fmt.Println(i, list)
	}

	total_cpu := float64(4000)
	total_cpu = total_cpu / 1000
	total_mem := float32(8266137600)
	used_mem := float32(6266150912)
	cpu_rate := float64(379)

	fmt.Println(total_mem / 1024 / 1024 / 1024)
	fmt.Println(used_mem / 1024 / 1024 / 1024)
	fmt.Println(fmt.Sprintf("%.1f", 100*used_mem/total_mem))
	fmt.Println(strconv.Atoi("8266137600"))
	fmt.Println(fmt.Sprint("%.1f", total_cpu))
	fmt.Println(string(fmt.Sprintf("%.3f", cpu_rate/1000)))

	var data []string
	data = append(data, "a")
	data = append(data, "b")
	data = append(data, "c")
	fmt.Println(cap(data))
	fmt.Println(len(data))

	//q := [3]int{1, 2, 3}
	//status := true
	//for i, num := range q {
	//	if num > 1 {
	//		fmt.Println("It's " + string(i))
	//		status = false
	//		break
	//	}
	//	if num > 2 {
	//		fmt.Println("It's " + string(i))
	//		status = false
	//		break
	//	}
	//}
	//
	//fmt.Println(status)
	rate := "355"
	ratefloat, _ := strconv.ParseFloat(rate, 64)
	fmt.Println(strings.Count(rate, "") - 1)
	ratebase := math.Pow10(strings.Count(rate, "") - 1)
	fmt.Println(ratefloat / ratebase)

	q := [3]string{"a", "b", "c"}
	var total_count int
	for i, num := range q {
		fmt.Println(i)
		fmt.Println(num)
		total_count = i
	}
	fmt.Println("---------")
	fmt.Println(total_count + 1)
}
