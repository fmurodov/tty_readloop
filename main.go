package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {

	buf := make([]byte, 1024)
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./tty_readloop TTY_PORT")
		return
	}

	c := &serial.Config{Name: os.Args[1], Baud: 9600, ReadTimeout: time.Millisecond * 500}
	// Opening port
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal("OpenPort:", err)
	}
	defer s.Close()

	for {
		// flush unreaded data on port
		s.Flush()

		//read new data
		n, err := s.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal("Read:", err)
		}

		//continue if no data
		if n == 0 {
			continue
		}
		log.Printf("%s", buf[:n])

	}
}
