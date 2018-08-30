package main

type game struct {
	Name   string
	Played bool
	Link   string
	Rating int
}

func getGamesFromDB() []game {
	return []game{
		{
			Name:   "Nefarious",
			Played: false,
			Link:   "https://www.youtube.com/watch?v=76e1ei-xyJg",
			Rating: 5,
		},
		{
			Name:   "Settlers of Catan",
			Played: false,
			Link:   "https://www.catan.com/",
			Rating: 5,
		},
	}
}
