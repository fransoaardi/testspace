package mapsyncplay

import (
	"sync"
)

type MMap struct{
	sync.Mutex
	data map[string]string
}

type RMMap struct{
	sync.RWMutex
	data map[string]string
}


type NoMMap struct{
	data map[string]string
}

func NewRMMap() *RMMap{
	return &RMMap{
		data: make(map[string]string),
	}
}

func NewNoMMap() *NoMMap{
	return &NoMMap{
		data: make(map[string]string),
	}
}

func NewMMap() *MMap{
	return &MMap{
		data: make(map[string]string),
	}
}

//func main(){
//
//	var target *MMap
//
//	a := NewMMap()
//	b := NewMMap()
//
//	a.Set("one", "1")
//	b.Set("two", "2")
//
//	target = a
//
//	go func(){
//		for i:=0; i<1000000; i++ {
//			if !checkEqual(target.Get("one"),"1" ){
//				fmt.Println("ERR", i)
//				break
//			}
//		}
//		fmt.Println("DONE")
//	}()
//
//	go func(){
//		target = b
//	}()
//
//	time.Sleep(time.Second)
//}
//
//func checkEqual (val1, val2 string) bool{
//	return val1 == val2
//}

func (n *NoMMap)Get(key string) string{
	return n.data[key]
}

func (m *RMMap)Get(key string) string{
	m.RLock()
	defer m.RUnlock()
	return m.data[key]
}

func (m *MMap)Get(key string) string{
	m.Lock()
	defer m.Unlock()
	return m.data[key]
}

func (m *MMap)Set(key, val string) {
	m.Lock()
	m.data[key] = val
	m.Unlock()
}





