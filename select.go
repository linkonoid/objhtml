package objhtml

import (
	"strconv"
)

//=============================================
//  Select  //
//=============================================

type (
	Select struct {
		*Element
		Optgroups []*Optgroup
		Options   []*Option
		Size      int
		Multiple  bool
		Name      string
		Autofocus bool
		Required  bool
		Form      *Element
		Disabled  bool
	}
	Optgroup struct {
		*Element
		Options  []*Option
		Label    string
		Disabled bool
	}
	Option struct {
		*Element
		Value    string
		Text     *Element
		Selected bool
		Label    string
		Disabled bool
	}
)

//Create new combobox, multiselection og list item.
func NewSelect(size int, multiple bool, options ...*Option) (*Select, error) {
	s := new(Select)
	s.Element = NewElement("select")
	s.SetSize(size)
	s.AddOptions(options...)
	if multiple {
		s.SetMultiple()
	} else {
		s.RemoveMultiple()
	}
	id := s.GetID()
	s.Name = id
	s.SetAttribute("name", id)
	return s, nil
}

func (s *Select) SetSize(size int) {
	if size > 0 {
		s.Size = size
		s.SetAttribute("size", strconv.Itoa(int(s.Size)))
	}
}

func (s *Select) SetMultiple() {
	s.Multiple = true
	s.SetAttribute("multiple", "")
}

func (s *Select) RemoveMultiple() {
	s.Disabled = false
	s.RemoveAttribute("multiple")
}

func (s *Select) Disable() {
	s.Disabled = true
	s.SetAttribute("disable", "")
}

func (s *Select) Enable() {
	s.Disabled = false
	s.RemoveAttribute("disable")
}

func NewOptgroup(label string) *Optgroup {
	g := new(Optgroup)
	g.Element = NewElement("optgroup")
	g.SetLabel(label)
	return g
}

func (g *Optgroup) AddOption(o *Option) {
	if o != nil {
		g.Options = append(g.Options, o)
		g.AddElement(o.Element)
	}
}

func (g *Optgroup) AddOptions(options ...*Option) {
	if options != nil {
		for _, o := range options {
			g.AddOption(o)
		}
	}
}

func (g *Optgroup) SetLabel(label string) {
	if label > "" {
		g.Label = label
		g.SetAttribute("label", g.Label)
	}
}

func (g *Optgroup) Disable() {
	g.Disabled = true
	g.SetAttribute("disable", "")
}

func (g *Optgroup) Enable() {
	g.Disabled = false
	g.RemoveAttribute("disable")
}

func NewOption(value, text string, selected bool) *Option {
	o := new(Option)
	o.Element = NewElement("option")
	o.SetValue(value)
	o.SetText(text)
	if selected {
		o.Select()
	}
	return o
}

func NewOptions(values ...string) []*Option {
	o := make([]*Option, len(values))
	for i, v := range values {
		o[i] = NewOption(v, v, false)
	}
	return o
}

func (o *Option) SetValue(value string) {
	if value > "" {
		o.Value = value
		o.SetAttribute("value", o.Value)
	}
}

func (o *Option) SetText(text string) {
	if text > "" {
		//o.text = NewStyledText(text, style)
		o.Text = NewText(text)
		o.Element.AddElement(o.Text)
	}
}

func (o *Option) SetLabel(label string) {
	if label > "" {
		o.SetAttribute("label", label)
	}
}

func (o *Option) Select() {
	o.Selected = true
	o.SetAttribute("selected", "")
}

func (o *Option) UnSelect() {
	o.Selected = false
	o.RemoveAttribute("selected")
}

func (o *Option) Disable() {
	o.Disabled = true
	o.SetAttribute("disabled", "")
}

func (o *Option) Enable() {
	o.Disabled = false
	o.RemoveAttribute("disable")
}

func (s *Select) AddOptgroup(g *Optgroup) error {
	if g != nil {
		s.Optgroups = append(s.Optgroups, g)
		s.AddElement(g.Element)
	}
	return nil
}

func (s *Select) AddOptgroups(optgroups ...*Optgroup) {
	if optgroups != nil {
		for _, g := range optgroups {
			s.AddOptgroup(g)
		}
	}
}

func (s *Select) AddOption(o *Option) {
	if o != nil {
		s.Options = append(s.Options, o)
		s.AddElement(o.Element)
	}
}

func (s *Select) AddOptions(options ...*Option) {
	if options != nil {
		for _, o := range options {
			s.AddOption(o)
		}
	}
}

func (s *Select) GetSelected() []string {
	var buf []string
	for _, o := range s.Options {
		_, sel := o.Element.GetAttribute("selected")
		if sel {
			val := o.Element.GetValue()
			buf = append(buf, val)
		}
	}
	return buf
}

//func (s *Selectform) HTML() (html string) {
//	if s.multiple {
//		html = fmt.Sprintf(`<select id="%s" size="%d" style="%s" multiple>`, s.id, s.size, s.style.Marshal())
//	} else {
//		html = fmt.Sprintf(`<select id="%s" size="%d" style="%s">`, s.id, s.size, s.style.Marshal())
//	}
//	for _, v := range s.options {
//		html += v.HTML() + "\n"
//	}
//	html += "</select>"
//	return
//}

//func (s *Selectform) SetOptions(o ...*Option) {
//	s.options = o
//	html := ""
//	for _, v := range s.options {
//		html += v.HTML()
//	}
//	events <- Event(jq(s.id, "html('"+escape(html)+"')"), nil)
//}

//func (s *Selectform) Selected() (string, []string) {
//	reply := make(chan string)
//	evt := Event(fmt.Sprintf(`reply = $("#%s").val(); if (reply == null) {reply = ""}`, s.id), reply)
//	events <- evt
//	if s.multiple {
//		return "", strings.Split(<-evt.reply, ",")
//	}
//	return <-evt.reply, nil
//}

//func NewOption(value, text string) *Option {
//	out := &Option{newWidget(), value, text}
//	return out
//}

//func NewOptions(values ...string) []*Option {
//	buf := make([]*Option, len(values))
//	for i, v := range values {
//		buf[i] = NewOption(v, v)
//	}
//	return buf
//}

//func (o *Option) HTML() string {
//	return fmt.Sprintf(`<option id="%s" value="%s" style="%s">%s</option>`, o.id, o.value, o.style.Marshal(), o.text)
//}
