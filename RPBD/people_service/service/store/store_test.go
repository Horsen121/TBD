package store

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	_, err := NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	if err != nil {
		t.Errorf("test NewStore: %v", err)
	}
}

func TestListPeople(t *testing.T) {
	st, _ := NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	_, err := st.ListPeople()
	if err != nil {
		t.Errorf("test ListPeople: %v", err)
	}
}

func TestGetPeopleById(t *testing.T) {
	st, _ := NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	_, err := st.GetPeopleByID("1")
	if err != nil {
		t.Errorf("test GetPeopleById: %v", err)
	}
}
