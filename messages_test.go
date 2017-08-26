package ari

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelConnectedLine(t *testing.T) {

	actual := ChannelConnectedLine{}
	parseJSON("ChannelConnectedLine.json", &actual)

	expected := ChannelConnectedLine{
		Event: Event{
			Message: Message{
				Type: "ChannelConnectedLine",
			},
			Application: "demo",
			Timestamp:   parseTime("2017-08-26T15:20:12.596+0200"),
		},
		Channel: &Channel{
			ID:           "1503753612.1674",
			AccountCode:  "12",
			CreationTime: parseTime("2017-08-26T15:20:12.595+0200"),
			Name:         "Local/2103@spaeter-00000072;1",
			State:        "Down",
			Connected:    &CallerID{Name: "Bob", Number: "012345678"},
			Caller:       &CallerID{Name: "Alice", Number: "123456e0"},
			Dialplan:     &DialplanCEP{Context: "spaeter", Exten: "123456e0", Priority: 1},
		},
	}

	assert.EqualValues(t, expected, actual)
}

func parseTime(str string) *Time {
	timestamp, err := time.Parse(timeFormat, str)
	if err != nil {
		panic(err)
	}
	ts := Time(timestamp)
	return &ts
}

func parseJSON(filename string, v interface{}) {
	file, err := os.Open("testdata/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&v)
	if err != nil {
		panic(err)
	}
}
