package wombat

import (
	"fmt"
	"math/rand"
)

type Wombat struct {
	holeLen int
	hp      int
	rep     int
	weight  float32
}

func New() *Wombat {
	return &Wombat{
		holeLen: 10,
		hp:      100,
		rep:     20,
		weight:  30,
	}
}

func (w *Wombat) Dig(intensity int) {
	switch intensity {
	case 1:
		w.holeLen += 5
		w.hp -= 30
	case 2:
		w.holeLen += 2
		w.hp -= 10
	}
}

func (w *Wombat) Eat(grass int) {
	switch grass {
	case 1:
		w.hp += 10
		w.weight += 15
	case 2:
		if w.rep >= 30 {
			w.hp += 30
			w.weight += 30
		} else {
			w.hp -= 30
		}
	}
}

func (w *Wombat) Fight(enemy float32) string {
	var rnd = rand.Float32()
	switch enemy {
	case 30:
		if w.weight/(w.weight+enemy) >= rnd {
			w.rep += rep_change(enemy, w.weight)
			return "Победа!"
		}
		w.hp += hp_change(enemy, w.weight)
		return "Поражение :("
	case 50:
		if w.weight/(w.weight+enemy) >= rnd {
			w.rep += rep_change(enemy, w.weight)
			return "Победа!"
		}
		w.hp += hp_change(enemy, w.weight)
		return "Поражение :("
	case 70:
		if w.weight/(w.weight+enemy) >= rnd {
			w.rep += rep_change(enemy, w.weight)
			return "Победа!"
		}
		w.hp += hp_change(enemy, w.weight)
		return "Поражение :("
	}
	return ""
}
func (w *Wombat) Sleep() {
	w.holeLen -= 2
	w.hp += 20
	w.rep -= 2
	w.weight -= 5
}

func (w *Wombat) Stats() string {
	if w.hp <= 0 || w.holeLen <= 0 || w.rep <= 0 || w.weight <= 0 {
		return "You loose"
	}

	if w.rep >= 100 {
		return "You win!!!"
	}

	return fmt.Sprintf("Your stats:\n\thp:%v\n\trep:%v\n\tweight:%v\n\tholeLength:%v", w.hp, w.rep, w.weight, w.holeLen)
}
func rep_change(enemy float32, wombat float32) int {
	score := enemy - wombat
	switch {
	case score <= -20:
		return 5
	case score >= -10 && score < 15:
		return 20
	case score >= 15:
		return 40
	}
	return 0
}
func hp_change(enemy float32, wombat float32) int {
	score := enemy - wombat
	switch {
	case score <= -20:
		return -15
	case score >= -10 && score < 15:
		return -40
	case score >= 15:
		return -70
	}
	return 0
}
