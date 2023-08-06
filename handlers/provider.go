package handlers

import (
	"github.com/atom-apps/dictionary/handlers/dictionaries"
	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: dictionaries.Provide},
	}
}
