package utils

import "github.com/charmbracelet/lipgloss"

var DownloadStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#881798")).
	Foreground(lipgloss.Color("#FFFFFF"))

var ExtractingStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#f7a900")).
	Foreground(lipgloss.Color("#FFFFFF"))

var ErrorStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#bf0105")).
	Foreground(lipgloss.Color("#FFFFFF"))

var SuccessStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#14cc04")).
	Foreground(lipgloss.Color("#FFFFFF"))
