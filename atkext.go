package ebgo

import (
  "github.com/visualfc/atk/tk"
  //"fmt"
)

type Text struct {
  *tk.Text
}

func (txt *Text) AddTag(tagname, tagstart,tagend string) {
  tk.MainInterp().Eval(txt.Id()+ " tag add "+tagname +" " +tagstart + " " + tagend )
}

func (txt *Text) ConfTag(tagname, flags  string){
  tk.MainInterp().Eval(txt.Id()+ ` tag configure ` + tagname + " " + flags)
}
