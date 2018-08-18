package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/app/model/response"
	"xiaoshi/app/model"
	"encoding/json"
	"xiaoshi/app/model/request"
	"github.com/gorilla/mux"
)

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-AccessToken")
	respBook := response.RespBook{}
	if hadToken, user := checkToken(db, token); hadToken {
		//get request book
		reqBook := request.ReqBooks{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&reqBook); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		defer r.Body.Close()
		//开启事务
		tx := db.Begin()
		//save book
		book := model.Books{}
		book.BookName = reqBook.BookName
		book.Author = reqBook.Author
		book.PageTotal = reqBook.PageTotal
		book.Image = reqBook.Image
		if err := tx.Save(&book).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			//回滚事务
			tx.Rollback()
			return
		}
		myBook := model.MyBooks{}
		myBook.BookId = book.ID
		myBook.UserId = user.ID
		myBook.Private = reqBook.Private
		myBook.Progress = "0"
		if err := tx.Save(&myBook).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			//回滚事务
			tx.Rollback()
			return
		}
		//提交事务
		tx.Commit()
		respBookData := response.RespBookData{}
		respBookData.UserId = user.ID
		respBookData.BookId = book.ID
		respBookData.BookAuthor = book.Author
		respBookData.BookName = book.BookName
		respBook.Data = respBookData
		respBook.Message = "pass"
		respBook.Success = "0"
		respondJSON(w, http.StatusCreated, respBook)
	} else {
		respBook.Message = "reject"
		respBook.Success = "1"
		respondJSON(w, http.StatusInternalServerError, respBook)
	}
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
		respondJSON(w, http.StatusCreated, respBook)
	} else {
		respBook.Message = "reject"
		respBook.Success = "1"
		respondJSON(w, http.StatusInternalServerError, respBook)
	}
}
