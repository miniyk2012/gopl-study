package main

import "fmt"


type Data struct {
	Value int
}
func main() {
	m := map[string]int {
		"1": 100,
	}

	x, y := m["10"]
	fmt.Println(x, y)

	x, y = m["1"]
	fmt.Println(x, y)

	x= m["0"]
	fmt.Println(x)

	datas := []Data{Data {10}, Data{10}}
	for _, data := range datas {
		data.Value += 10
	}
	for _, data := range datas {
		fmt.Println(data)
	}

	for i, _ := range datas {
		datas[i].Value += 10
	}
	for _, data := range datas {
		fmt.Println(data)
	}
}

