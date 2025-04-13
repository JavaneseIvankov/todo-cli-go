package cli

import (
	"fmt"
	"strings"
)

type Table struct {
	Rows [][]string
	Headers []string
	Padding int
}

func CreateTable(padding int) *Table {
	return &Table{Padding: padding}
}

func (t *Table) SetHeaders(headers []string) {
	t.Headers = headers
}

func (t *Table) Append(row []string) {
	t.Rows = append(t.Rows, row)
}

func (t *Table) Render() {
    // Calculate column widths
    colWidths := make([]int, len(t.Headers))
    for i, header := range t.Headers {
        colWidths[i] = len(header)
    }
    for _, row := range t.Rows {
        for i, cell := range row {
            if len(cell) > colWidths[i] {
                colWidths[i] = len(cell)
            }
        }
    }

    t.printRow(t.Headers, colWidths)
    t.printSeparator(colWidths)

    for _, row := range t.Rows {
        t.printRow(row, colWidths)
    }

}

func (t *Table) printRow(row []string, colWidths []int) {
    for i, cell := range row {
        fmt.Printf("%-*s", colWidths[i]+t.Padding, cell)
    }
    fmt.Println()
}


func (t *Table) printSeparator(colWidths []int) {
    for _, width := range colWidths {
        fmt.Print(strings.Repeat("-", width+t.Padding))
    }
    fmt.Println()
}