/*
	非常好的代码解读: https://www.cnblogs.com/yjf512/p/12121345.html
 */
package main

import "gopkg.in/macaron.v1"

import "github.com/go-macaron/inject"

import "fmt"

import "reflect"

type A struct {
	Name string
}

type B struct {
	Name string
}

func (b *B) GetName() string {
	return b.Name
}

type I interface {
	GetName() string
}

type C struct {
	AStruct A `inject`
	BStruct B `inject`
}

type MyFastInvoker func(arg1 A, arg2 I, arg3 string)

func (invoker MyFastInvoker) Invoke(args []interface{}) ([]reflect.Value, error) {
	if a, ok := args[0].(A); ok {
		fmt.Println(a.Name)
	}

	if b, ok := args[1].(I); ok {
		fmt.Println(b.GetName())
	}
	if c, ok := args[2].(string); ok {
		fmt.Println(c)
	}
	//return "hello"
	return []reflect.Value{reflect.ValueOf("a")}, nil
}

func Hello(arg1 A) string {
	fmt.Println(arg1.Name)
	return "why"
}

func main() {
	InjectDemo()

	a := &A{Name: "inject name"}
	m := macaron.Classic()
	m.Map(a)
	m.Get("/", func(a *A) string {
		return "Hello world!" + a.Name
	})
	m.Run()
}

func InjectDemo() {
	a := A{Name: "a name"}
	inject1 := inject.New()
	inject1.Map(a)
	inject1.MapTo(&B{Name: "b name"}, (*I)(nil))
	inject1.Set(reflect.TypeOf("string"), reflect.ValueOf("c name"))
	inject1.Invoke(func(arg1 A, arg2 I, arg3 string) {
		fmt.Println(arg1.Name)
		fmt.Println(arg2.GetName())
		fmt.Println(arg3)
	})
	fmt.Println()
	c := C{}
	inject1.Apply(&c)
	fmt.Println(c.AStruct.Name)
	fmt.Println()
	inject2 := inject.New()
	inject2.Map(a)
	inject2.MapTo(&B{Name: "b name"}, (*I)(nil))
	inject2.Set(reflect.TypeOf("string"), reflect.ValueOf("c name"))
	ret, _ := inject2.Invoke(MyFastInvoker(nil))  // 这里只是一个强转罢了, 把nil转换为MyFastInvoker
	fmt.Println(ret)
	
	fmt.Println()
	ret, _ = inject2.Invoke(Hello)
	fmt.Println(ret)
}
