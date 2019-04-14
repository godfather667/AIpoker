**AIpoker** is the initial concept design for an _**AI Based Poker Simulation**_.

The goal is to allow a human player to test their skills against eight AI players in Texas Holdem.

The user will be able to select a variety of AI players of different skills levels.

It also serves as a test bed for the **github.com/hajimehoshi/ebiten** Display Package.

* It displays a table with with a user and nine players.
* Currently, only the users cards are shown - The others only the card backs.
* The 'images' directory contains all the cards and other necessary images.
* At this time the check/fold/bet "Buttons" are displayed, only bet works.
* The functions to display *user chips* and *dealer chip* have bee created and aligned on display.
* There is a BUG in that *the cursor must be in the initial display box* for the complete image to be displayed. 
* This BUG is also in the example games I have compiled. The fix is on my TODO List.

The files were refactored into **ai.go, AIpoker.go, card.go  data.go, diag.go, images.go and text_input.go**

The next work will be setting up the functions to handle each AI player and the utility functions
for displaying chips and buttons and performing other display functions.
