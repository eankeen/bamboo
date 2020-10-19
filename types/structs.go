package types

type BarConfig struct {
	Width         int      `toml:"width"`
	Height        int      `toml:"height"`
	Center        bool     `toml:"center"`
	OffsetX       int      `toml:"offsetX"`
	OffsetY       int      `toml:"offsetY"`
	Widgets       []string `toml:"widgets"`
	WidgetSpacing string   `toml:"widgetSpacing"`
}

type Config struct {
	Bar     BarConfig
	Widgets []map[string]string `toml:"widgets"`
}
