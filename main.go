// Calvin Bullock
// BattleShip main file


/* TODO
Minimum project
- two player
- console boards (yours and radar)
- player switching
Extra
- two player with computer player.
- if a hit is found by computer.
- check in points around the hit until sunk.
- save state

------ TODO
DONE Finish planing doc
DONE Install go-lang build / run hello world. Set up git repro.
Set up boards, 2 for each user (player board, radar). Allow player modification and print them all out.
Get clean print out of boards, allow switching between players boards.
UI – start menu
Ship placement and storage, sinking, display hits, missus, etc.
Clean implement all rules and game over checks.
Start computer player if time, if not work on fixing bugs and finishing minimum requirements.
Find good algorithms used by battleship players for the computer.
More computer player work and clean up, QA
Fill out document, make video final clean up and formfitting.

*/

package main
import "fmt"

type Postion struct {
    X int
    Y int
}

type ShipType string
const (
  ShipTypeDestroyer  ShipType = "Destroyer"
  ShipTypeBattleship ShipType = "Battleship"
  ShipTypeSubmarine ShipType = "Submarine"
  ShipTypeCarrier   ShipType = "Carrier"
)

type Ship struct {
    Postions = []Postion
    Model string
    IsSunk bool
}

/*
Carrier: 5 squares long, the largest and most valuable ship.
Battleship: 4 squares long, a powerful ship.
Cruiser: 3 squares long, two of these are included in the fleet.
Submarine: 3 squares long, two of these are included in the fleet.
Destroyer: 2 squares long, the smallest and least valuable ship.
So, to answer your question, a player gets:

1 Carrier
1 Battleship
2 Cruisers
2 Submarines
2 Destroyers
*/

type Player struct {
    Name string
    Ships = []Ships
    Board = [10][10]int
    Radar = [10][10]int
}
/*
    p := {
    Name string
    Ships = []Ships
    Board = [10][10]int
    Radar 
} {}
*/

func main() {
    // display menu
    playerMode := displayMainMenu()

    if playerMode == 1 {
        onePlayerGame()

    } else if playerMode == 2 {
        twoPlayerGame()

    } else if playerMode == 3 {
        return 0

    } 
}

// displayMainMenu is the first menu displayed when running the program.
//      returns playerMode, the users game mode choice.
func displayMainMenu() {
    // TODO input validation.
    var playerMode int
    fmt.Println("1. Single Player")
    fmt.Println("2. Multiplayer ")
    fmt.Println("3. Exit game ")
    fmt.Scanln(&playerMode)

    return playerMode
}

// onePLayerGame contains the game loop for a single player game.
func onePlayerGame() {
    // create player 1

    // create computer

    for true {
        playerMove(player1, computer)
    }
}

// twoPLayerGame contains the game loop for a two player game.
func twoPlayerGame() {
    // create player 1

    // create player 2

    for true {
        playerMove(player1, player2)
        playerMove(player2, player1)
    }
}

func playerMove(Player activePlayer, PLayer idlePlayer) {
//          diaplay_player_board() // display player one board / radar
//          take players move
//
//          check_valid_move() // print hit or miss
//          update_board()
//          swtich player when ready
}

func moveUpdate() {

}

