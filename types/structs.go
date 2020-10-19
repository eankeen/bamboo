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

type ClockConfig struct {
	Kind             string `toml:"type"`
	Format           string `toml:"format"`
	ForegroundNormal string `toml:"foregroundNormal"`
	BackgroundNormal string `toml:"backgroundNormal"`
	ForegroundHover  string `toml:"foregroundHover"`
	BackgroundHover  string `toml:"backgroundHover"`
}

type Config struct {
	Bar     BarConfig
	Widgets []ClockConfig
}
