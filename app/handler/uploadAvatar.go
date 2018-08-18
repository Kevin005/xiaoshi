package handler

import (
	"net/http"
	"strings"
	"path"
	"errors"
	"os"
	"io"
	"io/ioutil"
)

func UploadAvatar(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	// 接收图片
	uploadFile, handle, err := req.FormFile("image")
	errorAvatarHandle(err, w)
	// 检查图片后缀
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext != ".jpg" && ext != ".png" {
		errorAvatarHandle(errors.New("只支持jpg/png图片上传"), w);
		return
	}
	// 保存图片
	os.MkdirAll("/home/kevin/Pictures/xiaoshi/avatar/", 0777)
	saveFile, err := os.OpenFile("/home/kevin/Pictures/xiaoshi/avatar/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666);
	errorAvatarHandle(err, w)
	io.Copy(saveFile, uploadFile);
	defer uploadFile.Close()
	defer saveFile.Close()
	// 上传图片成功
	w.Write([]byte("xiaoshi/avatar/" + handle.Filename));
}

func GetAvatar(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("/home/kevin/Pictures" + req.URL.Path)
	errorAvatarHandle(err, w);
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorAvatarHandle(err, w);
	w.Write(buff)
}

func errorAvatarHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
