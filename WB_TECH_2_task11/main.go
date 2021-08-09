package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type event struct {
	Day   int
	Month int
	Year  int
	Event string
}
type output struct {
	Result []event `json:"result,omitempty"`
}
type outputDay struct {
	ResultDay event `json:"result,omitempty"`
}
type resultAndError struct {
	Result string `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}
type repo struct {
	myMap    map[string]string
	arrayDay []string
}

func makeJSON(w http.ResponseWriter, i interface{}) {
	jSon, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(jSon)
}

func (r *repo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var bbb bool
	var result resultAndError
	var evv string
	start := time.Now()
	Day := req.FormValue("day")
	Month := req.FormValue("month")
	Year := req.FormValue("year")
	eventT := req.FormValue("event")
	evv = Year + "/" + Month + "/" + Day
	day, _ := strconv.Atoi(Day)
	month, _ := strconv.Atoi(Month)
	year, _ := strconv.Atoi(Year)
	if _, err := time.Parse("2006/1/2", evv); err != nil&& req.URL.Path!="/events_for_month" {
		w.WriteHeader(400)
		return
	}
	switch req.URL.Path {
	case "/create_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodPost {
			r.myMap[evv] = eventT
			result = resultAndError{Result: "Событие создано успешно!"}
			makeJSON(w, result)
		}
	case "/update_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodPost {
			_, ok := r.myMap[evv]
			if ok {
				r.myMap[evv] = eventT
				result = resultAndError{Result: "Событие обновлено успешно!"}
			} else {
				result = resultAndError{Err: "Значение не найдено!"}
			}
			makeJSON(w, result)
		}
	case "/delete_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodPost {
			_, ok := r.myMap[evv]
			if ok {
				delete(r.myMap, evv)
				result = resultAndError{Result: "Событие удалено успешно!"}
			} else {
				result = resultAndError{Err: "Значение не найдено!"}
			}
			makeJSON(w, result)
		}
	case "/events_for_day":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodGet {
			value, ok := r.myMap[evv]
			if ok {
				newEvent := event{Day: day, Month: month, Year: year, Event: value}
				newOutput := outputDay{ResultDay: newEvent}
				makeJSON(w, newOutput)
			} else {
				result := resultAndError{Err: "Значение не найдено!"}
				makeJSON(w, result)
			}
		}

	case "/events_for_month":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodGet {
			var events []event
			for _, vvv := range r.arrayDay {
				value, ok := r.myMap[fmt.Sprintf("%s/%s/%s", Year, Month, vvv)]
				vv, _ := strconv.Atoi(vvv)
				if ok {
					newEvent := event{Day: vv, Month: month, Year: year, Event: value}
					events = append(events, newEvent)
				}
			}
			NewOutput := output{Result: events}
			makeJSON(w, NewOutput)

		}
	case "/events_for_week":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
		}
		if req.Method == http.MethodGet {
			var events []event
			layout := "2006/1/2"
			t, err := time.Parse(layout, evv)
			if err != nil {
				fmt.Printf("%v", err)
			}
			nDay := int(t.Weekday())
			if nDay == 0 {
				nDay = 7
			}
			for i := 1 - nDay; i <= 7-nDay; i++ {
				time1 := t.AddDate(0, 0, i)
				value, ok := r.myMap[fmt.Sprintf("%d/%d/%d", time1.Year(), time1.Month(), time1.Day())]
				if ok {
					newEvent := event{Day: time1.Day(), Month: int(time1.Month()), Year: time1.Year(), Event: value}
					events = append(events, newEvent)
				}
			}
			NewOutput := output{Result: events}
			makeJSON(w, NewOutput)
		}
	default:
		http.NotFound(w, req)
		bbb = true
	}
	if !bbb {
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error", err)
	}
	port := os.Getenv("Port")
	myMap := &repo{
		myMap: make(map[string]string),
		arrayDay: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
			"15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
	}
	fmt.Println("Start server!!!")
	log.Fatal(http.ListenAndServe(port, myMap))

}
