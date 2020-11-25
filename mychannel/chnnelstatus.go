package mychannel

import "log"

/*
channel可进行3种操作：
1. 读
2. 写
4. 关闭
把这3种操作和3种channel状态可以组合出9种情况：
操作		nil的channel	正常channel	        已关闭channel
<- ch		阻塞			成功或阻塞	         读到零值
ch <-		阻塞			成功或阻塞	         panic
close(ch)	panic		    成功	           panic
有1个特殊场景：当nil的通道在select的某个case中时，这个case会阻塞，但不会造成死锁。
*/
/*
关于channel的实现原理，可以参考以下说明：
https://blog.csdn.net/u010853261/article/details/85231944
*/

func MychannelMain1() {

	// nil的channel
	go func() {
		var c chan int
		log.Println("<- ch开始阻塞")
		log.Println(<-c)
		log.Println("<- ch阻塞结束")
	}()

	go func() {
		var c chan int
		log.Println("ch <-开始阻塞")
		c <- 1
		log.Println("ch <-阻塞结束")
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("发现panic", err)
			}
		}()
		var c chan int
		log.Println("close(c) painc 开始")
		close(c)
		log.Println("close(c) painc 结束")
	}()

	// 已关闭channel
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("发现panic", err)
			}
		}()

		/*
			创建channel实际上就是在内存中实例化了一个hchan的结构体，并返回一个ch指针，
			我们使用过程中channel在函数之间的传递都是用的这个指针，这就是为什么函数传递中无需使用channel的指针，
			而直接用channel就行了，因为channel本身就是一个指针。
		*/
		c := make(chan int, 1)
		c <- 1
		log.Println("clsoe(c)")
		close(c)
		log.Println("first:", <-c)
		c <- 1
		log.Println("second:", <-c)
		defer func() {
			close(c)
		}()

	}()

	select {}

}
