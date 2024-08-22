### **Module 3: User Input and Control**

#### **1. Handling User Input**

**Reading Keyboard Input:**
- **Introduction to Keyboard Input:**
  - The keyboard is a common input device used in games for player control.
  - Ebiten provides functions to check for key presses and releases.

- **Using Ebiten for Keyboard Input:**
  - **Checking Key States:** Use `ebiten.IsKeyPressed` to determine if a specific key is pressed.

  ```go
  if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
      // Move player up
  }
  ```

- **Handling Key Presses and Releases:**
  - Implement logic to handle key events for player actions.

**Reading Mouse Input:**
- **Introduction to Mouse Input:**
  - The mouse is used for more precise input, such as aiming or clicking.

- **Using Ebiten for Mouse Input:**
  - **Getting Mouse Position:** Use `ebiten.MousePosition` to retrieve the current position of the mouse cursor.

  ```go
  x, y := ebiten.CursorPosition()
  ```

  - **Mouse Button States:** Use `ebiten.IsMouseButtonPressed` to check if a mouse button is pressed.

  ```go
  if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
      // Handle mouse click
  }
  ```

---

#### **2. Implementing Player Control**

**Controlling a Player Character:**
- **Moving the Player Character:**
  - Use input states to update the position of the player character.

  ```go
  func (g *Game) Update() error {
      if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
          g.playerY -= 2
      }
      if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
          g.playerY += 2
      }
      if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
          g.playerX -= 2
      }
      if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
          g.playerX += 2
      }
      return nil
  }
  ```

- **Handling Player Actions:**
  - Implement additional actions such as jumping or shooting based on input.

**Managing Movement and Actions:**
- **Updating Player Position:**
  - Continuously update the player’s position based on input within the game loop.

  ```go
  func (g *Game) Draw(screen *ebiten.Image) {
      // Draw the player at the updated position
      ebitenutil.DrawRect(screen, g.playerX, g.playerY, 20, 20, color.RGBA{255, 0, 0, 255})
  }
  ```

- **Implementing Actions:**
  - Trigger different actions like shooting or interacting with objects based on input.

---

#### **3. Input Responsiveness**

**Ensuring Responsive Controls:**
- **Immediate Feedback:**
  - Ensure that controls are responsive and provide immediate feedback to the player.

- **Handling Multiple Inputs:**
  - Implement logic to handle simultaneous inputs and prevent conflicts.

  ```go
  func (g *Game) Update() error {
      if ebiten.IsKeyPressed(ebiten.KeySpace) {
          // Perform an action
      }
      if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
          g.playerY -= 2
      }
      return nil
  }
  ```

**Improving Input Handling:**
- **Debouncing:**
  - Implement debouncing to handle rapid input events and avoid unintended actions.

- **Custom Input Configurations:**
  - Allow users to customize controls if applicable.

---

#### **4. Exercise: Implementing User Controls**

**Objective:**
Create a simple game where you can control a player character using keyboard and mouse inputs.

**Steps:**
1. **Setup:**
   - Install Ebiten if you haven’t already.
   - Create a new Go file for your project.

2. **Basic Player Control:**
   - Implement basic controls to move the player character using keyboard arrows.

3. **Mouse Interaction:**
   - Add functionality to handle mouse clicks, such as shooting or interacting with objects.

4. **Update and Draw:**
   - Continuously update the player’s position and draw it on the screen.

5. **Final Touches:**
   - Test the game to ensure controls are responsive.
   - Refine input handling to improve gameplay experience.

**Sample Code:**

```go
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
    "log"
)

type Game struct {
    playerX, playerY float64
}

func (g *Game) Update() error {
    if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
        g.playerY -= 2
    }
    if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
        g.playerY += 2
    }
    if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
        g.playerX -= 2
    }
    if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        g.playerX += 2
    }
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        x, y := ebiten.CursorPosition()
        // Handle mouse click (e.g., shoot or interact)
        log.Printf("Mouse clicked at (%d, %d)", x, y)
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.RGBA{0, 0, 255, 255}) // Fill background with blue
    ebitenutil.DrawRect(screen, g.playerX, g.playerY, 20, 20, color.RGBA{255, 0, 0, 255}) // Draw player
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 640, 480
}

func main() {
    game := &Game{}
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Player Control Example")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
```

---

This module focuses on handling user input and controlling a player character, providing a foundational understanding of how to integrate input into your game. If you need further details or adjustments, just let me know!