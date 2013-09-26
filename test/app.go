package main

import (
        "fmt"
        "labix.org/v2/mgo"
        "labix.org/v2/mgo/bson"
	    "strconv"
	    "time"
        "os"
)

type Person struct {
        Name string
        Phone string
	    Cinums  int
        Insertime string
        ProPid int
}



func toString(a interface{}) string{
	
	 if  v,p:=a.(int);p{
	 	return strconv.Itoa(v)
	 }
	
	if v,p:=a.(float64);  p{
	 return strconv.FormatFloat(v,'f', -1, 64)
	}
	
	if v,p:=a.(float32); p {
		return strconv.FormatFloat(float64(v),'f', -1, 32)
	}
	
	 if v,p:=a.(int16); p { 
	 	return strconv.Itoa(int(v))
	 }
	  if v,p:=a.(uint); p { 
	 	return strconv.Itoa(int(v))
	 }
	  if v,p:=a.(int32); p { 
	 	return strconv.Itoa(int(v))
	 }
	return "wrong"
}


func main() {

        //计算运行时间
        i1 := time.Now()

        session, err := mgo.Dial("192.168.212.161:40201,192.168.212.161:40202")
        if err != nil {
                panic(err)
        }
        defer session.Close()
	
        fmt.Println("run:");

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")

	
	    for i :=0; i< 1000000; i++ {
            t := time.Now().Unix()
            err = c.Insert(&Person{"Ale"+toString(i), "+55 53 8116 9639", i, time.Unix(t, 0).String(), os.Getppid()},
	               &Person{"eric"+toString(i), "+86 15618388681", i, time.Unix(t, 0).String(), os.Getppid()})
            if err != nil {
                panic(err)
            }
            if( (i%1000)==0 ) {	
                result := Person{}
                err = c.Find(bson.M{"name": "eric"+toString(i)}).One(&result)
                if err != nil {
                    panic(err)
                }
                fmt.Println("Cinums:", result.Cinums)
                fmt.Println(toString(i)+" ：" + time.Now().String()) 
            }
	    }

        //result := Person{}
        //err = c.Find(bson.M{"name": "eric"}).One(&result)
        //if err != nil {
        //        panic(err)
        //}

        //fmt.Println("Phone:", result.Phone)

        i2 := time.Now() 
        i3 := i2.Sub(i1)
        seconds := i3.Seconds()
        fmt.Printf("It took %v minute %v seconds. \n", int(seconds)/60, int(seconds)%60) 
        

}

