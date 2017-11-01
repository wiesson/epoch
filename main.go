package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) FromString(stamp string) error {
	ts, err := strconv.ParseInt(stamp, 10, 64)

	if err != nil {
		return err
	}

	t.Time = time.Unix(ts, 0)

	return nil
}

func (t* Timestamp) UtcString() string {
	return t.UTC().String()
}

func main() {
	for _, timeAsString := range os.Args[1:] {
		ts := Timestamp{}
		ts.FromString(timeAsString)
		fmt.Println(ts.UtcString())
	}
}
