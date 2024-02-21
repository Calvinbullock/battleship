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
DONE Set up boards, 2 for each user (player board, radar). Allow player modification and print them all out.
TODO Get clean print out of boards, allow switching between players boards.
DONE UI – start menu
DONE Ship placement and storage, 
TODO sinking, display hits, missus, etc.
TODO Clean implement all rules and game over checks.
---------------------------
TODO Start computer player if time, if not work on fixing bugs and finishing minimum requirements.
TODO Find good algorithms used by battleship players for the computer.
TODO More computer player work and clean up, QA
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
    icon rune
    model string
    length int
    isSunk bool
}

type Player struct {
    name string
    ships []Ship
    board [10][10]rune
    radar [10][10]rune
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

    // create default ships with blank positions.
    carrier := Ship{[]Position{pos, pos, pos, pos, pos}, 'A', "Aircraft carrier", 5, false}
    battleShip := Ship{[]Position{pos, pos, pos, pos}, 'B', "BattleShip", 4, false}
    cruiser := Ship{[]Position{pos, pos, pos}, 'C', "Cruiser", 3, false}
    submarine := Ship{[]Position{pos, pos, pos}, 'S', "Submarine", 3, false}
    destroyer := Ship{[]Position{pos, pos, pos}, 'D', "Destryoer", 3, false}

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

// Helper to getPosition takes the rune input and returns the index that rune is at
//	in letterRange.
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

	// set as pos object
	pos = Position{x: xIn, y: yIn, isHit: false}
	fmt.Println("")

	// check that the positions are in bonds, -1 is parseRuneInput error return.
	if (xIn < 10 && xIn > -1) && yIn != -1 {
	    return pos
	}
	fmt.Println("ERROR: The position you entered was not on the board.")
    }
    return pos // Should never reach....
}

// check that the ship is placed in a valid position.
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

func placeShipsOnBoard(player *Player, ship Ship, posStart Position, posEnd Position) {
    // added to board
    xDelta := posStart.x - posEnd.x
    yDelta := posStart.y - posEnd.y
    icon := ship.icon
    board := &player.board
    
    // NOTE x, y is backwards when sending to board -> [y][x] correct
    //	    place ship icons from startPos to endPos

    fmt.Println(posStart.x, posStart. y)
    fmt.Println(posEnd.x, posEnd. y)

    if xDelta != 0 {
	for x := posStart.x; x <= posEnd.x; x++ {
	    (*board)[posStart.y][x] = icon
	}

    } else if yDelta != 0 {
	for y := posStart.y; y <= posEnd.y; y++ {
	    (*board)[y][posStart.x] = icon
	}
    }
    displayBoardHalf((*board))
}

func placeShips(player *Player) {
    //board := player.board
    ships := player.ships

    // Loop through all of players ships
    for _, ship := range ships {
	length := ship.length
	shipEndPos := Position{0, 0, false}
	shipStartPos := Position{0, 0, false}
	// loop until valid ship placement.
	for true {
	    fmt.Println(fmt.Sprintf("Place your %s it is %d long.", ship.model, ship.length))

	    fmt.Println("Ship start Postion") // TODO pass into getPosition
	    shipStartPos = getPosition()
	    fmt.Println("Ship end Postion") // TODO pass into getPosition
	    shipEndPos = getPosition()
	    fmt.Println("")

	    if isShipPositionValid(shipStartPos, shipEndPos, length) {
		break
	    }
	}
	placeShipsOnBoard(player, ship, shipStartPos, shipEndPos)
    }
}

// twoPLayerGame contains the game loop for a two player game.
func twoPlayerGame() {
    // arrays in golang passed by value.
    board := [10][10]rune{
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
    }
    radar := [10][10]rune{
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
	{'~', '~', '~', '~', '~', '~', '~', '~', '~', '~'},
    }

    ships := makeShipsList()
    
    // create players.
    p1 := &Player{"player1", ships, board, radar}
    p2 := &Player{"player2", ships, board, radar}
    
    // Place ships.
    displayBoard(p1)
    placeShips(p1)

    displayBoard(p2)
    placeShips(p2)

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
    playerName := player.name

    fmt.Println("")
    fmt.Println(playerName)
    fmt.Println("Radar")
    displayBoardHalf(radar)
    fmt.Println("")
    fmt.Println("Your Ships")
    displayBoardHalf(board)
}

// Displays the players board and radar.
func displayBoardHalf(board [10][10]rune) {
    // Print x-axis labels
    fmt.Print("  ")
    for i := 1; i <= len(board[0]); i++ {
	fmt.Printf("%c ", rune(i+64)) // Convert number to uppercase letter
    }
    fmt.Println()

    // Print board with y-axis labels
    for i := range board {
	fmt.Printf("%d ", i+1) // Print row number
	for _, char := range board[i] {
	    fmt.Printf("%c ", char)
	}
	fmt.Println()
    }
}

// This function updates the board and radar for hits and missis.
func playerMove(activePlayer *Player, idlePlayer *Player) {
    activeRadar := &activePlayer.radar
    idleBoard := &idlePlayer.board

    for true { 
	pos := getPosition()
	// NOTE x, y is backwards when sending to board -> [y][x] correct
	targetRune := idleBoard[pos.y][pos.x]

	if targetRune == '~' {
	    (*idleBoard)[pos.y][pos.x] = 'M'
	    (*activeRadar)[pos.y][pos.x] = 'M'
	    fmt.Println("You missed.")
	    break

	} else if targetRune == 'M' || targetRune == 'H' {
	    // Invalid spot already hit.
	    fmt.Println("You already shot here.")

	} else {
	    // valid  hit if in list of A, B, C, S
	    (*idleBoard)[pos.y][pos.x] = 'M'
	    (*activeRadar)[pos.y][pos.x] = 'M'
	    fmt.Println("You hit a Ship!")
	    break
	}
    }   
}
