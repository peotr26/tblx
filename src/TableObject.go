package tablx

import (
	"fmt"
)

////////////////////////////
// Object for the table. //
//////////////////////////

// Struct of a table.
type Table struct {
	data []Column
}

// Method that returns the number of columns in the table.
func (t Table) NumberCol() int {
	return len(t.data)
}

// Method that returns the maximum depth of the table.
func (t Table) MaxDepthCol() int {
	tmp := []int{}
	for _, e := range t.data {
		tmp = append(tmp, e.Depth())
	}
	return max(tmp)
}

/////////////////////////////
// Object for the column. //
///////////////////////////

// Struct of a column of the table.
type Column struct {
	data []Cell
}

// Method that returns the depth of the column.
func (c Column) Depth() int {
	return len(c.data)
}

// Method that prints out all the cells of the column.
func (c Column) PrintCol() {
	for _, e := range c.data {
		e.PrintCell()
	}
}

///////////////////////////
// Object for the cell. //
/////////////////////////

// Struct of a cell of a column.
type Cell struct {
	data string
}

// Method that returns the length of the data contained in the cell.
func (c Cell) Length() int {
	return len(c.data)
}

// Method that prints out the cell.
func (c Cell) PrintCell() {
	fmt.Println(c.data)
}

////////////////
// Functions //
//////////////

// Returns the maximum value in an array of integer.
func max(x []int) int {
	maxi := x[0]
	for _, e := range x {
		if e > maxi {
			maxi = e
		}
	}
	return maxi
}
