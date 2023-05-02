package models

import (
	"fmt"

	"elbarbaro.com/tornament-app/handlers/responses"
)

type Coach struct {
	Id         int16
	Name       string
	Percentage float32
}

type Team struct {
	Id      int16
	Name    string
	Logo    string
	Coach   Coach
	Players []Player
}

type Player struct {
	Id       int16
	Name     string
	Number   int8
	Country  string
	Position string
	Role     string
}

type Game struct {
	Id      int16
	TeamA   Team
	TeamB   Team
	Stadium string
	Date    string
	Status  string
}

type Tournament struct {
	Id      int16
	Name    string
	StartAt string
	EndAt   string
	Games   []Game
}

type Score struct {
	Id     int16
	Game   Game
	GoalsA int8
	GoalsB int8
}

type TournamentTeam struct {
	Id           int16
	Team         Team
	Points       int16
	WonMatches   int16
	LostMatches  int16
	DrawnMatches int16
}

type Leaderboard struct {
	Tournament Tournament
	Teams      []TournamentTeam
}

func (t Team) String() string {
	return fmt.Sprintf("\nId: %v, \nName: %v, \nCoach: %s", t.Id, t.Name, t.Coach.Name)
}

func (g Game) String() string {
	return fmt.Sprintf("\n%s vs %s \nDate: %s \nStadium: %s", g.TeamA.Name, g.TeamB.Name, g.Date, g.Stadium)
}

func (s Score) String() string {
	return fmt.Sprintf("\n%s:	%v\n%s:   %v\n%s", s.Game.TeamA.Name, s.GoalsA, s.Game.TeamB.Name, s.GoalsB, s.Game.Date)
}

func (tt TournamentTeam) String() string {
	return fmt.Sprintf("\n%s |%v %v %v %v", tt.Team.Name, tt.Points, tt.WonMatches, tt.LostMatches, tt.DrawnMatches)
}

func (t *Team) AddPlayer(player Player) {
	t.Players = append(t.Players, player)
}

func (t Team) ShowPlayers() {
	for _, player := range t.Players {
		fmt.Println(player)
	}
}

func (t Tournament) ShowGames() {
	fmt.Printf("Tournament: %s\n\n", t.Name)
	for _, game := range t.Games {
		fmt.Println(game)
	}
}

func (t *Tournament) AddGame(game Game) {
	t.Games = append(t.Games, game)
}

func (l *Leaderboard) AddTournamentTeam(tournamentTeam TournamentTeam) {
	l.Teams = append(l.Teams, tournamentTeam)
}

func (l *Leaderboard) Show() {
	fmt.Println("\nEquipo|	Pts	PG PP PE")
	for _, team := range l.Teams {
		fmt.Println(team)
	}
}

func (t Team) ToTeamResponse() responses.TeamResponse {
	return responses.TeamResponse{
		Id:   t.Id,
		Name: t.Name,
		Logo: t.Logo,
	}
}

func (p Player) ToPlayerResponse() responses.PlayerResponse {
	return responses.PlayerResponse{
		Id:       p.Id,
		Name:     p.Name,
		Number:   p.Number,
		Country:  p.Country,
		Position: p.Position,
		Role:     p.Role,
	}
}

func (c Coach) ToCoachResponse() responses.CoachResponse {
	return responses.CoachResponse{
		Id:         c.Id,
		Name:       c.Name,
		Percentage: c.Percentage,
	}
}
