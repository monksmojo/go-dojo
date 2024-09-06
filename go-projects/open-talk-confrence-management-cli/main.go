package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var CONFERENCE_LIST []conference
var SPEAKER_LIST []speaker
var USER_LIST []user
var BOOKINGS []booking
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var choice string = "yes"
	for strings.ToLower(choice) == "yes" || strings.ToLower(choice) == "y" {

		var mainMenuOption int
		fmt.Println(strings.Repeat("#", 40), "-- WELCOME --", strings.Repeat("#", 40))
		fmt.Println()
		fmt.Println(strings.Repeat("#", 30), "-- OPEN TALK CONFERENCE MANAGER --", strings.Repeat("#", 30))
		fmt.Println()
		fmt.Println(strings.Repeat("#", 30), "-- Main Menu --", strings.Repeat("#", 30))
		fmt.Println("1. Register the conference venue: ")
		fmt.Println("2. view all registered conferences")
		fmt.Println("3. Register the speaker for the conference: ")
		fmt.Println("4. view all registered conferences: ")
		fmt.Println("6. Book tickets for the conference: ")
		fmt.Println("7. view all bookings ")
		fmt.Println("Select the option: ")
		mainMenuOption = takeIntegerInputFromUser()
		switch mainMenuOption {
		case 1:
			fmt.Println(strings.Repeat("#", 30), "-- Register A Conference --", strings.Repeat("#", 30))
			conf := createConference()
			fmt.Println(strings.Repeat("#", 30), "-- Confirmation To Register Venue --", strings.Repeat("#", 30))
			viewConference(conf)
			fmt.Printf("Are You Sure you want to add the conference: yes/no: ")
			commitChoice := takeStringInputFromUser()
			if strings.ToLower(commitChoice) == "yes" {
				CONFERENCE_LIST = append(CONFERENCE_LIST, conf)
				fmt.Println(strings.Repeat("#", 30), "-- Conference Is Registered --", strings.Repeat("#", 30))
			} else {
				fmt.Println("operation aborted!! ")
			}
		case 2:
			fmt.Println(strings.Repeat("#", 30), "-- All The Conferences --", strings.Repeat("#", 30))
			viewAllConferences()
			fmt.Println(strings.Repeat("#", 30), "-- End of Conference List --", strings.Repeat("#", 30))

		case 3:
			fmt.Println(strings.Repeat("#", 30), "-- Register A Speaker --", strings.Repeat("#", 30))
			spk := createSpeaker()
			fmt.Println(strings.Repeat("#", 30), "-- Confirmation To Register Speaker --", strings.Repeat("#", 30))
			viewSpeaker(spk)
			fmt.Println("Register speaker: yes/no")
			commitChoice := takeStringInputFromUser()
			if strings.ToLower(commitChoice) == "yes" {
				SPEAKER_LIST = append(SPEAKER_LIST, spk)
				fmt.Println(strings.Repeat("#", 30), "-- Speaker is added --", strings.Repeat("#", 30))
			} else {
				fmt.Println("operation aborted!! ")
			}

		case 4:
		default:

		}
		fmt.Println(strings.Repeat("#", 100))
		fmt.Printf("press y/Yes to continue or n/no to exit: ")
		fmt.Scanln(&choice)
		fmt.Println(strings.Repeat("#", 100))

	}
}

// conference object
type conference struct {
	conferenceId            string
	conferenceName          string
	conferenceVenue         string
	conferenceVenueCapacity int
	conferenceSlotUnitPrice float64
	conferenceStartDate     time.Time
	conferenceEndDate       time.Time
}

func (conf conference) isActive() bool {
	currentDate := time.Now()
	if currentDate.Before(conf.conferenceEndDate) && currentDate.After(conf.conferenceStartDate) {
		return true
	}
	return false
}

// speaker object
type speaker struct {
	speakerId      string
	speakerName    string
	speakerTopic   string
	conferenceName conference
}

// user object
type user struct {
	userId       string
	userFullName string
	userEmail    string
	bookings     []booking
}

// booking object
type booking struct {
	bookingId          string
	bookingFor         user
	selectedConference conference
	bookedSlots        int
	totalAmount        float64
	bookedDate         time.Time
	confirmationSent   bool
}

// #################### conference start

// creating conference
func createConference() conference {
	layout := "2006-01-02"
	conferenceInfoMap := getConferenceInfo()
	cvc, _ := strconv.Atoi(conferenceInfoMap["conferenceVenueCapacity"])
	slotUnitPrice, _ := strconv.ParseFloat(conferenceInfoMap["conferenceVenueCapacity"], 64)
	conferenceStartDate := dateParser(conferenceInfoMap["conferenceStartDate"], layout)
	conferenceEndDate := dateParser(conferenceInfoMap["conferenceEndDate"], layout)

	return conference{
		conferenceId:            uuid.NewString(),
		conferenceName:          conferenceInfoMap["conferenceName"],
		conferenceVenue:         conferenceInfoMap["conferenceVenue"],
		conferenceVenueCapacity: cvc,
		conferenceSlotUnitPrice: slotUnitPrice,
		conferenceStartDate:     conferenceStartDate,
		conferenceEndDate:       conferenceEndDate,
	}

}

