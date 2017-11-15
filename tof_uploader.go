package main

import (
	"fmt"
	"log"
	"net"
)

type tofRefine struct {
	mid     string
	mask    string
	range0  float32
	range1  float32
	range2  float32
	range3  float32
	nranges int
	rseq    int
	delay   string
	ta      string
}

var ch = make(chan tofRefine)

func startSender() {
	conn, err := net.Dial("udp", "121.42.196.133:5555")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	json := "{\"id\":\"monitor\",\"target\":\"indoordemo\",\"logType\":\"monitor control\"," +
		"\"strategy\":\"trilateration\",\"quality\":0,\"timestamp\":\"0\",\"contentBean\":" +
		"{\"command\":\"setDistances\",\"args\":[%f, %f, %f, %f]}}"

	var fi filter
	fi.setR(3)
	for {
		tof := <-ch

		var vgroup [4]float64
		vgroup[0] = float64(tof.range0)
		vgroup[1] = float64(tof.range1)
		vgroup[2] = float64(tof.range2)
		vgroup[3] = float64(tof.range3)
		vgroup = fi.filterValue(vgroup)
		str := fmt.Sprintf(json, vgroup[0], vgroup[1], vgroup[2], vgroup[3])
		conn.Write([]byte(str))
		//fmt.Println(str)
		log.Println(str)
	}
}
