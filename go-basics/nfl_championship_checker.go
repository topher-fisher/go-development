package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TeamInfo holds the total number of wins and the specific calendar years they won
type TeamInfo struct {
	Wins  int
	Years []int
}

func main() {
	team := favoriteNFLTeam()
	fmt.Printf("\nYour favorite NFL team are the %v\n", team)

	info, hasWon := superBowlWins[team]
	if hasWon && info.Wins > 0 {
		fmt.Printf("\nThe %s have won the Super Bowl %d time(s).\n", team, info.Wins)

		// Prompt the user if they want to see the specific winning years
		if askForYears() {
			fmt.Printf("\nThe year(s) they won the Super Bowl is/are: %v\n", info.Years)
		}
	} else if hasWon {
		fmt.Printf("Unfortunately, the %s have never won a Super Bowl.\n", team)
	} else {
		fmt.Printf("Sorry, '%s' was not recognized as a valid NFL team name. Please check spelling and capitalization.\n", team)
	}

	fmt.Println("\nThank you for your interest!")
}

func favoriteNFLTeam() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is your favorite NFL team? ")
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func askForYears() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nWould you like to know the years of those Super Bowl wins? (y/n): ")
	scanner.Scan()
	input := strings.ToLower(strings.TrimSpace(scanner.Text()))
	return input == "y" || input == "yes"
}

// Global data map tracking Super Bowl wins and years for all 32 current NFL teams
var superBowlWins = map[string]TeamInfo{
	"Arizona Cardinals":     {Wins: 0, Years: []int{}},
	"Atlanta Falcons":       {Wins: 0, Years: []int{}},
	"Baltimore Ravens":      {Wins: 2, Years: []int{2000, 2012}},
	"Buffalo Bills":         {Wins: 0, Years: []int{}},
	"Carolina Panthers":     {Wins: 0, Years: []int{}},
	"Chicago Bears":         {Wins: 1, Years: []int{1985}},
	"Cincinnati Bengals":    {Wins: 0, Years: []int{}},
	"Cleveland Browns":      {Wins: 0, Years: []int{}},
	"Dallas Cowboys":        {Wins: 5, Years: []int{1971, 1977, 1992, 1993, 1995}},
	"Denver Broncos":        {Wins: 3, Years: []int{1997, 1998, 2015}},
	"Detroit Lions":         {Wins: 0, Years: []int{}},
	"Green Bay Packers":     {Wins: 4, Years: []int{1966, 1967, 1996, 2010}},
	"Houston Texans":        {Wins: 0, Years: []int{}},
	"Indianapolis Colts":    {Wins: 2, Years: []int{1970, 2006}},
	"Jacksonville Jaguars":  {Wins: 0, Years: []int{}},
	"Kansas City Chiefs":    {Wins: 4, Years: []int{1969, 2019, 2022, 2023}},
	"Las Vegas Raiders":     {Wins: 3, Years: []int{1976, 1980, 1983}},
	"Los Angeles Chargers":  {Wins: 0, Years: []int{}},
	"Los Angeles Rams":      {Wins: 2, Years: []int{1999, 2021}},
	"Miami Dolphins":        {Wins: 2, Years: []int{1972, 1973}},
	"Minnesota Vikings":     {Wins: 0, Years: []int{}},
	"New England Patriots":  {Wins: 6, Years: []int{2001, 2003, 2004, 2014, 2016, 2018}},
	"New Orleans Saints":    {Wins: 1, Years: []int{2009}},
	"New York Giants":       {Wins: 4, Years: []int{1986, 1990, 2007, 2011}},
	"New York Jets":         {Wins: 1, Years: []int{1968}},
	"Philadelphia Eagles":   {Wins: 2, Years: []int{2017, 2025}},
	"Pittsburgh Steelers":   {Wins: 6, Years: []int{1974, 1975, 1978, 1979, 2005, 2008}},
	"San Francisco 49ers":   {Wins: 5, Years: []int{1981, 1984, 1988, 1989, 1994}},
	"Seattle Seahawks":      {Wins: 2, Years: []int{2013, 2026}},
	"Tampa Bay Buccaneers":  {Wins: 2, Years: []int{2002, 2020}},
	"Tennessee Titans":      {Wins: 0, Years: []int{}},
	"Washington Commanders": {Wins: 3, Years: []int{1982, 1987, 1991}},
}
