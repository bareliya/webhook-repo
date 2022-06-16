package controllers

import (
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/webhook-repo/models"
	"github.com/webhook-repo/utilities"
)

func GetAllData() (returndata utilities.ResponseJSON) {
	resp, err := models.MongoDbConnection.FindAllDoc()
	if err != nil {
		utilities.ErrorResponse(&returndata, err.Error())
		return
	}
	utilities.SuccessResponse(&returndata, resp)
	return
}

func AddData(playload PlayLoad) (returndata utilities.ResponseJSON) {

	doc := models.Doc{}
	doc.Author = playload.Repository.Owner.Login
	doc.Time = time.Now()
	if playload.PullRequest.ID != 0 {
		doc.Action = "PULL_REQUEST"
		comapre := strings.Split(playload.PullRequest.Head.Label, ":")
		doc.FromBranch = comapre[len(comapre)-1]
		base := strings.Split(playload.PullRequest.Base.Label, ":")
		doc.ToBranch = base[len(base)-1]
		doc.RequestId = cast.ToString(playload.PullRequest.ID)
	} else {
		doc.Action = "PUSH"
		doc.RequestId = cast.ToString(playload.Ref)

	}
	err := models.MongoDbConnection.InserOneDoc(doc)
	if err != nil {
		utilities.ErrorResponse(&returndata, err.Error())
		return
	}
	utilities.SuccessResponse(&returndata, doc)
	return
}
