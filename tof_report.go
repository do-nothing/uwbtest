package main

import (
	"fmt"
	"github.com/bastengao/bytesparser"
	"strconv"
	"strings"
)

type tofData struct {
	data  [128]byte
	index int
}

type tofInfo struct {
	Head0        byte   `byte:"len:1,equal:0x6D"`
	Data_MID     []byte `byte:"len:2"`
	Data_MASK    []byte `byte:"len:3"`
	Data_RANGE0  []byte `byte:"len:9"`
	Data_RANGE1  []byte `byte:"len:9"`
	Data_RANGE2  []byte `byte:"len:9"`
	Data_RANGE3  []byte `byte:"len:9"`
	Data_NRANGES []byte `byte:"len:5"`
	Data_RSEQ    []byte `byte:"len:3"`
	Data_DEBUG   []byte `byte:"len:9"`
	Data_TA      []byte `byte:"len:4"`
	Tail         []byte `byte:"len:2,equal:0x0D0A"`
}

func (t *tofData) setData(data byte) {
	//fmt.Printf("TofData[%d]=%d\n", t.index, data);
	t.data[t.index] = data
	t.index++

	if data == 10 {
		if t.index == 65 {
			//fmt.Printf("%s\t", t.data[:t.index-2])
			parse(t.data[:t.index])
		}
		t.index = 0
	}
}

func parse(data []byte) {
	//fmt.Printf("%X\n", data)
	packet := tofInfo{}
	offset, err := bytesparser.Parse(data, &packet)
	if err != nil {
		fmt.Println(err)
	}
	if offset > 0 {
		refine := refineTof(packet)
		//fmt.Printf("%v\n", refine)
		if refine.mid == "mc" {
			ch <- refine
		}
	}
}

func refineTof(info tofInfo) tofRefine {
	var refine tofRefine

	refine.mid = "m" + strings.TrimSpace(string(info.Data_MID))
	refine.mask = strings.TrimSpace(string(info.Data_MASK))
	refine.range0 = float32(getIntFromBytes(info.Data_RANGE0)) / 1000
	refine.range1 = float32(getIntFromBytes(info.Data_RANGE1)) / 1000
	refine.range2 = float32(getIntFromBytes(info.Data_RANGE2)) / 1000
	refine.range3 = float32(getIntFromBytes(info.Data_RANGE3)) / 1000
	refine.nranges = getIntFromBytes(info.Data_NRANGES)
	refine.rseq = getIntFromBytes(info.Data_RSEQ)
	refine.delay = string(info.Data_DEBUG)
	refine.ta = string(info.Data_TA)

	return refine
}

func getIntFromBytes(bytes []byte) int {
	str := strings.TrimSpace(string(bytes))
	if i, err := strconv.ParseUint(str, 16, 32); err == nil {
		return int(i)
	} else {
		return -1
	}
}
