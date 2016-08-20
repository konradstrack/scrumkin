package nlp

type Token struct {
	Index         int    `json:"index"`
	Word          string `json:"word"`
	OriginalText  string `json:"originalText"`
	Lemma         string `json:"lemma"`
	POS           string `json:"pos"`
	NER           string `json:"ner"`
	NormalizedNER string `json:"normalizedNER"`
	Speaker       string `json:"speaker"`
}

type Sentence struct {
	Tokens []Token `json:"tokens"`
}

type Text struct {
	Sentences []Sentence `json:"sentences"`
}
