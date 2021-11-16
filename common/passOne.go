package common


import "sync"

//计数
var sum int64 = 0
// 互斥锁
var mutex sync.Mutex
//获取一个
func PassOne(hasNum int64) bool {
	// 总数
	//var hasNum int64
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	// 判断是否超限
	if sum < hasNum {
		sum += 1
		return true
	}
	return false
}