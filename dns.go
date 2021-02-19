package main

import (
	"fmt"
	"net"
)
 
func main() {
	cname, err := net.LookupCNAME("joshbrookshire.com")
	if err != nil {
		panic(err)
	}
        // dig +short research.swtch.com cname
	fmt.Printf("%s\n", cname)
 
//ghs.google.com.
 
	txts, err := net.LookupTXT("redcross.nl")
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		//dig +short gmail.com txt
		fmt.Printf("%s\n", txt)
	}
 
// v=spf1 redirect=_spf.google.com
 
}


/* 	Resources:
	https://golang.org/pkg/net/#pkg-variables
	http://networkbit.ch/golang-dns-loo
*/


