package pn25

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func LoyalCustomer() {
	customerLogs, err := fetchCustomerLogs(`D:\geek-practice\go-dojo\go-practicals\pn25\logs_both_days.txt`)
	if err != nil {
		
	}

}

type customerLog struct {
	logDate    time.Time
	page       string
	customerId string
}

func fetchCustomerLogs(path string) ([]customerLog, error) {
	customerLogs := make([]customerLog, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cl := strings.Split(scanner.Text(), ",")
		newCustomerLog := customerLog{
			logDate:    parseDate(cl[0]),
			page:       cl[1],
			customerId: cl[2],
		}
		customerLogs = append(customerLogs, newCustomerLog)
	}
	return customerLogs, nil
}

func parseDate(strDate string) time.Time {
	dateLayout := "2006-01-02"
	date, err := time.Parse(dateLayout, strDate)
	if err != nil {
		log.Fatal("error parsing the date")
	}
	return date
}
