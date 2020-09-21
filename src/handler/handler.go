package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
	data,err :=	ioutil.ReadFile("D:/go-work/bb-pan/src/static/view/index.html")
	if err != nil{
		io.WriteString(w,"internel server error:"+err.Error())
		return
	}
	io.WriteString(w,string(data))
	}else if r.Method == "POST"{

	}
}