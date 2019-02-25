package roman

type glyph struct {
	limit int
	roman string
}

type glyphs []glyph

var romanGlyphs = glyphs {

	glyph{limit:10, roman:"X"},
	glyph{limit:9, roman:"IX"},
	glyph{limit:5, roman:"V"},
	glyph{limit:4, roman:"IV"},
	glyph{limit:1, roman:"I"},
}

func toRoman(n int) (s string ) {

	for _, g := range romanGlyphs {
		for ; n >= g.limit; n-= g.limit {
			s += g.roman
		}
	}

	return
}