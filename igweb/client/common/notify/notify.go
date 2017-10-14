package notify

import "github.com/gopherjs/gopherjs/js"

func Success(message string) {
	js.Global.Get("alertify").Call("success", message)
}

func Error(message string) {
	js.Global.Get("alertify").Call("error", message)
}
