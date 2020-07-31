package main

import "./gadget"

func main() {
	mixtape := []string{
		"Walk Like an Egyptian", "Bohemian Rhapsody", "Space Oddity"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
