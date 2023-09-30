package main

import (
	"errors"
)

type Airpod struct {
	//TODO
	Status      string
	ch          chan byte
	BTConnected bool
}

type AirpodCase struct {
	//TODO
	RightAirpod Airpod
	LeftAirpod  Airpod
	Connected   bool
}

func NewAirpodCase() *AirpodCase {
	return &AirpodCase{
		RightAirpod: Airpod{
			Status:      "Docked",
			BTConnected: false,
			ch:          make(chan byte, 100),
		},
		LeftAirpod: Airpod{
			Status:      "Docked",
			BTConnected: false,
			ch:          make(chan byte, 100),
		},
	}
}

func (a *AirpodCase) GetRightAirpod() *Airpod {
	return &a.RightAirpod
}

func (a *AirpodCase) GetLeftAirpod() *Airpod {
	return &a.LeftAirpod

}
func (a *Airpod) GetState() string {
	return a.Status
}

func (a *AirpodCase) UndockLeft() *Airpod {
	if a.LeftAirpod.GetState() == "Docked" {
		a.LeftAirpod.Status = "Connected"
		a.LeftAirpod.ch = make(chan byte, 100)
		return &a.LeftAirpod
	}
	return nil
}

func (a *AirpodCase) UndockRight() *Airpod {
	if a.RightAirpod.GetState() == "Docked" {
		a.RightAirpod.Status = "Connected"
		a.RightAirpod.ch = make(chan byte, 100)
		return &a.RightAirpod
	}
	return nil
}

func (a *AirpodCase) DockLeft() error {
	if a.LeftAirpod.GetState() == "Docked" {
		return errors.New("already docked")
	}
	a.LeftAirpod.Status = "Docked"
	a.LeftAirpod.ch = nil
	return nil

}
func (a *AirpodCase) DockRight() error {
	if a.RightAirpod.GetState() == "Docked" {
		return errors.New("already docked")
	}
	a.RightAirpod.Status = "Docked"
	a.RightAirpod.ch = nil
	return nil
}

func (a *Airpod) GetChannel() chan byte {
	return a.ch
}

func (c *AirpodCase) ConnectBluetooth(ch chan byte) error {

	if c.Connected {
		return errors.New("already connected")
	}
	c.Connected = true
	if c.LeftAirpod.GetState() == "Disconnected" {
		c.LeftAirpod.Status = "Connected"
		c.LeftAirpod.BTConnected = true
	}
	if c.RightAirpod.GetState() == "Disconnected" {
		c.RightAirpod.Status = "Connected"
		c.RightAirpod.BTConnected = true
	}
	go func() {
		for {
			for data := range ch {

				if c.LeftAirpod.Status == "Connected" {
					c.LeftAirpod.GetChannel() <- data
				}
				if c.RightAirpod.Status == "Connected" {
					c.RightAirpod.GetChannel() <- data
				}
			}
		}
	}()

	return nil
}
