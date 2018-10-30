package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/model/response"
	"xiaoshi/model"
	"encoding/json"
	"time"
	"xiaoshi/conf"
	"xiaoshi/util"
	log "github.com/alecthomas/log4go"
)

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Info("register start")
	user := &model.Users{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(user); err != nil {
		respondError(w, conf.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	user.CreateTime = time.Now().Unix()
	if err := db.Save(user).Error; err != nil {
		respondError(w, conf.StatusInternalServerError, err.Error())
		return
	}
	//设置到redis
	util.SetTokenToCache(user.Token)
	respUser := &response.RespUser{
		RespModel: &response.RespModel{
			Success: "0",
			Message: "pass",
		},
		Data: user,
	}
	respondJSON(w, conf.StatusCreated, respUser)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, conf.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if hadToken, _ := checkToken(db, user.Token); hadToken {
		//设置到redis
		util.SetTokenToCache(user.Token)
		db.First(&user, "token = ?", user.Token)
		respUserData := response.RespUserData{}
		respUserData.User = user
		respUser.Data = respUserData
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, conf.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, conf.StatusGone, respUser)
	}
}

func EditUserInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	headerToken := r.Header.Get("X-AccessToken")
	reqUser := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqUser); err != nil {
		respondError(w, conf.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if hadToken, dbUser := checkToken(db, headerToken); hadToken {
		db.Model(&dbUser).Where("Token = ?", headerToken).Updates(model.Users{Age: reqUser.Age, NickName: reqUser.NickName, Gender: reqUser.Gender})
		respUserData := response.RespUserData{}
		respUserData.User = dbUser
		respUser.Data = respUserData
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, conf.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, conf.StatusGone, respUser)
	}
}

func EditPwd(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	reqUser := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqUser); err != nil {
		respondError(w, conf.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if reqUser.Token == "" {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "token not empty"
		respondJSON(w, conf.StatusGone, respUser)
	} else if hadPhone, dbUser := checkPhoneNumber(db, reqUser.PhoneNumber); hadPhone {
		db.Model(&dbUser).Update("Token", reqUser.Token)
		respUser.Data = reqUser.Token
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, conf.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, conf.StatusGone, respUser)
	}
}
