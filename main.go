package main

import(
  "fmt"
  "github.com/dchest/uniuri"
  "time"
  "sync"
)

func main(){
  finished := make(chan bool)
  mu := new(sync.Mutex)
  guessedSet := map[string]bool{}
  maxRoutines := 20
  password := "luoo"
  start := time.Now()

  for i:=0; i < maxRoutines; i++{
    go crackPassword(password, finished, guessedSet, mu)
  }
  <-finished

  mu.Lock()
  numberGuesses := len(guessedSet)
  mu.Unlock()

  timeTaken := time.Now().Sub(start)
  fmt.Println("Number of go routines:", maxRoutines)
  fmt.Println("The time elapsed:", timeTaken)
  fmt.Println("Total number of guesses:", numberGuesses)
  fmt.Println("The guesses per milliseconds:", float64(numberGuesses)/float64(timeTaken.Milliseconds()))
}

func crackPassword( password string, finished chan bool, guessedSet map[string]bool, mu *sync.Mutex){
  //fmt.Println("Starting new routine!")
  lenPassword := len(password)

  guess := uniuri.NewLen(lenPassword)
  for guess != password {
   guess = uniuri.NewLen(lenPassword)
   mu.Lock()
   guessedSet[guess] = true
   mu.Unlock()
  }
  finished <- true
}
