package token

const (
	Hiragana Type = iota
	Katakana
	Kanji
	Okurigana
	Q
)

type Type int

func (i Type) String() (ret string) {
	switch i {
	case Hiragana:
		ret = "Hiragana"
	case Katakana:
		ret = "Katakana"
	case Kanji:
		ret = "Kanji"
	case Okurigana:
		ret = "Okurigana"
	case Q:
		ret = "Q"
	default:
		ret = "?"
	}
	return
}

type Tokenizer struct {
	b                  []byte
	start, end, offset int
	ttype              Type
}

func NewTokenizer(b []byte) *Tokenizer {
	return &Tokenizer{b: b}
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func (t *Tokenizer) Next() bool {
	if t.end >= len(t.b) {
		return false
	}

	t.start = t.offset
	t.end = t.offset

	if t.b[t.end] == 'q' {
		t.ttype = Q
		t.end++
		t.offset = t.end
	} else if isUpper(t.b[t.end]) {
		t.end++
		t.advToTokenEnd()
	} else {
		t.ttype = Hiragana
		for ; t.end < len(t.b) && t.b[t.end] != 'q' && !isUpper(t.b[t.end]); t.end++ {
		}
		t.offset = t.end
	}
	return true
}

func (t *Tokenizer) advToTokenEnd() {
	t.ttype = Kanji
forloop:
	for ; t.end < len(t.b); t.end++ {
		switch {
		case t.b[t.end] == ' ':
			t.offset = t.end + 1
			return
		case t.b[t.end] == 'q':
			t.offset = t.end + 1
			t.ttype = Katakana
			break forloop
		case isUpper(t.b[t.end]):
			t.end++
			t.offset = t.end
			t.ttype = Okurigana
			break forloop
		}
	}
}

type Token struct {
	Value []byte
	Type  Type
}

func (t *Tokenizer) Token() *Token {
	end := t.end
	if end > len(t.b) {
		end = len(t.b)
	}
	return &Token{Value: t.b[t.start:end], Type: t.ttype}
}
