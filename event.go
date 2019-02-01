package objhtml

import (
	"sync"

	velox "github.com/linkonoid/velox"
)

type (
	//EventElement represents the DOM element sending an event
	EventElement struct {
		Properties map[string]string `json:"properties"`
	}

	//Event represents a DOM event
	Event struct {
		//Сlid   string       `json:"clientid"`
		Name   string       `json:"name"`
		Sender EventElement `json:"sender"`
		//Change Element        `json:"change"`
		Inputs []EventElement `json:"inputs"`
		Data   string         `json:"data"`
		//Id      string         `json:"id"`
		//Jscript string         `json:"jscript"`
		//Reply   bool           `json:"reply"`
	}

	Sse struct {
		velox.State
		sync.Mutex
		Type     string      `json:"type"` // "innerhtml", "setvalue", "jseval", "jsalert", "json", "jqid"
		SenderId string      `json:"senderid"`
		ChangeId string      `json:"changeid"`
		Data     interface{} `json:"data"`
	}
)

var (
	SseData = &Sse{}
)

const (
	OnClick     = "onclick"
	OnChange    = "onchange"
	OnKeyPress  = "onkeypress"
	OnAbort     = "onabort"
	OnBlur      = "onblur"
	OnDblclick  = "ondblclick"
	OnError     = "onerror"
	OnFocus     = "onfocus"
	OnKeydown   = "onkeydown"
	OnKeyup     = "onkeyup"
	OnLoad      = "onload"
	OnMousedown = "onmousedown"
	OnMousemove = "onmousemove"
	OnMouseout  = "onmouseout"
	OnMouseover = "onmouseover"
	OnMouseup   = "onmouseup"
	OnReset     = "onreset"
	OnResize    = "onresize"
	OnSelect    = "onselect"
	OnSubmit    = "onsubmit"
	OnUnload    = "onunload"
)

//GetID get the id of the event sender.
func (e *EventElement) GetID() string {
	id, _ := e.Properties["id"]
	return id
}

//GetValue gets the value of the event sender.
func (e *EventElement) GetValue() string {
	id, _ := e.Properties["value"]
	return id
}

//Makes an event that can be send to the browser
func EventSend(clientid, js string, reply chan string) event {
	return event{clientid, replyid.New(""), js, reply}
}

//Sends and Event to the browser
func SendEvent(clientid, js string, reply chan string) {
	events <- EventSend(clientid, js, reply)
}

//func sendEventDataToClient(eventtype string, data interface{}, elementrender *wgui.Element, sender, change *wgui.Element) {
func SendEventDataToClient(eventtype string, data interface{}, renderelement *Element, sender, change *Element) {
	SseData.Lock()
	SseData.Type = eventtype
	SseData.SenderId = sender.GetID()
	SseData.ChangeId = change.GetID()
	switch eventtype {
	case "innerhtml":
		buf := Output
		Output.Reset()
		err := renderelement.Render()
		if err == nil {
			SseData.Data = Output.String()
		}
		Output = buf
		buf.Reset()
	default:
		SseData.Data = data
	}
	SseData.Unlock()
	SseData.Push()
}
