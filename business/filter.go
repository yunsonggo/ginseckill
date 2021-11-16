package business

import (
	"fmt"
	"net/http"
	"strings"
)

// cookie独立验证的中间件
// 声明一个函数类型
type FilterHandle func(rw http.ResponseWriter,req *http.Request) error

// 存储需要拦截的URL
type Filter struct {
	// key:url value:对此url的验证方法
	FilterMap map[string]FilterHandle
}

func NewFileter() *Filter {
	return &Filter{
		FilterMap: make(map[string]FilterHandle),
	}
}

// 注册拦截器中间件
func (f *Filter) RegisterUri(url string,handle FilterHandle) {
	f.FilterMap[url] = handle
}

// 获取handle
func (f *Filter) GetHandle(url string) FilterHandle {
	return f.FilterMap[url]
}

// 声明一个函数类型 方便功能扩展
type WebHandle func(rw http.ResponseWriter,req *http.Request)

// 这里实现对请求的URL是否拦截的方法
// 传入WebHandle类型 返回通用函数 避免返回数据类型error而受限
func (f *Filter) Handle(webHandle WebHandle) func(rw http.ResponseWriter,req *http.Request) {
	fmt.Println("执行业务逻辑")
	return func(rw http.ResponseWriter,req *http.Request) {
		for path,handle := range f.FilterMap {
			fmt.Printf("path:%v,handle:%v\n",path,handle)
			// 如果请求的URL 在我们存储的FilterMap里 就进行拦截
			// path 在 req.RequestURI中
			if strings.Contains(req.RequestURI,path) {
				// 执行拦截业务 注册拦截器时传入的处理函数 也是map中的value
				err := handle(rw,req)
				if err != nil {
					_,_ = rw.Write([]byte(err.Error()))
					return
				}
			}
		}
		// 请求过滤以后执行真正的业务函数 执行传入的正常注册的函数
		fmt.Println("执行业务逻辑")
		webHandle(rw,req)
	}
}