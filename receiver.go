package main

import (
    "os"
    "bytes"
    "fmt"
)

const DATA_READ       = "DRD"
const DATA            = "DAT"
const DATA_SEND_READY = "DSR"

func main() {
    var bitRead string
    var buffer bytes.Buffer

    //Remove DATA_READ
    os.Remove(DATA_READ)
    //Loop
    for {
        //loop
        for {
            //If DATA_SEND_READY exists
            if _, err := os.Stat(DATA_SEND_READY); err != nil {
                break
            }
        }
        //Read data "If DATA exists, bit is 1
        if _, err := os.Stat(DATA); err != nil {
            bitRead = "1"
        } else {
        //else, bit is zero
            bitRead = "0"
        }

        //Touch DATA_READ
        os.Create(DATA_READ)
        //loop
        for {
            //if DATA_SEND_READY exists
            if _, err := os.Stat(DATA_SEND_READY); err != nil {
                continue
            } else {
                break
            }
        }

        //Write bit to bitstream
        buffer.WriteString(bitRead)
        //print the bit that was received
        fmt.Println("Bit Received:", bitRead)
        fmt.Println("Cumulative Message:", buffer.String())
        
        //Remove DATA_READ
        os.Remove(DATA_READ)
    }
}
