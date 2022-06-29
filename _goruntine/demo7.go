
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Response struct {
	body   string        //响应内容
	elapse time.Duration //响应耗时
}

func request(timeout time.Duration) (Response, error) {
	//用于触发超时的计时器
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	start := time.Now()

	var respData Response
	var err error

	//同一个协程中执行IO操作
	respData.body = doIO()

	//如果超时计时器已处在就绪状态，则生成一个error的实例，否则继续往下执行。
	select {
	case <-timer.C:
		err = errors.New("request timeout1")
	default:
	}

	//记录总耗时
	respData.elapse = time.Now().Sub(start)

	return respData, err
}

func doIO() string{
	//随机产生一个[0,100)毫秒的延迟，以模拟IO延时延迟
	rand.Seed(time.Now().UnixNano())
	delay := time.Duration(rand.Intn(100)) * time.Millisecond
	fmt.Printf("delay=%v\n", delay)
	time.Sleep(delay)

	return "Hello World"
}

func main() {
	resp, err := request(50 * time.Millisecond)
	if err != nil {
		fmt.Printf("error: %s elapse=%s\n", err.Error(), resp.elapse)
		return
	}

	fmt.Printf("response: body=%s elapse=%s\n", resp.body, resp.elapse)
}