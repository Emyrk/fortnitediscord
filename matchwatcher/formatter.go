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

func FormatMatch(match *Match) string {
	// Format Match Header
	str := ""
	str += fmt.Sprintf(":game_die: Looks like someone just played a %s game!\n", match.MatchTypeString())
	if match.Win {
		str += " The match was a victory!\n"
	} else {
		str += " The match was a defeat :cry:\n"
	}

	totalKills := 0
	playerstr := "Player Stats:\n```\n"
	for _, p := range match.Players {
		playerstr += fmt.Sprintf("Player %s", p.Name) + "\n"
		playerstr += " " + line("Kills:", fmt.Sprintf("%d", p.Kills)) + "\n"
		totalKills += p.Kills
	}
	playerstr += "```\n"

	str += "Game Stats \n```\n"
	placestr := ">25"
	if match.TopN != -1 {
		placestr = fmt.Sprintf("Top %d", match.TopN) + "\n"
	}
	str += line("Place:", placestr)
	if match.MatchTypeString() != "solo" {
		str += line("TotalKills:", fmt.Sprintf("%d", totalKills)) + "\n"
	}

	str += "```\n"
	str += playerstr
	return str
}

func line(label string, value string) string {
	return fmt.Sprintf("%-15s: %s", label, value)
}
