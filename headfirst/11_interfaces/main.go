package main

import "./gadget"

func main() {
	mixtape := []string{
		"Walk Like an Egyptian", "Bohemian Rhapsody", "Space Oddity"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
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

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}
