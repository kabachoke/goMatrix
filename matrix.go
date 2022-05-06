package matrix

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
)

// функция транспонирования матрицы
func T(a [][]float64) [][]float64 {
	m := MakeMatrix(len(a[0]), len(a), 0)
	for i, row := range a {
		for j := range row {
			m[j][i] = a[i][j]
		}
	}
	return m
}

// функция умножения матриц
func Dot(a, b [][]float64) ([][]float64, error) {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_col_count == b_row_count {
		m := MakeMatrix(a_row_count, b_col_count, 0)

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
func fake_Dot(a, b [][]float64) ([][]float64, error) {
	a_row_count, a_col_count := len(a), len(a[0])
	b_row_count, b_col_count := len(b), len(b[0])

	if a_col_count == b_row_count {
		m := MakeMatrix(a_row_count, b_col_count, 0)
		var ch []chan string
		for i := range m {
			ch = append(ch, make(chan string, 1))
			go dot_threading(a, b, m, i, ch[i])
		}
		for i := range m {
			<-ch[i]
		}
		return m, nil
	} else {
		msg := fmt.Sprintf("Can't multiply matrices, "+
			"because a_col_count(%d) != b_row_count(%d)", a_col_count, b_row_count)
		return nil, errors.New(msg)
	}
}
func dot_threading(a, b, m [][]float64, rowIndex int, channel chan string) {
	for j := range m[rowIndex] {
		m[rowIndex][j] = sum(a, b, rowIndex, j)
	}
	channel <- ""
}
func Subtract(a, b [][]float64) [][]float64 {
	if len(a) == len(b) && len(a[0]) == len(b[0]) {
		res := MakeMatrix(len(a), len(a[0]), 0)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				res[i][j] = a[i][j] - b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("Can't subtract matrixes, because len(a) != len(b)"))
		return nil
	}
}
func SubtractFromConst(c float64, a [][]float64) [][]float64 {
	res := MakeMatrix(len(a), len(a[0]), 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			res[i][j] = c - a[i][j]
		}
	}
	return res
}
func Sum(a, b [][]float64) [][]float64 {
	if len(a) == len(b) && len(a[0]) == len(b[0]) {
		res := MakeMatrix(len(a), len(a[0]), 0)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				res[i][j] = a[i][j] + b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("Can't sum matrixes, because len(a) != len(b)"))
		return nil
	}
}
func Multiply(a, b [][]float64) [][]float64 {
	if len(a) == len(b) && len(a[0]) == len(b[0]) {
		res := MakeMatrix(len(a), len(a[0]), 0)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				res[i][j] = a[i][j] * b[i][j]
			}
		}
		return res
	} else {
		log.Fatal(errors.New("Can't multiply matrixes, because len(a) != len(b)"))
		return nil
	}
}
func MultiplyOnConst(a [][]float64, c float64) [][]float64 {
	res := MakeMatrix(len(a), len(a[0]), 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			res[i][j] = a[i][j] * c
		}
	}
	return res
}

// функция вывода матрицы в консоль
func PrintMatrix(m [][]float64) {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println()
}

// функция, создающая и возвращающая матрицу
func MakeMatrix(n, m int, initvalue float64) [][]float64 {
	matrix := make([][]float64, n)
	rows := make([]float64, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = initvalue
		}
	}
	return matrix
}

// функция сигмоиды
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
func sum(a, b [][]float64, a_row, b_col int) (sum float64) {
	for i := 0; i < len(b); i++ {
		h := a[a_row][i] * b[i][b_col]
		sum += h
	}
	return sum
}

// функция заполнения матрицы случайными значениями
func FillMatrix(m [][]float64) [][]float64 {
	for i, row := range m {
		for j := range row {
			m[i][j] = rand.Float64()*0.2 - 0.1
		}
	}
	return m
}

// функция нормального распределения
func normalDistribution(mu, sigma, x float64) float64 {
	return math.Exp(-math.Pow(x-mu, 2)/(2*math.Pow(sigma, 2))) / (sigma * math.Pow(2*math.Pi, 0.5))
}
