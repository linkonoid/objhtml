package objhtml

import (
	"fmt"
	//"golang.org/x/net/html"
)

//=============================================
//  <input> Input elements //
//=============================================

//=============================================
//  <input> Text Input  //
//=============================================

// Input types
const (
	InputButton        = "button"
	InputCheckbox      = "checkbox"
	InputColor         = "color"
	InputDate          = "date"
	InputDatetime      = "datetime"
	InputDatetimeLocal = "datetime-local"
	InputEmail         = "email"
	InputFile          = "file"
	InputHidden        = "hidden"
	InputImage         = "image"
	InputMonth         = "month"
	InputNumber        = "number"
	InputPassword      = "password"
	InputRadio         = "radio"
	InputRange         = "range"
	InputReset         = "reset"
	InputSearch        = "search"
	InputSubmit        = "submit"
	InputTel           = "tel"
	InputText          = "text"
	InputTime          = "time"
	InputUrl           = "url"
	InputWeek          = "week"
)

//Checkbox represents a checkbox with label
type Input struct {
	*Element
	value *Element
}

//NewTextInput creates a new "text" input
func NewInput(typeinput string) *Input {
	i := new(Input)
	i.Element = NewElement("input")
	i.SetAttribute("type", typeinput)
	return i
}

//NewTextInput creates a new "text" input
func NewTextInput(value string) *Element {
	input := NewElement("input")
	input.SetAttribute("type", "text")
	if value != "" {
		input.SetAttribute("value", value)
	}
	return input
}

//=============================================
//  <input> Checkbox Input  //
//=============================================

//Checkbox represents a checkbox with label
type CheckBoxInput struct {
	*Element
	chkbox *Element
	text   *Element
}

//NewCheckBox creates a bootstrap checkbox with label
func NewCheckBoxInput(caption string, checked bool) *CheckBoxInput {
	cb := new(CheckBoxInput)
	cb.Element = NewElement("label")
	cb.chkbox = NewElement("input")
	cb.chkbox.SetAttribute("type", "checkbox")
	if checked {
		cb.chkbox.SetAttribute("checked", "")
	}
	cb.text = NewText(caption)
	cb.Element.AddElement(cb.chkbox)
	cb.Element.AddElement(cb.text)
	cb.Element.SetAttribute("for", cb.chkbox.GetID())
	return cb
}

//Checked is true if the checkbox is checked
func (cb *CheckBoxInput) Checked() bool {
	_, exists := cb.chkbox.GetAttribute("checked")
	return exists
}

//=============================================
//  Radiobuttons checkboxes  //
//=============================================

//type Radiocheckbox struct {
//	*Element
//	group    string
//	radiobox bool
//}

////Creates either new radiobox or checkbox.
////Checkboxes are not affected by the grouping
//func NewRadioCheckbox(radiobox bool, group string, styles ...Style) *Radiocheckbox {
//	return &Radiocheckbox{newWidget(styles...), group, radiobox}
//}

////Get the state of a radiobox/checkbox
//func (t *Radiocheckbox) Checked() bool {
//	reply := make(chan string)
//	evt := Event(`reply = $("#`+t.id+`").prop("checked")`, reply)
//	events <- evt
//	if <-reply == "true" {
//		return true
//	}
//	return false
//}

////Checks the checkbox/radiobox
//func (t *Radiocheckbox) Check() {
//	events <- Event(jq(t.id, `prop("checked", "checked")`), nil)
//}

////Unchecks the checkbox/radiobox
//func (t *Radiocheckbox) Uncheck() {
//	events <- Event(jq(t.id, `prop("checked", false)`), nil)
//}

//func (t *Radiocheckbox) HTML() string {
//	if t.radiobox {
//		return fmt.Sprintf(`<input type="radio" id="%s" name="%s" style="%s"/>`, t.id, t.group, t.style.Marshal())
//	}
//	return fmt.Sprintf(`<input type="checkbox" id="%s" style="%s"/>`, t.id, t.style.Marshal())
//}

//=============================================
//  <input> File Input  //
//=============================================

//FileButton is an file-input linked to a button
type FileInput struct {
	*Element
	btn   *Element
	input *Element
}

//NewFileButton creates new 'file' input with a button
func NewFileInput(buttontype string, caption string, foldersOnly bool) *FileInput {
	fb := new(FileInput)
	//fb.Element = wgui.NewElement("div")
	fb.btn = NewButton(buttontype, caption)
	fb.input = NewElement("input")
	fb.input.SetAttribute("type", "file")
	fb.input.SetAttribute("style", "display:none;")
	if foldersOnly {
		fb.input.SetAttribute("nwdirectory", "")
	}
	fb.btn.SetAttribute(OnClick, fmt.Sprintf("%s.click()", fb.input.GetID()))
	fb.AddElement(fb.input)
	fb.AddElement(fb.btn)
	return fb

}

////OnChange registers the onchange event
//func (fb *FileInput) OnChange(handler.EventHandler) {
//	fb.input.OnEvent(OnChange, handler)
//}

//GetValue returns the selected file value
func (fb *FileInput) GetValue() string {
	return fb.input.GetValue()
}
