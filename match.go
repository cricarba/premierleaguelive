package main

import "encoding/json"

func UnmarshalMatch(data []byte) (Match, error) {
	var r Match
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Match) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Match struct {
	Fixture Fixture `json:"fixture"`
	Events  Events  `json:"events"`
}

type Events struct {
	PageInfo PageInfo  `json:"pageInfo"`
	Content  []Content `json:"content"`
}

type Content struct {
	ID        int64   `json:"id"`
	Time      Clock   `json:"time"`
	Type      string  `json:"type"`
	Text      string  `json:"text"`
	PlayerIDS []int64 `json:"playerIds,omitempty"`
}

type Clock struct {
	Secs  int64  `json:"secs"`
	Label string `json:"label"`
}

type PageInfo struct {
	Page       int64 `json:"page"`
	NumPages   int64 `json:"numPages"`
	PageSize   int64 `json:"pageSize"`
	NumEntries int64 `json:"numEntries"`
}

type Fixture struct {
	Gameweek           Gameweek      `json:"gameweek"`
	Kickoff            Kickoff       `json:"kickoff"`
	ProvisionalKickoff Kickoff       `json:"provisionalKickoff"`
	Teams              []TeamElement `json:"teams"`
	Replay             bool          `json:"replay"`
	Ground             Ground        `json:"ground"`
	NeutralGround      bool          `json:"neutralGround"`
	Status             string        `json:"status"`
	Phase              string        `json:"phase"`
	Outcome            string        `json:"outcome"`
	Attendance         int64         `json:"attendance"`
	Clock              Clock         `json:"clock"`
	FixtureType        string        `json:"fixtureType"`
	ExtraTime          bool          `json:"extraTime"`
	Shootout           bool          `json:"shootout"`
	BehindClosedDoors  bool          `json:"behindClosedDoors"`
	ID                 int64         `json:"id"`
}

type Gameweek struct {
	ID       int64 `json:"id"`
	Gameweek int64 `json:"gameweek"`
}

type Ground struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Source string `json:"source"`
	ID     int64  `json:"id"`
}

type Kickoff struct {
	Completeness int64  `json:"completeness"`
	Millis       int64  `json:"millis"`
	Label        string `json:"label"`
	GmtOffset    int64  `json:"gmtOffset"`
}

type TeamElement struct {
	Team  TeamTeam `json:"team"`
	Score int64    `json:"score"`
}

type TeamTeam struct {
	Name      string `json:"name"`
	Club      Club   `json:"club"`
	TeamType  string `json:"teamType"`
	ShortName string `json:"shortName"`
	ID        int64  `json:"id"`
}

type Club struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	ID   int64  `json:"id"`
}
