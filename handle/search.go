package handle

import (
	"strconv"
	"strings"
)

var filteredData []Band

func (groupie *Groupie) filteredData(searchInserted string) {
	search := strings.ToLower(searchInserted)
	for i, data := range groupie.Artists {
		place := strings.ToLower(searchInserted)
		if strings.Contains(searchInserted, "-") {
			place = strings.ToLower(fixedLocation(searchInserted))
		}
		for city := range data.Relationships {
			if strings.Contains(strings.ToLower(city), place) {
				if Check(data.Id) {
					filteredData = append(filteredData, groupie.Artists[i])
				}
			}
		}
		if strings.Contains(strings.ToLower(data.Name), search) ||
			strings.Contains(strings.ToLower(data.FirstAlbum), search) ||
			strings.Contains(strings.ToLower(strconv.Itoa(data.CreationDate)), search) {
			if Check(data.Id) {
				filteredData = append(filteredData, groupie.Artists[i])
			}
		}
		for _, members := range data.Members {
			if strings.Contains(strings.ToLower(members), search) {
				if Check(data.Id) {
					filteredData = append(filteredData, groupie.Artists[i])
				}
			}
		}
	}
}

func Check(id int) bool {
	for _, artists := range filteredData {
		if artists.Id == id {
			return false
		}
	}
	return true
}
