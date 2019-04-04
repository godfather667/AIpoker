//
// Display Table
//
package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

type Bits uint8 // Bit Mask for updateMode

const (
	cardDeal = 1 << iota // Update Mask Encoding
	cardFlop
	cardTurn
	cardRiver
	betValue
	betInput
	betEnable
	inputWait
)

const (
	mPlus = iota // mplus fonts
	mSmall
	mTiny
	aPlus // Arcade fonts
	aSmall
	aTiny
)
const (
	unhide    = iota // Display Card Value
	hide             // Hide Card Value (Show Card Back)
	display          // Display card at position
	undisplay        // Display Nothing
)

const (
	screenWidth  = 1024
	screenHeight = 768

	dpi = 72
)

var updateMode Bits // Update Mask

var valueResult int64 // Value input value

var mplusNormalFont font.Face // Font Variables
var smallNormalFont font.Face
var tinyNormalFont font.Face
var arcadeFont font.Face
var smallArcadeFont font.Face
var tinyArcadeFont font.Face

var table *ebiten.Image // Table Image
var card *ebiten.Image  // Card Image

var err error // Error

var cardTable = "images/table.png"

var cardBack = "images/playing-cards-back.png"

var deck = map[int]string{
	1:  "images/10_of_clubs.png",
	2:  "images/10_of_diamonds.png",
	3:  "images/10_of_hearts.png",
	4:  "images/10_of_spades.png",
	5:  "images/2_of_clubs.png",
	6:  "images/2_of_diamonds.png",
	7:  "images/2_of_hearts.png",
	8:  "images/2_of_spades.png",
	9:  "images/3_of_clubs.png",
	10: "images/3_of_diamonds.png",
	11: "images/3_of_hearts.png",
	12: "images/3_of_spades.png",
	13: "images/4_of_clubs.png",
	14: "images/4_of_diamonds.png",
	15: "images/4_of_hearts.png",
	16: "images/4_of_spades.png",
	17: "images/5_of_clubs.png",
	18: "images/5_of_diamonds.png",
	19: "images/5_of_hearts.png",
	20: "images/5_of_spades.png",
	21: "images/6_of_clubs.png",
	22: "images/6_of_diamonds.png",
	23: "images/6_of_hearts.png",
	24: "images/6_of_spades.png",
	25: "images/7_of_clubs.png",
	26: "images/7_of_diamonds.png",
	27: "images/7_of_hearts.png",
	28: "images/7_of_spades.png",
	29: "images/8_of_clubs.png",
	30: "images/8_of_diamonds.png",
	31: "images/8_of_hearts.png",

	32: "images/8_of_spades.png",
	33: "images/9_of_clubs.png",
	34: "images/9_of_diamonds.png",
	35: "images/9_of_hearts.png",
	36: "images/9_of_spades.png",
	37: "images/ace_of_clubs.png",
	38: "images/ace_of_diamonds.png",
	39: "images/ace_of_hearts.png",
	40: "images/ace_of_spades.png",
	41: "images/jack_of_clubs.png",
	42: "images/jack_of_diamonds.png",
	43: "images/jack_of_hearts.png",
	44: "images/jack_of_spades.png",
	45: "images/king_of_clubs.png",
	46: "images/king_of_diamonds.png",
	47: "images/king_of_hearts.png",
	48: "images/king_of_spades.png",
	49: "images/queen_of_clubs.png",
	50: "images/queen_of_diamonds.png",
	51: "images/queen_of_hearts.png",
	52: "images/queen_of_spades.png",
}

func init() { // Initialize Normal Fonts

	// Seup a new fonts for text display

	// Initialize Normal Fonts

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	smallNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	tinyNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Arcade Fonts

	tt, err = truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    18,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	tinyArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	// Insure Good Random Number - Initialize Seed with Nano-second time value!
	rand.Seed(time.Now().UnixNano())

}

// Bit Wise Functions
//
func Set(b, flag Bits) Bits { return b | flag }

func Clear(b, flag Bits) Bits { return b &^ flag }

func Toggle(b, flag Bits) Bits { return b ^ flag }

func Has(b, flag Bits) bool { return b&flag != 0 }

