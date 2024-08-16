package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/praserx/ipconv"
)

// var mp = make(map[string]int, 10000000000)
var slc = make([]uint32, 100)

// uyuyuyuyuyææ

func main() {
	unique := 0
	t := time.Now()
	scan() //116971867 470.900543541 // just read 2.211524667
	sort.Slice(slc, func(i, j int) bool {
		return slc[i] < slc[j]
	})
	l := len(slc)
	for i := 1; i < l; i++ {
		if slc[i] != slc[i-1] {
			unique++
		}
	}
	fmt.Println(len(slc), cap(slc), time.Since(t).Seconds())
	fmt.Println(unique)

}

func scan() {
	file, _ := os.Open("./ip.txt")

	fileScanner := bufio.NewScanner(file)

	// read line by line
	it := 0
	t := time.Now()
	for fileScanner.Scan() {
		if it%10000000 == 0 {
			fmt.Println(it, time.Since(t).Seconds())
			t = time.Now()
		}
		it++
		// mp[fileScanner.Text()]++
		// a := net.ParseIP(fileScanner.Text())
		ip, version, err := ipconv.ParseIP(fileScanner.Text())
		if err != nil && version == 4 {
			fmt.Println(ipconv.IPv4ToInt(ip))
		}
		intIP, _ := ipconv.IPv4ToInt(ip)
		slc = append(slc, intIP)
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

}

// func gen() {
// 	for {
// 		fmt.Printf("%d.%d.%d.%d\n", rand.IntN(255), rand.IntN(255), rand.IntN(255), rand.IntN(255))
// 	}
// }
