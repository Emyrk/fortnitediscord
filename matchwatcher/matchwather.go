package matchwatcher

import (
	"encoding/json"
	"fmt"
	"time"
)

// MatcheWatcher will look for a new match and record it
type MatcheWatcher struct {
	Players      []string
	PlayerDeets  map[string]*PlayerStats
	MatchChannel chan *Match
}

func NewMatchWatcher(players []string) *MatcheWatcher {
	w := new(MatcheWatcher)
	w.Players = players
	w.MatchChannel = make(chan *Match, 1000)

	return w
}

func (m *MatcheWatcher) Run() {
	index := 0
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		name := m.Players[index]
		match := m.DetectGame(name)
		if match != nil {
			matches := m.QuickPlayerCycle(index)
			combinedmatches := m.CombineMatches(append(matches, match...))
			for _, mat := range combinedmatches {
				m.MatchChannel <- mat
			}
			// TODO: Maybe sleep some more?
		}
		index++
	}
}

func (m *MatcheWatcher) CombineMatches(matches []*Match) []*Match {
	// For now just saw all duos are together, all squads are together
	// The key is the topN
	duos := make(map[int]*Match)
	squads := make(map[int]*Match)
	resp := make([]*Match, 0)

	for _, m := range matches {
		switch m.GameType {
		case Solo_GameType:
			resp = append(resp, m)
		case Duo_GameType:
			if duos[m.TopN] != nil {
				duos[m.TopN].Players = append(duos[m.TopN].Players, m.Players...)
			} else {
				duos[m.TopN] = m
			}
		case Squad_GameType:
			if squads[m.TopN] != nil {
				squads[m.TopN].Players = append(squads[m.TopN].Players, m.Players...)
			} else {
				squads[m.TopN] = m
			}
		default:
			resp = append(resp, m)
		}
	}

	for _, duo := range duos {
		resp = append(resp, duo)
	}
	for _, squad := range squads {
		resp = append(resp, squad)
	}
	return resp
}

func (m *MatcheWatcher) DetectGame(player string) []*Match {
	prev, ok := m.PlayerDeets[player]
	if !ok {
		// Eh
	}
	curr, err := GetStatisics(player)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if prev != nil && curr.LifetimeStats.Matches > prev.LifetimeStats.Matches {
		var resp []*Match

		// Played a game!
		//		What kind of game? It can be more than 1!
		if prev.Group.Squad.Matches != curr.Group.Squad.Matches {
			// Squad game!
			ma := curr.Group.Squad.GetMatch(prev.Group.Squad)
			ma.GameType = Squad_GameType
			ma.AddPlayer(player, curr.Group.Squad.Kills-prev.Group.Squad.Kills)
			resp = append(resp, ma)
		}

		if prev.Group.Duo.Matches != curr.Group.Duo.Matches {
			// Duo game!
			ma := curr.Group.Duo.GetMatch(prev.Group.Duo)
			ma.GameType = Duo_GameType
			ma.AddPlayer(player, curr.Group.Duo.Kills-prev.Group.Duo.Kills)
			resp = append(resp, ma)

		}

		if prev.Group.Solo.Matches != curr.Group.Solo.Matches {
			// Single game!
			ma := curr.Group.Solo.GetMatch(prev.Group.Solo)
			ma.GameType = Solo_GameType
			ma.AddPlayer(player, curr.Group.Solo.Kills-prev.Group.Solo.Kills)
			resp = append(resp, ma)

		}

		pdata, _ := json.Marshal(prev)
		ndata, _ := json.Marshal(curr)
		fmt.Printf("New Match!: \n %s \n %s\n", string(pdata), string(ndata))

		return resp
	}
	return nil
}

// QuickPlayerCycle will grab all matches that are played so we can start aggregating
func (m *MatcheWatcher) QuickPlayerCycle(index int) (matches []*Match) {
	for i := 0; i < len(m.Players); i++ {
		time.Sleep(100 * time.Millisecond)
		if i == index {
			continue
		}
	}

	return
}
