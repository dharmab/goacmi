package colors

type Color string

const (
	Red    Color = "Red"
	Orange Color = "Orange"
	Yellow Color = "Yellow"
	Green  Color = "Green"
	Cyan   Color = "Cyan"
	Blue   Color = "Blue"
	Violet Color = "Violet"
	Grey   Color = "Grey"
)

func (c Color) String() string {
	return string(c)
}
