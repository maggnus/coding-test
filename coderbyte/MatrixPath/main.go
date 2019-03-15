package main

import "fmt"

// Step represent intermediate state
//   matrix - 2d array of string elements [0, 1, x], where x is indicator of visited point
//   m - current row position
//   n - current column position
type Step struct {
	matrix [][]string
	m      int
	n      int
}

// Left step
func (c Step) Left() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.n = c.n - 1
	return c
}

// Right step
func (c Step) Right() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.n = c.n + 1
	return c
}

// Up step
func (c Step) Up() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.m = c.m - 1
	return c
}

// Down step
func (c Step) Down() Step {
	matrix := CopyMatrix(c.matrix)
	c.matrix = matrix
	c.m = c.m + 1
	return c
}

// Check if current step position allowed and has "1" element
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

// Mark defined step as visited
func (c Step) Mark() Step {
	c.matrix[c.m][c.n] = "x"
	return c
}

// IsMarked check if step visited
func (c Step) IsMarked() bool {
	return c.matrix[c.m][c.n] == "x"
}

// IsFinal check if current position is final
func (c Step) IsFinal() bool {
	rows := len(c.matrix)
	cols := len(c.matrix[1])
	return c.m == rows-1 && c.n == cols-1
}

// CopyMatrix makes copy of 2d array
func CopyMatrix(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

// MatrixPath return string value
//   - "true" if path of 1's exists
//   - number of position where if a single 0 is replaced with a 1, a path of 1's will be created successfully
//   - "not possible" in all other cases
func MatrixPath(strArr []string) string {

	// Parse string array to matrix
	matrix := BuildMatrix(strArr)

	// Initial step with [0,0] position and parsed matrix
	initStep := Step{
		matrix: matrix,
		m:      0,
		n:      0,
	}

	// Default return
	res := "not possible"

	// Check if paths exist
	if initStep.Check() {
		initStep.Mark()
		steps := FindPath([]Step{initStep})
		if len(steps) > 0 {
			return "true"
		}
	}

	// Paths not exists let's try find all "0" replace them with "1" and check again
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
		return fmt.Sprintf("%d", count) // return count as string
	}

	return res

}

// BuildMatrix convert array of strings to matrix object
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

// FindPath return array of all possible steps with successful path of 1's
// Function recursively iterating input array of steps and filling up new one while
// checking if the step left, up, right or down is available.
func FindPath(steps []Step) []Step {
	// Create new empty steps array
	newSteps := make([]Step, 0)

	f := false // changes flag

	for _, step := range steps {
		if step.Check() {
			step.Mark()
		}

		if step.IsFinal() {
			newSteps = append(newSteps, step)
		} else {

			stepLeft := step.Left()
			if stepLeft.Check() {
				stepLeft.Mark()
				newSteps = append(newSteps, stepLeft)
			}

			stepUp := step.Up()
			if stepUp.Check() {
				stepUp.Mark()
				newSteps = append(newSteps, stepUp)
			}

			stepRight := step.Right()
			if stepRight.Check() {
				stepRight.Mark()
				newSteps = append(newSteps, stepRight)
			}

			stepDown := step.Down()
			if stepDown.Check() {
				stepDown.Mark()
				newSteps = append(newSteps, stepDown)
			}

			f = true

		}

	}

	if f == false { // no new changes, return input steps
		return steps
	}

	return FindPath(newSteps)
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(MatrixPath(readline()))

}
