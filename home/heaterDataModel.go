package home

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
LIMIT max amount of history data calues
*/
const LIMIT = 100000

/*
HData is set of home data values, can be parsed from json
*/
type HData struct {
	TempInside    float32 `json:"TempInside, float"`
	TempOutside   float32 `json:"TempOutside, float"`
	TempHeater    float32 `json:"TempHeater, float"`
	TempReverse   float32 `json:"TempReverse, float"`
	TempEntryRoom float32 `json:"TempEntryRoom, float"`

	PumpState   bool `json:"PumpState, bool"`
	HeaterState bool `json:"HeaterState, bool"`
	Timestamp   int  `json:"Timestamp, int"`
	Index       int  `json:"index, int"`
}

/*
HistoryData is storage for recent states
*/
type HistoryData []*HData

/*
Len of HistoryData
*/
func (q HistoryData) Len() int { return len(q) }

/*
func (q historyData) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}


func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
*/

/*
Push HomeData to HistoryData
*/
func (q *HistoryData) Push(x interface{}) {
	n := len(*q)
	item := x.(*HData)
	item.Index = n
	*q = append(*q, item)
	if n > LIMIT {
		old := *q
		item := old[n-1]
		item.Index = -1 // for safety
		*q = old[0 : n-1]
		item = nil
	}
}

/*
Pop HomeData from HistoryData
*/
func (q *HistoryData) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

/*
ToJSON returns serialized hash
*/
func (q *HData) ToJSON() (d []byte, err error) {
	//now := int(time.Now().Unix())

	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/*
ToJSON returns serialized hash
*/
func (q *HistoryData) ToJSON(from int) (d []byte, err error) {
	old := *q
	sl := HistoryData{}
	now := int(time.Now().Unix())

	if from > 0 {
		var interval = 1

		if from > 60*60 && old.Len() > 50 {
			interval = 5
		}

		if from > 60*60*24 && old.Len() > 100 {
			interval = 10
		}

		for i := 0; i < old.Len(); i = i + interval {
			index := i
			if i >= old.Len() {
				index = old.Len() - 1
			}

			item := old[index]
			if item.Timestamp > (now - from) {
				sl.Push(item)
			}
		}
	} else {
		sl = old
	}

	b, err := json.Marshal(sl)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/*
// update modifies the priority and value of an Item in the queue.
func (q *historyData) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
*/

/*
ToGOB64 encodes to string
*/
func (q *HistoryData) ToGOB64() (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(&q)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

/*
FromGOB64 decodes from string
*/
func (q *HistoryData) FromGOB64(str string) error {
	//q := &HistoryData{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
		return err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(q)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return nil
}

/*
SerializeToFile writes slice to file
*/
func (q *HistoryData) SerializeToFile(name string) error {
	f, err := os.Create(name)
	defer f.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	w := bufio.NewWriter(f)
	str, err := q.ToGOB64()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	x, err := w.WriteString(str)
	log.Println("bytes written: ", x)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	w.Flush()
	return nil
}

/*
RestoreFromFile writes slice to file
*/
func (q *HistoryData) RestoreFromFile(name string) error {
	dat, err := ioutil.ReadFile(name)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = q.FromGOB64(bytes.NewBuffer(dat).String())

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
