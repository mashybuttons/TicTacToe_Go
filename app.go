package main
import "fmt"
//const board
//make board function
  //toggle piece function

func board () [][]string {
  temp := make([][]string, 3)
  for i := 0; i < len(temp); i++ {
    temp[i] = append(temp[i], "_", "_", "_")
  }
  return temp
}

func togglePiece (i, j int) (err string, won bool) {

  if legalMove(i, j) {
    b[i][j] = player
    oldPlayer := player
    if player == "X" {
      player = "O"
    } else {
      player = "X"
    }
    return "Move placed. Checking win: ", checkWin(i, j, oldPlayer)
  }

  return "Not valid move. Checking win: ", false

}

func legalMove (i, j int) bool {
  if i < 0 || i > len(b) || j < 0 || j > len(b) || b[i][j] != "_" {
    return false
  } 
  return true
}

func horizontal(i, j int, player string) bool {
  if b[i][0] == player && b[i][1] == player && b[i][2] == player {
    return true
  }
  return false
}

func vertical(i, j int, player string) bool {
  if b[0][j] == player && b[1][j] == player && b[2][j] == player {
    return true
  }
  return false
}

func rightDiag(i, j int, player string) bool {
  if b[0][0] == player && b[1][1] == player && b[2][2] == player {
    return true
  }
  return false
}

func leftDiag(i, j int, player string) bool {
  if b[0][2] == player && b[1][1] == player && b[2][0] == player {
    return true
  }
  return false
}

func checkWin(i, j int, player string) bool {
  if (i == 0 && j == 0) || (i == 2 && j == 0) {
    if rightDiag(i, j, player) || vertical(i, j, player) || horizontal(i, j, player) {
      return true
    }
    return false
  }

  if (i == 0 && j == 2) || (i == 2 && j == 2) {
    if leftDiag(i, j, player) || vertical(i, j, player) || horizontal(i, j, player) {
      return true
    }
    return false
  }

  if horizontal(i, j, player) || vertical(i, j, player) {
    return true
  }

  return false

}

var b = board()
var player = "X"

func main() {
  fmt.Println(b)

  
  fmt.Println(togglePiece(0,0))
  fmt.Println(b)

  fmt.Println(togglePiece(1,1))
  fmt.Println(b)
  
  fmt.Println(togglePiece(0,2))
  fmt.Println(b)

  fmt.Println(togglePiece(0,1))
  fmt.Println(b)

  fmt.Println(togglePiece(2,1))
  fmt.Println(b)

  fmt.Println(togglePiece(1,0))
  fmt.Println(b)

  fmt.Println(togglePiece(2,2))
  fmt.Println(b)

  fmt.Println(togglePiece(2,2))
  fmt.Println(b)

  fmt.Println(togglePiece(1,2))
  fmt.Println(b)

}