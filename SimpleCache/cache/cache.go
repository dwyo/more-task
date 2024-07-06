package cache

import (
	"fmt"
	"time"
)

type MCache interface {
	SetMaxMemory(size string)
	Set(key string, value interface{}, exp int)
	Get(key string) interface{}
	Del(key string)
	Exist(key string) bool
	Flush()
}

type memCache struct {
	// 最大内存
	MaxMemorySize int64
	// B,KB,MB,GB
	MaxMemorySizeStr string
	// 当前使用内存
	CurUseMemory int64

	Value map[string]mValue
}

type mValue struct {
	val interface{}
	exp time.Time
}

func NewMemoryCache() MCache {
	return &memCache{}
}

func (mc *memCache) SetMaxMemory(size string) {
	s, u := ParseSize(size)
	mc.MaxMemorySize = s
	mc.MaxMemorySizeStr = u
	// init container
	mc.Value = make(map[string]mValue)
}
func (mc *memCache) Set(key string, value interface{}, exp int) {
	v := &mValue{
		val: value,
		exp: time.Now().Add(time.Second * time.Duration(exp)),
	}
	//todo size cal
	//sizeof := int64(unsafe.Sizeof(v))
	sizeof := CalSize(key, v)
	fmt.Println("内存对比: ", mc.CurUseMemory+sizeof, mc.MaxMemorySize)
	if (mc.CurUseMemory + sizeof) > mc.MaxMemorySize {

		// 内存不足，删除OR 清除过期数据 OR LRU释放内存
		fmt.Println("memory no space!")
		return
	}
	mc.Value[key] = *v
}
func (mc *memCache) Get(key string) interface{} {
	value, ok := mc.Value[key]
	if ok {
		if value.exp.Before(time.Now()) {
			mc.Del(key)
			return nil
		}
	}
	return value.val
}
func (mc *memCache) Del(key string) {
	delete(mc.Value, key)
}

func (mc *memCache) Exist(key string) bool {
	if _, ok := mc.Value[key]; !ok {
		return false
	}
	return true
}

func (mc *memCache) Flush() {
	mc.Value = make(map[string]mValue)
	mc.CurUseMemory = 0
}
