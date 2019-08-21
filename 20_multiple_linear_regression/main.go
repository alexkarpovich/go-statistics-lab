package main

import (
	"fmt"
	"math"
	"os"

	"github.com/alexkarpovich/go-statistics-lab/library"
)

func Round(val float64, roundOn float64, places int) float64 {

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)

	var round float64
	if val > 0 {
		if div >= roundOn {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div >= roundOn {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}

	return round / pow
}

func main() {
	var rn, cn, rtn int

	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(f, "%d %d", &cn, &rn)
	cn++

	x := make([]float64, rn*cn)
	y := make([]float64, rn)

	for i := 0; i < rn; i++ {
		for j := 0; j < cn; j++ {
			if j == 0 {
				x[i*cn+j] = 1
			} else {
				fmt.Fscan(f, &x[i*cn+j])
			}
		}
		fmt.Fscan(f, &y[i])
	}

	fmt.Fscanf(f, "%d", &rtn)
	tx := make([]float64, rtn*cn)

	for i := 0; i < rtn; i++ {
		for j := 0; j < cn; j++ {
			if j == 0 {
				tx[i*cn+j] = 1
			} else {
				fmt.Fscan(f, &tx[i*cn+j])
			}
		}
	}

	ty := library.ComputeMultRegressionOutcome(cn, x, y, tx)
	for i := 0; i < rtn; i++ {
		fmt.Printf("%.2f\n", Round(ty[i], .2, 2))
	}
}
