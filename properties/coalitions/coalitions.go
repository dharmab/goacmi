package coalitions

type Coalition string

const (
	// Allies is the Red coalition in DCS World.
	Allies Coalition = "Allies"
	// Enemeis is the Blue coalition in DCS World.
	Enemies Coalition = "Enemies"
	// Neutrals is the Green coalition in DCS World.
	Neutrals Coalition = "Neutrals"
)

func (c Coalition) String() string {
	return string(c)
}
