# Pokedex CLI

A Golang CLI for interacting with the PokeAPI.

## Features

- **Explore**: Explore different areas and find Pokemon encounters.
- **Catch**: Attempt to catch a Pokemon by name.
- **Inspect**: Inspect the stats of a caught Pokemon.
- **Map**: Display names of locations.
- **Pokedex**: Display names of caught Pokemon.

## Installation

To install the productivity-counter, clone the repository and build the project:

```sh
git clone https://github.com/GauravC4/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
```

## Usage

Run the `pokedexcli` executable to start the REPL (Read-Eval-Print Loop) interface:
```./pokedexcli```

## Commands

- **help**: Display a help message.  
- **exit**: Exit the Pokedex.  
- **map**: Display names of the next 20 locations.  
- **mapb**: Display names of the previous 20 locations.  
- **explore `<area-name>`**: Display Pokemon encounters in the specified area.  
- **catch `<pokemon-name>`**: Attempt to catch the specified Pokemon.  
- **inspect `<pokemon-name>`**: Show stats of the specified caught Pokemon.  
- **pokedex**: Display names of caught Pokemon.  

## Cache

Implements inmemory cache by default with a garbage collection loop to remove expired keys.
Optionally it also has a redis implementation which can be used by invloking `NewRedisCache` function.
