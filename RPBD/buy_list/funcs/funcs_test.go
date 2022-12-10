package funcs

import (
	"github/Horsen121/TBD/RPBD/buy_list/service/conn"
	"testing"
)

func TestStart(t *testing.T) {
	s, err := conn.NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	if err != nil {
		panic(err)
	}

	res := Start(s, "TestUser", 77770000)
	if res != `Hello! I am a bot that will help you keep track of products from purchase to use.
	To start using me, click on the appropriate button and follow the instructions.` {
		panic("PANIC!!!")
	}
}
