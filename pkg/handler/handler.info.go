package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/alisyahbana/tax-calculator/pkg/common/app"
	"github.com/alisyahbana/tax-calculator/pkg/common/env"
	"github.com/alisyahbana/tax-calculator/pkg/common/log"
	"net/http"
	"time"
)

type Info struct {
	Version   string `json:"version"`
	Env       string `json:"env"`
	StartTime string `json:"start_time"`
}

var startTime string

func init() {
	startTime = time.Now().Format("2006-01-02 15:04:05")
}

func InfoHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	info, err := json.Marshal(Info{
		Version:   app.GetConfig().Version,
		Env:       env.GetEnv(),
		StartTime: startTime,
	})

	if err != nil {
		log.Error(err.Error())
	}

	writer.Write(info)
}
