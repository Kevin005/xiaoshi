package handler

import (
	"encoding/json"
	log "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/conf"
	"xiaoshi/model"
)

type RespI interface {
	Normal(payload interface{})
	ServerError(payload interface{})
	BadRequest(payload interface{})
}

type RespJs struct {
	w http.ResponseWriter
}

func (r *RespJs) Normal(payload interface{}) {
	respondJSON(r.w, conf.STATUS_CREATED, payload)
}

func (r *RespJs) ServerError(payload interface{}) {
	respondJSON(r.w, conf.STATUS_INTERNAL_SERVER_ERROR, payload)
}

func (r *RespJs) BadRequest(payload interface{}) {
	respondJSON(r.w, conf.STATUS_BAD_REQUEST, payload)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(conf.STATUS_INTERNAL_SERVER_ERROR)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
	log.Info(string(response[:]))
	//todo 可以统一返回json {"success":0}格式
}

func respondReject() {
	//todo 可以统一返回json {"success":1}格式
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
	//todo 可以统一返回json {"success":0}格式
}

func checkToken(db *gorm.DB, rToken string) (bool, model.Users) {
	user := model.Users{}
	db.First(&user, "token = ?", rToken)
	if user.Token != "" {
		return true, user
	}
	return false, user
}

func checkPhoneNumber(db *gorm.DB, phoneNum string) (bool, model.Users) {
	user := model.Users{}
	db.First(&user, "phone_number = ?", phoneNum)
	if user.PhoneNumber != "" {
		return true, user
	}
	return false, user
}
