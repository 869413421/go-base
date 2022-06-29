package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type DataBucket struct {
	buffer *bytes.Buffer
	mutex  *sync.RWMutex
	cond   *sync.Cond
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

func (db *DataBucket) Put(b []byte) (int, error) {
	// 写入时加锁，完成时解锁
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n, err := db.buffer.Write(b)

	// 向所有阻塞的携程发送通知
	db.cond.Broadcast()
	return n, err
}

func (db *DataBucket) Read(i int) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	var data = make([]byte, 0)
	var d byte
	var err error
	for {
		d, err = db.buffer.ReadByte()
		if err != nil {
			if err == io.EOF {
				//读取完成缓冲区为空
				if string(data) != "" {
					fmt.Printf("reader-%d: %s \n", i, data)
				}
				db.cond.Wait()  // 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				data = data[:0] // 将 data 清空
				continue
			}
		}
		data = append(data, d)
	}
}

func main() {
	db := NewDataBucket()
	for i := 1; i < 3; i++ { // 启动多个读取器
		go db.Read(i)
	}
	for j := 0; j < 100; j++ { // 启动多个写入器
		d := fmt.Sprintf("data-%d", j)
		go db.Put([]byte(d))

	}
	time.Sleep(100 * time.Millisecond) // 每次启动一个写入器暂停100ms，让读取器阻塞
}
