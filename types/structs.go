package types

import (
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

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
	Widgets []WidgetBase
}

type WidgetBase struct {
	Widget
}

func (w *WidgetBase) UnmarshalText(text []byte) error {
	var widgetType struct {
		Type string
	}

	if err := toml.Unmarshal(text, &widgetType); err != nil {
		return err
	}

	fn, ok := WidgetTypes[widgetType.Type]
	if !ok {
		return errors.New("Invalid widget type: " + widgetType.Type)
	}
	widgetStruct := fn()

	if err := toml.Unmarshal(text, &widgetStruct); err != nil {
		return err
	}

	w.Widget = widgetStruct
	return nil
}

type Widget interface{}

type TimeWidget struct {
	Name             string
	Type             string
	Format           string
	ForegroundNormal string
	BackgroundNormal string
	ForegroundHover  string
	BackgroundHover  string
}

var WidgetTypes = map[string]func() Widget{
	"Time": func() Widget { return new(TimeWidget) },
}
