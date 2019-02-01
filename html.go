package objhtml

import (
	//	"fmt"
	"strconv"
)

// Content types
const (
	TypeTextHtml   = "text/html"
	TypeJavascript = "text/javascript"
	TypeCss        = "text/css"
)

//=============================================
//  <html>  Html //
//=============================================

type Html struct {
	*Element
	Lang string `json:"lang"`
	Head *Head  `json:"head"`
	Body *Body  `json:"body"`
}

//Create new html element
func NewHtml(lang string, head *Head, body *Body) *Html {
	h := new(Html)
	h.Element = NewElement("html")
	if lang != "" {
		h.Lang = lang
		h.SetAttribute("lang", h.Lang)
	}
	if head != nil {
		h.Head = head
		h.AddElement(head.Element)
	}
	if body != nil {
		h.Body = body
		h.AddElement(body.Element)
	}
	return h
}

//=============================================
//  <body>  Body //
//=============================================

type Body struct {
	*Element
}

//Create new header element
func NewBody() *Body {
	b := new(Body)
	b.Element = NewElement("body")
	return b
}

//=============================================
//  <head>  Head //
//=============================================
type Head struct {
	*Element
	Charset string     `json:"charset"`
	Meta    []*Element `json:"meta"`
	Css     []*Css     `json:"css"`
	Script  []*Script  `json:"script"`
}

//Create new header element
func NewHead(charset string, metas ...*Element) *Head {
	h := new(Head)
	h.Element = NewElement("head")
	if charset != "" {
		h.Charset = charset
	}
	h.AddMetas(metas...)
	return h
}

//AddMeta add Meta attr/value to header
func (h *Head) AddMeta(meta *Element) {
	if meta != nil {
		h.Meta = append(h.Meta, meta)
		h.AddElement(meta)
	}
}

//AddMeta add new Meta attr/value to header
func (h *Head) AddNewMeta(meta map[string]string) {
	m := NewMeta(meta)
	h.AddMeta(m)
}

//AddMetas adds a Metas array to header
func (h *Head) AddMetas(metas ...*Element) {
	//	m := make([]*map[string]string, len(metas))
	for _, v := range metas {
		h.AddMeta(v)
	}
}

//AddCss add a Css to header
func (h *Head) AddCss(css *Css) {
	if css != nil {
		h.Css = append(h.Css, css)
		h.AddElement(css.Element)
	}
}

//AddNewCss add new Css to header
func (h *Head) AddNewCss(href string) {
	c := NewCss(href)
	h.AddCss(c)
}

//AddCss adds a Css array to header
func (h *Head) AddCsss(csss ...*Css) {
	//	m := make([]*Css, len(csss))
	for _, v := range csss {
		h.AddCss(v)
	}
}

//AddScript add a Script to header
func (h *Head) AddScript(script *Script) {
	if script != nil {
		h.Script = append(h.Script, script)
		h.AddElement(script.Element)
	}
}

//AddScript add new Script to header
func (h *Head) AddNewScript(src string) {
	s := NewScript(src)
	h.AddElement(s.Element)
}

//AddScripts adds a Script array to header
func (h *Head) AddScripts(scripts ...*Script) {
	//	m := make([]*Script, len(scripts))
	for _, v := range scripts {
		h.AddScript(v)
	}
}

//=============================================
//  <meta>  Meta //
//=============================================
//Create Meta element
func NewMeta(meta map[string]string) *Element {
	m := NewElement("meta")
	//m := make([]*Element, len(meta))
	for k, v := range meta {
		m.SetAttribute(k, v)
	}
	return m
}

//=============================================
//  <script> Script element //
//=============================================

type Script struct {
	*Element
	Type string
	Src  string
	Text string
}

//NewScript creates a scrypt
func NewScript(src string) *Script {
	s := new(Script)
	s.Element = NewElement("script")
	s.Type = TypeJavascript
	s.SetAttribute("type", s.Type)
	if src != "" {
		s.AddSrc(src)
	}
	return s
}

//AddSrc adds a src
func (s *Script) AddSrc(src string) {
	if src > "" {
		s.Src = src
		s.SetAttribute("src", s.Src)
	}
}

//AddSrc adds a text
func (s *Script) AddText(text string) {
	if text > "" {
		//s.SetText(html.EscapeString(text))
		s.SetText(text)
	}
}

//=============================================
//  <css>  Css //
//=============================================