// Update called each Run Cycle
//
func update(screen *ebiten.Image) error {

	// Logical Operations to Setup Rendering

	if Has(updateMode, cardDeal) {

		if ebiten.IsDrawingSkipped() {
			// When the game is running slowly, the rendering result
			// will not be adopted.
			return nil
		}
		initTable(screen)

		deal(updateMode, screen)                // Deal Cards
		updateMode = Set(updateMode, betEnable) // Enable Message Boxes
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && Has(updateMode, betValue) {
		x, y := ebiten.CursorPosition()
		if x > 320 && x < 470 && y > 600 && y < 690 {
			updateMode = Set(updateMode, betValue) // Process various Bet Options
			updateMode = Set(updateMode, betInput)
		}
	}
	if Has(updateMode, betInput) {
		txt, hasCR := TextInput(screen)
		if hasCR {
			Clear(updateMode, betInput)
			Clear(updateMode, betValue)
		} else {
			charDisplay(1, txt, 550, 650, screen)
		}
	}

	return nil
}

func instruct() {
	fmt.Println("\n                        A I  P O K E R")
	fmt.Println("\nAIpoker is the initial concept design for an AI Based Poker Simulation")
	fmt.Println("Only one 'bug' at the moment: You must put the cursor in the image box for")
	fmt.Println("the entire image to be completely displayed.")
	fmt.Println("\nThis 'Feature' also shows up in their example games. May take Big Think!")
}

// Initialize Poker Table Display

func initTable(screen *ebiten.Image) {

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

	if Has(updateMode, betEnable) {
		messageSquare(150, 90, 20, 600, color.NRGBA{0xff, 0xff, 0x00, 0xff}, screen)
		charDisplay(aSmall, "CHECK", 40, 650, screen)

		messageSquare(150, 90, 170, 600, color.NRGBA{0xff, 0x00, 0x00, 0xff}, screen)
		charDisplay(aSmall, "FOLD", 210, 650, screen)

		messageSquare(150, 90, 320, 600, color.NRGBA{0x00, 0xff, 0x00, 0xff}, screen)
		charDisplay(aSmall, " BET", 360, 650, screen)

		if Has(updateMode, betValue) {
			charDisplay(aTiny, "Value: ", 500, 650, screen)
			Set(updateMode, betInput)
		}
	}
}

func shuffle() {
	deck[52], deck[rand.Intn(52)] = deck[rand.Intn(52)], deck[52]

	rand.Shuffle(52, func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func cardDisplay(x, y float64, cardValue, h, d int, screen *ebiten.Image) {

	// Create Card image
	if d == undisplay {
		return
	}
	if h == hide {
		card, _, err = ebitenutil.NewImageFromFile(cardBack, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		card, _, err = ebitenutil.NewImageFromFile(deck[cardValue], ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Display Image
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(x, y)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(card, opts)
}

func deal(mode Bits, screen *ebiten.Image) {

	if Has(updateMode, cardDeal) {
		cardDisplay(0, 250, 1, hide, display, screen)
		cardDisplay(70, 80, 2, hide, display, screen)
		cardDisplay(378, 20, 3, hide, display, screen)
		cardDisplay(560, 20, 4, hide, display, screen)
		cardDisplay(800, 80, 5, hide, display, screen)
		cardDisplay(850, 250, 6, hide, display, screen)
		cardDisplay(70, 440, 7, hide, display, screen)
		cardDisplay(800, 440, 8, hide, display, screen)
		cardDisplay(420, 460, 9, unhide, display, screen)

		cardDisplay(64, 250, 10, hide, display, screen)
		cardDisplay(134, 80, 11, hide, display, screen)
		cardDisplay(314, 20, 12, hide, display, screen)
		cardDisplay(624, 20, 13, hide, display, screen)
		cardDisplay(864, 80, 14, hide, display, screen)
		cardDisplay(914, 250, 15, hide, display, screen)
		cardDisplay(134, 440, 16, hide, display, screen)
		cardDisplay(864, 440, 17, hide, display, screen)
		cardDisplay(484, 460, 18, unhide, display, screen)
		return
	}

	if (mode & cardFlop) != 0 { // Burn Card = 20
		cardDisplay(319, 250, 21, hide, undisplay, screen)
		cardDisplay(393, 250, 22, hide, undisplay, screen)
		cardDisplay(467, 250, 23, hide, undisplay, screen)
		return
	}

	if (mode & cardTurn) != 0 { // Burn Card = 24
		cardDisplay(541, 250, 25, hide, undisplay, screen)
		return
	}

	if (mode & cardRiver) != 0 { // Burn Card = 26
		cardDisplay(615, 250, 27, hide, undisplay, screen)
		return
	}

	return
}

func charDisplay(font int, msg string, x, y int, screen *ebiten.Image) {
	// Add Text
	switch font {
	case 0:
		text.Draw(screen, msg, mplusNormalFont, x, y, color.Black)
	case 1:
		text.Draw(screen, msg, smallNormalFont, x, y, color.Black)
	case 2:
		text.Draw(screen, msg, tinyNormalFont, x, y, color.Black)
	case 3:
		text.Draw(screen, msg, arcadeFont, x, y, color.Black)
	case 4:
		text.Draw(screen, msg, smallArcadeFont, x, y, color.Black)
	case 5:
		text.Draw(screen, msg, tinyArcadeFont, x, y, color.Black)
	default:
		messageError("Text Error: Unknown Font Code", screen)
	}
}

func messageError(msg string, screen *ebiten.Image) {
	text.Draw(screen, msg, smallArcadeFont, 250, 760, color.NRGBA{0xff, 0x00, 0x00, 0xff}) // Color Red
}

func messageSquare(sx, sy int, px, py float64, colorCode color.NRGBA, screen *ebiten.Image) {

	square, _ := ebiten.NewImage(sx, sy, ebiten.FilterNearest)
	square.Fill(colorCode)

	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct.
	opts.GeoM.Translate(px, py)
	// Draw the card image to the screen with an empty option
	screen.DrawImage(square, opts)
}

func main() {
	instruct() // Display current notes on project progress
	shuffle()  // Shuffle Deck

	updateMode = cardDeal // Clear Deal Mode

	//
	// Run Loop
	//
	if err := ebiten.Run(update, 1024, 768, 1, "AI POKER - You against eight AI Players!"); err != nil {
		log.Fatal(err)
	}
}
