package main

import (
	"fmt"
	"github.com/yunsonggo/loggo"
	"net/http"
	"sync"
)

var sum int64 = 0

// 预存商品数量
var productNum int64 = 100

// 互斥锁
var mutex sync.Mutex

// 获取秒杀商品
func getOneProduct() bool {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	// 判断数据是否超限
	if sum < productNum {
		sum += 1
		fmt.Println("商品已抢购", sum)
		return true
	}
	return false
}
// 失败返还
func backOne() bool {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	// 判断数据是否超限
	if 0 < sum && sum <= productNum {
		sum -= 1
		fmt.Println("商品抢购出错,sum数量返还1", sum)
		return true
	}
	return false
}

func GetProduct(w http.ResponseWriter, req *http.Request)  {
	if getOneProduct() {
		_,_ = w.Write([]byte("true"))
		return
	}
	_,_ = w.Write([]byte("false"))
	return
}

func BackOneProduct(w http.ResponseWriter, req *http.Request)  {
	if backOne() {
		_,_ = w.Write([]byte("true"))
		return
	}
	_,_ = w.Write([]byte("false"))
	return
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	http.HandleFunc("/backOne",BackOneProduct)
	err := http.ListenAndServe(":8084", nil)
	if err!=nil {
		loggo.Error("Err:", err)
	}
	fmt.Printf("listenAndServer:%s\n","0.0.0.0:8084")
}
