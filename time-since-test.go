package main
  
import (
    "fmt"
    "time"
)

func main() {
  
    t0 := time.Now() 
    t1 := time.Now() 
	since :=  time.Since(t0)

    fmt.Println("t0:",t0)	
    fmt.Println("t1:",t1)	
    fmt.Println("time.Since(t0):", since)

    for i := 0; i < 10; i++ {
        t := time.Now()   
		fmt.Println( t )
    }   


}
