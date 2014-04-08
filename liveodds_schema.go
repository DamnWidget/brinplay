// Copyright 2013 Oscar Campos <oscar.campos@member.fsf.org>
// See LICENSE file for details.

// Package DamnWidget/brinplay is a Go implementation of BetRadar InPlay protocol
//
// This package provides a BetRadar InPlay XML Feeds client and high level
// convenience methods for common operations with BetRadar InPlay
//
// This package is just a project to learn Go and is not intended to be used
// in production environments.

package liveodds

import (
	"encoding/xml"
	"time"
)

// BetRadarLiveOdds is a Struct ready to XML Unmarshal a BetRadarLiveOdds XML
// message from BetRadar in play live XML Feeds. It can be used just as:
//
// v := BetRadarLiveOdds{}
// xml_msg := `
//            <BetRadarLiveOdds status="alive" timestamp="1386870302430"
//                  xmlns="http://www.betradar.com/BetradarLiveOdds">
//            `
// err := xml.Unmatshal([]byte(xml_data), &v)
// if err != nil {
//     fmt.Printf("error: %v", err)
// }
type BetRadarLiveOdds struct {
	XMLName   xml.Name `xml:"BetradarLiveOdds"`
	Status    string   `xml:"status,attr"`
	Timestamp int64    `xml:"timestamp,attr"`
	XMLNS     string   `xml:"xmlns,attr"`
	Matches   []Match  `xml:"Match"`
	OddsType  []OddsType
}

// BetRadar does not follows the RFC3339 for Epoch Timestamps
func (t *BetRadarLiveOdds) Epoch() (epoch time.Time) {
	epoch = time.Unix(t.Timestamp/1000, 0)
	return epoch
}

type Match struct {
	Active    bool   `xml:"active,attr"`
	BetStatus string `xml:"betstatus,attr,omitempty"`
	MatchID   uint32 `xml:"matchid,attr"`
	MatchTime uint8  `xml:"matchtime,attr,omitempty"`
	MsgNR     uint16 `xml:"msgnr,attr,omitempty"`
	GameScore string `xml:"gamescore,attr,omitempty"`
	Score     string `xml:"score,attr,omitempty"`
	Status    string `xml:"status,attr,omitempty"`
	SetScores string `xml:"setscores,attr,omitempty"`
	Odds      []Odd
	Card      []Card
	Scores    []Score `xml:"Score"`
}

type Odd struct {
	OddsID           uint32 `xml:"id,attr"`
	Active           bool   `xml:"active,attr"`
	Changed          string `xml:"changed,attr"`
	Combination      uint8  `xml:"combination,attr"`
	FreeText         string `xml:"freetext,attr"`
	SpecialOddsValue string `xml:"specialoddsvalue,attr"`
	SubType          uint16 `xml:"subtype,attr"`
	Type             string `xml:"type,attr"`
	TypeID           uint16 `xml:"typeid,attr"`
	OddsField        []OddsField
}

type OddsField struct {
	Value  float32 `xml:",chardata"`
	Active bool    `xml:"active,attr,omitempty"`
	Type   string  `xml:"type,attr"`
}

type OddsType struct {
	Type      string `xml:"type,attr"`
	FreeText  string `xml:"freetext,attr,omitempty"`
	TypeID    uint16 `xml:"typeid,attr"`
	OddsField []TranslationOddsField
	Name      []Name
}

type TranslationOddsField struct {
	Type string `xml:"type,attr"`
	Name []Name
}

type Name struct {
	Value string `xml:",chardata"`
	Lang  string `xml:"lang,attr"`
}

type Card struct {
	CardID uint32 `xml:"id,attr"`
	Player string `xml:"player,attr"`
	Team   string `xml:"team,attr"`
	Time   uint8  `xml:"time,attr"`
	Type   string `xml:"type,attr"`
}

type Score struct {
	ScoreID     uint32 `xml:"id,attr"`
	Away        bool   `xml:"away,attr"`
	Home        bool   `xml:"home,attr"`
	Player      string `xml:"player,attr,omitempty"`
	ScoringTeam string `xml:"scoringteam,attr"`
	Time        int8   `xml:"time,attr"`
	Type        string `xml:"type,attr"`
}

type BookMakerStatus struct {
	XMLName     xml.Name `xml:"BookMakerStatus"`
	Timestamp   int64    `xml:"timestamp,attr"`
	Type        string   `xml:"type,attr"`
	BookmakerID uint16   `xml:"bookmakerid,attr"`
	Match       []Match  `xml:"Match,omitempty"`
}
