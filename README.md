# objhtml
Library for generating html-elements on native Golang with application of object model. Implemented full generation of all html elements

Usage:
```go
import (
	"github.com/linkonoid/objhtml"
)

func main() {

	body := objhtml.NewBody()

	//Link demo
	link := objhtml.NewLink("StartLink", "")
	link2 := objhtml.NewLink("StartLink2", "")

	//List demo
	list := objhtml.NewList(objhtml.ListUnordered)
	list.AddItem(link)
	list.AddItem(link2)

	//Checkbox demo
	checkbox := objhtml.NewCheckBoxInput("StartCheckbox", true)

	//Textarea demo
	textarea := objhtml.NewTextarea(10, 50, "Textarea demo")

	//Select create demo
	select1, _ := objhtml.NewSelect(10, false, objhtml.NewOptions("value1", "value2", "value3", "value4", "value5")...)
	select2, _ := objhtml.NewSelect(1, false, nil)
	optgroup1 := objhtml.NewOptgroup("optgroup1")
	optgroup1.AddOptions(
		objhtml.NewOption("value1", "text1", false),
		objhtml.NewOption("value2", "text2", false),
		objhtml.NewOption("value3", "text3", false),
	)
	optgroup2 := objhtml.NewOptgroup("optgroup2")
	optgroup2.AddOptions(
		objhtml.NewOption("value4", "text4", false),
		objhtml.NewOption("value5", "text5", false),
		objhtml.NewOption("value6", "text6", false),
	)
	select2.AddOptgroups(optgroup1, optgroup2)

	select2.AddOptions(
		objhtml.NewOption("value7", "text7", false),
		objhtml.NewOption("value8", "text8", false),
		objhtml.NewOption("value9", "text9", false),
	)

	//Table create demo
	table := objhtml.NewTable()
	tablecaption := objhtml.NewTableCaption("Caption demo")
	table.AddCaption(tablecaption)
	col1 := objhtml.NewTableCol(0, "background-color: #0f0")
	col2 := objhtml.NewTableCol(2, "")
	tablecolgroup := objhtml.NewTableColgroup(0, col1, col2)
	table.AddColgroup(tablecolgroup)
	row1 := objhtml.NewTableRow(
		objhtml.NewTableCell(true, 0, 0, objhtml.NewText("Head1")),
		objhtml.NewTableCell(true, 0, 0, objhtml.NewText("Head2")),
		objhtml.NewTableCell(true, 0, 0, objhtml.NewText("Head3")),
	)
	tablehead := objhtml.NewTableHead(row1)
	table.AddHead(tablehead)
	row2 := objhtml.NewTableRow(
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell4")),
		objhtml.NewTableCell(false, 1, 0, objhtml.NewText("cell5")),
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell6")),
	)
	row3 := objhtml.NewTableRow()
	row3.AddCells(
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell7")),
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell8")),
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell9")),
	)
	tablebody := objhtml.NewTableBody(row2, row3)
	table.AddBody(tablebody)
	row4 := objhtml.NewTableRow()
	row4.AddCells(
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell10")),
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell11")),
		objhtml.NewTableCell(false, 0, 0, objhtml.NewText("cell12")),
	)
	tablefooter := objhtml.NewTableFooter(row4)
	table.AddFooter(tablefooter)
	//Progress
	progress := objhtml.NewProgress(70, 100)
	//Form create
	form := objhtml.NewForm("http://linkonoid.com", "post")
	btn2 := objhtml.NewButton("Start", objhtml.ButtonSubmit)
	form.AddElement(btn2)
	//Button demo
	btn := objhtml.NewButton("Start", objhtml.ButtonButton)
	btn.OnEvent(objhtml.OnClick, BtnClicked, nil)
	//Fieldset demo
	fieldset := objhtml.NewFieldset("Fieldset", btn, list.Element)

	body.AddElement(list.Element)
	body.AddElement(checkbox.Element)
	body.AddElement(fieldset.Element)
	body.AddElement(textarea.Element)
	body.AddElement(select1.Element)
	body.AddElement(select2.Element)
	body.AddElement(table.Element)
	body.AddElement(progress.Element)
	body.AddElement(form.Element)
  //OR in one: body.AddElements(checkbox.Element, fieldset.Element, textarea.Element...)

	//Js data create
	script := objhtml.NewScript("")
	script.AddText("")
	body.AddElement(script.Element)
  
  //Create full html page
  head := objhtml.NewHead("en",
		objhtml.NewMeta(map[string]string{"charset": "utf-8"}),
		objhtml.NewMeta(map[string]string{"name": "viewport", "content": "width=device-width, initial-scale=1"}),
	)
	head.AddScript(objhtml.NewScript("js"))
	head.AddCss(objhtml.NewCss("css"))
	//body.Element.SetAttribute("onload", `runApplication("`+app.Id+`");`)
	html := objhtml.NewHtml("en", head, body)

  //Render html to 
  err := html.Element.Render()
	if err != nil {
		println("Error rendering")
		//return err
	} else {
		fmt.Fprintln(objhtml.Output.Bytes()) //Output html page to io.Writer
      		//For non cashing outs: objhtml.Output.Reset()
	}
}
```
 

