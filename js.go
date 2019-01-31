//!!!!!!!!!https://godoc.org/github.com/aichaos/rivescript-go/lang/javascript
//https://github.com/aichaos/rivescript-go/blob/master/lang/javascript/javascript.go
//https://golang.org/pkg/html/template/
//https://github.com/gopherjs/gopherjs/wiki/bindings
//https://github.com/gopherjs/jquery
//https://github.com/flimzy/jqeventrouter
//https://github.com/vishen/go-jang
//https://github.com/goaltools/jq
//https://github.com/vishen/go-jang/blob/master/main.go
//https://github.com/flimzy/onload
//https://github.com/flimzy/go-sql.js

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
