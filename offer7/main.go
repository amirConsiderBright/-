package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var filedir = "D:/golandWorkapace/src/cn.basecommon/offer/offer7/string.json"

func main() {
	http.HandleFunc("/read", Read)
	http.HandleFunc("/writer", Writer)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Read(w http.ResponseWriter, r *http.Request) {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open(filedir)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	//split := strings.Split(string(content), "^^")
	w.Write(content)

}
func Writer(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	get := data.Get("str")
	if get == "" {
		w.Write([]byte("操作失败"))
		return
	}
	err := ioutil.WriteFile(filedir, []byte(get), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte("修改成功"))

}
