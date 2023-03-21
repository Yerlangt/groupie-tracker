package handle

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type DataToArtistPage struct {
	Concerts Coords
	Artist   Band
}

var (
	index, indParse         = template.ParseFiles("web/template/main.html")
	errPage, errParse       = template.ParseFiles("web/template/error.html")
	artistPage, artParse    = template.ParseFiles("web/template/artist.html")
	searchPage, searchParse = template.ParseFiles("web/template/search.html")
)

func (groupie *Groupie) MainHandler(w http.ResponseWriter, r *http.Request) {
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		ErrorPageExecute(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageExecute(w, http.StatusMethodNotAllowed)
		return
	}
	IndexPageExecute(w, groupie.Artists)
}

func (groupie *Groupie) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	log.Println(r.Method, "/artist")
	link := r.URL.Path
	linkList := strings.Split(link, "/")
	id, err := strconv.Atoi(linkList[len(linkList)-1])
	// fmt.Println(id)
	if err != nil || id > 52 || id < 1 {
		ErrorPageExecute(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPageExecute(w, http.StatusMethodNotAllowed)
		return
	}

	points := groupie.GetLocationPoints(groupie.Artists[id-1].LocationsList)
	res := DataToArtistPage{points, groupie.Artists[id-1]}
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	ArtistPageExecute(w, res)
}

func (groupie *Groupie) SearchHandler(w http.ResponseWriter, r *http.Request) {
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/search" {
		log.Println(r.URL.Path)
		ErrorPageExecute(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		log.Println(r.Method)
		ErrorPageExecute(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	for key := range r.PostForm {
		// fmt.Println(key, values)
		if !(key == "search") {
			ErrorPageExecute(w, http.StatusBadRequest)
			log.Println("Wrong key in search field")
			return
		}
	}
	filteredData = []Band{}
	searchInserted := r.FormValue("search")
	groupie.filteredData(searchInserted)
	allData := AllData{groupie.Artists, filteredData}
	SearchPageExecute(w, allData)
}

func ErrorPageExecute(writer http.ResponseWriter, status int) {
	writer.WriteHeader(status)
	if errParse == nil {
		if err := errPage.Execute(writer, status); err == nil {
			return
		}
	}
	http.Error(writer, errParse.Error(), http.StatusInternalServerError)
}

func IndexPageExecute(writer http.ResponseWriter, data []Band) {
	if indParse == nil {
		if err := index.Execute(writer, data); err == nil {
			return
		}
	}
	ErrorPageExecute(writer, http.StatusInternalServerError)
}

func ArtistPageExecute(writer http.ResponseWriter, data DataToArtistPage) {
	if artParse == nil {
		if err := artistPage.Execute(writer, data); err == nil {
			return
		}
	}
	ErrorPageExecute(writer, http.StatusInternalServerError)
}

func SearchPageExecute(writer http.ResponseWriter, data AllData) {
	if searchParse == nil {
		if err := searchPage.Execute(writer, data); err == nil {
			return
		}
	}
	ErrorPageExecute(writer, http.StatusInternalServerError)
}
