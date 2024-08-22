### **Module 7: Final Project**

#### **1. Project Overview**

**Objective:**
- Create a complete, functioning game that incorporates all the concepts learned in previous modules: game loops and timing, basic graphics, collision detection and physics, game states and menus.

**Game Concept:**
- **Example Game:** A simple 2D platformer where a player controls a character that can move, jump, and interact with objects, with a main menu, gameplay state, and a game over screen.

**Requirements:**
- **Main Menu:** Allows the player to start the game, view instructions, or quit.
- **Gameplay:** Includes player movement, jumping, collisions with platforms, and obstacles.
- **Game Over Screen:** Shows when the player loses and gives options to restart or quit.

---

#### **2. Setting Up the Final Project**

**Project Structure:**
- **Directories and Files:**
  - `main.go` – Main game loop and entry point.
  - `game.go` – Game logic, state management, and updates.
  - `assets/` – Folder for images, sounds, and other assets.

**Example File Structure:**
```
/final_project
    /assets
        player.png
        platform.png
    main.go
    game.go
```

---

#### **3. Implementing the Main Menu**

**Menu Design:**
- **Features:**
  - Start Game
  - Instructions
  - Quit

**Sample Code:**

```go
func (g *Game) Update() error {
    switch g.state {
    case StateMenu:
        if ebiten.IsKeyPressed(ebiten.KeyEnter) {
            g.state = StatePlaying
        } else if ebiten.IsKeyPressed(ebiten.KeyI) {
            g.state = StateInstructions
        } else if ebiten.IsKeyPressed(ebiten.KeyQ) {
            return fmt.Errorf("Game exited")
        }
    case StatePlaying:
        // Gameplay logic
    case StateGameOver:
        if ebiten.IsKeyPressed(ebiten.KeyR) {
            g.state = StatePlaying
            // Reset game state
        } else if ebiten.IsKeyPressed(ebiten.KeyQ) {
            return fmt.Errorf("Game exited")
        }
    case StateInstructions:
        if ebiten.IsKeyPressed(ebiten.KeyEsc) {
            g.state = StateMenu
        }
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    switch g.state {
    case StateMenu:
        ebitenutil.DebugPrint(screen, "Press Enter to Start\nPress I for Instructions\nPress Q to Quit")
    case StatePlaying:
        // Draw gameplay
    case StateGameOver:
        ebitenutil.DebugPrint(screen, "Game Over! Press R to Restart or Q to Quit")
    case StateInstructions:
        ebitenutil.DebugPrint(screen, "Instructions:\nUse Arrow Keys to Move\nPress Space to Jump\nPress Esc to Return to Menu")
    }
}
```

---

#### **4. Implementing Gameplay**

**Player Movement and Jumping:**
- **Player Controls:**
  - Arrow keys to move.
  - Spacebar to jump.

**Sample Code:**

```go
func (g *Game) Update() error {
    if g.state == StatePlaying {
        if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
            g.playerX -= 2
        }
        if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
            g.playerX += 2
        }
        if ebiten.IsKeyPressed(ebiten.KeySpace) && g.onGround {
            g.velocityY = -10 // Jump
        }
        g.velocityY += 0.5 // Gravity
        g.playerY += g.velocityY

        // Handle collisions with platforms
        if g.playerY > screenHeight - 20 {
            g.playerY = screenHeight - 20
            g.velocityY = 0
            g.onGround = true
        } else {
            g.onGround = false
        }
    }
    return nil
}
```

**Drawing the Player and Platforms:**

```go
func (g *Game) Draw(screen *ebiten.Image) {
    if g.state == StatePlaying {
        playerImage := ebiten.NewImageFromImage(loadImage("assets/player.png"))
        platformImage := ebiten.NewImageFromImage(loadImage("assets/platform.png"))
        
        // Draw player
        screen.DrawImage(playerImage, &ebiten.DrawImageOptions{GeoM: ebiten.GeoM.Translate(g.playerX, g.playerY)})
        
        // Draw platforms
        screen.DrawImage(platformImage, &ebiten.DrawImageOptions{GeoM: ebiten.GeoM.Translate(100, screenHeight-20)})
    }
}
```

---

#### **5. Implementing the Game Over Screen**

**Game Over Logic:**
- **Trigger Game Over:**
  - Set the game state to `StateGameOver` when the player loses.

**Sample Code:**

```go
func (g *Game) Update() error {
    if g.state == StatePlaying {
        if g.playerY > screenHeight {
            g.state = StateGameOver
        }
    }
    return nil
}
```

---

#### **6. Testing and Refinement**

**Testing:**
- **Playtest the Game:**
  - Ensure all game states transition correctly.
  - Verify that player controls work as intended.
  - Check that collisions and physics behave as expected.

**Refinement:**
- **Adjust Parameters:**
  - Tweak values for movement speed, jump height, and gravity for a balanced experience.
- **Optimize Code:**
  - Refactor code for clarity and efficiency.

---

#### **7. Finalizing the Project**

**Documentation:**
- **Add Comments:**
  - Document your code with comments explaining key parts of the implementation.

**Deployment:**
- **Build and Package:**
  - Package your game for distribution or deployment.

**Sample Code for Building:**

```bash
go build -o final_project
```

---

This module guides you through creating a complete game project, integrating all concepts learned in previous modules. If you need more details or have specific questions, feel free to ask!