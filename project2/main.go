// Assignment name: Project 2
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// [COMPLETE]
// Utility function. Clears the console screen.
func ClearScreen() {
	fmt.Println("\033[2J\033[H")
}

// [COMPLETE]
// Utility function. "Pauses" the program until the user presses enter.
func Pause() {
	var throwAway string
	fmt.Printf("Press enter when you are ready...\n")
	fmt.Scanf("%s\r\n", &throwAway)
}

// [INCOMPLETE]
// Number guessing game. The computer picks a random number between
// 1 and 100 and the player tries to guess what it is. The parameter is the
// player's name, which can be used to personalize messages.
//
// Experiment with the number of guesses you allow the player to adjust
// your game to the desired difficulty.
//
// Returns -1 if the player fails to guess the number and 1 if the player succeeds.
func NumberGuesser(player string) int {

}

// [INCOMPLETE]
// Helper function. Uses the name and choice (1=rock, 2=paper, 3=scissors)
// to display an appropriate message.
//
// EX: Computer plays paper.
func ShowRSPChoice(choice int, name string) {

}

// [INCOMPLETE]
// Rock-paper-scissors game. The computer plays randomly. The parameter is the
// player's name, which can be used to personalize messages.
//
// Returns -1 if the player loses, 0 for a draw, and 1 if the player wins.
func RockPaperScissors(player string) int {

}

// [INCOMPLETE]
// main contains the program loop with main menu, and manages status variables
// for plays/wins/losses/etc. for each game.
func main() {
	// Only need to apply this seed to the random number generator one time
	rand.Seed(time.Now().UnixNano())

	// variables (loop control, player name, status)

	// this is a good place to prompt for and read the player's name,
	// so that messages can be personalized

	// main program loop -- main menu and switching to proper game function

	// after exiting program loop, display summary

}
