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
Finish planing doc
Install go-lang build / run hello world. Set up git repro.
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

// TODO step one how to store the boards and ships
//      boards can be a 2d int array
// NOTE make a struct that holds pl
// 10 x 10 board

type Player struct {
    Designation string
    Ships = []Ships
    Board = [][]int
    Radar = [][]int
}

type Postion struct {
    X int
    Y int
}

ype ShipType string
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

func main() {
    fmt.Println("Hello, World!")
}

