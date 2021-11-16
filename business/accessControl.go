package business

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	// 负载均衡设置 一致性哈希算法
	hashConsistent = NewConsistent()
	localIp        string
	requestPort    string
	blackList      = &BlackList{List: make(map[int]bool)}
	// 延时
	interval = 10
)

// 添加节点 v 是服务器节点IP
func AddHashConsistent(localhost string, port string, hostArray []string) {
	localIp = localhost
	requestPort = port
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}
	return
}

// 访问控制
// 存储用户访问信息
type AccessControl struct {
	// 这里演示存储访问时间
	sourcesArray map[int]time.Time
	sync.RWMutex
}

// 设置记录
func (ac *AccessControl) SetNewRecord(uid int) {
	ac.RWMutex.Lock()
	ac.sourcesArray[uid] = time.Now()
	ac.RWMutex.Unlock()
}

// 获取记录
func (ac *AccessControl) GetNewRecord(uid int) time.Time {
	ac.RWMutex.RLock()
	defer ac.RWMutex.RUnlock()
	return ac.sourcesArray[uid]
}

func (ac *AccessControl) GetDistributedRight(req *http.Request) bool {
	// 获取用户ID
	uid, err := req.Cookie("uid")
	if err != nil {
		return false
	}
	// 采用一致性hash算法，根据用户ID,判断获取具体机器
	// 服务器节点哈希值大于用户ID哈希值中的最小值为分配节点
	hostRequest, err := hashConsistent.Get(uid.Value)
	if err != nil {
		return false
	}
	// 判断是否为本机
	if hostRequest == localIp {
		// 执行本机数据读取和校验
		return ac.GetDataFromMap(uid.Value)
	} else {
		// 不是本机充当代理访问数据返回结果
		return ac.GetDataFromOtherMap(hostRequest, req)
	}
}

// 获取本机map， 并且处理业务逻辑，返回的结果类型为bool类型
func (ac *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	// 查询是否黑名单
	if blackList.GetBlackListByID(uidInt) {
		return false
	}
	// 获取用户访问信息
	dataRecord := ac.GetNewRecord(uidInt)
	// 不是第一次访问
	if !dataRecord.IsZero() {
		if dataRecord.Add(time.Duration(interval) * time.Second).After(time.Now()) {
			// 这里设置 10 秒后删除该用户时间记录并返回false
			// 可以再次请求
			delete(ac.sourcesArray, uidInt)
			return false
		}
	}
	// 是第一次访问 设置访问时间信息
	ac.SetNewRecord(uidInt)
	return true
}

// 获取其他服务器map 并且处理业务逻辑，返回的结果类型为bool类型
func (ac *AccessControl) GetDataFromOtherMap(host string, request *http.Request) bool {
	hostUrl := "http://" + host + ":" + requestPort + "/checkRight"
	resp, body, err := GetCurl(hostUrl, request)
	if err != nil {
		return false
	}
	// 判断状态
	if resp.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}
