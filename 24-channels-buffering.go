// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import (
    "fmt"
    "time"
    "github.com/adzest/myutils"
)

func main() {
    defer myutils.TimeTrack(time.Now(), "24-channels-buffering.go")

    // Here we `make` a channel of strings buffering up to
    // 2 values.
    messages := make(chan string, 2)

    // Because this channel is buffered, we can send these
    // values into the channel without a corresponding
    // concurrent receive.
    messages <- "buffered"
    messages <- "channel"

    // Later we can receive these two values as usual.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
