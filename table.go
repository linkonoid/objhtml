package objhtml

import (
	"fmt"
	"strconv"
	//	"golang.org/x/net/html"
)

//=============================================
//  Table  //
//=============================================

type (
	//Table represents <table>
	Table struct {
		*Element
		Head     *TableHead
		Body     *TableBody
		Footer   *TableFooter
		Colgroup *TableColgroup
		Caption  *Element
		Rows     []*TableRow
	}
	//TableColgroup represents a <colgroup>
	TableColgroup struct {
		*Element
		Span int
		Cols []*TableCol
	}

	//TableCol represents a <col>
	TableCol struct {
		*Element
		Span  int
		Style string
	}

	//TableHead
	TableHead struct {
		*Element
		Rows []*TableRow
	}
	//TableBody represents
	TableBody struct {
		*Element
		Rows []*TableRow
	}
	//TableFooter
	TableFooter struct {
		*Element
		Rows []*TableRow
	}
	//TableRow represents a <tr>
	TableRow struct {
		*Element
		Cells []*TableCell
	}
	//TableCell represents a <td>
	TableCell struct {
		*Element
		Header  bool
		Colspan int
		Rowspan int
		Content *Element
	}
)

//NewTable creates a new table with type
func NewTable(rows ...*TableRow) *Table {
	t := new(Table)
	t.Element = NewElement("table")
	t.AddRows(rows...)
	return t
}

//AddColgroup adds a head
func (t *Table) AddColgroup(c *TableColgroup) {
	if c != nil {
		t.Colgroup = c
		t.AddElement(c.Element)
	}
}

//AddCaption adds a caption
func (t *Table) AddCaption(c *Element) {
	if c != nil {
		t.Caption = c
		t.AddElement(c)
	}
}

//AddHead adds a head
func (t *Table) AddHead(h *TableHead) {
	if h != nil {
		t.Head = h
		t.AddElement(h.Element)
	}
}

//AddBody adds a body
func (t *Table) AddBody(b *TableBody) {
	if b != nil {
		t.Body = b
		t.AddElement(b.Element)
	}
}

//AddFooter adds a footer
func (t *Table) AddFooter(f *TableFooter) {
	if f != nil {
		t.Footer = f
		t.AddElement(f.Element)
	}
}

//AddRow adds a row
func (t *Table) AddRow(row *TableRow) {
	if row != nil {
		t.Rows = append(t.Rows, row)
		t.AddElement(row.Element)
	}
}

//AddRows adds a rows
func (t *Table) AddRows(rows ...*TableRow) {
	if rows != nil {
		for _, row := range rows {
			t.AddRow(row)
		}
	}
}

//NewTableColgroup adds an colgroup row
func NewTableColgroup(span int, cols ...*TableCol) *TableColgroup {
	c := new(TableColgroup)
	c.Element = NewElement("colgroup")
	if span > 0 {
		c.Span = span
		c.SetAttribute("span", strconv.Itoa(int(c.Span)))
	}
	c.AddColgroupCols(cols...)
	return c
}

//AddColgroupCol adds a col
func (c *TableColgroup) AddColgroupCol(col *TableCol) {
	if col != nil {
		c.Cols = append(c.Cols, col)
		c.AddElement(col.Element)
	}
}

//AddColgroupCols adds a cols
func (c *TableColgroup) AddColgroupCols(cols ...*TableCol) {
	if cols != nil {
		for _, col := range cols {
			c.AddColgroupCol(col)
		}
	}
}

//NewTableCol creates a new table colgroup col
func NewTableCol(span int, style string) *TableCol {
	c := new(TableCol)
	c.Element = NewElement("col")
	if span > 0 {
		c.Span = span
		c.SetAttribute("span", strconv.Itoa(int(c.Span)))
	}
	if style > "" {
		c.Style = style
		c.SetAttribute("style", style)
	}
	return c
}

