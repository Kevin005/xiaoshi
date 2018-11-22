package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/conf"
	"xiaoshi/model"
	"xiaoshi/model/request"
	"xiaoshi/model/response"
)

type handBooks struct {
	rjs RespI
}

func (h *handBooks) CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("X-AccessToken")
	h.rjs = &RespJs{w}
	hadToken, user := checkToken(db, t)
	if !hadToken {
		rbk := response.RespBook{
			RespModel: response.RespModel{
				Message: "reject",
				Success: "1",
			},
		}
		h.rjs.ServerError(rbk)
		return
	}
	//get request book
	reqBook := request.ReqBooks{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBook); err != nil {
		h.rjs.BadRequest(err.Error())
	}
	defer r.Body.Close()

	//开启事务
	tx := db.Begin()
	//save book
	bk := model.Books{
		BookName:  reqBook.BookName,
		Author:    reqBook.Author,
		PageTotal: reqBook.PageTotal,
		Image:     reqBook.Image,
	}
	if err := tx.Save(&bk).Error; err != nil {
		respondError(w, conf.STATUS_INTERNAL_SERVER_ERROR, err.Error())

		h.rjs.ServerError(err.Error())
		//回滚事务
		tx.Rollback()
		return
	}
	mbk := model.MyBooks{
		BookId:   bk.ID,
		UserId:   user.ID,
		Private:  reqBook.Private,
		Progress: "0",
	}
	if err := tx.Save(&mbk).Error; err != nil {
		h.rjs.ServerError(err.Error())
		//回滚事务
		tx.Rollback()
		return
	}
	//提交事务
	tx.Commit()

	rbk := response.RespBook{
		Data: response.RespBookData{
			UserId:     user.ID,
			BookId:     bk.ID,
			BookAuthor: bk.Author,
			BookName:   bk.BookName,
		},
		RespModel: response.RespModel{
			Message: "pass",
			Success: "0",
		},
	}
	h.rjs.Normal(rbk)
}

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("X-AccessToken")
	rsp := &RespJs{w}
	hadToken, user := checkToken(db, t)
	if !hadToken {
		rbk := response.RespBook{
			RespModel: response.RespModel{
				Message: "reject",
				Success: "1",
			},
		}
		rsp.ServerError(rbk)
		return
	}
	//get request book
	reqBook := request.ReqBooks{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBook); err != nil {
		rsp.BadRequest(err.Error())
	}
	defer r.Body.Close()

	//开启事务
	tx := db.Begin()
	//save book
	bk := model.Books{
		BookName:  reqBook.BookName,
		Author:    reqBook.Author,
		PageTotal: reqBook.PageTotal,
		Image:     reqBook.Image,
	}
	if err := tx.Save(&bk).Error; err != nil {
		respondError(w, conf.STATUS_INTERNAL_SERVER_ERROR, err.Error())

		rsp.ServerError(err.Error())
		//回滚事务
		tx.Rollback()
		return
	}
	mbk := model.MyBooks{
		BookId:   bk.ID,
		UserId:   user.ID,
		Private:  reqBook.Private,
		Progress: "0",
	}
	if err := tx.Save(&mbk).Error; err != nil {
		rsp.ServerError(err.Error())
		//回滚事务
		tx.Rollback()
		return
	}
	//提交事务
	tx.Commit()

	rbk := response.RespBook{
		Data: response.RespBookData{
			UserId:     user.ID,
			BookId:     bk.ID,
			BookAuthor: bk.Author,
			BookName:   bk.BookName,
		},
		RespModel: response.RespModel{
			Message: "pass",
			Success: "0",
		},
	}
	rsp.Normal(rbk)
}

func GetMyBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	vars := mux.Vars(r)
	//获取url中的值
	userId := vars["user_id"]
	respBook := response.RespBook{}
	if hadToken, _ := checkToken(db, token); hadToken {
		myBooks := []model.MyBooks{}
		//查出我所有读的书 todo 需要判断正在读
		db.Find(&myBooks, "user_id = ?", userId)
		respBookData := []response.RespBookData{}
		//根据myBook外键BookId查出书籍信息 todo 不用for,可以用sql join一次查出
		for _, myBook := range myBooks {
			book := model.Books{}
			db.First(&book, "id = ?", myBook.BookId)
			bookData := response.RespBookData{}
			bookData.BookName = book.BookName
			bookData.BookId = book.ID
			bookData.BookAuthor = book.Author
			respBookData = append(respBookData, bookData)
		}
		respBook.Data = respBookData
		respBook.Success = "0"
		respondJSON(w, conf.STATUS_CREATED, respBook)
	} else {
		respBook.Message = "reject"
		respBook.Success = "1"
		respondJSON(w, conf.STATUS_INTERNAL_SERVER_ERROR, respBook)
	}
}
