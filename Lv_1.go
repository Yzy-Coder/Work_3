package main

import "fmt"

const FliterSize = 1000


//本次算法基本思想为利用素数作为筛子来筛选新的素数，如7不能被2、3、5整除，而2、3、5是素数，所以7也是素数，再将7加入筛子，继续筛选
//初始化函数
func Initialization ( Maxnum int) []int {

	//存放等待被判断的数的管道
	In := make(chan int)

	//用来通知协程结束运行的管道
	Final := make(chan int)

	//用来存放素数的切片
	PrimeNumber := make([]int , 0 , 10)

	//创建协程
	go OutPut ( In , Final , &PrimeNumber )


	In <- 2
	for i := 3 ; i < Maxnum ; i = i + 2{

		In <- i

	}

	In <- -1
	<- Final
	return PrimeNumber
}

//输出协程，判断和存放最新的一组素数，当存满一组后，生成一个筛子（Fliter）协程
func OutPut( In chan int , Final chan int , PrimeNumber *[]int) {
	for {

		n := <- In
		if n == -1 {
			Final <- 0
			return
		}

		//利用当前最新一组的素数来判断最新的一个数是不是素数
		for i := len(*PrimeNumber) / FliterSize * FliterSize ; i < len(*PrimeNumber) ; i++ {
			if  n%(*PrimeNumber)[i] == 0 {
				goto end
			}
		}

		//如果是素数，加入素数数组
		*PrimeNumber = append(*PrimeNumber , n)
		fmt.Println(n , "为素数")

		//当长度够新建一个组时，生成一个筛子协程，并将这一组素数素组分配给它
		if 	len(*PrimeNumber)%FliterSize == 0 {
			NewIn := make(chan int)
			go Fliter( In , NewIn , (*PrimeNumber)[ (len(*PrimeNumber)/FliterSize -1) * FliterSize :])
			In = NewIn
		}
		end: 
	}
}

//筛子协程
func Fliter(In chan int, Out chan int, ThisPrimeNumber []int ) {
	for true {
		n := <- In
		for _ , Sieve := range ThisPrimeNumber{

			if n%Sieve == 0 {
				goto end
			}
		}

		Out <- n
	end:
	}
}

func main() {
	Initialization(50000)
}