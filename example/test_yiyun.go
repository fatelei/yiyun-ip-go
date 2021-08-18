package main

import (
	"flag"
	"fmt"
	yiyun "github.com/fatelei/yiyun-ip-go/pkg"
)

func main() {
	var appCode string
	var ip string
	flag.StringVar(&appCode, "app_code", "", "app_code")
	flag.StringVar(&ip, "ip", "", "ip")
	flag.Parse()
	if len(appCode) == 0 {
		fmt.Println("app_code is required")
		return
	}

	if len(ip) == 0 {
		fmt.Println("ip is required")
		return
	}

	cli := yiyun.NewYiYunClient(appCode)
	data, err := cli.GetLocationByIP(ip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
