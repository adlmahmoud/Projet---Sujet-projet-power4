package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Player struct {
	Name  string
	Color string
}

type Game struct {
	Board       [7][6]string
	Player1     Player
	Player2     Player
	CurrentTurn int
	Winner      string
	Status      string
	MoveCount   int
	StartTime   time.Time
}

type HistoryEntry struct {
	Player1 string
	Player2 string
	Winner  string
	Moves   int
	Date    string
}

type GameState struct {
	Game    *Game
	History []HistoryEntry
	mu      sync.Mutex
}

var state = &GameState{
	History: []HistoryEntry{},
}

func NewGame(p1Name, p1Color, p2Name, p2Color string) *Game {
	return &Game{
		Board:       [7][6]string{},
		Player1:     Player{Name: p1Name, Color: p1Color},
		Player2:     Player{Name: p2Name, Color: p2Color},
		CurrentTurn: 1,
		Status:      "en cours",
		MoveCount:   0,
		StartTime:   time.Now(),
	}
}

func (g *Game) ColorOf(playerName string) string {
	if playerName == g.Player1.Name {
		return g.Player1.Color
	}
	return g.Player2.Color
}

func (g *Game) CurrentPlayer() Player {
	if g.CurrentTurn == 1 {
		return g.Player1
	}
	return g.Player2
}

func (g *Game) PlayMove(col int) bool {
	if col < 0 || col >= 7 || g.Status != "en cours" {
		return false
	}

	for row := 5; row >= 0; row-- {
		if g.Board[col][row] == "" {
			g.Board[col][row] = g.CurrentPlayer().Name
			g.MoveCount++

			if g.checkWin(col, row) {
				g.Winner = g.CurrentPlayer().Name
				g.Status = "terminé"
				return true
			}

			if g.MoveCount >= 42 {
				g.Status = "nul"
				return true
			}

			if g.CurrentTurn == 1 {
				g.CurrentTurn = 2
			} else {
				g.CurrentTurn = 1
			}
			return true
		}
	}
	return false
}

func (g *Game) checkWin(col, row int) bool {
	player := g.Board[col][row]
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	for _, dir := range directions {
		count := 1
		for i := 1; i < 4; i++ {
			c, r := col+dir[0]*i, row+dir[1]*i
			if c >= 0 && c < 7 && r >= 0 && r < 6 && g.Board[c][r] == player {
				count++
			} else {
				break
			}
		}
		for i := 1; i < 4; i++ {
			c, r := col-dir[0]*i, row-dir[1]*i
			if c >= 0 && c < 7 && r >= 0 && r < 6 && g.Board[c][r] == player {
				count++
			} else {
				break
			}
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

var templates *template.Template

func loadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "templates/accueil/page_d_accueil.html")
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/init/init.html")
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/init", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	p1Name := r.FormValue("player1Name")
	p1Color := "red"
	p2Name := r.FormValue("player2Name")
	p2Color := "yellow"

	if p1Name == "" {
		p1Name = "Joueur 1"
	}
	if p2Name == "" {
		p2Name = "Joueur 2"
	}
	if p1Color == "" {
		p1Color = "red"
	}
	if p2Color == "" {
		p2Color = "yellow"
	}

	state.mu.Lock()
	state.Game = NewGame(p1Name, p1Color, p2Name, p2Color)
	state.mu.Unlock()

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	state.mu.Lock()
	defer state.mu.Unlock()

	if state.Game == nil {
		http.Redirect(w, r, "/init", http.StatusSeeOther)
		return
	}

	data := struct {
		Game              *Game
		Player1Active     bool
		Player2Active     bool
		CurrentPlayerName string
		Finished          bool
	}{
		Game:              state.Game,
		Player1Active:     state.Game.CurrentTurn == 1,
		Player2Active:     state.Game.CurrentTurn == 2,
		CurrentPlayerName: state.Game.CurrentPlayer().Name,
		Finished:          state.Game.Status != "en cours",
	}

	templates.ExecuteTemplate(w, "game.html", data)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	state.mu.Lock()
	if state.Game != nil {
		state.Game.PlayMove(col)

		if state.Game.Status != "en cours" {
			entry := HistoryEntry{
				Player1: state.Game.Player1.Name,
				Player2: state.Game.Player2.Name,
				Winner:  state.Game.Winner,
				Moves:   state.Game.MoveCount,
				Date:    time.Now().Format("02/01/2006 15:04"),
			}
			state.History = append(state.History, entry)
		}
	}
	state.mu.Unlock()

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	state.mu.Lock()
	if state.Game != nil {
		state.Game = NewGame(
			state.Game.Player1.Name,
			state.Game.Player1.Color,
			state.Game.Player2.Name,
			state.Game.Player2.Color,
		)
	}
	state.mu.Unlock()
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	state.mu.Lock()
	defer state.mu.Unlock()

	data := struct {
		History []HistoryEntry
	}{
		History: state.History,
	}

	templates.ExecuteTemplate(w, "History", data)
}

func main() {
	loadTemplates()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/init", initHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/move", moveHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/history", historyHandler)

	fmt.Println("🎮 Serveur Puissance 4 démarré sur http://localhost:8080")
	fmt.Println("📂 Ouvrez votre navigateur à l'adresse: http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erreur serveur: %v\n", err)
	}
}
