package main
 
import (
    "fmt"
    "time"
    "os"
    "math/rand"
)

 
func main() {   
    

    r := rand.New( rand.NewSource(99) )
    fmt.Println( r.Intn(10) )
    //time.Sleep(10000 * time.Millisecond)
    fmt.Println(os.Getppid())

    fmt.Println(os.Getpagesize())

    //时间戳
    t := time.Now().Unix()
    fmt.Println(t)
     
    //时间戳到具体显示的转化
    fmt.Println(time.Unix(t, 0).String()+"\n====")
     
    //带纳秒的时间戳
    t = time.Now().UnixNano()
    fmt.Println(t)
    fmt.Println("------------------")
    
    for i :=0; i< 11; i++ {
        fmt.Println(i%10);
    }
    
    t2 := time.Now().UnixNano();
    
    fmt.Println(t2-t)
    fmt.Println(t2)
 
    //基本格式化的时间表示
    fmt.Println(time.Now().String())
     
    fmt.Println(time.Now().Format("2006year 01month 02day"))
 
}
