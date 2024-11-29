package main

import (
	"fmt"
	"sync"
	"time"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type Cat struct {
	Name   string
	Age    int64
	Weight float64
}

type CatCoef struct {
	Name         string
	Coef         float64
	TagGoroutine int
}

var cats = []Cat{
	{"Vaska", 9, 5.04},
	{"Sofa", 10, 1.75},
	{"Vasilisa", 5, 2.05},
	{"Khudochavii", 2, 1.05},
	{"Tolstovatii", 3, 4.05},
	{"Krasivii", 3, 3.25},
	{"Objorniy", 3, 7.25},
}

var once = sync.Once{}

func HandleExtraRecords(extraRecords int) []Cat {
	Printfln("HandleExtraRecords works")
	maxRow := len(cats)
	return cats[maxRow-extraRecords : maxRow]
}

func devideCollectionOnBuckets(bucketNumber int, amountofGoroutine int, cats []Cat) []Cat {
	cntRecords := len(cats) / amountofGoroutine
	extraRecords := len(cats) % amountofGoroutine
	var extraCats []Cat

	var hashRes = make(map[int][]Cat)
	var l_maxRowInBucket = 0
	for i := 0; i < amountofGoroutine; i++ {
		l_minRowInBucket := l_maxRowInBucket
		l_maxRowInBucket = cntRecords * (i + 1)
		if l_maxRowInBucket > len(cats)+extraRecords {
			l_maxRowInBucket = len(cats) + extraRecords
			// hashRes[i] = []int{l_minRowInBucket,l_maxRowInBucket}
			hashRes[i] = cats[l_minRowInBucket:l_maxRowInBucket]
			break
		}
		hashRes[i] = cats[l_minRowInBucket:l_maxRowInBucket]
	}
	if extraRecords > 0 {
		once.Do(func() {
			extraCats = HandleExtraRecords(extraRecords)
		})
	}

	if resCat, ok := hashRes[bucketNumber]; !ok {
		fmt.Errorf("There's no such bucket. Probably you specified buchet and total threads not correctly.")
		return nil
	} else {
		if len(extraCats) > 0 {
			resCat = append(resCat, extraCats...)
		}
		return resCat
	}
}

func CalcAgeWeightCoef(chCoefCat chan CatCoef, numgoroutine int, totalgoroutines int) {
	Printfln("Goroutine - %v started", numgoroutine)

	resCats := devideCollectionOnBuckets(numgoroutine, totalgoroutines, cats)
	for _, cat := range resCats {
		Printfln("Goroutine - %v processes with cat=%v", numgoroutine, cat.Name)
		chCoefCat <- CatCoef{cat.Name, float64(cat.Age) * cat.Weight, numgoroutine}
	}
	Printfln("Goroutine - %v finished", numgoroutine)
}

func main() {
	var chCoefCat chan CatCoef = make(chan CatCoef, 2)
	var totalGoRoutines = 2

	for i := 0; i < 2; i++ {
		go CalcAgeWeightCoef(chCoefCat, i, totalGoRoutines)
	}

	for i := 0; i < len(cats); i++ {
		time.Sleep(time.Second * 1)
		Printfln("Weird coef = %v", <-chCoefCat)
	}
}
