package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDockingSample(t *testing.T) {
	a := NewAirpodCase()
	l := a.UndockLeft()
	r := a.UndockRight()
	assert.NotNil(t, l)
	assert.NotNil(t, r)
}

func TestLeftPodSample(t *testing.T) {
	a := NewAirpodCase()
	ch := make(chan byte, 200)
	data := []byte{1, 2, 3, 4}
	var ap *Airpod
	if err := a.ConnectBluetooth(ch); err != nil {
		panic("error")
	}
	ap = a.UndockLeft()
	time.Sleep(50 * time.Millisecond)
	for _, msg := range data {
		ch <- msg
	}

	resultSlice := []byte{}
	assert.NotNil(t, ap.GetChannel())
	for i := 0; i < len(data); i++ {
		select {
		case msg := <-ap.GetChannel():
			resultSlice = append(resultSlice, msg)
		case <-time.After(500 * time.Millisecond):
			t.Fatal(i)
		}
	}
	assert.EqualValues(t, data, resultSlice)
}
