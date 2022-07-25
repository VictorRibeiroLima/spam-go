package main

import (
	"strings"
	"time"

	"github.com/VictorRibeiroLima/spam-go/keys"
	"github.com/VictorRibeiroLima/spam-go/scripts"
	"github.com/VictorRibeiroLima/spam-go/utils"
	"github.com/micmonay/keybd_event"
)

func main() {
	kb, _ := keybd_event.NewKeyBonding()
	time.Sleep(time.Second * 5)
	sherekArray := strings.Split(scripts.Sherek(), "\n")
	for _, sherek := range sherekArray {
		letters := strings.Split(sherek, "")
		for _, letter := range letters {
			var keyboardEvent int
			if letter == " " {
				keyboardEvent = keys.KeyboardEvents["VK_SPACE"]
			} else {
				keyboardEvent = keys.KeyboardEvents["VK_"+strings.ToUpper(letter)]
			}
			kb.SetKeys(keyboardEvent)
			if utils.IsUpper(letter) {
				kb.HasSHIFT(true)
			} else {
				kb.HasSHIFT(false)
			}
			err := kb.Launching()
			if err != nil {
				panic(err)
			}
			time.Sleep(10 * time.Millisecond)
		}
		kb.HasSHIFT(false)
		kb.SetKeys(keybd_event.VK_ENTER)
		err := kb.Launching()
		if err != nil {
			panic(err)
		}
	}

}
