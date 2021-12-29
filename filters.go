package main

import "time"

type Filter func(*Record) bool

// getFilters creates filters based on command line arguments
func getFilters(flags cmdFlags) []Filter {
	var filters []Filter
	if *flags.minTrackCount > 0 {
		filters = append(filters, func(r *Record) bool {
			return len(r.TrackListing) >= *flags.minTrackCount
		})
	}
	if len(*flags.releasedBefore) > 0 {
		beforeDate, err := convertStringToDate(*flags.releasedBefore)
		if err == nil {
			filters = append(filters, func(r *Record) bool {
				t := time.Time(r.ReleaseDate)
				return t.Before(beforeDate)
			})
		}

	}
	return filters
}

func applyFilters(records []Record, filters []Filter) []Release {
	filteredRecords := FilterRecords(records, filters...)
	matchedRelases := make([]Release, len(filteredRecords))
	for i := range filteredRecords {
		matchedRelases[i] = Release{
			Name:       filteredRecords[i].Name,
			TrackCount: len(filteredRecords[i].TrackListing),
		}
	}
	return matchedRelases
}
