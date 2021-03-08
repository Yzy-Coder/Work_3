package main

import (
	"fmt"
	"sync"
	"time"
)

type Context interface {

	NewContext()
	Done()

}


type MyContext struct {

	value	string				//存储的信息
	mu 		sync.Mutex			//确保安全
	done 	<- chan struct{}	//用来通知关闭的管道

}

//设置Context内的value
func (c *MyContext)	Setvalue (value string)  {

	c.value = value
}

//初始化Context的通知关闭管道
func (c *MyContext) Done() chan struct{} {
	c.mu.Lock()
	d := make(chan struct{})
	if c.done == nil {


		c.done = d
	}

	c.mu.Unlock()
	return d
}

func Test(TestContex MyContext, Ticker *time.Ticker) {

	for true {

		select {

		case <-TestContex.done:
			fmt.Println("协程收到关闭信号，协程关闭")
			return
		case <-Ticker.C:
			fmt.Println("协程内的消息为",TestContex.value)
		}
	}
}

func main() {

	var TestContext MyContext
	var Ticker = time.NewTicker(time.Second * 1)
	TestContext.Setvalue("这是一个测试Context")
	Cancle := TestContext.Done()
	go Test(TestContext,Ticker)
	time.Sleep(5*time.Second)
	close(Cancle)
	time.Sleep(2*time.Second)
}