package main

import(
  "time"
  "log"
)

func main(){
  START:
    time.Sleep(time.Minute*15)
    log.Println("\a")
    log.Println("\a")
    log.Println("\a")
    log.Println("\a")
    log.Println("\a")

    time.Sleep(time.Minute)
  goto START

}
