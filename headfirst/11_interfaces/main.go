package main

import "./gadget"

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Walk Like an Egyptian", "Bohemian Rhapsody", "Space Oddity"}
	playList(player, mixtape)
}

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
