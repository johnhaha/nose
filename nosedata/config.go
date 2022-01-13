package nosedata

var (
	noseColors  []string = []string{"default", "gray", "brown", "orange", "yellow", "green", "blue", "purple", "pink", "red", "gray_background", "brown_background", "orange_background", "yellow_background", "green_background", "blue_background", "purple_background", "pink_background", "red_background"}
	boldTag              = "bold"
	textCodeTag          = "text-code"
)

func isNoseColor(color string) bool {
	for _, c := range noseColors {
		if c == color {
			return true
		}
	}
	return false
}
