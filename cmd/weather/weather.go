package main

import (
	"fmt"
	"os"

	"github.com/nero-dv/weather/pkg/cmd/weather"
)

func main() {
	if err := weather.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
