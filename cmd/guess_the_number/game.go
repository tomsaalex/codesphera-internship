package main

type Game struct {
	Secret, Tries, MaxTries int
	isWon                   bool
}

func NewGame(secret, maxTries int) *Game {
	return &Game{Secret: secret, MaxTries: maxTries}
}

func (g *Game) Guess(n int) string {
	g.Tries++

	// Game has been won previously
	if g.isWon {
		return "You have won!"
	}

	// Game has been lost previously
	if g.Tries > g.MaxTries {
		return "You have lost!"
	}

	// Game has been lost right now
	if g.Tries == g.MaxTries && n != g.Secret {
		return "You have lost!"
	}

	// Wrong attempt, but the game continues
	if n > g.Secret {
		return "Too high!"
	}
	if n < g.Secret {
		return "Too low!"
	}

	// Game has been won right now
	g.isWon = true
	return "You have won!"
}
