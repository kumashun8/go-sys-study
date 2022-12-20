package main

import "fmt"

type Event struct {
	data    string
	eventID int
}

type Shard struct {
	events  []Event
	shardID string
}

func newEvent(data string, id int) Event {
	return Event{
		data:    data,
		eventID: id,
	}
}

func (s Shard) getAllData(c chan []string) {
	list := make([]string, 0, len(s.events))
	for _, v := range s.events {
		list = append(list, v.data)
	}
	c <- list
}

func main() {
	shardList := []Shard{
		{
			shardID: "shard-0001",
			events: []Event{
				newEvent("hoge", 1),
			},
		},
		{
			shardID: "shard-0002",
			events: []Event{
				newEvent("fuga", 2),
				newEvent("pugi", 3),
			},
		},
		{
			shardID: "shard-0003",
			events:  []Event{},
		},
	}
	c := make(chan []string)
	for _, shard := range shardList {
		go shard.getAllData(c)
	}
	fmt.Println(<-c, <-c, <-c)
}
