package slots

import (
	"math/rand/v2"
)

var emojis = []string{
	"🍒",    // Cereja
	"🍋",    // Limão
	"🔔",    // Sino
	"🍇",    // Uva
	"🍉",    // Melancia
	"⭐",    // Estrela
	"💎",    // Diamante
	"7️⃣ ", // Número sete
	"🍀",    // Trevo de quatro folhas
}

var random [3]int

func RodarSlots() []string {
	random = [3]int{rand.IntN(len(emojis)), rand.IntN(len(emojis)), rand.IntN(len(emojis))}

	return []string{emojis[random[0]], emojis[random[1]], emojis[random[2]]}
}

func ChecarResultado() bool {
	if random[0] == random[1] && random[0] == random[2] {
		return true
	}

	return false
}
