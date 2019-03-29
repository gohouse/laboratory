package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = `あ Ю 〓 § ♤ ⊱ ⋛ ⋌ ⋚ ⊰ ⊹ ≈ ۞ ۩ ๑ ✲ ❈ ➹ ~ - 【 】 ┱ ┲ ❣ ✚ ✪ ✣ ✤ ✥ ✦ ❉ ❥ ❦ ❧ ❃ ❂ ❁ ❀ ☭ ღ ▷ ◀ ◁ ☁ ☂ ☃ ☄ ☇ ☈ ☊ ☋ ☌ ☍ ⓛ ⓞ ⓥ ⓔ ╬ 『 』 ∴ ☀ . ∷ ﹌ ▶ ↘ の → ♧ ぃ ￡ ‿ ◕ ｡ ✎ ✟ ஐ ❤ • ۰ ● ○ ① ⊕ Θ ㊣ ★ ☆ ◆ ◇ ◣ ◢ ▲ ▼ △ ▽ ⊿ ◤ ◥ ✐ ✌ ✍ ✡ ✓ ✔ ✕ ✖ ♀ ☜ ☞ ⊙ ◎ ☺ ☻ ► ◄ ▧ ▨ ◐ ◑ ↔ ↕ ♡ ▪ ▫ ▀ ▄ █ ▌ ▐ ░ ▒ ▬ ♦ ◊ ◦ ☼ ♠ ♣ ▣ ▤ ▥ ▦ ▩ ◘ ◙ ◈ ✄ ☣ ☢ ☠ ♯ ♩ ♪ ♫ ♬ ♭ ♮ ☎ ☏ ☪ ♈ ♨ º ₪ ¤ 큐 « » ™ ♂ ✿ ♥`
	a = strings.Replace(a, " ", "", -1)
	b := []byte(a)
	for _, item := range b {
		fmt.Println(string(item))
	}
}
