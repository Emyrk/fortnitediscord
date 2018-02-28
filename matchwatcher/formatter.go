package matchwatcher

import "fmt"

//func FormatPlayerStats(name string, s *PlayerStats) string {
//	stats := s
//
//	str := "```\n"
//	str += fmt.Sprintf("%s LifetimeStats \n", name)
//	str += line("Wins", fmt.Sprintf("%s", stats["Wins"])) + "\n"
//	str += line("Kills", fmt.Sprintf("%s", stats["Kills"])) + "\n"
//	str += line("Matches", fmt.Sprintf("%s", stats["Matches Played"])) + "\n"
//
//	str += "```"
//	return str
//}

func FormatMatch(match Match) string {
	player := match.Players[0]
	win := "lost"
	if match.Win {
		win = "win"
	}

	str := "```\n"
	str += fmt.Sprintf("%s Just %s a game\n", player.Name, win)
	str += line("Kills", fmt.Sprintf("%d", player.Kills)) + "\n"

	str += "```"
	return str
}

func line(label string, value string) string {
	return fmt.Sprintf("%-15s: %s", label, value)
}
