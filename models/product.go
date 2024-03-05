package models

type Product struct {
	Name, Image, Link                               string
	Price, Bonuses, BonusesPercentage, VirtualPrice int
}