// getting conference info from the user called by create conference()
func getConferenceInfo() map[string]string {
	fmt.Printf("enter the conference name: ")
	conferenceName := takeStringInputFromUser()
	conferenceVenue := getConferenceVenue()
	fmt.Printf("enter the capacity of the venue hall: ")
	conferenceVenueCapacity := takeStringInputFromUser()
	fmt.Printf("enter the conference booked slot unit price($): ")
	conferenceSlotUnitPrice := takeStringInputFromUser()
	fmt.Printf("enter conference start date 'yyyy:mm:dd' : ")
	conferenceStartDate := takeStringInputFromUser()
	fmt.Printf("enter conference end date 'yyyy:mm:dd' : ")
	conferenceEndDate := takeStringInputFromUser()
	return map[string]string{
		"conferenceName":          conferenceName,
		"conferenceVenue":         conferenceVenue,
		"conferenceVenueCapacity": conferenceVenueCapacity,
		"conferenceSlotUnitPrice": conferenceSlotUnitPrice,
		"conferenceStartDate":     conferenceStartDate,
		"conferenceEndDate":       conferenceEndDate,
	}
}

// getting venue of the conference from getConferenceInfo()
func getConferenceVenue() (conferenceVenue string) {
	fmt.Printf("enter the country where conference is held : ")
	country := takeStringInputFromUser()
	fmt.Printf("enter the city where conference is held : ")
	city := takeStringInputFromUser()
	fmt.Printf("enter the street address where conference is held : ")
	streetAddress := takeStringInputFromUser()
	fmt.Printf("enter the pin code where conference is held: ")
	pinCode := takeStringInputFromUser()
	return strings.Join([]string{country, city, streetAddress, pinCode}, "-")
}

// selecting a conference
func selectConference() conference {
	var conferenceSelection int
	var validSelection bool = false
	fmt.Println("select the current active and ongoing conferences")

	viewActiveConferences()
	for !validSelection {
		fmt.Println("enter the conference number: ")
		conferenceSelection = takeIntegerInputFromUser()
		if conferenceSelection > len(CONFERENCE_LIST) {
			fmt.Println("invalid conference selection:")
		} else {
			validSelection = true
		}
	}
	return CONFERENCE_LIST[conferenceSelection-1]
}

func viewActiveConferences() {
	record_present := false
	for i, v := range CONFERENCE_LIST {
		if v.isActive() {
			record_present = true
			fmt.Printf("%v \t", i)
			viewConference(v)
		}
	}

	if !record_present {
		fmt.Println("there are currently no active conference")
	}
}

func viewArchivedConferences() {
	record_present := false
	for i, v := range CONFERENCE_LIST {
		if !v.isActive() {
			record_present = true
			fmt.Printf("%v \t", i)
			viewConference(v)
		}
	}

	if !record_present {
		fmt.Println("all conference are active")
	}
}

func isConferenceListEmpty() bool {
	defer func() {
		if len(CONFERENCE_LIST) == 0 {
			fmt.Println("no registered conferences")
		}
	}()

	return len(CONFERENCE_LIST) == 0
}

func viewAllConferences() {
	if !isConferenceListEmpty() {
		for i, v := range CONFERENCE_LIST {
			fmt.Printf("%v \t", i)
			viewConference(v)
		}
	}
}

func viewConference(conf conference) {
	fmt.Printf("conference name: %v \nwith capacity: %d \n and booking price %.2f \n", conf.conferenceName, conf.conferenceVenueCapacity,
		conf.conferenceSlotUnitPrice)
	fmt.Println("conference venue: ")
	var venue []string = strings.Split(conf.conferenceVenue, "-")
	fmt.Println("conference country: ", venue[0])
	fmt.Println("conference city: ", venue[1])
	fmt.Println("conference address: ", venue[3], " - pinCode : ", venue[4])
	fmt.Println("conference start date", conf.conferenceStartDate)
	fmt.Println("conference end date", conf.conferenceEndDate)

}

// #################### conference end

// #################### speaker start

// creating of speaker model
func createSpeaker() speaker {
	speakerName, speakerTopic, conference := getSpeakerInfo()
	return speaker{
		speakerId:      uuid.NewString(),
		speakerName:    speakerName,
		speakerTopic:   speakerTopic,
		conferenceName: conference,
	}
}

// getting speaker info used in createSpeaker()
func getSpeakerInfo() (string, string, conference) {
	fmt.Printf("enter the speakerName: ")
	speakerName := takeStringInputFromUser()
	fmt.Printf("enter the speaker's Topic: ")
	speakerTopic := takeStringInputFromUser()
	conf := selectConference()
	return speakerName, speakerTopic, conf
}

func isSpeakerListEmpty() bool {
	defer func() {
		if len(SPEAKER_LIST) == 0 {
			fmt.Println("no registered speaker")
		}
	}()

	return len(SPEAKER_LIST) == 0
}

func viewAllSpeakers() {
	if !isConferenceListEmpty() {
		for i, v := range SPEAKER_LIST {
			fmt.Println("%d: \t", i+1)
			viewSpeaker(v)
		}
	}
}

func viewSpeaker(s speaker) {
	fmt.Printf("speaker : %v \n is speaking at: %v \n on topic : %v \n", s.speakerName, s.conferenceName.conferenceVenue, s.speakerTopic)
}

// #################### speaker end

// custom input functions to take input as string
func takeStringInputFromUser() string {
	val := readeStringFromReader()
	return val
}

// custom input functions to take input as int
func takeIntegerInputFromUser() int {
	val := readeStringFromReader()
	intVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Printf("error converting string %v to integer: %v\n", val[:len(val)-1], err)
	}
	return intVal
}

// reader to read string from the
func readeStringFromReader() string {
	val, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error taking input from the user")
	}
	return strings.TrimSpace(val[:len(val)-1])
}

// customDateParser
func dateParser(strDate, layout string) time.Time {
	date, err := time.Parse(layout, strDate)
	if err != nil {
		fmt.Printf("error parsing the date %v into the format %v \n", strDate, layout)
	}
	return date
}
