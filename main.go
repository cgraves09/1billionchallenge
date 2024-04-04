package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Service struct{}

func main() {
	Run()
}

type CityTemperature struct {
	min   int64
	max   int64
	count int
	sum   int64
	avg   int
}

func Run() error {
	f, err := os.Open("./data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	start := time.Now()

	data, err := scanFile(f)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, v := range data {
		v.avg = int(v.sum) / v.count
	}

	end := time.Now()

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

	return nil
}

func scanFile(f *os.File) (map[string]CityTemperature, error) {
	data := make(map[string]CityTemperature)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ";")

		rInt, err := strconv.ParseFloat(text[1], 64)
		if err != nil {
			return nil, err
		}

		city := text[0]
		temp := int64(math.Round(rInt))
		if t, ok := data[city]; ok {
			if temp > t.max {
				t.max = temp
			}

			if temp < t.min {
				t.min = temp
			}

			t.count++
			t.sum += temp
			continue
		}

		data[city] = CityTemperature{
			max:   temp,
			min:   temp,
			count: 1,
			sum:   temp,
		}
	}

	return data, nil
}
