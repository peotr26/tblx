package tablx

import (
	"fmt"
)

////////////////////////////
// Object for the table. //
//////////////////////////

// A table.
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

func (t Table) NewCol() {
	t.data = append(t.data, Column{})
}

// Method that purges all empty cells in the table.
func (t Table) Purge() {
	for _, e := range t.data {
		e.Purge()
	}
}

/////////////////////////////
// Object for the column. //
///////////////////////////

// A column of the table.
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

// Method that creates a new cell in the column.
func (c Column) NewCell() {
	c.data = append(c.data, Cell{})
}

// Method that fills a column with data from an array of strings.
func (c Column) Fill(data_array []string) {
	for c.Depth() <= len(data_array) {
		c.NewCell()
	}
	for i, e := range data_array {
		c.data[i].Write(e)
	}
}

// Method that appends an array of strings into the column.
func (c Column) Append(data_array []string) {
	for _, e := range data_array {
		c.NewCell()
		c.data[c.Depth()].Write(e)
	}
}

// Method that purges all empty cells of the column.
func (c Column) Purge() {
	toBePurged := []int{}
	for i, e := range c.data {
		if e.data == "" {
			toBePurged = append(toBePurged, i)
		}
	}
	for i, e := range toBePurged {
		index := e-i
		c.data = append(c.data[:index], c.data[index+1:]...)
	}
}

///////////////////////////
// Object for the cell. //
/////////////////////////

// A cell of a column.
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

// Method that overwrites the data in the cell.
func (c Cell) Write(input string) {
	c.data = input
}

// Method that appends data in the cell.
func (c Cell) Append(input string) {
	c.data += input
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
