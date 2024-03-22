package Err

import "log"

func LogErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
