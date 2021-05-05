package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Tried with an EventBox surrounding the gtk.Fixed but it did not help
	eventBox, _ := gtk.EventBoxNew()
	fixed, _ := gtk.FixedNew()
	textView, _ := gtk.TextViewNew()
	textView.SetHasWindow(false)  // When set to false, the coords are correct, but doesn't show text
											// When set to true, the coords are incorrect, but shows text
	textView.SetSizeRequest(100,100)
	buffer, _ :=textView.GetBuffer()
	buffer.SetText("Test")
	fixed.Put(textView, 100,100)

	eventBox.Connect("button-press-event", func(eventBox *gtk.EventBox, e *gdk.Event) {
		eventButton := gdk.EventButtonNewFromEvent(e)
		if eventButton.Button() != gdk.BUTTON_PRIMARY {
			return
		}

		eventMotion := gdk.EventMotionNewFromEvent(e)
		x, y := eventMotion.MotionVal()

		fmt.Println(x,y)
	})

	eventBox.Add(fixed)
	win.Add(eventBox)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}