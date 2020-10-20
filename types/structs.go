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
	Widgets []map[string]string
}
