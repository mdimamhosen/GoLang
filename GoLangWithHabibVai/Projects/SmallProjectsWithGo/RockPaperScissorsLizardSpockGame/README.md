## Building a game of Rock, Paper, Scissors, Lizard, Spock

### Part One

1. **Setup**: Initialize a new Go project.

   - Create a new directory for your project.
   - Initialize a new Go module using `go mod init <module-name>`.
   - Create a main.go file.

2. **Constants**: Define constants for Rock, Paper, Scissors, Lizard, and Spock.

   - Use `const` to define the choices as strings.

3. **User Input**: Create a function to get the user's choice.

   - Use `fmt.Scanln` to read input from the user.
   - Validate the input to ensure it is one of the valid choices.
   - If the input is invalid, print "Invalid move. Game over."

4. **Computer Choice**: Create a function to randomly select the computer's choice.

   - Use `math/rand` to generate a random number.
   - Map the random number to one of the valid choices.
   - Store the computer's choice in a variable.

5. **Determine the Winner**: Compare the user's choice and the computer's choice to determine the winner.
   - Create a function that takes the user's choice and the computer's choice as arguments.
   - Use a series of `if` statements to compare the choices and determine the winner.
   - Print "You won! Congratulations!" if the user wins.
   - Print "You lost! hahaha, you suck!" if the computer wins.
   - Print "Damn, it's a draw. Let's play again!" if it's a draw and loop back to ask for the user's input again.

### Part Two

1. **Expand the Game**: Modify the program to include Lizard and Spock in addition to Rock, Paper, and Scissors.

   - Update the constants to include Lizard and Spock.
   - Update the user input validation to include the new choices.
   - Update the computer choice function to include the new choices.

2. **Input and Output**: Use the same input and output methods as before.

   - Ensure the input validation and random choice generation include all five choices.

3. **Loop on Draw**: Ensure the program loops back to ask for the user's input again if it's a draw.
   - Use a loop to repeatedly ask for the user's input until a winner is determined.

### Game Rules

1. **Winning Conditions**:

   - Rock crushes Scissors and Lizard.
   - Paper covers Rock and disproves Spock.
   - Scissors cuts Paper and decapitates Lizard.
   - Lizard eats Paper and poisons Spock.
   - Spock smashes Scissors and vaporizes Rock.

2. **Losing Conditions**:

   - The opposite of the winning conditions.

3. **Draw Conditions**:
   - If both the user and the computer choose the same option.

### Points Calculation

1. **Win**: Award 1 point to the user.
2. **Loss**: Award 1 point to the computer.
3. **Draw**: No points awarded.

### Example

- User chooses Rock, Computer chooses Scissors: User wins.
- User chooses Paper, Computer chooses Rock: User wins.
- User chooses Spock, Computer chooses Lizard: Computer wins.
- User chooses Scissors, Computer chooses Scissors: Draw.
