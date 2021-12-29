package main

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

type (
	Date   time.Time
	Record struct {
		Title        string   `xml:"title"`
		Name         string   `xml:"name"`
		Genre        string   `xml:"genre"`
		ReleaseDate  Date     `xml:"releasedate"`
		Label        string   `xml:"label"`
		Formats      string   `xml:"formats"`
		TrackListing []string `xml:"tracklisting>track"`
	}
	Records struct {
		XMLName xml.Name `xml:"records"`
		Records []Record `xml:"record"`
	}
	Release struct {
		Name       string `xml:"name"`
		TrackCount int    `xml:"trackCount"`
	}
	MatchingReleases struct {
		XMLName  xml.Name  `xml:"matchingReleases"`
		Releases []Release `xml:"release"`
	}
)

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var relaseDate string
	err := dec.DecodeElement(&relaseDate, &start)
	if err != nil {
		return err
	}
	date, err := convertStringToDate(relaseDate)
	if err != nil {
		return err
	}
	*d = Date(date)
	return nil
}

func (r *Records) Unmarshal(data []byte) error {
	return xml.Unmarshal(data, r)
}

func FilterRecords(records []Record, filterFn ...Filter) []Record {
	resp := make([]Record, len(records))
	j := 0
	for i := 0; i < len(records); i++ {
		matched := true
		for _, v := range filterFn {
			if !v(&records[i]) {
				matched = false
				break
			}
		}
		if matched {
			resp[j] = records[i]
			j += 1
		}
	}
	return resp[:j]
}

func convertStringToDate(date string) (time.Time, error) {
	var year, month, day int
	dataComponents := strings.Split(date, ".")
	year, err := strconv.Atoi(dataComponents[0])
	if err != nil {
		return time.Time{}, err
	}
	month, err = strconv.Atoi(strings.TrimLeft(dataComponents[1], "0"))
	if err != nil {
		return time.Time{}, err
	}
	day, err = strconv.Atoi(strings.TrimLeft(dataComponents[2], "0"))
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil

}
