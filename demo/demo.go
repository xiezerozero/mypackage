package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/go-redis/redis/v7"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("readFile fail: ", err.Error())
	}
	fmt.Println("read bytes: ", len(b))
	return string(b)
}

func compress() {
	originalStr := readFile("compact_data.json")

	// 压缩字符串
	var compressedData bytes.Buffer
	writer := gzip.NewWriter(&compressedData)
	_, err := writer.Write([]byte(originalStr))
	if err != nil {
		fmt.Println("Failed to compress data:", err)
		return
	}
	writer.Close()

	redisErr := connectLocalRedis().Set("test", compressedData.Bytes(), 0).Err()
	if redisErr != nil {
		fmt.Println("redis set fail: ", redisErr.Error())
	}
}

func decompress() {
	b, err := connectLocalRedis().Get("test").Bytes()
	if err != nil {
		fmt.Println("redis get fail: ", err.Error())
		return
	}
	fmt.Println("read bytes: ", len(b))
	reader, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Failed to create gzip reader:", err)
		return
	}
	decompressedData, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("Failed to decompress data:", err)
		return
	}
	reader.Close()
	fmt.Println("Decompressed data:", string(decompressedData))
}

func requestGzip() {
	client := new(http.Client)

	request, err := http.NewRequest("GET", "http://127.0.0.1:8198/3rd/gprograms?token=d1606c8aca9fab32583d1b26d8b1fe3d&from=171066&tm=1690859374503&date=20230731&cid=zj-ckvz4pkryuio", nil)
	request.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to get url:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	rawByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("read raw bytes: ", len(rawByte))
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		//reader, err = gzip.NewReader(resp.Body)
		reader, err = gzip.NewReader(bytes.NewReader(rawByte))
		defer reader.Close()
	default:
		fmt.Println("not gzip")
		//reader = resp.Body
		reader = io.NopCloser(bytes.NewReader(rawByte))
	}

	n, _ := io.Copy(os.Stdout, reader) // print html to standard out
	fmt.Println("ungzip read bytes: ", n)
}

func UnGzip(compressContent string) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader([]byte(compressContent)))
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	decompressedData, err := ioutil.ReadAll(zr)
	if err != nil {
		fmt.Println("Failed to decompress data:", err)
		return nil, err
	}
	return decompressedData, nil
}

func main() {
	requestGzip()
}

func connectLocalRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "47.95.250.60:6799",
		Password: "reEW#@323Vfeoiv0lkaaa",
		DB:       11,
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	return rdb
}
