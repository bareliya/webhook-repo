package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/webhook-repo/models"
	"log"
	"strings"
	"time"
)

func StartUI() {
	a := app.New()
	w := a.NewWindow("UI to Show Github Action")

	clock := widget.NewLabel("")
	UpdateData(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			UpdateData(clock)
		}
	}()
	w.ShowAndRun()
}

func UpdateData(clock *widget.Label) {
	resp, err := models.MongoDbConnection.FindAllDoc()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(resp)
	}
	formatted := ""
	current := resp[len(resp)-1]

	if current.Action == "PUSH" {
		arr := strings.Split(current.RequestId, "/")
		br := arr[len(arr)-1]
		formatted = current.Author + " pushed to " + br + " on " + current.Time.String()
	} else if current.Action == "PULL_REQUEST" {
		formatted = current.Author + " submitted a pull request from " + current.FromBranch + " to " + current.ToBranch + " on " + current.Time.String()
	} else {
		formatted = current.Author + " merged branch " + current.FromBranch + " to " + current.ToBranch + " on " + current.Time.String()
	}
	clock.SetText(formatted)
}
