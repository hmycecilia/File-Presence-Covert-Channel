package main
import (
    "fmt"
    "time"
    "os"
    "bytes"
)

const TRANSMIT string = "TX"
const DATA     string = "DATA"

func main() {
    var bitRead string
    var buffer bytes.Buffer

    //check for Transmit flag
    for {
        if _, err := os.Stat(TRANSMIT); err == nil {
            break
        }
    }

    for {
        t := time.Now()
        if t.Second() % 2 == 0 { //Read on the read seconds
            //If bit is 1, 
            if _, err := os.Stat(DATA); os.IsNotExist(err) {
                bitRead = "0"
            } else {
                bitRead = "1"
            }

            time.Sleep(1 * time.Second)
            if _, err :=os.Stat(TRANSMIT); os.IsNotExist(err) {
                fmt.Println("Transmission stopped")
                break
            }
            buffer.WriteString(bitRead)
            //print the bit that was received
            fmt.Println("Bit Received:", bitRead)
            fmt.Println("Cumulative Message:", buffer.String())
        }
    }
}
