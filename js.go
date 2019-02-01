package objhtml

import (
	"fmt"

	"golang.org/x/net/html"
)

func JqFuncCont(textinner string) string {
	//return html.EscapeString(fmt.Sprintf(`$(function(){%s});`, textinner))
	return fmt.Sprintf(`$(function(){%s});`, textinner)
}

func JqIdPropCont(id, property, textinner string) string {
	//return html.EscapeString(fmt.Sprintf(`$('#%s').%s({%s});`, id, property, textinner))
	return fmt.Sprintf(`$("#%s").%s(%s);`, id, property, textinner)
}

func JqIdProp(id, property string) string {
	return html.EscapeString(fmt.Sprintf(`$("#%s").%s`, id, property))
}

func JsAlert(alert string) string {
	return html.EscapeString(fmt.Sprintf("alert("+html.EscapeString(alert)+");", alert))
}

//NewJs send function
func JsSendFunc(send func()) {
	//events <- Event(clientid, nil, nil, send, nil, nil)
}

//NewJs send data
func JsSendJson(json string) {
	//events <- Event(clientid, nil, nil, send, nil, nil)
}
