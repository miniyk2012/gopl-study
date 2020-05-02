package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var sumType = flag.String("checkSum", "SHA256", "加密")
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	flag.Parse()
	var c []byte
	switch *sumType {
	case "SHA256":
		temp := sha256.Sum256(b)
		c = temp[:]
	case "SHA384":
		temp := sha512.Sum384(b)
		c = temp[:]
	case "SHA512":
		temp := sha512.Sum512(b)
		c = temp[:]
	default:
		log.Fatal("Unexpected width specified.")
	}
	fmt.Printf("%x\n", c)
}
