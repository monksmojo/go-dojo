package main

import "fmt"

func main() {
	fmt.Println("Welcome to the type assertion")
	/**
	type assertion is the way to type casting an interface \
	back to type implementing the interface
	**/
	/**
	we will also see how multiple type interface are implemented by a single struct
	**/
	e1 := email{
		isSubscribed: true,
		msg:          "How are you doing",
		toEmail:      "blue.boy@gmal.com",
	}
	invoice(e1, e1)

	m1 := sms{
		isSubscribed:  false,
		msg:           "have you seen me in dark rooms",
		toPhoneNumber: "88775544",
	}
	invoice(m1, m1)

	e2 := email{
		isSubscribed: false,
		msg:          "Night owl is a dark theme",
		toEmail:      "dark.boy@gmal.com",
	}
	invoice(e2, e2)
	m2 := sms{
		isSubscribed:  true,
		msg:           "shake it move it don't settle !!!",
		toPhoneNumber: "8877442277",
	}
	invoice(m2, m2)
	reportGenerationV1(e1)
	reportGenerationV1(m1)
	reportGenerationV2(e2)
	reportGenerationV2(m2)
}

type expense interface {
	getMsgLength() float64
	cost(textLength float64) (charges float64)
}

type print interface {
	printMsg()
}

type email struct {
	isSubscribed bool
	msg          string
	toEmail      string
}

func (e email) getMsgLength() float64 {
	return float64(len(e.msg))
}

func (e email) cost(textLength float64) (charges float64) {
	if e.isSubscribed {
		return 0.01 * textLength
	} else {
		return 0.05 * textLength
	}
}

func (e email) printMsg() {
	fmt.Println("the message you sent is")
	fmt.Println(e.msg)
}

type sms struct {
	isSubscribed  bool
	msg           string
	toPhoneNumber string
}

func (s sms) getMsgLength() float64 {
	return float64(len(s.msg))
}

func (s sms) cost(textLength float64) (charges float64) {
	if s.isSubscribed {
		return 0.01 * textLength
	} else {
		return 0.05 * textLength
	}
}

func (s sms) printMsg() {
	fmt.Println("the message you sent is")
	fmt.Println(s.msg)
}

func invoice(e expense, p print) {
	p.printMsg()
	var mgsLength float64 = e.getMsgLength()
	fmt.Printf("the cost of the message is %.3f \n", e.cost(mgsLength))
}

func reportGenerationV1(e expense) {
	em, ok := e.(email)
	if ok {
		msgLength := em.getMsgLength()
		fmt.Printf("the email to %s  of length %d \t", em.toEmail, len(em.msg))
		fmt.Printf("cost %v \n", em.cost(msgLength))
	}

	sm, ok := e.(sms)
	if ok {
		msgLength := sm.getMsgLength()
		fmt.Printf("the sms to %s of length %d \t", sm.toPhoneNumber, len(sm.msg))
		fmt.Printf("cost %v \n", sm.cost(msgLength))
	}
}

func reportGenerationV2(e expense) {	
	switch v := e.(type) {
	case email:
		msgLength := v.getMsgLength()
		fmt.Printf("the email to %s  of length %d \t", v.toEmail, len(v.msg))
		fmt.Printf("cost %v \n", v.cost(msgLength))
	case sms:
		msgLength := v.getMsgLength()
		fmt.Printf("the sms to %s of length %d \t", v.toPhoneNumber, len(v.msg))
		fmt.Printf("cost %v \n", v.cost(msgLength))
	default:
		fmt.Println("type not recognized")

	}
}

// note an empty interface is implemented by all types
