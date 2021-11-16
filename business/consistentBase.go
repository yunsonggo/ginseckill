package business

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// 一致性哈希环数据结构

// 声明uints切片 存储哈希值
type uints []uint32

// uints类型内置方法:返回切片长度
// (实现sort包中的Sort接口,需要这三个方法)
func (x uints) Len() int {
	return len(x)
}
// uints类型内置方法:	比较俩值大小
// (实现sort包中的Sort接口,需要这三个方法)
func (x uints) Less(i,j int) bool {
	return x[i] < x[j]
}
// uints类型内置方法:	交换两值
// (实现sort包中的Sort接口,需要这三个方法)
func (x uints) Swap(i,j int) {
	x[i],x[j] = x[j],x[i]
}

var errEmpty = errors.New("hash环没有数据")

// 哈希环实体结构体
type Consistent struct {
	// hash环 key 为 hash值 value 为节点的信息
	Circle map[uint32]string
	// 排序后的哈希值key切片
	SortedHashes uints
	// 设置虚拟节点数量 用于平衡hash环节点
	VirtualNode int
	// 读写锁
	sync.RWMutex
}

// 生成hash环中的key 节点信息拼接虚拟节点排位
func (c *Consistent) generateKey(element string,index int) string {
	return element + strconv.Itoa(index)
}

// 根据generateKey计算hash值做为key
func (c *Consistent) hashKey(key string) uint32 {
	// 不够64位 使用copy函数 填充到64位字节数组
	if len(key) < 64 {
		var srcatch [64]byte
		copy(srcatch[:],key)
		// 国际标准 多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

// 顺时针查找节点在排序后的c.SortedHashes[uints]位置
func (c *Consistent) search(key uint32) int {
	f := func(x int) bool {
		return c.SortedHashes[x] > key
	}
	// 使用“二分查找”查找老搜索指定切片满足条件的最小值
	// 满足f返回为true 的 最小值
	i := sort.Search(len(c.SortedHashes),f)
	// 超出范围 就归零 回归哈希环首位
	if i >= len(c.SortedHashes) {
		i = 0
	}
	return i
}

// 添加节点信息
func (c *Consistent) add(element string) {
	// 循环虚拟节点 设置副本
	for i := 0; i < c.VirtualNode; i++ {
		key := c.hashKey(c.generateKey(element,i))
		c.Circle[key] = element
	}
	// 添加完节点后更新排序
	c.updateSortedHashes()
}
// 对添加的节点信息 进行排序 方便查找使用
func (c *Consistent) updateSortedHashes() {
	// 获取 0 个 等同于空切片
	hashes := c.SortedHashes[:0]
	// 判断切片容量 过大则重置
	if cap(c.SortedHashes) / (c.VirtualNode * 4) > len(c.Circle) {
		hashes = nil
	}
	// 添加已存入的节点
	for k := range c.Circle {
		hashes = append(hashes,k)
	}
	// 对hashes排序 方便查找使用
	sort.Sort(hashes)
	c.SortedHashes = hashes
}
// 删除节点
func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		delete(c.Circle,c.hashKey(c.generateKey(element,i)))
	}
	c.updateSortedHashes()
}
