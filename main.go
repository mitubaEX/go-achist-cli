package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"strconv"
	"os"
)

type RateData struct {
	performance int
	newRate int
	oldRate int
	subset int
}

func Create(performance int, newRate int, subset int) *RateData {
	newRateData := new(RateData)
	newRateData.performance = performance
	newRateData.newRate = newRate
	newRateData.oldRate = newRate - subset
	newRateData.subset = subset
	return newRateData
}

func GetRateSlice(doc *goquery.Document) []int{
	var rateSlice = []int{}
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		if i >= 3 && i <= 5{

			// oh my god
			replaceStr := strings.Replace(strings.Replace(s.Text(), "\t", "", -1), "\n", "", -1)

			num, err := strconv.Atoi(replaceStr)
			if err != nil {
				log.Fatal(err)
			}
			rateSlice = append(rateSlice, num)
		}
	})
	return rateSlice
}

func GetDoc(url string) *goquery.Document{
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func PrintRate(rateData *RateData) {
	var transition = "Highest"
	if rateData.subset < 0 {
		transition = "Lowest"
	}
	fmt.Printf("%d->%d(%d) %s\nPerformance %d\n",
		rateData.oldRate, rateData.newRate, rateData.subset, transition, rateData.performance)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ユーザ名を入力してちょ")
	} else {
		url := fmt.Sprintf("http://atcoder.jp/user/%s/history", os.Args[1])
		rateSlise := GetRateSlice(GetDoc(url))
		if len(rateSlise) <= 0 {
			fmt.Println("そのユーザ名は知らん")
			os.Exit(0)
		}
		rateData := Create(rateSlise[0], rateSlise[1], rateSlise[2])
		PrintRate(rateData)
	}
}
