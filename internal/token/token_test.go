package token_test

import (
	"testing"

	"github.com/rmatsuoka/sked/internal/token"
)

func Test(t *testing.T) {
	x := []struct {
		str string
	}{
		{""},
		{"ohayougozaimasu"},
		{"konnichihaSekai"},
		{"Wagahai haNeko dearu"},
		{"TakaKuNoboRuTaiyou"},
		{"Boku rahaAndoromedaqGinga woMezaShiteiru"},
		{"tuginoBunshou karakatakanadeKaKu.qkorehakatakanadesuka.qkokohahiraganadesu"},
	}

	for _, x := range x {
		t.Log(x.str)
		tr := token.NewTokenizer([]byte(x.str))
		/*
			tr.Next()
			tk := tr.Token()
			t.Logf("type: %v\tvalue: %s", tk.Type, tk.Value)
		*/
		for tr.Next() {
			tk := tr.Token()
			t.Logf("type: %v\tvalue: %s", tk.Type, tk.Value)
		}
	}
}