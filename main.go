package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {

	//if len(os.Args) > 0 {
	//	for index, val := range os.Args {
	//
	//		fmt.Printf("arg[%d] is %v \n", index, val)
	//
	//	}
	//}

	////定义命令行参数方式1
	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	//
	////解析命令行参数
	//flag.Parse()
	//fmt.Println(name, age, married, delay)
	////返回命令行参数后的其他参数
	//fmt.Println(flag.Args())
	////返回命令行参数后的其他参数个数
	//fmt.Println(flag.NArg())
	////返回使用的命令行参数个数
	//fmt.Println(flag.NFlag())

	var ipn string
	var address string

	flag.StringVar(&ipn, "ip", "-ip", "ip tag")
	flag.StringVar(&address, "address", "127.0.0.1", "address")

	fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.Func("ip", "`IP address` to parse", func(s string) error {
		ip = net.ParseIP(s)
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	})
	err := fs.Parse([]string{ipn, address})
	if err != nil {
		return
	}
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

}
