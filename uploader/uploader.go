package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		ty := request.URL.Query().Get("type")
		if ty == "base64" {
			base64Encoding(writer, request)
			return
		}
		if ty == "parse" {
			b, _ := ioutil.ReadFile("test.txt")
			data, e := base64.StdEncoding.DecodeString(string(b))
			if e != nil {
				writer.Write([]byte(e.Error()))
				return
			}
			f, err := os.OpenFile("a.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			f.Write(data)
			writer.Write([]byte("ok"))
			return
		}
		if request.Method != "POST" {
			writer.Write([]byte("method not allowed"))
			return
		}
		//url := "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=eETzdH6LEFJrfWxxYmlW-AnaM8DjYcORJ61SFqoFyLE3tdJml0A1Xbp3S05R4Miw-kcawn5Dqs98FnnjFfy7_o_wwAVqAJu91sJG_NyokMD-hWvYLZ0hvpfiNkB3yTir5NAaD2uPfAik6hVXtRgoYZgruuYz-Lvhv54Iw93Ta_9BWNcCdTqEbScP2BvMnb4n7jeOFyTEyuEmq4vM7xIUQQ&type=image"
		file, _, e := request.FormFile("file")
		if e != nil {
			writer.Write([]byte("read file error: " + e.Error()))
			return
		}
		//fileContentType := h.Header.Get("Content-Type")
		//fileContentLength := h.Header.Get("Content-Length")
		//fmt.Println(h.Header)
		//writer.Write([]byte("content-type: " + fileContentType + ";" + fileContentLength + ";" + h.Filename + ";" + strconv.FormatInt(h.Size, 10)))
		//return

		b, e := ioutil.ReadAll(file)
		if e != nil {
			writer.Write([]byte(e.Error()))
			return
		}
		base64Str := base64.StdEncoding.EncodeToString(b)
		fmt.Println(base64Str)
		//f, err := os.Open("text.txt")
		f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write([]byte(base64Str))
		writer.Write([]byte("ok"))
		return
		//resp, e := fileUpload(header.Filename, file, url)
		//if e != nil {
		//	writer.Write([]byte(e.Error()))
		//	return
		//}
		//writer.Write(resp)
	})

	http.ListenAndServe(":8080", nil)
}

func base64Encoding(writer http.ResponseWriter, request *http.Request) {
	b, e := ioutil.ReadFile("test.txt")
	if e != nil {
		writer.Write([]byte("readfile fail:" + e.Error()))
		return
	}
	dist, _ := base64.StdEncoding.DecodeString(string(b))
	url := "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=G4VkDGZp5bhmD0rR8RawrLqBJFbhnRsEQCdCBQQyqtdY3InHCzs-3Oxe3lB6pnfGKXh9npAEc3pKeL7xige1lfh9tvN3JOzSE78OHuR7ZUoZ2wbyzmo0xirPoGB_IcIyp0yLMpSxjIJAYIXwVP0fmYaiQPtpKEGIBZshe_wlLGNMp7kVI1jVWo0F2bnZDy4ovMKZX8uwGTVNUZTGryUFPg&type=image"
	resp, e := fileUpload("a.jpg", bytes.NewReader(dist), url)
	if e != nil {
		writer.Write([]byte(e.Error()))
		return
	}
	writer.Write(resp)
	//f, err := os.OpenFile("a.jpg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//e = ioutil.WriteFile("a.jpg", dist, os.ModePerm)
	////defer f.Close()
	////_, e = f.Write(b)
	//if e != nil {
	//	writer.Write([]byte("write b fail" + e.Error()))
	//	return
	//}
}

func fileUpload(filename string, reader io.Reader, url string) ([]byte, error) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer) //创建multipart.Writer, 后续写入文件内容,存放在bodyBuffer中
	fileWriter, e := bodyWriter.CreateFormFile("file", filename)

	if e != nil {
		fmt.Println("CreateFormFile error", e.Error())
		return nil, e
	}
	io.Copy(fileWriter, reader)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, e := http.Post(url, contentType, bodyBuffer)
	if e != nil {
		fmt.Println("http.post fail", e.Error())
		return nil, e
	}
	defer resp.Body.Close()
	responseBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println("ioutil.ReadAll fail", e.Error())
		return nil, e
	}
	return responseBody, nil
}
