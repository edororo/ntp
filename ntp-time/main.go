package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp&quot;
)

const ntpServer = "pool.ntp.org"

func main() {
	n, err := getTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(n.Format(time.RFC3339))
}

func getTime() (time.Time, error) {
	n, err := ntp.Time(ntpServer)
	if err != nil {
		return time.Time{}, err
	}
	return n, nil
}
