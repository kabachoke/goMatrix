package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func main() {
	m1 := MakeMatrix[float64](10000, 80)
	m2 := MakeMatrix[float64](80, 10000)

	fillMatrix(m1)
	fillMatrix(m2)

	start := time.Now()
	m, _ := Dot(m1, m2)
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(m[0][0])
}

// функция транспонирования матрицы
func Transpose[T any](a [][]T) [][]T {
	m := MakeMatrix[T](len(a[0]), len(a))
	for i, row := range a {
		for j := range row {
			m[j][i] = a[i][j]
		}
	}
	return m
}

// функция умножения матриц
func Dot[T Number](a, b [][]T) ([][]T, error) {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_col_count == b_row_count {
		m := MakeMatrix[T](a_row_count, b_col_count)
		for i, row := range m {
			for j := range row {
				m[i][j] = sum(a, b, i, j)
			}
		}
		return m, nil
	} else {
		msg := fmt.Sprintf("Can't multiply matrices, "+
			"because a_col_count(%d) != b_row_count(%d)", a_col_count, b_row_count)
		return nil, errors.New(msg)
	}
}

func Sum[T Number](a, b [][]T) [][]T {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_row_count == b_row_count && a_col_count == b_col_count {
		res := MakeMatrix[T](a_row_count, a_col_count)
		for i := 0; i < a_row_count; i++ {
			for j := 0; j < a_col_count; j++ {
				res[i][j] = a[i][j] + b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("can't sum matrixes, because len(a) != len(b)"))
		return nil
	}
}
func Subtract[T Number](a, b [][]T) [][]T {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_row_count == b_row_count && a_col_count == b_col_count {
		res := MakeMatrix[T](a_row_count, a_col_count)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				res[i][j] = a[i][j] - b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("can't subtract matrixes, because len(a) != len(b)"))
		return nil
	}
}
func SumConst[T Number](a [][]T, c T) [][]T {
	res := MakeMatrix[T](len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			res[i][j] = a[i][j] + c
		}
	}
	return res
}

func Multiply[T Number](a, b [][]T) [][]T {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_row_count == b_row_count && a_col_count == b_col_count {
		res := MakeMatrix[T](a_row_count, a_col_count)
		for i := 0; i < a_row_count; i++ {
			for j := 0; j < a_col_count; j++ {
				res[i][j] = a[i][j] * b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("can't multiply matrixes, because len(a) != len(b)"))
		return nil
	}
}
func MultiplyOnConst[T Number](a [][]T, c T) [][]T {
	res := MakeMatrix[T](len(a), len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			res[i][j] = a[i][j] * c
		}
	}
	return res
}

// функция вывода матрицы в консоль
func PrintMatrix[T any](m [][]T) {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println()
}

// функция, создающая и возвращающая матрицу
func MakeMatrix[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func sum[T Number](a, b [][]T, a_row, b_col int) (sum T) {
	for i := 0; i < len(b); i++ {
		h := a[a_row][i] * b[i][b_col]
		sum += h
	}
	return sum
}

// функция сигмоиды
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// функция нормального распределения
func normalDistribution(mu, sigma, x float64) float64 {
	return math.Exp(-math.Pow(x-mu, 2)/(2*math.Pow(sigma, 2))) / (sigma * math.Pow(2*math.Pi, 0.5))
}

func fillMatrix(a [][]float64) [][]float64 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			a[i][j] = rand.Float64()*2 - 1
		}
	}
	return a
}
