package main

import (
	"errors"
	"fmt"
    "math/rand"
)
func callbackCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("No pokemon name provided")
    }
    pokemonName := args[0]

    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
        //fmt.Println("ERRORED IN ASSIGNMENT")
        return err
    }
    
    //fmt.Println("TESTING before rand BaseExperience is:", pokemon.BaseExperience)
    const threshold = 50
    randNum := rand.Intn(pokemon.BaseExperience)
    //fmt.Println("TESTING", pokemon.BaseExperience, randNum, threshold)
    if randNum > threshold {
        return fmt.Errorf("failed to catch %s", pokemonName)
    }

    cfg.caughtPokemon[pokemonName] = pokemon
    fmt.Printf("%s was caught!\n", pokemonName)
    return nil
}
