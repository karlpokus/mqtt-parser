package main

import (
	"strings"
	"testing"

  "github.com/karlpokus/bufw"
)

const input = `#
T 127.0.0.1:36186 -> 127.0.0.1:1883 [AP] #1
  e0 00                                                 ..
########`

func TestStart(t *testing.T) {
  r := strings.NewReader(input)
  w := bufw.New()
  done := make(chan bool)
  go func(){
    err := <-start(r, w)
    if err != nil {
      t.Fatal(err)
    }
    done <-true
  }()
  err := w.Wait()
  if err != nil {
    t.Fatal(err)
  }
  output := w.String()
  expected := "T 127.0.0.1:36186 -> 127.0.0.1:1883 [AP] #1: DISCONNECT"
  if output != expected {
		t.Fatalf("%s does not match %s", output, expected)
	}
  <-done
}
