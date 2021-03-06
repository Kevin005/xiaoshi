package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"xiaoshi/conf"
	"xiaoshi/model"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *conf.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.DBName,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
		return
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/xiaoshi/nsq/msg", a.GetNsqMsg)
	a.Get("/xiaoshi/feedback", a.GetAllFeedback)
	/** 意见反馈 */
	a.Post("/xiaoshi/feedback", a.CreateFeedback)
	a.Get("/xiaoshi/feedback", a.GetAllFeedback)
	/** 用户 */
	a.Post("/xiaoshi/avatar", a.UploadAvatar)
	a.Get("/xiaoshi/avatar/{name}", a.GetAvatar)
	a.Post("/xiaoshi/user/register", a.Register)
	a.Post("/xiaoshi/user/login", a.Login)
	a.Post("/xiaoshi/user/edit_pwd", a.EditPwd)
	a.Post("/xiaoshi/user/edit", a.EditUserInfo)
	/** 书籍 */
	a.Post("/xiaoshi/book", a.CreateBook)
	a.Get("/xiaoshi/books/{user_id}", a.getMyBooks)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

/**
反馈
 */
func (a *App) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	CreateFeedback(a.DB, w, r)
}

/**
查询所有反馈
 */
func (a *App) GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	GetAllFeedback(a.DB, w, r)
}

/**
上传用户头像
 */
func (a *App) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	UploadAvatar(w, r)
}

/**
获取用户头像
 */
func (a *App) GetAvatar(w http.ResponseWriter, r *http.Request) {
	GetAvatar(w, r)
}

/**
用户注册
 */
func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	Register(a.DB, w, r)
}

/**
用户登录
 */
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	Login(a.DB, w, r)
}

/**
修改密码
 */
func (a *App) EditPwd(w http.ResponseWriter, r *http.Request) {
	EditPwd(a.DB, w, r)
}

/**
修改用户信息
 */
func (a *App) EditUserInfo(w http.ResponseWriter, r *http.Request) {
	EditUserInfo(a.DB, w, r)
}

/**
发布一本书
 */
func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook(a.DB, w, r)
}

/**
获取我所有在读的书
 */
func (a *App) getMyBooks(w http.ResponseWriter, r *http.Request) {
	GetMyBooks(a.DB, w, r)
}

// just test GET
func (a *App) GetNsqMsg(w http.ResponseWriter, r *http.Request) {
	GetNsqMsg(w, r)
}
