package main

import(
     "fmt"
     "math/rand"
     "time"                                  
)

func main() {
     var u int

     nano := time.Now().Nanosecond()              
     rand.Seed( int64(nano) )  

     i1 := time.Now()                                            //把for语句开始循环的时间放入i1变量。
     for i := 0; i < 100; i++ {
          r := rand.Int() % 25 + 65

          fmt.Printf("%d What is the order of %c: ", i + 1, r)
          fmt.Scanf("%d\n", &u)

          if u == r - 64 {
               fmt.Print("You are right. \n")
          } else {
               fmt.Printf("You are wrong, the correct answer is %d. \n", r - 64)
          }
     }
     i2 := time.Now()                                            //把for语句循环结束的时间放入i2变量。
     i3 := i2.Sub(i1)                                              //把i2减去i1的结果放入i3变量盒子。Sub：减。
     seconds := i3.Seconds()                               //把i3变量里的时间转化为秒数放入seconds变量。
     fmt.Printf("It took %v minute %v seconds. \n", int(seconds)/60, int(seconds)%60)  //算出分钟和秒数的方法
     fmt.Scanf("%d", &u)
}
