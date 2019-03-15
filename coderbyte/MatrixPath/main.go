package main

import "fmt"

type Step struct {
	matrix [][]string
	m      int
	n      int
}

func (c Step) Left() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.n = c.n - 1
	return c
}

func (c Step) Right() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.n = c.n + 1
	return c
}

func (c Step) Up() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.m = c.m - 1
	return c
}

func (c Step) Down() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.m = c.m + 1
	return c
}

func (c Step) Check() bool {

	matrix := c.matrix
	m := c.m
	n := c.n

	rows := len(matrix)
	cols := len(matrix[1])

	if m >= rows || m < 0 || n >= cols || n < 0 {
		return false
	}

	return c.matrix[m][n] == "1"
}

func (c Step) Mark() Step {
	c.matrix[c.m][c.n] = "x"
	return c
}

func (c Step) IsMarked() bool {
	return c.matrix[c.m][c.n] == "x"
}

func (c Step) IsFinal() bool {
	rows := len(c.matrix)
	cols := len(c.matrix[1])
	return c.m == rows-1 && c.n == cols-1
}

func CopyMatrix(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func MatrixPath(strArr []string) string {

	matrix := BuildMatrix(strArr)

	initStep := Step{
		matrix: matrix,
		m:      0,
		n:      0,
	}

	res := "not possible"

	if initStep.Check() {
		initStep.Mark()
		steps := FindPath([]Step{initStep})
		if len(steps) > 0 {
			return "true"
		}
	}

	count := 0
	for m := range matrix {
		for n := range matrix[m] {
			if matrix[m][n] == "0" {
				k := CopyMatrix(matrix)
				k[m][n] = "1"
				initStep := Step{
					matrix: k,
					m:      0,
					n:      0,
				}
				steps := FindPath([]Step{initStep})
				if len(steps) > 0 {
					count = count + 1
				}
			}

		}
	}

	if count > 0 {
		return fmt.Sprintf("%d", count)
	}

	return res

}

func BuildMatrix(strArr []string) [][]string {
	arr := [][]string{}
	for _, row := range strArr {
		strRow := string(row)
		intRow := []string{}
		for _, s := range strRow {
			intRow = append(intRow, string(s))
		}
		arr = append(arr, intRow)
	}
	return arr
}

func FindPath(steps []Step) []Step {
	mas := make([]Step, 0)
	f := false

	for _, step := range steps {
		if step.Check() {
			step.Mark()
		}

		if step.IsFinal() {
			mas = append(mas, step)
		} else {

			stepLeft := step.Left()
			if stepLeft.Check() {
				stepLeft.Mark()
				mas = append(mas, stepLeft)
			}

			stepUp := step.Up()
			if stepUp.Check() {
				stepUp.Mark()
				mas = append(mas, stepUp)
			}

			stepRight := step.Right()
			if stepRight.Check() {
				stepRight.Mark()
				mas = append(mas, stepRight)
			}

			stepDown := step.Down()
			if stepDown.Check() {
				stepDown.Mark()
				mas = append(mas, stepDown)
			}

			f = true

		}

	}

	if f == false {
		return steps
	}

	return FindPath(mas)
}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(MatrixPath(readline()))

}
