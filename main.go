package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/praserx/ipconv"
)

// var mp = make(map[string]int, 10000000000)
var slc = [10]map[uint32]bool{}

// uyuyuyuyuyææ

func main() {
	for i := 0; i < 10; i++ {
		slc[i] = make(map[uint32]bool, 10000000)
	}
	unique := 0
	t := time.Now()
	scan() //116971867 470.900543541 // just read 2.211524667
	for i := 0; i < 10; i++ {
		unique += len(slc[i])
	}
	fmt.Println(unique, time.Since(t).Seconds())
	// fmt.Println(unique)

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

		// octs := strings.Split(fileScanner.Text(), ".")
		// _, _ = strconv.Atoi(octs[0])

		ip, version, err := ipconv.ParseIP(fileScanner.Text())
		if err != nil && version == 4 {
			fmt.Println(ipconv.IPv4ToInt(ip))
		}
		intIP, _ := ipconv.IPv4ToInt(ip)
		k := intIP % 10
		slc[k][intIP] = true
		// slc = append(slc, intIP)

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
