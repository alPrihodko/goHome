package home

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
)

/*
SmsEndpoint url to send sms
*/
const SmsEndpoint = "http://letsads.com/api"

/*
ReportAlert sends message to the system owner "emailTo"
*/
func ReportAlert(b string, s string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)

	m.SetHeader("Subject", s)
	m.SetBody("text/html", b)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, googleAccountName, googleAPIKey)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

/*
ExeCmd execs shell script
*/
func ExeCmd(cmd string, wg *sync.WaitGroup) {
	log.Println(cmd)
	parts := strings.Fields(cmd)
	var out []byte
	var err error
	if len(parts) > 2 {
		out, err = exec.Command(parts[0], parts[1], parts[2]).Output()
	} else if len(parts) == 2 {
		out, err = exec.Command(parts[0], parts[1]).Output()
	} else if len(parts) == 1 {
		out, err = exec.Command(parts[0]).Output()
	} else {
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println("error occured")
		log.Println(err.Error())
		log.Println(string(out))
	}
	log.Println(string(out))
	wg.Done()
}

/*
ExeCmdNoWait execs shell script
*/
func ExeCmdNoWait(cmd string) {
	log.Println(cmd)
	parts := strings.Fields(cmd)
	var out []byte
	var err error
	if len(parts) > 2 {
		out, err = exec.Command(parts[0], parts[1], parts[2]).Output()
	} else if len(parts) == 2 {
		out, err = exec.Command(parts[0], parts[1]).Output()
	} else if len(parts) == 1 {
		out, err = exec.Command(parts[0]).Output()
	} else {
		log.Println("Invalid arguments")
	}
	if err != nil {
		log.Println("error occured")
		log.Println(err.Error())
		log.Println(string(out))
	}
	log.Println(string(out))
}

/*
Exists returns file exists
*/
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)

	//log.Println(v)

	if err != nil {
		log.Println(err.Error())
	}

	if os.IsNotExist(err) {
		log.Println("seems no file: ", name)
		return false, nil
	}

	return err == nil, err
}

/*
Sms sends sms to recipient
*/
func Sms(from string, msg string, recipients []string) error {
	text := "<request><auth><login>380939760324</login><password>fktrc-fk</password>" +
		"</auth><message><from>" + from + "</from><text>" + msg + "</text>"

	for _, rpt := range recipients {
		text += "<recipient>" + rpt + "</recipient>"
	}

	text += "</message></request>"
	//log.Println(text)

	req, err := http.NewRequest("POST", SmsEndpoint, bytes.NewBuffer([]byte(text)))
	//req.Header.Set("X-Custom-Header", "myvalue")
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Cannot send alert: " + err.Error())
		return err
	}

	log.Println("response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	sbody := string(body)
	log.Println("response Body:", sbody)

	if strings.IndexAny(sbody, "Complete") < 0 {
		return errors.New(sbody)
	}

	return nil
}

/*
IsChipTimeZone - returns true between 23 - 7 local time where we have the chipest electricity
*/
func IsChipTimeZone() bool {
	hour := time.Now().Hour()
	ret := false
	if hour >= ElectroOnFrom || hour < ElectroOnTo {
		ret = true
	}

	return ret
}
