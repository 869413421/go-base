package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type DataBucket struct {
	buffer *bytes.Buffer // 字节缓冲区
	mutex  *sync.RWMutex // 互斥锁
	cond   *sync.Cond    // 条件变量
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

// Put 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()         // 打开写锁
	defer db.mutex.Unlock() // 结束关闭写锁

	//写入数据
	n, err := db.buffer.Write(d)

	// 写入数据后通过 Signal 通知处于阻塞状态的读取器已经完成写操作
	db.cond.Signal()
	return n, err
}

func (db *DataBucket) Read(i int) {
	// 加读取锁
	db.mutex.RLock()
	// 结束后释放读锁
	defer db.mutex.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		//每次读取一个字节
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF {
				//读取完成缓冲区为空
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.cond.Wait()  // 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				data = data[:0] // 将 data 清空
				continue
			}
		}
		data = append(data, d) // 将读取到的数据添加到 data 中
	}
}

func main() {
	db := NewDataBucket()
	go db.Read(1)
	go func(i int) {
		d := fmt.Sprintf("data-%d", i)
		db.Put([]byte(d))
	}(1)
	time.Sleep(100 * time.Millisecond)
}