type Css struct {
	*Element
	Type string `json:"type"`
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

//NewTextInput creates a new "text" input
func NewCss(href string) *Css {
	c := new(Css)
	c.Element = NewElement("link")
	c.Type = TypeCss
	c.SetAttribute("type", c.Type)
	c.AddRel("stylesheet")
	c.AddHref(href)
	return c
}

//AddSrc adds a src
func (c *Css) AddHref(href string) {
	if href != "" {
		c.SetAttribute("href", href)
	}
}

//AddSrc adds a text
func (c *Css) AddRel(rel string) {
	if rel != "" {
		c.Type = rel
		c.SetAttribute("rel", c.Type)
	}
}

//=============================================
//  <button>  Button //
//=============================================

// Button types
const (
	ButtonButton = "button"
	ButtonReset  = "reset"
	ButtonSubmit = "submit"
)

//NewButton creates <button> element
func NewButton(caption, buttontype string) *Element {
	btn := NewElement("button")
	if buttontype != "" {
		btn.SetAttribute("type", buttontype)
	}
	if caption != "" {
		btn.SetText(caption)
	}
	return btn
}

//=============================================
//  <a> Link  //
//=============================================

//NewLink creates <a> link element
func NewLink(caption, href string) *Element {
	link := NewElement("a")
	if caption != "" {
		link.SetAttribute("href", href)
	} else {
		link.SetAttribute("href", "#")
	}
	if caption != "" {
		link.AddElement(NewText(caption))
	}
	return link
}

//=============================================
//  <ul> ListItems  //
//=============================================

//ListItems a list of elements
type ListItems []*Element

//List is a struct for <ul>, <ol>, <dl>
type List struct {
	*Element
	Items ListItems
}

//Lists types
const (
	//ListUnordered is <ul>
	ListUnordered = "ul"
	//ListOrdered is <ol>
	ListOrdered = "ol"
	//DescriptionList is <dl>
	DescriptionList = "dl"
)

//NewList creates a new list
func NewList(listType string) *List {
	l := new(List)
	l.Element = NewElement(listType)
	l.Items = make(ListItems, 0)
	return l
}

//AddItem creates new LI, adds the elem to the li and returns the li to the caller.
func (l *List) AddItem(elem *Element) *Element {
	item := NewElement("li")
	item.AddElement(elem)
	l.AddElement(item)
	l.Items = append(l.Items, elem)
	return item
}

//=============================================
//  Textarea  //
//=============================================

type Textarea struct {
	*Element
	Rows  int
	Cols  int
	Value string
}

//Multiline text input
func NewTextarea(rows, cols int, value string) *Textarea {
	t := new(Textarea)
	t.Element = NewElement("textarea")
	if rows > 0 {
		t.Rows = rows
		t.SetAttribute("rows", strconv.Itoa(int(t.Rows)))
	}
	if cols > 0 {
		t.Cols = cols
		t.SetAttribute("cols", strconv.Itoa(int(t.Cols)))
	}
	if value > "" {
		t.Value = value
		t.SetValue(value)
	}
	return t
}

//func (t *Textarea) GetValue() string {
//	v := t.GetValue()
//	return v
//}

func (t *Textarea) SetValue(s string) {
	t.SetText(s)
}

//=============================================
//  Fieldset  //
//=============================================

type Fieldset struct {
	*Element
	legend  *Element
	content []*Element
}

func NewFieldset(legend string, content ...*Element) *Fieldset {
	fs := new(Fieldset)
	fs.Element = NewElement("fieldset")
	//fs.SetText(fs.content)
	if legend != "" {
		fs.legend = NewText(legend)
		lgd := NewElement("legend")
		lgd.AddElement(fs.legend)
		fs.AddElement(lgd)
	}
	//fs.content = make([]*Element, 0)
	fs.content = content
	for _, c := range content {
		fs.AddElement(c)
	}
	//sd, _ := fs.AddHTML(fs.content, nil)
	//fs.Element = sd[0]
	return fs
}

//=============================================
//  Image  //
//=============================================

type Image struct {
	*Element
	Src string
	Alt string
}

func NewImage(src, alt string) *Image {
	img := new(Image)
	img.Element = NewElement("img")
	if src != "" {
		img.SetAttribute("src", src)
	} else {
		img.SetAttribute("src", "#")
	}
	if alt != "" {
		img.SetAttribute("alt", alt)
	}
	return img
}

//=============================================
//  Progress  //
//=============================================

type Progress struct {
	*Element
	Value int
	Max   int
}

func NewProgress(value, max int) *Progress {
	pgs := new(Progress)
	pgs.Element = NewElement("progress")
	if max > 0 {
		pgs.Max = max
		pgs.SetAttribute("max", strconv.Itoa(int(pgs.Max)))
	}
	if value > max {
		value = max
	}
	if value >= 0 {
		pgs.Value = value
		pgs.SetValue(strconv.Itoa(int(pgs.Value)))
	}
	return pgs
}

//=============================================
//  Label  //
//=============================================

type Label struct {
	*Element
	For *Element
}

func NewLabel(elem *Element) *Label {
	lbl := new(Label)
	lbl.Element = NewElement("label")
	id := elem.GetID()
	lbl.SetAttribute("for", id)
	return lbl
}

//=============================================
//  Form  //
//=============================================

type Form struct {
	*Element
	Name          string
	Action        string
	Method        string
	Target        string
	Enctype       string
	Autocomplete  string
	AcceptCharset string
}

//// Method types
//const (
//	FormGet  = "get"
//	FormPost = "post"
//)

func NewForm(action, method string) *Form {
	f := new(Form)
	f.Element = NewElement("form")
	if action == "" {
		action = "#"
	}
	f.Action = action
	f.SetAttribute("action", f.Action)
	if method == "" {
		method = "post"
	}
	f.Method = method
	f.SetAttribute("method", f.Method)
	id := f.GetID()
	f.Name = id
	f.SetAttribute("name", id)
	return f
}
