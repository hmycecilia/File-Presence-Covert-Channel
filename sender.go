package main
import (
    "os"
    "log"
    "fmt"
    "strconv"
    "bytes"
    "time"
)

const TRANSMIT string = "TX"
const DATA     string = "DATA"

func main() {
    os.Remove(TRANSMIT)
    inputString := "test"
    var messageString string
    var buffer bytes.Buffer

    for i := 0; i < len(inputString); i++ {
        buffer.WriteString(strconv.FormatInt(int64(inputString[i]), 2))
    }
    messageString = buffer.String()

    os.Create(TRANSMIT)

    for i := 0; i < len(messageString); i++ {
        for {
            t := time.Now()
            if t.Second() % 2 == 1 { //Write on the odd seconds
                //If bit is 1, 
                if string(messageString[i]) == "1" {
                    //touch DATA
                    os.Create(DATA)
                    fmt.Println("Create DATA")
                    fmt.Println("Sending: 1")
                    time.Sleep(1 * time.Second)
                    break
                } else if string(messageString[i]) == "0" {
                //else remove DATA
                    os.Remove(DATA)
                    fmt.Println("Remove DATA")
                    fmt.Println("Sending: 0")
                    time.Sleep(1 * time.Second)
                    break
                } else {
                    l := log.New(os.Stderr, "", 0)
                    l.Println("Character is not 1 or 0.  Something went horribly wrong.")
                    os.Exit(1)
                }
            } else { //Sleep on the even seconds
                continue
            }
        }
    }
    os.Remove(TRANSMIT)
    os.Remove(DATA)
}
