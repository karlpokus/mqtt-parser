package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	err := start(os.Stdin, os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
}

func start(r io.Reader, w io.Writer) error {
	log.SetOutput(w)
	log.SetFlags(0)

	textChan := make(chan string)
	packetChan := make(chan Packet)
	errChan := make(chan error)

	go format(packetChan, errChan)
	go parse(textChan, packetChan)
	go listen(r, textChan, errChan)

	select {
	case err := <-errChan:
		return err
	case <-interrupt():
		return nil
	}
}

// listen reads from r and writes to textChan
// when reading is done - it closes textChan and
// if any error occured push this error on errChan
func listen(r io.Reader, textChan chan string, errChan chan error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		textChan <- scanner.Text()
	}
	err := scanner.Err()
	// only push non-nil err since this will exit main
	if err != nil {
		errChan <- err
	}
	close(textChan)
}

// parse reads textChan and writes to packetChan
// and closes packetChan when done
func parse(textChan chan string, packetChan chan Packet) {
	var p Packet
	var pending bool
	for s := range textChan {
		if pending {
			s = strings.TrimSpace(s)
			// first char will be ASCII so this is ok
			p.ControlPacket = string(s[0])
			packetChan <- p
			// reset
			p = Packet{}
			pending = false
		}
		if strings.HasPrefix(s, "T") {
			p.Header = s
			pending = true
		}
	}
	close(packetChan)
}

// format reads-, and dumps from packetChan
// and writes to errChan when done
func format(packetChan chan Packet, errChan chan error) {
	for p := range packetChan {
		log.Printf("%s: %s", p.Header, parsePacketType(p.ControlPacket))
	}
	errChan <- nil
}

// interrupt returns a channel that recieves interrupt signals
func interrupt() <-chan os.Signal {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	return sigc
}
