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
	ad := Ad{&a}
	bd := Bd{a}  // 做了复制, 因此不是原来的a咯
	fmt.Println(ad.AppName == &bd.AppName)
	fmt.Println(*(ad.AppName) == bd.AppName)
	fmt.Printf("%v\n%v\n%v\n", &a, ad.AppName, &bd.AppName)
}
