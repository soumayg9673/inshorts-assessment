package gemini

type geminiMdlString int

const (
	Gemini_2_0_Flash geminiMdlString = iota
	Gemini_2_0_Flash_Lite
)

func (gm geminiMdlString) String() string {
	return [...]string{"gemini-2.0-flash", "gemini-2.0-flash-lite"}[gm]
}

var (
	GeminiMdl = &geminiModel{
		Gemini_2_0_Flash: Gemini_2_0_Flash.String(),
	}
)

type geminiModel struct {
	Gemini_2_0_Flash      string
	Gemini_2_0_Flash_Lite string
}
