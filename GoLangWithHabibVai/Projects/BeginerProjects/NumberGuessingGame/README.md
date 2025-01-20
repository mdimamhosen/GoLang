# Number Guessing Game

This project is a simple Number Guessing Game implemented in GoLang. The game involves the player guessing a number, and the computer providing feedback until the correct number is guessed.

## Game Logic

The game follows these steps:

1. **Game Initialization**: Set up the game environment, including generating a random number for the player to guess.
2. **Game Console UI**: Display the game interface to the player, using characters like `.` and `-` to represent different states.
3. **Result by Computer**: The computer evaluates the player's guess and provides feedback.
4. **Game Logic**: The core logic of the game, including handling player input and determining if the guess is correct.
5. **Game Result**: Display the result of the game, indicating whether the player has won or lost.

## Methods

### gameInit

```go
func gameInit() {
    // Game initialization : Number Guessing Game
}
```

- **Description**: Initializes the game by setting up necessary variables and generating a random number for the player to guess based on the chosen difficulty level.

### gameConsoleUI

```go
func gameConsoleUI() {
    // Game Console UI : Number Guessing Game by . and -
}
```

- **Description**: Displays the game interface to the player, using characters like `.` and `-` to represent different states.

### resultByComputer

```go
func resultByComputer() {
    // Result by Computer : Number Guessing Game
}
```

- **Description**: The computer generates a random number based on the chosen difficulty level.

### gameLogic

```go
func gameLogic() {
    // Game Logic : Number Guessing Game
}
```

- **Description**: Contains the core logic of the game, including handling player input, checking the guess against the target number, and providing feedback.

### gameResult

```go
func gameResult() {
    // Game Result : Number Guessing Game
}
```

- **Description**: Displays the result of the game, indicating whether the player has won or lost.

## Main Function

The `main` function orchestrates the flow of the game by calling the above methods in sequence.

```go
func main() {
    rand.Seed(time.Now().UnixNano())
    gameInit()
    if choice >= 1 && choice <= 3 {
        gameConsoleUI()
        resultByComputer()
        gameLogic()
        gameResult()
    }
}
```
