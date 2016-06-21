package web

import "time"

//go:generate astTools -input-dir .

type Tour struct {
	UID      string    `json:"uid"`
	Year     int       `json:"year"`
	Etappes  []Etappe  `json:"etappes"`
	Cyclists []Cyclist `json:"cyclists"`
}

type Cyclist struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Etappe struct {
	UID            string        `json:"uid"`
	Day            time.Time     `json:"day"`
	StartLocation  string        `json:"startLocation"`
	FinishLocation string        `json:"finishLocation"`
	EtappeResult   *EtappeResult `json:"etappeResult"`
}

type EtappeResult struct {
	EtappeUID      string   `json:"etappeUid"`
	DayRankings    []string `json:"dayRankings"`
	YellowRankings []string `json:"yellowRankings"`
	ClimbRankings  []string `json:"climbRankings"`
	SprintRankings []string `json:"sprintRankings"`
}

// {"Annotation":"RestService","With":{"Path":"/api"}}
// @RestService(Path = "/api")
type TourService struct {
}

// {"Annotation":"RestOperation","With":{"Method":"GET", "Path":"/tour/{year}"}}
// @RestOperation(Method = "GET", Path = "/tour/{year}")
func (ts TourService) getTourOnUid(year string) (Tour, error) {
	return Tour{
		UID:      year,
		Year:     2016,
		Cyclists: []Cyclist{},
		Etappes:  []Etappe{},
	}, nil
}

// {"Annotation":"RestOperation","With":{"Method":"POST", "Path":"/tour/{year}/etappe"}}
// @RestOperation(Method = "POST", Path = "/tour/{year}/etappe")
func (ts *TourService) createEtappe(year string, etappe Etappe) (Etappe, error) {
	dateString := "2016-07-14"
	day, _ := time.Parse(dateString, dateString)
	return Etappe{
		UID:            "2",
		Day:            day,
		StartLocation:  "Paris",
		FinishLocation: "Roubaix",
	}, nil
}

// {"Annotation":"RestOperation","With":{"Method":"POST", "Path":"/tour/{year}/etappe/{etappeUid}"}}
// @RestOperation(Method = "PUT", Path = "/tour/:year/etappe/:etappeUid")
func (ts *TourService) addEtappeResults(year string, etappeUid string, results EtappeResult) error {
	return nil
}

// {"Annotation":"RestOperation","With":{"Method":"POST", "Path":"/tour/{year}/cyclist"}}
// @RestOperation(Method = "POST", Path = "/tour/{year}/cyclist")
func (ts *TourService) createCyclist(year string, cyclist Cyclist) (Cyclist, error) {
	return Cyclist{
		UID:    "42",
		Name:   "Boogerd, Michael",
		Points: 180,
	}, nil
}

// {"Annotation":"RestOperation","With":{"Method":"DELETE", "Path":"/tour/{year}/cyclist/{cyclistUid}"}}
// @RestOperation(Method = "DELETE", Path = "/tour/{year}/cyclist/{cyclistUid}")
func (ts *TourService) markCyclistAbondoned(year string, cyclistUid string) error {
	return nil
}