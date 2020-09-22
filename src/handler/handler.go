package handler

import (
	"bb-pan/src/meta"
	"bb-pan/src/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		file,head,err := r.FormFile("file")
		if err != nil{
			fmt.Printf("Failed to get data,err:%s\n",err.Error())
			return
		}

		defer file.Close()


		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"D:/tmp/"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile,err := os.Create(fileMeta.Location)

		if err != nil{
			fmt.Printf("Failed to create fle,err:%s\n",err)
			return
		}
		defer newFile.Close()

		fileMeta.FileSize,err = io.Copy(newFile,file)
		if err != nil{
			fmt.Printf("Failed to sae data into file,err:%s",err.Error())
			return
		}

		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		//meta.UpdateFileMeta(fileMeta)

		_ = meta.UpdateFileMetaDB(fileMeta)

		//上传成功跳转
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

//上传完成操作
func UploadSucHandler(w http.ResponseWriter, r * http.Request){
	io.WriteString(w,"Upload finished")
}


func GetFileMetaHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(filehash)
	fMeta,err := meta.GetFileMetaDB(filehash)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data,err := json.Marshal(fMeta)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)


}


//下载
func DownloadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fsha1 :=r.Form.Get("filehash")

	fm := meta.GetFileMeta(fsha1)

	f,err := os.Open(fm.Location)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data,err := ioutil.ReadAll(f)

	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("content-disposition","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)

}


//修改文件元信息
func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method != "POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data,err := json.Marshal(curFileMeta)
	if err !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}


//删除文件
func FileDelHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	fileSha1 := r.Form.Get("filehash")

	fMeta := meta.GetFileMeta(fileSha1)

	os.Remove(fMeta.Location)

	meta.RemoveFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)
}
























