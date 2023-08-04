package handlers

import "github.com/rogeecn/atom/container"

func Providers() container.Providers {
	return container.Providers{
		{Provider: dictProvider},
	}
}
