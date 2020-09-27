package main

import (
	"bb-pan/src/handler"
	"fmt"
	"net/http"
)

func main(){
	 http.HandleFunc("/file/upload",handler.UploadHandler)
	 http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	 http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	 http.HandleFunc("/file/download",handler.DownloadHandler)
	 http.HandleFunc("/file/update",handler.FileUpdateMetaHandler)
	 http.HandleFunc("/file/delete",handler.FileDelHandler)


	 http.HandleFunc("/user/signup",handler.SignupHandler)
	 http.HandleFunc("/user/signin",handler.SignInHandler)
	 http.HandleFunc("/user/info",handler.HTTPInterceptor( handler.UserInfoHandler))


	// 秒传接口
	http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(
		handler.TryFastUploadHandler))

	http.HandleFunc("/file/downloadurl", handler.HTTPInterceptor(
		handler.DownloadURLHandler))

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init",
		handler.HTTPInterceptor(handler.InitialMultipartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart",
		handler.HTTPInterceptor(handler.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete",
		handler.HTTPInterceptor(handler.CompleteUploadHandler))




	 err := http.ListenAndServe(":8080",nil)
	 if err != nil{
	 	fmt.Printf("Failed to start server,err:%s",err.Error())
	 }
}
