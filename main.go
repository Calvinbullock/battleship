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

package main

import (
	"fmt"
	"unicode"
)

/* ==================================== *\
|*		    Consts	        *|
\* ==================================== */
var lettersRange = [...]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}


/* ==================================== *\
|*              Program Structs         *|
\* ==================================== */

type Position struct {
    x int
    y int
    isHit bool // NOTE not sure if this is best here
}

type Ship struct {
    positions []Position
    model string
    length int
    isSunk bool
}

type Player struct {
    name string
    ships []Ship
    board [10][10]int
    radar [10][10]int
}

/* ==================================== *\
|*              Main Program            *|
\* ==================================== */

func main() {
    // display menu
    playerMode := displayMainMenu()

    if playerMode == 1 {
        onePlayerGame()

    } else if playerMode == 2 {
        twoPlayerGame()

    } else if playerMode == 3 {
        return

    } 
}

// displayMainMenu is the first menu displayed when running the program.
//      returns playerMode, the users game mode choice.
func displayMainMenu() int {
    // TODO input validation.
    var playerMode int
    fmt.Println("1. Single Player")
    fmt.Println("2. Multiplayer ")
    fmt.Println("3. Exit game ")
    fmt.Scanln(&playerMode)

    return playerMode
}

// returns the list of default ships
func makeShipsList() []Ship {
    pos := Position{0, 0, false}
    
    // create defult ships with blank positions.
    carrier := Ship{[]Position{pos, pos, pos, pos, pos}, "carrier", 5, false}
    battleShip := Ship{[]Position{pos, pos, pos, pos}, "battleShip", 4, false}
    cruiser := Ship{[]Position{pos, pos, pos}, "cruiser", 3, false}
    submarine := Ship{[]Position{pos, pos, pos}, "submarine", 3, false}
    destroyer := Ship{[]Position{pos, pos, pos}, "destryoer", 3, false}

    ships := []Ship{carrier, battleShip, cruiser, cruiser, submarine, submarine, destroyer, destroyer}
    return ships
}

// onePLayerGame contains the game loop for a single player game.
func onePlayerGame() {
    // create player 1

    // create computer

    for true {
        // playerMove(player1, computer)
    }
}

// Helper to getPosition Parses a rune postion input to a matching int.
//	-1 return means char not in list.
func parseRuneInput(char rune) int {
    upperChar := unicode.ToUpper(char)

    for i, letter := range lettersRange {
	if letter == upperChar {
	    return i
	}
    }
    // not in the letterRange
    return -1
} 

// Take user input return a potions.
func getPosition() Position {
	pos := Position {}
	var xRune rune
	var xIn int
	var yIn int

    for true {
	fmt.Println("give an x position (ex. A): ")
	fmt.Scanf("%c\n", &xRune)
	xIn = parseRuneInput(xRune)

	fmt.Println("give an y position (ex. 1): ")
	fmt.Scanf("%d\n", &yIn)
	yIn-- // match index 0

	pos = Position{x: xIn, y: yIn, isHit: false}
	fmt.Println("")
	
	// check that the postions are in bonds, -1 is parseRuneInput error return.
	if (xIn < 10 && xIn > -1) && yIn != -1 {
	    return pos
	}
	fmt.Println("ERROR: The position you entered was not on the board.")
    }
    return pos // Should never reach....
}

// check that the ship is placed in a valid postion.
func isShipPositionValid(startPosition Position, endPosition Position, shipLength int) bool {
    xDelta := startPosition.x + endPosition.x + 1
    yDelta := startPosition.y + endPosition.y + 1

    if xDelta != shipLength && yDelta == 0 {
	fmt.Println("ERROR: Your ship is not the right length.")
	return false
    }
    if yDelta != shipLength && xDelta == 0 {
	fmt.Println("ERROR: Your ship is not the right length.")
	return false
    }
    if yDelta == xDelta {
	fmt.Println("ERROR: Your ship can not be diagnal.")
	return false
    }
    return true
}

func placeShips(player *Player) {
    //board := player.board
    ships := player.ships
    
    // Loop through all of players ships
    for _, ship := range ships {
	
	// loop until valid ship placement.
	for true {
	    fmt.Println(fmt.Sprintf("Place your %s it is %d long.", ship.model, ship.length))

	    // TODO need to valid that all pos are +-1 from each other.
	    fmt.Println("Ship start Postion") // TODO pass into getPosition
	    shipStartPos := getPosition()
	    fmt.Println("Ship end Postion") // TODO pass into getPosition
	    shipEndPos := getPosition()
	    fmt.Println("")

	    length := ship.length

	    if isShipPositionValid(shipStartPos, shipEndPos, length) {
		break
	    }
	}
    }
}

// twoPLayerGame contains the game loop for a two player game.
func twoPlayerGame() {
    // arrays in golang passed by value.
    board := [10][10]int{}
    radar := [10][10]int{}
    ships := makeShipsList()

    // create players.
    p1 := &Player{"player1", ships, board, radar}
    p2 := &Player{"player2", ships, board, radar}
    
    // TODO Place ships.
    placeShips(p1)

    // game loop.
    for true {
	displayBoard(p1)
        playerMove(p1, p2)

	displayBoard(p2)
        playerMove(p2, p1)
    }
}

// display a players radar and board.
func displayBoard(player *Player) {
    board := player.board
    radar := player.radar
    fmt.Print("")
    fmt.Print("Radar")
    displayBoardHalf(radar)
    fmt.Print("")
    fmt.Print("Your Ships")
    displayBoardHalf(board)
}

// Displays the players board and radar.
func displayBoardHalf(board [10][10]int) {
    // Print x-axis labels
	fmt.Print(" ")
	for i := 1; i <= len(board[0]); i++ {
		fmt.Printf("%c ", rune(i+64)) // Convert number to uppercase letter
	}
	fmt.Println()

	// Print board with y-axis labels
	for i := range board {
		fmt.Printf("%d ", i+1) // Print row number
		for _, v := range board[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func playerMove(activePlayer *Player, idlePlayer *Player) {
//          diaplay_player_board() // display player one board / radar
//          take players move
//
//          check_valid_move() // print hit or miss
//          update_board() // NOTE might put this in gameloop insted of here
}

func moveUpdate() {

}

