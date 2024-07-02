package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	filename := "/Users/xiezero/code/go/src/mypackage/wechat/1572944593312.jpg"
	f, err := os.Open(filename)
	defer f.Close()
	fmt.Println("open error:", err)
	fileds := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: "media",
			Reader:    f,
			Filename:  f.Name(),
		},
	}
	b, err := PostMultipartForm("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=&type=image", fileds)
	fmt.Println("PostMultipartForm fail:", err)
	fmt.Println(string(b))
}

type MultipartFormField struct {
	IsFile    bool
	Fieldname string
	Reader    io.Reader
	Filename  string
}

// 文件上传处理
func PostMultipartForm(uri string, fields []MultipartFormField) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			if _, err = io.Copy(fileWriter, field.Reader); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.Fieldname)
			if e != nil {
				err = e
				return
			}
			if _, err = io.Copy(partWriter, field.Reader); err != nil {
				return
			}
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = io.ReadAll(resp.Body)
	return
}
