package main

import "fmt"

// struct can have field that compose from another struct
type User struct {
	Id             int
	Name, Location string
}

type Player struct {
	// Player compose User
	User
	GameId int
}

func main() {
	// first way to fill composition struct value
	p1 := Player{}
	p1.Id = 42
	p1.Name = "John"
	p1.Location = "LA"
	p1.GameId = 989
	fmt.Printf("%+v", p1)

	// second way to use struct literal to fill value
	p2 := Player{
		User{Id: 43, Name: "Matt", Location: "LA"},
		999,
	}
	fmt.Println()
	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, GameId: %d\n",
		p2.Id, p2.Name, p2.Location, p2.GameId,
	)
	p2.Id = 111
	fmt.Printf("%+v", p2)
}
