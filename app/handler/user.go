package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/app/model/response"
	"xiaoshi/app/model"
	"encoding/json"
	"time"
)

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	user.CreateTime = time.Now().Unix()
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	//设置到redis
	setTokenToCache(user.Token)
	respUser := response.RespUser{}
	respUser.Data = user
	respUser.Message = "pass"
	respUser.Success = "0"
	respondJSON(w, http.StatusCreated, respUser)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if hadToken, _ := checkToken(db, user.Token); hadToken {
		//设置到redis
		setTokenToCache(user.Token)
		db.First(&user, "token = ?", user.Token)
		respUserData := response.RespUserData{}
		respUserData.User = user
		respUser.Data = respUserData
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, http.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, http.StatusGone, respUser)
	}
}

func EditUserInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	headerToken := r.Header.Get("X-AccessToken")
	reqUser := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqUser); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if hadToken, dbUser := checkToken(db, headerToken); hadToken {
		db.Model(&dbUser).Where("Token = ?", headerToken).Updates(model.Users{Age: reqUser.Age, NickName: reqUser.NickName, Gender: reqUser.Gender})
		respUserData := response.RespUserData{}
		respUserData.User = dbUser
		respUser.Data = respUserData
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, http.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, http.StatusGone, respUser)
	}
}

func EditPwd(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	reqUser := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqUser); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if reqUser.Token == "" {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "token not empty"
		respondJSON(w, http.StatusGone, respUser)
	} else if hadPhone, dbUser := checkPhoneNumber(db, reqUser.PhoneNumber); hadPhone {
		db.Model(&dbUser).Update("Token", reqUser.Token)
		respUser.Data = reqUser.Token
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, http.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, http.StatusGone, respUser)
	}
}
