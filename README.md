```md
# Football League Table CLI (Go)

A command-line football league simulation built in Go.  
It supports team registration, match recording, automatic standings calculation, fixture validation, and league completion detection.

---

## 🚀 Features

- Create a league with a fixed number of teams
- Register teams interactively via CLI
- Case-insensitive team input
- Alias support (e.g. ARS, WOL, MAN U)
- Double round-robin enforcement (each team plays others twice)
- Prevents duplicate or invalid fixtures
- Live league table updates after each match
- Automatic detection of league completion
- Automatic champion declaration

---

## 📊 League Rules

If there are N teams:

- Each team plays every other team twice (home & away)
- Maximum games per team = (N - 1) × 2
- League ends automatically when all teams reach max games
- Winner is determined by:
  1. Points
  2. Goal Difference
  3. Goals Scored

---

## 🧠 Example

### Input

Enter number of teams in the league: 4

Enter team 1 name: Arsenal  
Enter team 2 name: Chelsea  
Enter team 3 name: Liverpool  
Enter team 4 name: Wolves  

Enter match: ARS, CHE, 2, 1  
Enter match: LIV, WOL, 1, 1  

---

### Output

UPDATED TABLE

Team                      P   W   D   L   GD   GF   Pts  
Arsenal                   1   1   0   0   1    2    3  
Chelsea                   1   0   0   1  -1    1    0  
...

---

## 🛠️ How to Run

### 1. Clone repo

git clone https://github.com/Perfected1/football-league-table.git  
cd football-league-table

### 2. Run program

go run .

---

## 📁 Project Structure

football-league-table/

├── main.go  
├── go.mod  
│  
├── team/  
│   └── team.go  
│  
└── league/  
    └── league.go  

---

## ⚙️ Core Concepts Used

- Struct-based modeling (Team, League)
- Maps for fast lookup
- Slice sorting for standings
- Input parsing with bufio
- String normalization for robustness
- Business rules enforcement (fixture limits, league completion)

---

## 🏆 League Completion

When all teams reach maximum matches:

🏁 LEAGUE HAS ENDED  
🏆 CHAMPION: Arsenal (XX pts)

---

## 📌 Author

Chike Jerry Nnamadim

---

## 📜 License

This project is open-source and free to use for learning and development purposes.
```
