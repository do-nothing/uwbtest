package main

import (
	"container/list"
	"math"
	"fmt"
)

type filter struct {
	r         int
	template  []float64
	valueList *list.List
}

func (f *filter) setR(r int) {
	f.r = r
	f.template = make([]float64, 2*r-1)

	var sum float64
	for i, _ := range f.template {
		i = i + 1
		f.template[i-1] = math.Exp(float64(-math.Pow(float64((i - f.r)), 2))/2) / math.Sqrt(2*math.Pi)
		sum += f.template[i-1]
	}

	var testSum float64
	for i, _ := range f.template {
		f.template[i] /= sum
		fmt.Println(f.template[i])
		testSum += f.template[i]
	}
	fmt.Println(testSum)
	f.valueList = list.New()
}

func (f *filter) filterValue(value [4]float64) [4]float64 {
	l := 2*f.r - 1
	f.valueList.PushBack(value)
	if (f.valueList.Len() > l) {
		f.valueList.Remove(f.valueList.Front())
	}

	index := 0
	var result [4]float64
	for e := f.valueList.Front(); e != nil; e = e.Next() {
		for i := 0; i < 4; i++ {
			if (value[i] != 0) {
				tempV := e.Value.([4]float64)[i] - 0.61
				result[i] += f.template[index] * tempV
				if(result[i] < 0){
					result[i] = 0
				}
			}
		}
		index ++
	}
	//fmt.Println(f.valueList.Len())

	return result
}
