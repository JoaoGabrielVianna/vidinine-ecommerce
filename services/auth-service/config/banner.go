package config

import (
	"fmt"

	"github.com/vidinine-ecommerce/auth-service/fonts"
)

func ShowWelcomeBanner() {
	render("AUTH-SERVICE")

	systemlogger.System("┌────────────────────────────────────────────────────┐")
	systemlogger.System("│                INICIALIZANDO SISTEMA               │")
	systemlogger.System("└────────────────────────────────────────────────────┘\n")
}

func render(text string) {
	lines := make([]string, 7)

	for _, char := range text {
		if glyph, ok := fonts.AnsiShadow[char]; ok {
			for i := 0; i < 7; i++ {
				lines[i] += glyph[i]
			}
		} else {
			for i := 0; i < 7; i++ {
				lines[i] += "        "
			}
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
