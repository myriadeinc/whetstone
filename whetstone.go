package main

import (
	"github.com/myriadeinc/whetstone/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msgf("Starting %s", os.Getenv("service__local__name"))
	r := server.New()
	go func() {
		r.ListenHTTP()
	}()
	r.Listen()
}
