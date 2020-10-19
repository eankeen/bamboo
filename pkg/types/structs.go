package types

type BarConfig struct {
	Width         int      `toml:"width"`
	Height        int      `toml:"height"`
	Center        bool     `toml:"center"`
	OffsetX       int      `toml:"offset-x"`
	OffsetY       int      `toml:"offset-y"`
	Widgets       []string `toml:"widgets"`
	WidgetSpacing string   `toml:"widget-spacing"`
}

type ClockConfig struct {
	Kind             string `toml:"type"`
	Format           string `toml:"format"`
	ForegroundNormal string `toml:"foreground-normal"`
	BackgroundNormal string `toml:"background-normal"`
	ForegroundHover  string `toml:"foreground-hover"`
	BackgroundHover  string `toml:"background-hover"`
}

type Config struct {
	Bar     BarConfig
	Widgets []ClockConfig
}
