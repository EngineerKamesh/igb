package notify

import (
	"errors"
	"reflect"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"
	"github.com/uxtoolkit/cog"
)

var cogType reflect.Type
var alreadyStarted bool
var JS = js.Global

type Notify struct {
	cog.UXCog
	alertify                         *js.Object
	successNotificationEventListener func(*js.Object)
	errorNotificationEventListener   func(*js.Object)
	divContainerID                   string
}

func NewNotify() *Notify {
	n := &Notify{}
	n.SetCogType(cogType)
	n.SetCleanupFunc(n.Cleanup)
	return n
}

func (n *Notify) StartListening() {

	alreadyStarted = true
	D := dom.GetWindow()
	n.successNotificationEventListener = D.AddEventListener("displaySuccessNotification", false, func(event dom.Event) {
		message := event.Underlying().Get("detail").String()
		n.notifySuccess(message)
	})

	n.errorNotificationEventListener = D.AddEventListener("displayErrorNotification", false, func(event dom.Event) {
		message := event.Underlying().Get("detail").String()
		n.notifyError(message)
	})
}

func (n *Notify) StopListening() {
	D := dom.GetWindow()
	if n.successNotificationEventListener != nil {
		D.RemoveEventListener("displaySuccessNotification", false, n.successNotificationEventListener)
	}

	if n.errorNotificationEventListener != nil {
		D.RemoveEventListener("displayErrorNotification", false, n.errorNotificationEventListener)
	}

}

func (n *Notify) Cleanup() {
	n.StopListening()
}

func (n *Notify) notifySuccess(message string) {
	n.alertify.Call("success", message)
}

func (n *Notify) notifyError(message string) {
	n.alertify.Call("error", message)
}

func (n *Notify) Start() error {
	if alreadyStarted == true {
		return errors.New("The notification cog can be instantiated only once.")
	}

	n.SetID(n.divContainerID)
	n.alertify = js.Global.Get("alertify")

	n.StartListening()
	return nil
}

func init() {
	cogType = reflect.TypeOf(Notify{})
}

// Success broadcasts a success notification message, it's meant to be used by the cog user
func Success(message string) {
	var eventDetail = js.Global.Get("Object").New()
	eventDetail.Set("detail", message)
	customEvent := js.Global.Get("window").Get("CustomEvent").New("displaySuccessNotification", eventDetail)
	js.Global.Get("window").Call("dispatchEvent", customEvent)
}

// Error broadcasts an error notification message, it's meant to be used by the cog user
func Error(message string) {
	var eventDetail = js.Global.Get("Object").New()
	eventDetail.Set("detail", message)
	customEvent := js.Global.Get("window").Get("CustomEvent").New("displayErrorNotification", eventDetail)
	js.Global.Get("window").Call("dispatchEvent", customEvent)
}
