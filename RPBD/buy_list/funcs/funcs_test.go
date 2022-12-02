package funcs

import (
	"github/Horsen121/TBD/RPBD/buy_list/service/conn"
	"testing"
)

func TestStart(t *testing.T) {
	_, err := conn.NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	if err != nil {
		panic(err)
	}
}
