// card.go supplies functions that manipulate virtual cards.
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// init function for card manipulation that insures a good Random Number.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// createTable sets up basic table images.
func createTable(screen *ebiten.Image) {

	// Fill the Screen with the white color
	screen.Fill(color.White)

	//  Insert Table Image
	if table == nil {
		// Fill the Screen with the white color
		screen.Fill(color.White)
		// Create an Table image
		table, _, err = ebitenutil.NewImageFromFile(cardTable, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get Options Structure
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(15, 70)
	// Draw the table image to the screen with an empty option
	screen.DrawImage(table, opts)

	// if waiting for Deal
	if has(mode, dealWait) {
		messageSquare(150, 90, 10, 600, color.NRGBA{0x50, 0x50, 0xff, 0xff}, screen)
		messageDisplay(aSmall, "DEAL", 40, 650, screen)

		if dealPost == 8 { // Button(dealer) Chip Position updated
			dealPost = 0
		} else {
			dealPost++
		}
	}

	// If the user has a bet/fold/check decision - Setup up the Buttons.
	if has(mode, betEnable) {

		messageSquare(150, 90, 150, 600, color.NRGBA{0xff, 0xff, 0x00, 0xff}, screen)
		messageDisplay(aSmall, "CHECK", 180, 650, screen)

		messageSquare(150, 90, 300, 600, color.NRGBA{0xff, 0x00, 0x00, 0xff}, screen)
		messageDisplay(aSmall, "FOLD", 340, 650, screen)

		messageSquare(150, 90, 450, 600, color.NRGBA{0x00, 0xff, 0x00, 0xff}, screen)
		messageDisplay(aSmall, "BET", 490, 650, screen)

		if has(mode, betValue) {
			// Display Error Message for 30 seconds.
			messageError(screen)
			// Blink the cursor once a second
			t := blinkCounter()
			messageDisplay(aTiny, "Value:"+result+t, 620, 650, screen)
			mode = set(mode, betInput)
		}
	}
}

// dealCards place cards images on the displayed table.
//   The actual cards dealt depend on the mode bits
func dealCards(mode Bits, screen *ebiten.Image) {
	// displayAll(players)
	//	undisplayAll(players)

	card := 0
	ui := 0
	for pi := dealPost; pi < dealPost+9; pi++ {
		id := pi * 9 // Compute the table index
		imageDisplay(float64(players[id]), float64(players[id+1]), card,
			players[id+7], players[id+8], screen)
		users[ui][0] = card
		ui++
		card++
	}
	ui = 0
	for pi := dealPost; pi < dealPost+9; pi++ {
		id := pi * 9 // Compute the table index
		imageDisplay(float64(players[id+2]), float64(players[id+3]), card,
			players[id+7], players[id+8], screen)
		users[ui][1] = card
		ui++
		card++
	}

	fmt.Println(users)

	if has(mode, cardFlop) { // Burn Card = 20
		imageDisplay(319, 250, 21, unhide, display, screen)
		imageDisplay(393, 250, 22, unhide, display, screen)
		imageDisplay(467, 250, 23, unhide, display, screen)
		return
	}

	if has(mode, cardTurn) { // Burn Card = 24
		imageDisplay(541, 250, 25, unhide, display, screen)
		return
	}

	if has(mode, cardRiver) { // Burn Card = 26
		imageDisplay(615, 250, 27, unhide, display, screen)
		return
	}

	return
}

func resetPlayersTable(players []int) {
	for i := 0; i < 18; i++ {
		j := i * 9
		if i == 8 || i == 17 {
			players[j+7] = unhide
			players[j+8] = display
		} else {
			players[j+7] = hide
			players[j+8] = display
		}
	}
}

// changeCard function that changes the hide/display parameters
// of a particular card. Primarily to show a card front on screen.
func changeCard(player, hide, display int, screen *ebiten.Image) {
	fmt.Println("Player = ", player)
	comp := []int{}
	for _, v := range players {
		comp = append(comp, v)
	}
	postIndex := player * 9 // Compute position of player parameters
	// Change Players Table to reflect current status
	players[postIndex+7] = hide
	players[postIndex+8] = display
	// Display Chared player status
	for i, v := range players {
		j := i % 9
		if comp[j] != v {
			fmt.Print("row = ", j)
			for k := j; k < j+9; k++ {
				fmt.Print("  k = ", comp[j])
			}
		}
		fmt.Println("")
	}
	imageDisplay(float64(players[postIndex]), float64(players[postIndex+1]),
		players[postIndex+4], players[postIndex+7], players[postIndex+8], screen)
	imageDisplay(float64(players[postIndex+2]), float64(players[postIndex+3]),
		players[postIndex+5], players[postIndex+7], players[postIndex+8], screen)
}

// Mode Bit Functions - The "mode" value is used to control Deal Operationns,
//   Set, Clear, Toggle Mode Bits modify the "mode" word.
//   Has tests the state a particular bit in the "mode" word.
func set(b, flag Bits) Bits { return b | flag }

func clear(b, flag Bits) Bits { return b &^ flag }

func toggle(b, flag Bits) Bits { return b ^ flag }

func has(b, flag Bits) bool { return b&flag != 0 }

// shuffle randomly rearranges the cards in the virtual deck.
//   Due to a tendency of cards at the first and last locations
//   to stay thru a number of shuffles.  The cards at the locations
//   first(1) and last(51) are swapped with random middle positions.
//
func shuffle() {
	//  Initialize Deck Values
	for i, v := range deckInitial {
		deck[i] = v
	}
	// Three Shuffles for each new hand!
	deck[51], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[51]
	deck[0], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[0]

	rand.Shuffle(51, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// Shuffle Two
	deck[51], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[51]
	deck[0], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[0]

	rand.Shuffle(51, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// Shuffle Three
	deck[51], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[51]
	deck[0], deck[rand.Intn(51)] = deck[rand.Intn(51)], deck[0]

	rand.Shuffle(51, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}
