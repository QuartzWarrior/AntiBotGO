package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/valyala/fasthttp"
)

func main() {
	go server()

	client, err := discordgo.New("Bot ")
	if err != nil {
		panic(err)
	}

	err = client.Open()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func server() {
	err := fasthttp.ListenAndServe(":443", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("Hello, world!")
	})
	if err != nil {
		panic(err)
	}
}
