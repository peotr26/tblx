package tblx

import (
	"fmt";
	"unicode"
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

// Method that erases a lign from the table.
func (t Table) EraseLign(index int) {
	
}

/////////////////////////////
// Object for the column. //
///////////////////////////

// A column of the table.
// The table contains a header and a dataset.
type Column struct {
	header string
	data []Cell
}

// Method that returns the depth of the column without the header.
func (c Column) Depth() int {
	return len(c.data)
}

// Method that returns the maximum width of the column.
func (c Column) MaxWidth() int {
	tmp := []int{len(c.header)}
	for _, e := range c.data {
		tmp = append(tmp, e.Length())
	}
	return max(tmp)
}

// Method that prints out all the cells of the column.
func (c Column) PrintCol() {
	fmt.Println(c.header)
	for _, e := range c.data {
		e.PrintCell()
	}
}

// Method that creates a new cell in the column.
func (c Column) NewCell() {
	c.data = append(c.data, Cell{})
}

// Method that erases a cell in the column.
func (c Column) EraseCell(index int) {
	c.data = append(c.data[:index], c.data[index+1:]...)
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

// Method that returns the header of the column.
func (c Column) Header() string {
	return c.header
}

// Method that renames the header of the column.
func (c Column) RenameHeader(header string) {
	c.header = header
}

// Method that returns if all cells of the column are floats.
func (c Column) IsFloat() bool {
	for _, e := range c.data {
		if e.IsFloat() == false {
			return false
		}
	}
	return true
}

// Method that returns if all cells of the column are integers.
func (c Column) IsInt() bool {
	for _, e := range c.data {
		if e.IsInt() == false {
			return false
		}
	}
	return true
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

// Method that returns if the data of the cell is a float.
func (c Cell) IsFloat() bool {
	nb_dot := 0
	for _, e := range c.data {
		if unicode.IsNumber(e) == false{
			if e != '.' {
				return false
			} else if e == '.' {
				nb_dot++
			}
		} 
	}
	if nb_dot != 1 {
		return false
	}
	return true
}

// Method that returns if the data of the cell is an integer.
func (c Cell) IsInt() bool {
	for _, e := range c.data {
		if unicode.IsDigit(e) == false {
			return false
		}
	}
	return true
}
