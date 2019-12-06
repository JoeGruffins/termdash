package main

import (
	"context"

	"github.com/joegruffins/termdash"
	"github.com/joegruffins/termdash/container"
	"github.com/joegruffins/termdash/linestyle"
	"github.com/joegruffins/termdash/terminal/termbox"
	"github.com/joegruffins/termdash/terminal/terminalapi"
	"github.com/joegruffins/termdash/widgets/menu"
)

func main() {
	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	menu, err := menu.New()
	if err != nil {
		panic(err)
	}
	if err := menu.Write("one\ntwo\nthree\nfour"); err != nil {
		panic(err)
	}

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.PlaceWidget(menu),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		panic(err)
	}
}
