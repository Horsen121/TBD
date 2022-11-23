package wombat

import "testing"

func TestEat_1(t *testing.T) {
	w := New()
	w.Eat(1)
	gotHp := w.hp
	expHp := 110
	gotW := w.weight
	var expW float32 = 45
	if gotHp != expHp {
		t.Errorf("got %v hp, but exp %v", gotHp, expHp)
	}
	if gotW != expW {
		t.Errorf("got %v weight, but exp %v", gotW, expW)
	}
}
func TestEat_2_False(t *testing.T) {
	w := New()
	w.Eat(2)
	got := w.hp
	exp := 70
	if got != exp {
		t.Errorf("got %v hp, but exp %v", got, exp)
	}
}

func TestDig_Int(t *testing.T) {
	w := New()
	w.Dig(1)
	gotHp := w.hp
	expHp := 70
	gotL := w.holeLen
	expL := 15
	if gotHp != expHp {
		t.Errorf("got %v hp, but exp %v", gotHp, expHp)
	}
	if gotL != expL {
		t.Errorf("got %v length, but exp %v", gotL, expL)
	}
}
func TestDig_Lazy(t *testing.T) {
	w := New()
	w.Dig(2)
	gotHp := w.hp
	expHp := 90
	gotL := w.holeLen
	expL := 12
	if gotHp != expHp {
		t.Errorf("got %v hp, but exp %v", gotHp, expHp)
	}
	if gotL != expL {
		t.Errorf("got %v length, but exp %v", gotL, expL)
	}
}

func TestSleep(t *testing.T) {
	w := New()
	w.Sleep()
	got := w.holeLen
	exp := 8
	if got != exp {
		t.Errorf("got %v,but exp %v", got, exp)
	}
}
