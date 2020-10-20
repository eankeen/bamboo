package types

type BarConfig struct {
	Width         int
	Height        int
	Center        bool
	OffsetX       int
	OffsetY       int
	Widgets       []string
	WidgetSpacing string
}

type Config struct {
	Bar     []BarConfig
	Widgets map[string]interface{}
}

type TimeWidget struct {
	Name             string
	Type             string
	Format           string
	ForegroundNormal string
	BackgroundNormal string
	ForegroundHover  string
	BackgroundHover  string
}
