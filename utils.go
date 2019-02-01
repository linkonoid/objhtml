package objhtml

import (
	//	"fmt"
	//	"html"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

//=============================================
//  Variable Declartions  //
//=============================================
type (
	//Threadsafe unique list of strings
	unique struct {
		list  []string
		mutex sync.Mutex
	}
	event struct {
		clid    string
		id      string
		jscript string
		reply   chan string
	}
)

var (
	events      = make(chan event, 1000) //event channel
	replyid     unique                   //Id for replies
	replyqueue  = make(chan event, 1000) //replies
	Ids         unique                   //Id for widgets
	renderMutex sync.Mutex
	handlers    = map[string]func(){}
	resources   = map[string][]byte{}
)

//=============================================
//  Practical  //
//=============================================
//When you compile a file, be it image, or page or whatever, to a []byte, it can be used with this map.
//when the page is requested on the server, fx. /img/cat.jpg, it will write the bytes in
//		hgui.SetResource(map[string][]byte{"/img/cat.jpg", catpicvar})
//back to the client.
func SetResource(files map[string][]byte) {
	resources = files
}

func escape(s string) string {
	s = strings.Replace(s, `"`, `\"`, -1)
	s = strings.Replace(s, `'`, `\'`, -1)
	return s
	//return html.EscapeString(s)
}

func (i *unique) New(prefix string) string {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.list == nil {
		i.list = make([]string, 0, 100)
	}
	var id string
	for {
		id = prefix + strconv.Itoa(rand.Int())
		for _, v := range i.list {
			if v == id {
				continue
			}
		}
		break
	}
	i.list = append(i.list, id)
	//i.mutex.Unlock()
	return id
}

func (i *unique) Remove(id string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.list == nil {
		return
	}

	for j, v := range i.list {
		if v == id {
			i.list = append(i.list[:j], i.list[j+1:]...)
		}
	}
	//i.mutex.Unlock()
}
