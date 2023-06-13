// Package asn1 implements parsing of DER-encoded ASN.1 data structures, as defined in ITU-T Rec X.690.

package main

import (
	"encoding/asn1"
	"fmt"
	"math/big"
)

type MyStruct struct {
	BIGINT string
}

func exploreASN1() {
	bi := big.NewInt(292929292929292929)
	for i := 0; i < 10; i++ {
		bi = bi.Mul(bi, bi)
	}
	ms := MyStruct{BIGINT: bi.String()}
	// fmt.Println(ms)
	msb, err := asn1.Marshal(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("msb: %x", msb)
}

func main() {
	exploreASN1()
}
