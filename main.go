package main

import(
  "fmt"
  "github.com/dchest/uniuri"
  "time"
)

func main(){
  finished := make(chan bool)
  maxRoutines := 100000
  password := "alss"
  start := time.Now()

  for i:=0; i < maxRoutines; i++{
    go crackPassword(password, finished)
  }
  <-finished

  end := time.Now()
  fmt.Println("The time elapsed:", end.Sub(start))
}

func crackPassword( password string, finished chan bool){
  //fmt.Println("Starting new routine!")
  lenPassword := len(password)

  guess := uniuri.NewLen(lenPassword)
  for guess != password {
   guess = uniuri.NewLen(lenPassword)
   fmt.Println(guess)
  }
  finished <- true
}
