package utilities

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

type ResponseJSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}


func UnprocessableResponse(returnData *ResponseJSON)  {
	returnData.Code = 422
	returnData.Msg = "Failure: Unprocessable data error"
	returnData.Model = nil
}


func ErrorResponse(returnData *ResponseJSON, msg string) {
	returnData.Msg = msg
	returnData.Code = 400
}

func SuccessResponse(returnData *ResponseJSON, data interface{}) {
	returnData.Msg = "success"
	returnData.Code = 200
	returnData.Model = data
}

func URLReturnResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Del("currentUser")
	spew.Dump("================ API RESPONSE ================", data, "======== END RESPONSE ========")
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
