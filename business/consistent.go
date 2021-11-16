package business

// 一致性哈希环CRUD接口
// 基于consistentBase.go的数据结构算法
// 这里用于存放服务器节点信息
// 方便查询使用

func NewConsistent() *Consistent {
	return &Consistent{
		// hash环 key 为 hash值 value 为节点的信息
		Circle:       make(map[uint32]string),
		VirtualNode:  20,
	}
}

// 向hash环添加一个节点
func (c *Consistent) Add(element string) {
	c.Lock()
	defer c.Unlock()
	c.add(element)
}

// 删除一个节点
func (c *Consistent) Remove(element string) {
	// 枷锁
	c.Lock()
	defer c.Unlock()
	c.remove(element )
}

// 获取节点
// 根据数据标识 获取最近服务器节点信息
// 服务器IP的哈希值跟用户ID的哈希值比较，大于用户ID的服务器节点中的最小值
// 以最接近用户ID哈希值的服务器为准
func (c *Consistent) Get(name string) (string,error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.Circle) == 0 {return "",errEmpty}
	// 生成hash值
	key := c.hashKey(name)
	// 根据hash值查找最近的节点 如果超出范围 就归零
	i := c.search(key)
	return c.Circle[c.SortedHashes[i]],nil
}