//NewTableCaption adds an caption
func NewTableCaption(text string) *Element {
	c := NewElement("caption")
	c.SetText(text)
	return c
}

//NewTableHead adds an header row
func NewTableHead(rows ...*TableRow) *TableHead {
	thead := new(TableHead)
	thead.Element = NewElement("thead")
	thead.AddHeadRows(rows...)
	return thead
}

//AddHeadRow adds a row
func (h *TableHead) AddHeadRow(row *TableRow) {
	if row != nil {
		h.Rows = append(h.Rows, row)
		h.AddElement(row.Element)
	}
}

//AddHeadRows adds a rows
func (h *TableHead) AddHeadRows(rows ...*TableRow) {
	if rows != nil {
		for _, row := range rows {
			h.AddHeadRow(row)
		}
	}
}

//NewTableBody adds body
func NewTableBody(rows ...*TableRow) *TableBody {
	tbody := new(TableBody)
	tbody.Element = NewElement("tbody")
	tbody.AddBodyRows(rows...)
	return tbody
}

//AddBodyRow adds a row
func (b *TableBody) AddBodyRow(row *TableRow) {
	if row != nil {
		b.Rows = append(b.Rows, row)
		b.AddElement(row.Element)
	}
}

//AddBodyRows adds a rows
func (b *TableBody) AddBodyRows(rows ...*TableRow) {
	if rows != nil {
		for _, row := range rows {
			b.AddBodyRow(row)
		}
	}
}

//NewTableFooter adds footer
func NewTableFooter(rows ...*TableRow) *TableFooter {
	tfoot := new(TableFooter)
	tfoot.Element = NewElement("tfoot")
	tfoot.AddFooterRows(rows...)
	return tfoot
}

//AddFooterRow adds a row
func (f *TableFooter) AddFooterRow(row *TableRow) {
	if row != nil {
		f.Rows = append(f.Rows, row)
		f.AddElement(row.Element)
	}
}

//AddFooterRows adds a rows
func (f *TableFooter) AddFooterRows(rows ...*TableRow) {
	if rows != nil {
		for _, row := range rows {
			f.AddFooterRow(row)
		}
	}
}

//NewTableRow creates a new table row
func NewTableRow(cells ...*TableCell) *TableRow {
	tr := new(TableRow)
	tr.Element = NewElement("tr")
	tr.AddCells(cells...)
	return tr
}

//AddCells adds cell to a table row
func (tr *TableRow) AddCell(cell *TableCell) {
	if cell != nil {
		tr.Cells = append(tr.Cells, cell)
		tr.AddElement(cell.Element)
	}
}

//AddCells adds cells to a table row
func (tr *TableRow) AddCells(cells ...*TableCell) {
	if cells != nil {
		for _, cell := range cells {
			tr.AddCell(cell)
		}
	}
}

//NewTableCell creates and adds new cell
func NewTableCell(header bool, colspan, rowspan int, content *Element) *TableCell {
	td := new(TableCell)
	//td := NewElement("td")
	if header {
		td.Element = NewElement("th")
	} else {
		td.Element = NewElement("td")
	}
	if colspan > 0 {
		td.Colspan = colspan
		td.SetAttribute("colspan", strconv.Itoa(int(td.Colspan)))

	}
	if rowspan > 0 {
		td.Rowspan = rowspan
		td.SetAttribute("rowspan", strconv.Itoa(int(td.Rowspan)))

	}
	td.Content = content
	td.AddElement(content)
	return td
}

//=============================================
//  Table utils //
//=============================================

//QuickTable creates a table from a given map with key-value pairs
func QuickTable(data map[string]interface{}) *Table {
	t := NewTable()
	for key, value := range data {
		cell1 := NewTableCell(false, 0, 0, NewText(key))
		cell2 := NewTableCell(false, 0, 0, NewText(fmt.Sprintf("%v", value)))
		row := NewTableRow([]*TableCell{cell1, cell2}...)
		t.AddRow(row)
	}
	return t
}
