package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/obviyus/markovIRC/markov"
	irc "github.com/thoj/go-ircevent"
)

const channel = "#darkscience"
const serverssl = "irc.darkscience.net:6697"

func main() {
	ircnick1 := "markov"
	irccon := irc.IRC(ircnick1, "markov")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	err := irccon.Connect(serverssl)

	markov.Init()

	// Callback to execute when a message is received either on the channel or
	// directly to the bot (query/msg for example)
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		m := e.Message()
		if strings.HasPrefix(m, "!") {
			if strings.HasPrefix(m, "!mk") {
				irccon.Privmsg("#bots", markov.MainChain.Generate())
			}
		} else if strings.HasPrefix(m, "markov") {
			irccon.Privmsg("#bots", markov.MainChain.Generate())
		} else {
			markov.MainChain.Build(m)
		}
	})

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
