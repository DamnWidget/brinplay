package liveodds

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
	"time"
)

// Common message for failed tests
const failed_msg = "%s: expected %v, got %v"

// We need this to can use whatever result we need
type CurrentResult interface{}
type ExpectedResult interface{}

type xmlTest struct {
	n        CurrentResult
	expected ExpectedResult
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadXMLFixture(fixture string) (f BetRadarLiveOdds) {
	msg, err := ioutil.ReadFile(fixture)
	check(err)

	feed := BetRadarLiveOdds{}
	unmarshal_err := xml.Unmarshal(msg, &feed)
	if unmarshal_err != nil {
		panic(unmarshal_err)
	}
	return feed
}

func TestSimpleAlive(t *testing.T) {

	feed := LoadXMLFixture("fixtures/alive.xml")

	ti := time.Date(2013, time.December, 15, 17, 56, 30, 0, time.Local)
	var xmlTests = []xmlTest{
		{feed.Status, "alive"},
		{len(feed.Matches), 1},
		{ti, feed.Epoch()},
		{feed.XMLNS, "http://www.betradar.com/BetradarLiveOdds"},
	}

	for _, tt := range xmlTests {
		if tt.n != tt.expected {
			t.Errorf(failed_msg, "TestSimpleAlive", tt.expected, tt.n)
		}
	}

}

func TestSimpleChange(t *testing.T) {

	feed := LoadXMLFixture("fixtures/change.xml")
	var xmlTests = []xmlTest{
		{feed.Status, "change"},
		{len(feed.Matches), 1},
		{feed.Matches[0].Active, true},
		{feed.Matches[0].BetStatus, "stopped"},
		{feed.Matches[0].MatchID, uint32(867278)},
		{feed.Matches[0].MsgNR, uint16(2)},
		{feed.Matches[0].Score, "-:-"},
		{feed.Matches[0].Status, "not_started"},
		{len(feed.Matches[0].Odds), 5},

		// Odds[0]
		{feed.Matches[0].Odds[0].Active, true},
		{feed.Matches[0].Odds[0].Changed, "false"},
		{feed.Matches[0].Odds[0].Combination, uint8(0)},
		{feed.Matches[0].Odds[0].FreeText, "Next goal"},
		{feed.Matches[0].Odds[0].OddsID, uint32(78557)},
		{feed.Matches[0].Odds[0].SpecialOddsValue, "0:0"},
		{feed.Matches[0].Odds[0].SubType, uint16(13)},
		{feed.Matches[0].Odds[0].Type, "ft3w"},
		{feed.Matches[0].Odds[0].TypeID, uint16(6)},

		// OddsField[0]
		{feed.Matches[0].Odds[0].OddsField[0].Active, true},
		{feed.Matches[0].Odds[0].OddsField[0].Type, "1"},
		{feed.Matches[0].Odds[0].OddsField[0].Value, float32(1.4)},
		// OddsField[1]
		{feed.Matches[0].Odds[0].OddsField[1].Active, true},
		{feed.Matches[0].Odds[0].OddsField[1].Type, "x"},
		{feed.Matches[0].Odds[0].OddsField[1].Value, float32(7.0)},
		// OddsField[2]
		{feed.Matches[0].Odds[0].OddsField[2].Active, true},
		{feed.Matches[0].Odds[0].OddsField[2].Type, "2"},
		{feed.Matches[0].Odds[0].OddsField[2].Value, float32(4.05)},

		// Odds[1]
		{feed.Matches[0].Odds[1].Active, true},
		{feed.Matches[0].Odds[1].Changed, "false"},
		{feed.Matches[0].Odds[1].Combination, uint8(0)},
		{feed.Matches[0].Odds[1].FreeText, ""},
		{feed.Matches[0].Odds[1].OddsID, uint32(78558)},
		{feed.Matches[0].Odds[1].SpecialOddsValue, "2.5"},
		{feed.Matches[0].Odds[1].SubType, uint16(0)},
		{feed.Matches[0].Odds[1].Type, "to"},
		{feed.Matches[0].Odds[1].TypeID, uint16(5)},

		// OddsField[0]
		{feed.Matches[0].Odds[1].OddsField[0].Active, true},
		{feed.Matches[0].Odds[1].OddsField[0].Type, "o"},
		{feed.Matches[0].Odds[1].OddsField[0].Value, float32(2.4)},
		// OddsField[1]
		{feed.Matches[0].Odds[1].OddsField[1].Active, true},
		{feed.Matches[0].Odds[1].OddsField[1].Type, "u"},
		{feed.Matches[0].Odds[1].OddsField[1].Value, float32(1.45)},

		// Odds[2]
		{feed.Matches[0].Odds[2].Active, true},
		{feed.Matches[0].Odds[2].Changed, "false"},
		{feed.Matches[0].Odds[2].Combination, uint8(0)},
		{feed.Matches[0].Odds[2].FreeText, "Halftime - Who wins the rest?"},
		{feed.Matches[0].Odds[2].OddsID, uint32(79538)},
		{feed.Matches[0].Odds[2].SpecialOddsValue, "0:0"},
		{feed.Matches[0].Odds[2].SubType, uint16(20)},
		{feed.Matches[0].Odds[2].Type, "ft3w"},
		{feed.Matches[0].Odds[2].TypeID, uint16(6)},

		// OddsField[0]
		{feed.Matches[0].Odds[2].OddsField[0].Active, true},
		{feed.Matches[0].Odds[2].OddsField[0].Type, "1"},
		{feed.Matches[0].Odds[2].OddsField[0].Value, float32(2.0)},
		// OddsField[1]
		{feed.Matches[0].Odds[2].OddsField[1].Active, true},
		{feed.Matches[0].Odds[2].OddsField[1].Type, "x"},
		{feed.Matches[0].Odds[2].OddsField[1].Value, float32(2.15)},
		// OddsField[2]
		{feed.Matches[0].Odds[2].OddsField[2].Active, true},
		{feed.Matches[0].Odds[2].OddsField[2].Type, "2"},
		{feed.Matches[0].Odds[2].OddsField[2].Value, float32(6.75)},

		// Odds[3]
		{feed.Matches[0].Odds[3].Active, true},
		{feed.Matches[0].Odds[3].Changed, "false"},
		{feed.Matches[0].Odds[3].Combination, uint8(0)},
		{feed.Matches[0].Odds[3].FreeText, "Who wins the rest of the match?"},
		{feed.Matches[0].Odds[3].OddsID, uint32(78560)},
		{feed.Matches[0].Odds[3].SpecialOddsValue, "0:0"},
		{feed.Matches[0].Odds[3].SubType, uint16(4)},
		{feed.Matches[0].Odds[3].Type, "ft3w"},
		{feed.Matches[0].Odds[3].TypeID, uint16(6)},

		// OddsField[0]
		{feed.Matches[0].Odds[3].OddsField[0].Active, true},
		{feed.Matches[0].Odds[3].OddsField[0].Type, "1"},
		{feed.Matches[0].Odds[3].OddsField[0].Value, float32(1.45)},
		// OddsField[1]
		{feed.Matches[0].Odds[3].OddsField[1].Active, true},
		{feed.Matches[0].Odds[3].OddsField[1].Type, "x"},
		{feed.Matches[0].Odds[3].OddsField[1].Value, float32(3.65)},
		// OddsField[2]
		{feed.Matches[0].Odds[3].OddsField[2].Active, true},
		{feed.Matches[0].Odds[3].OddsField[2].Type, "2"},
		{feed.Matches[0].Odds[3].OddsField[2].Value, float32(7.25)},

		// Odds[4]
		{feed.Matches[0].Odds[4].Active, true},
		{feed.Matches[0].Odds[4].Changed, "true"},
		{feed.Matches[0].Odds[4].Combination, uint8(0)},
		{feed.Matches[0].Odds[4].FreeText, "Which team has kick off?"},
		{feed.Matches[0].Odds[4].OddsID, uint32(78559)},
		{feed.Matches[0].Odds[4].SpecialOddsValue, "-1"},
		{feed.Matches[0].Odds[4].SubType, uint16(2)},
		{feed.Matches[0].Odds[4].Type, "ft2w"},
		{feed.Matches[0].Odds[4].TypeID, uint16(7)},

		// OddsField[0]
		{feed.Matches[0].Odds[4].OddsField[0].Active, true},
		{feed.Matches[0].Odds[4].OddsField[0].Type, "1"},
		{feed.Matches[0].Odds[4].OddsField[0].Value, float32(1.8)},
		// OddsField[1]
		{feed.Matches[0].Odds[4].OddsField[1].Active, true},
		{feed.Matches[0].Odds[4].OddsField[1].Type, "2"},
		{feed.Matches[0].Odds[4].OddsField[1].Value, float32(1.8)},
	}

	for _, tt := range xmlTests {
		if tt.n != tt.expected {
			t.Errorf(failed_msg, "TestSimpleChange", tt.expected, tt.n)
		}
	}
}

func TestSimpleTranslation(t *testing.T) {
	feed := LoadXMLFixture("fixtures/translation.xml")
	var xmlTests = []xmlTest{
		{feed.Status, "translation"},
		{len(feed.OddsType), 2},
		{feed.OddsType[0].Type, "3w"},
		{feed.OddsType[0].TypeID, uint16(2)},
		{feed.OddsType[0].Name[0].Lang, "en"},
		{feed.OddsType[0].Name[0].Value, "3way"},
		{feed.OddsType[0].OddsField[0].Type, "1"},
		{feed.OddsType[0].OddsField[0].Name[0].Lang, "en"},
		{feed.OddsType[0].OddsField[0].Name[0].Value, "1"},
		{feed.OddsType[0].OddsField[1].Type, "x"},
		{feed.OddsType[0].OddsField[1].Name[0].Lang, "en"},
		{feed.OddsType[0].OddsField[1].Name[0].Value, "x"},
		{feed.OddsType[0].OddsField[2].Type, "2"},
		{feed.OddsType[0].OddsField[2].Name[0].Lang, "en"},
		{feed.OddsType[0].OddsField[2].Name[0].Value, "2"},

		{feed.OddsType[1].Type, "hc"},
		{feed.OddsType[1].TypeID, uint16(4)},
		{feed.OddsType[1].Name[0].Lang, "en"},
		{feed.OddsType[1].Name[0].Value, "Handicap"},
		{feed.OddsType[1].OddsField[0].Type, "1"},
		{feed.OddsType[1].OddsField[0].Name[0].Lang, "en"},
		{feed.OddsType[1].OddsField[0].Name[0].Value, "1"},
		{feed.OddsType[1].OddsField[1].Type, "x"},
		{feed.OddsType[1].OddsField[1].Name[0].Lang, "en"},
		{feed.OddsType[1].OddsField[1].Name[0].Value, "x"},
		{feed.OddsType[1].OddsField[2].Type, "2"},
		{feed.OddsType[1].OddsField[2].Name[0].Lang, "en"},
		{feed.OddsType[1].OddsField[2].Name[0].Value, "2"},
	}

	for _, tt := range xmlTests {
		if tt.n != tt.expected {
			t.Errorf(failed_msg, "TestSimpleTranslation", tt.expected, tt.n)
		}
	}
}

func TestSimpleScore(t *testing.T) {
	feed := LoadXMLFixture("fixtures/score.xml")
	var xmlTests = []xmlTest{
		{feed.Status, "score"},
		{len(feed.Matches), 1},
		{feed.Matches[0].Active, true},
		{feed.Matches[0].BetStatus, "stopped"},
		{feed.Matches[0].MatchTime, uint8(3)},
		{feed.Matches[0].MsgNR, uint16(10)},
		{feed.Matches[0].Score, "0:1"},
		{feed.Matches[0].SetScores, "0:1"},
		{feed.Matches[0].Status, "1p"},
		{len(feed.Matches[0].Card), 0},
		{feed.Matches[0].Scores[0].Away, true},
		{feed.Matches[0].Scores[0].Home, false},
		{feed.Matches[0].Scores[0].ScoreID, uint32(66664)},
		{feed.Matches[0].Scores[0].Player, ""},
		{feed.Matches[0].Scores[0].ScoringTeam, "away"},
		{feed.Matches[0].Scores[0].Time, int8(-1)},
		{feed.Matches[0].Scores[0].Type, "live"},
	}

	for _, tt := range xmlTests {
		if tt.n != tt.expected {
			t.Errorf(failed_msg, "TestSimpleScore", tt.expected, tt.n)
		}
	}
}

func TestSimpleCard(t *testing.T) {
	feed := LoadXMLFixture("fixtures/card.xml")
	var xmlTests = []xmlTest{
		{feed.Status, "score"},
		{len(feed.Matches), 1},
		{feed.Matches[0].Active, true},
		{feed.Matches[0].BetStatus, "stopped"},
		{feed.Matches[0].MatchID, uint32(1355389)},
		{feed.Matches[0].Score, "3:0"},
		{feed.Matches[0].Status, "ended"},
		{len(feed.Matches[0].Card), 2},
		{feed.Matches[0].Card[0].CardID, uint32(111556)},
		{feed.Matches[0].Card[0].Player, "Ramires"},
		{feed.Matches[0].Card[0].Team, "home"},
		{feed.Matches[0].Card[0].Time, uint8(70)},
		{feed.Matches[0].Card[0].Type, "yellow"},
		{feed.Matches[0].Card[1].CardID, uint32(111555)},
		{feed.Matches[0].Card[1].Player, "Fuentes, Ismael"},
		{feed.Matches[0].Card[1].Team, "away"},
		{feed.Matches[0].Card[1].Time, uint8(67)},
		{feed.Matches[0].Card[1].Type, "yellow"},
	}

	for _, tt := range xmlTests {
		if tt.n != tt.expected {
			t.Errorf(failed_msg, "TestSimpleCard", tt.expected, tt.n)
		}
	}
}

func TestBookMakerstatus(t *testing.T) {
	v := &BookMakerStatus{Type: "error", BookmakerID: 1234}
	v.Timestamp = time.Date(
		2013, time.December, 15, 17, 56, 30, 0, time.Local).Unix()

	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	feed := BookMakerStatus{}
	unmarshal_err := xml.Unmarshal(output, &feed)
	if unmarshal_err != nil {
		panic(unmarshal_err)
	}

	if feed.Timestamp != v.Timestamp || feed.BookmakerID != v.BookmakerID {
		t.Errorf("%v and %v are not equal", feed, v)
	}

	m := Match{Active: true, MatchID: 12345678}
	v.Match = append(v.Match, m)
	output, err = xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	feed = BookMakerStatus{}
	unmarshal_err = xml.Unmarshal(output, &feed)
	if unmarshal_err != nil {
		panic(unmarshal_err)
	}

	if feed.Match[0].MatchID != m.MatchID {
		t.Error("Matches ID does not match")
	}
}
