package main

import "fmt"

type Ad struct {
	AppName *string
}

type Bd struct {
	AppName string
}

func main() {
	a := "abc"
	b := "abc"
	ad := Ad{&a}
	bd := Bd{b}
	fmt.Println(ad.AppName == &bd.AppName)
	fmt.Println(*(ad.AppName) == bd.AppName)
}
