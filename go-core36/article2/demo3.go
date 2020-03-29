package main

import (
	"flag"
	"fmt"
	"strings"
	// 需在此处添加代码。[1]
	"os"
)

var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

type customFlag struct {
	name string
	hobby string
}

//实现Set接口,Set接口决定了如何解析flag的值
func (c *customFlag) Set(str string) error {
	str = strings.TrimSpace(str)
	s := strings.Split(str, ",")
	c.name = s[0]
	c.hobby = s[1]
	return nil
}

//实现String接口
func (c *customFlag) String() string {
	return fmt.Sprintf("%v", *c)
}

var flagvar customFlag

func init() {
	// 需在此处添加代码。[2]
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.Var(&flagvar, "flag", "help message for flagname")
	// flag.StringVar(&name, "name", "every one", "The greeting object.")
}

func main() {
	// 需在此处添加代码。[3]
	// flag.Usage = func() { 
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question") 
	// 	flag.PrintDefaults()
	// }
	// flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("%v\n", flagvar)
}

// go run demo3.go -flag yangkai,play
