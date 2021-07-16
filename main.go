package main

import "fmt"

// Normal represents a Normal type move or Pokemon
type Normal struct{}

// Fighting represents a Fighting type move or Pokemon
type Fighting struct{}

// Flying represents a Flying type move or Pokemon
type Flying struct{}

// Poison represents a Poison type move or Pokemon
type Poison struct{}

// Ground represents a Ground type move or Pokemon
type Ground struct{}

// Rock represents a Rock type move or Pokemon
type Rock struct{}

// Bug represents a Bug type move or Pokemon
type Bug struct{}

// Ghost represents a Ghost type move or Pokemon
type Ghost struct{}

// Steel represents a Steel type move or Pokemon
type Steel struct{}

// Fire represents a Fire type move or Pokemon
type Fire struct{}

// Water represents a Water type move or Pokemon
type Water struct{}

// Electric represents a Electric type move or Pokemon
type Electric struct{}

// Psychic represents a Psychic type move or Pokemon
type Psychic struct{}

// Ice represents a Ice type move or Pokemon
type Ice struct{}

// Dragon represents a Dragon type move or Pokemon
type Dragon struct{}

// Dark represents a Dark type move or Pokemon
type Dark struct{}

// Fairy represents a Fairy type move or Pokemon
type Fairy struct{}

// Grass represents a Grass type move or Pokemon
type Grass struct{}

//Here we will know the modifier that has to be used depending on the attacked pokemon
func (g Grass) Modifier(types ...interface{}) float64 {

	//1.0 means that the damage of the attack is the normal ammount.
	modifier := 1.0

	for _, t := range types {

		//The normal way would be use a lot of 'if' conditionals, 1 for every of existing pokemon, so
		//to save our hands of typping so many times instead we use a 'switch case' on the
		//reserved keyword 'type', relaying the dependency of the condittional on an abstraction
		//(the 'type' itself) and not a concretion (every single one of the types of pokemon).
		//Using this way we save a tons of lines of code and do more readable our code :) .

		switch t.(type) {

		//In this case, Grass is uneffective vs this types of pokemon, so we assign '/2' value to modifier
		case Fire, Flying, Bug, Poison, Steel, Grass, Dragon:

			modifier /= 2

		//But here type grass is very effective vs this types, so instead of dividing, we multiplie the
		//value x2
		case Rock, Ground, Water:
			modifier *= 2
		}
	}

	return modifier
}

func main() {

	//We choose a grass type for the example.
	myPokemon := Grass{}

	//'Cause our rival is not stupid, he choose a strong type vs our grass pokemon
	rival := Fire{}

	//But Brock loves Rock/Grounds pokemons, so our pokemon will ve very strong vs him.
	brock := []interface{}{Rock{}, Ground{}}

	mod := myPokemon.Modifier(rival)

	fmt.Printf("My %#v pokemon has a %.2f modifier against %#v\n", myPokemon, mod, rival)

	mod2 := myPokemon.Modifier(brock...)

	fmt.Printf("My %#v pokemon has a %.2f modifier against %#v\n", myPokemon, mod2, brock)

}
