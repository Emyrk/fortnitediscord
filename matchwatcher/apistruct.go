package matchwatcher

const (
	// GameTypes
	Uknown_GameType int = iota
	Solo_GameType
	Duo_GameType
	Squad_GameType
)

type Match struct {
	Players []PlayerMatch
	Win     bool

	// Type of game
	//	0 --> Uknown
	//	1 --> Solos
	//  2 --> Duo
	//  3 --> Squad
	GameType int

	TopN int
}

func NewMatch() *Match {
	m := new(Match)

	return m
}

func (m *Match) AddPlayer(name string, kills int) {
	m.Players = append(m.Players, NewPlayerMatch(name, kills))
}

type PlayerMatch struct {
	Name  string
	Kills int
}

func NewPlayerMatch(name string, kills int) PlayerMatch {
	var pm PlayerMatch
	pm.Name = name
	pm.Kills = kills

	return pm
}

type GroupDetails struct {
	Wins          int    `json:"wins"`
	Top3          int    `json:"top3"`
	Top5          int    `json:"top5"`
	Top6          int    `json:"top6"`
	Top10         int    `json:"top10"`
	Top12         int    `json:"top12"`
	Top25         int    `json:"top25"`
	KD            string `json:"k/d"`
	Win           string `json:"win%"`
	Matches       int    `json:"matches"`
	Kills         int    `json:"kills"`
	TimePlayed    string `json:"timePlayed"`
	KillsPerMatch string `json:"killsPerMatch"`
	KillsPerMin   string `json:"killsPerMin"`
	Score         int    `json:"score"`
}

func (curr GroupDetails) GetMatch(prev GroupDetails) *Match {
	m := NewMatch()
	if curr.Wins > prev.Wins {
		m.Win = true
	} else if curr.Top3 > prev.Top3 {
		m.TopN = 3
	} else if curr.Top5 > prev.Top5 {
		m.TopN = 5
	} else if curr.Top6 > prev.Top6 {
		m.TopN = 6
	} else if curr.Top10 > prev.Top10 {
		m.TopN = 10
	} else if curr.Top12 > prev.Top12 {
		m.TopN = 12
	} else if curr.Top25 > prev.Top25 {
		m.TopN = 25
	} else {
		m.TopN = -1
	}

	return m
}

type PlayerStats struct {
	Group struct {
		Solo  GroupDetails
		Duo   GroupDetails `json:"duo"`
		Squad GroupDetails `json:"squad"`
	} `json:"group"`
	Info struct {
		AccountID string `json:"accountId"`
		Username  string `json:"username"`
		Platform  string `json:"platform"`
	} `json:"info"`
	LifetimeStats struct {
		Wins          int    `json:"wins"`
		Top3S         int    `json:"top3s"`
		Top5S         int    `json:"top5s"`
		Top6S         int    `json:"top6s"`
		Top10S        int    `json:"top10s"`
		Top12S        int    `json:"top12s"`
		Top25S        int    `json:"top25s"`
		KD            string `json:"k/d"`
		Win           string `json:"win%"`
		Matches       int    `json:"matches"`
		Kills         int    `json:"kills"`
		KillsPerMin   string `json:"killsPerMin"`
		TimePlayed    string `json:"timePlayed"`
		Score         int    `json:"score"`
		KillsPerMatch string `json:"killsPerMatch"`
	} `json:"lifetimeStats"`
}

/*
{
   "group":{
      "solo":{
         "wins":0,
         "top3":0,
         "top5":0,
         "top6":0,
         "top10":4,
         "top12":0,
         "top25":16,
         "k/d":"1.19",
         "win%":"0.00",
         "matches":48,
         "kills":57,
         "timePlayed":"5h 56m",
         "killsPerMatch":"1.19",
         "killsPerMin":"0.16",
         "score":2319
      },
      "duo":{
         "wins":2,
         "top3":0,
         "top5":23,
         "top6":0,
         "top10":0,
         "top12":51,
         "top25":0,
         "k/d":"1.44",
         "win%":"1.16",
         "matches":172,
         "kills":245,
         "timePlayed":"22h 11m",
         "killsPerMatch":"1.42",
         "killsPerMin":"0.18",
         "score":10421
      },
      "squad":{
         "wins":48,
         "top3":114,
         "top5":0,
         "top6":178,
         "top10":0,
         "top12":0,
         "top25":0,
         "k/d":"1.87",
         "win%":"8.32",
         "matches":577,
         "kills":989,
         "timePlayed":"3d 18h 6m",
         "killsPerMatch":"1.71",
         "killsPerMin":"0.18",
         "score":91191
      }
   },
   "info":{
      "accountId":"10c20c0b316b4b29b362dd2ba6fe2de5",
      "username":"Emyrks",
      "platform":"pc"
   },
   "lifetimeStats":{
      "wins":50,
      "top3s":114,
      "top5s":23,
      "top6s":178,
      "top10s":4,
      "top12s":51,
      "top25s":16,
      "k/d":"1.73",
      "win%":"6.27",
      "matches":797,
      "kills":1291,
      "killsPerMin":"0.18",
      "timePlayed":"4d 22h 13m",
      "score":103931,
      "killsPerMatch":"1.62"
   }
}


*/
