package main

import (
    "os"
    "log"
    "fmt"
    "strconv"
    "bytes"
)

const DATA_READ       = "DRD"
const DATA            = "DAT"
const DATA_SEND_READY = "DSR"

func main() {

    //convert bitstream into bits
    inputString := "test"
    var messageString string
    var buffer bytes.Buffer

    for i := 0; i < len(inputString); i++ {
        buffer.WriteString(strconv.FormatInt(int64(inputString[i]), 2))
    }
    messageString = buffer.String()

    //Display data to user
    fmt.Println(messageString)
    
    //remove DATA
    os.Remove(DATA)
    
    //loop through the bitstream
    for i := 0; i < len(messageString); i++ {
        //If bit is 1, 
        if string(messageString[i]) == "1" {
            //touch DATA
            os.Create(DATA)
        } else if string(messageString[i]) == "0" {
        //else remove DATA
            os.Remove(DATA)
        } else {
            l := log.New(os.Stderr, "", 0)
            l.Println("Character is not 1 or 0.  Something went horribly wrong.")
            os.Exit(1)
        }

        //touch DATA_SEND_READY
        os.Create(DATA_SEND_READY)
        //loop
        for {
            //if DATA_READ exists,
            if _, err := os.Stat(DATA_READ); err != nil {
                break
            }
        }
        os.Remove(DATA_SEND_READY)
    }

    
}
