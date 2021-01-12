package main

import(
  "time"
  "log"
  "os"
  "bufio"
  "io"
  "strconv"
  "strings"
)

func MustReadStdin() string {
  r := bufio.NewReader(os.Stdin)

  var in string
  for {
    var err error
    in, err = r.ReadString('\n')
    if err != io.EOF {
      if err != nil {
        panic(err)
      }
    }
    in = strings.TrimSpace(in)
    if len(in) > 0 {
      break
    }
  }
  log.Println("")

  return in
}

func main(){
  log.Printf("How many minutes to Sleep?:")

  x, err := strconv.Atoi( MustReadStdin() )
  if err != nil {
    panic(err)
  }

  START:
    for i := 0; i < x; i++{
      time.Sleep(time.Minute)
      log.Printf( strconv.Itoa(x-i) )
    }

    log.Println("\a")
    time.Sleep(time.Second)
    log.Println("\a")
    time.Sleep(time.Second)
    log.Println("\a")

    log.Printf("End of timer! One minute until the timer resetarts.")
    time.Sleep(time.Minute)
  goto START

}
