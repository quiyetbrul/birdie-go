# Flappy Bird in Go

This is a simple clone of the popular game **Flappy Bird** written in **Golang** using the **Ebiten** game development library. The game follows the same basic gameplay mechanics: the player controls a bird, attempting to fly between obstacles (pipes) without hitting them.

I also have it in CPP using the SFML library [here](https://github.com/quiyetbrul/flappy_bird), that's why repo is named birdie-go

## Features

- Press the spacebar to make the bird "jump" upwards.
- Fly through an ever-moving series of pipes without colliding.
- The game keeps track of the score based on how many pipes you successfully pass.
- A game-over screen is shown when the bird collides with a pipe, allowing you to restart the game.

## Gameplay Instructions

- **Press Space** to start the game.
- **Press Space** to make the bird jump.
- Try to avoid hitting the pipes!

## Open Source Libraries Used

This game uses the following open-source libraries and packages:

- [Ebiten](https://github.com/hajimehoshi/ebiten) (v2.5 or higher): A simple and efficient game library for Go, which provides graphics, input, and audio utilities.

## How to Build and Run the Game

### Prerequisites

- **Go**: Make sure you have Go installed. You can download it from [here](https://golang.org/dl/).
- **Ebiten**: Ebiten is a Go game library, and it will be automatically installed when running the commands below, as long as your `go.mod` file is set up correctly.

### Setup

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yourusername/flappybird-go
    cd flappybird-go
    ```

2. **Initialize Go Modules**:

    If your project isn’t already initialized, run the following command to create a `go.mod` file:

    ```bash
    go mod init flappybird
    ```

3. **Install Dependencies**:

    Ensure that all required dependencies are installed by running:

    ```bash
    go mod tidy
    ```

### Running the Game

To run the game on your machine, execute the following command:

```bash
go run main.go
```

### Building the Game

To build the game into an executable binary, run:

```bash
go build -o flappybird main.go
```

This will create an executable called flappybird (or flappybird.exe on Windows) that you can run directly:

```bash
./flappybird # On Unix-based systems
flappybird.exe # On Windows
```
