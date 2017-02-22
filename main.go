package main

import (
	"fmt"
	ioutil "io/ioutil"
	"strconv"
	"strings"
)

func main() {
	PrintIps("192.168.1.", 3, 20)
}

func PrintIps(prefix string, start, end int) {
	IPs := []string{""}
	fullFileName := "./ips.txt"
	for i := start; i <= end; i++ {
		IPs = append(IPs, prefix)
		IPs = append(IPs, strconv.Itoa(i)+"\n")
	}
	IPsStr := strings.Join(IPs, "")
	fmt.Printf("%v", IPsStr)
	data := []byte(IPsStr)
	ioutil.WriteFile(fullFileName, data, 644)

}
