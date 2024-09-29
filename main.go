package main

import (
	"br/bitset"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/praserx/ipconv"
)

var mem runtime.MemStats = runtime.MemStats{}
var totalIPNumber = int64(math.Pow(2, 32)) //4,294,967,296
var slcLen = int(totalIPNumber / 32)       //134,217,728
var slc = make(bitset.BitSet, slcLen)

func main() {
	t := time.Now()
	unique := scan()
	runtime.GC()
	fmt.Println(unique, time.Since(t).Seconds())
	runtime.ReadMemStats(&mem)
	fmt.Printf("Total allocated memory (in bytes): %d\n", mem.Alloc)
}

func scan() int64 {
	unique := int64(0)
	file, _ := os.Open("./ips.txt")
	fileScanner := bufio.NewScanner(file)

	// read line by line
	it := 0
	t := time.Now()
	for fileScanner.Scan() {
		//Just to track the proggress, prints progress for each 10 million lines
		if it%10000000 == 0 {
			runtime.ReadMemStats(&mem)
			fmt.Println(it, time.Since(t).Seconds())
			fmt.Printf("Total allocated memory (in bytes): %d\n", mem.Alloc)
			// fmt.Printf("Number of garbage collections: %d\n", mem.NumGC)
			t = time.Now()
		}
		it++

		ip, version, err := ipconv.ParseIP(fileScanner.Text())
		if err != nil && version == 4 {
			fmt.Println(ipconv.IPv4ToInt(ip))
		}
		intIP, _ := ipconv.IPv4ToInt(ip)
		if !slc.IsSet(intIP) {
			slc.Set(intIP)
			unique++
		}
		if unique == totalIPNumber {
			return unique
		}
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()
	return unique
}
