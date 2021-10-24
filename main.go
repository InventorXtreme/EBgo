package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/visualfc/atk/tk"
	"fmt"
  //"github.com/faiface/mainthread"
	"time"
	//"strconv"
  //"runtime"
	vlc "github.com/adrg/libvlc-go/v3"
	//"strings"
  "errors"
)

func SimpleGet(url string) (string,error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", errors.New("HTTP Get Error")
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", errors.New("Body Parse Error")
  }
  bodytext := string(body)
  return bodytext, nil
}

type Window struct {
	*tk.Window
}




func (w *Window) winfo_id() string {
	interp := tk.MainInterp()
	val, _ := interp.EvalAsString(fmt.Sprintf("winfo id %v", w.Id()))
	return val
}

func WidFStr(id string) string {
	interp := tk.MainInterp()
	val, _:= interp.EvalAsString(fmt.Sprintf("winfo id %v", id))
	return val
}

func GetChat(server, port, channel string) string{
  url := server + ":" + port + "?get"
  if channel != "" {
    url = url + "&channel=" + channel
  }
  data, err := SimpleGet(url)
  if err != nil{
    log.Fatal(err)
  }
  return data
}

func (txt *Text) UpdateText(server,port, channel string, value *string,queue *Queue ){
  fmt.Println("e")
  *value = GetChat(server,port,channel)
  time.Sleep(4*time.Second)
  queue.Insert("g")

  //fmt.Println("eee")
  //fmt.Println(root.Id())
  //x := tk.SendEvent(root,"<<test>>")
  //fmt.Println(x)

  fmt.Println("eeeee")
  //fmt.Println(z)
  //txt.SetText(z)
    // need to get return value and pass back out of current go call back to the call where I can use mainthread

}

func NewWindow(queue *Queue) *Window {
  root := &Window{tk.RootWindow()}
  root.ResizeN(300,200)
  bigtset := ""
  fmt.Println(queue.Remove())

  //widgets
  // text widget
  help := &Text{tk.NewText(root)}
  help.SetText("eee")
  help.SetHeight(5)
  help.AddTag("pog", "1.0", "1.4")
  help.ConfTag("pog","-background red")

  //Refresh Button
  refr := tk.NewButton(root,"refresh")
  root.BindEvent("<<test>>",func(_ *tk.Event){
    help.SetText(bigtset)
    fmt.Println("fff")
  })

  refr.OnCommand(func(){
    go help.UpdateText("http://server.whirlwind.run","81","",&bigtset,queue)
    //fmt.Println(bigtset)
    //tk.SendEvent(root,"<<test>>")
  })
  //tk.Pack(help)
  //tk.Pack(refr)
  tk.NewVPackLayout(root).AddWidgets(help,tk.NewLayoutSpacer(root,0,true),refr)

  tk.SendEvent(root,"<<test>>")

  return root
}


func main() {
  fmt.Println("Started")
  queue := CreateQueue(100)
  queue.Insert("testpoggers")
  if err := vlc.Init(); err != nil {
    fmt.Println("ERR VLC INIT")
  }
  fmt.Println("tset")
  tk.Init()
  root := gui(queue)
  h := ""
  for {
  tk.Update()
  //time.Sleep(100000*time.Nanosecond)
  h, _ = queue.Remove()
  switch h {
  case "g":
    tk.SendEvent(root,"<<test>>")
  case "h":
    fmt.Println("d")
  }

}
  //tk.MainLoop(gui)
  //gui()


}

func gui(queue *Queue) *Window {
  root := NewWindow(queue)
  root.SetTitle("Test")
  root.Center(nil)
  root.ShowNormal()
  return root
}
