package main

import (
	"flag"
	"github.com/larspensjo/config"
	"github.com/tarm/goserial"
	"log"
	"os"
)

var (
	conFile = flag.String("configfile", "/config.ini", "config file")
)

func main() {
	//获取当前路径
	file, _ := os.Getwd()
	cfg, err := config.ReadDefault(file + *conFile)

	//获取配置文件中的配置项
	id, err := cfg.String("COM", "COMID")
	if err != nil {
		log.Fatal(err)
	}

	//设置串口编号
	c := &serial.Config{Name: id, Baud: 9600}
	//打开串口
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	go startSender()

	buf := make([]byte, 128)
	var data tofData
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		for _, d := range buf[:n] {
			data.setData(d)
		}
	}
}
