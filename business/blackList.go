package business

import "sync"

// 设置黑名单 存储恶意uid
// 可以通过前端过滤筛查后提交指定接口添加
type BlackList struct {
	List map[int]bool
	sync.RWMutex
}

// 获取黑名单
func (m *BlackList)GetBlackListByID(uid int) bool  {
	m.RLock()
	defer m.RUnlock()
	return m.List[uid]
}

// 添加黑名单
func (m *BlackList) SetBloackListByID(uid int) bool {
	m.Lock()
	defer m.Unlock()
	m.List[uid] = true
	return true
}
