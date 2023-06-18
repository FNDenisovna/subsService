package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// map, key - trigger fields, value - list of subs ids
var SubTrigger map[string][]int

// Storage of offers
var Offers map[string]Offer

// List of subscribers for current msg
var ToPrint []bool

// Count of Subs
var nSubs int

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	line, err := buf.ReadString('\n')
	if err != nil { //file is finished
		fmt.Printf("File is empty. %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("1 line = %v\n", line[2:])

	n, _ := strconv.Atoi(line[:1])
	m, _ := strconv.Atoi(line[2:3])

	fmt.Printf("n = %v, m = %v\n", n, m)

	//Read Subscribers
	for i := 0; i < n; i++ {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		//fmt.Print(line)
		//Successed reading sub from buffer
		ParseSub(line, i)
	}

	//Read Msgs
	for i := 0; i < m; i++ {
		line, err := buf.ReadBytes('\n') //buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		//Successed reading sub from buffer
		ParseMsg(line)
	}
}

func ParseSub(line string, i int) {
	line, _ = strings.CutSuffix(line, "\r\n")
	fields := strings.Fields(line)
	if len(fields) < 0 {
		return
	}

	//Get count of triggers
	triggerN, err := strconv.Atoi(fields[0])
	if err != nil || triggerN <= 0 {
		return
	}

	//Init list of triggers
	if len(SubTrigger) <= 0 {
		SubTrigger = make(map[string][]int)
	}
	//For n triggers do:
	for t := 2; t < len(fields); t++ {
		//Check new trigger in map of triggers
		sub, ok := SubTrigger[fields[t]]
		if !ok {
			//Add new trigger
			SubTrigger[fields[t]] = []int{i}
		} else {
			//Add new id_subscriber to array of trigger
			SubTrigger[fields[t]] = append(sub, i)
		}
		//To do, shipment is not used now
	}

	nSubs++
}

func ParseMsg(line []byte) {
	//Parse msg to model
	var msg Msg
	err := json.Unmarshal(line, &msg)
	if err != nil {
		return
	}

	//Clean subs
	ToPrint = make([]bool, nSubs)

	//Init of storage
	if len(Offers) <= 0 {
		Offers = make(map[string]Offer)
	}
	//Search offer witch have current event
	exOffer, ok := Offers[msg.ID]
	if !ok {
		exOffer = Offer{ID: msg.ID}
	}

	//Analize type Offer to getting count of struct fields
	//as well as find json tag of current field
	//to compare with SubTrigger map on key
	typeOffer := reflect.TypeOf(exOffer)

	for f := 0; f < typeOffer.NumField(); f++ {
		//Get value of every field in old offer and new
		exValue := reflect.ValueOf(&exOffer).Elem().Field(f)
		newValue := reflect.ValueOf(&msg.Offer).Elem().Field(f)
		//Compare
		if exValue.Interface() != newValue.Interface() {
			//If field have changes
			//Get his json-tag
			tag, ok := typeOffer.Field(f).Tag.Lookup("json")
			if ok {
				//Check any subsribers in map on json-tag = key
				trigger, ok := SubTrigger[tag]
				if ok {
					//Have any sub-s
					for sub := range trigger {
						ToPrint[sub] = true
					}
				}
			}

			exValue.Set(newValue)
		}
	}

	//Send msg as many subscribers this event have
	for _, send := range ToPrint {
		if send {
			fmt.Print(line)
		}
	}
	fmt.Print(line)
}

/*
///Version 1.0
///Idea: parse json as map and then search of changed fields witch have any subscribers

var msgMap map[string]interface{}

		json.Unmarshal(line, &msg)
		offer, ok := msgMap["offer"]
		if ok {
			changed, ok := offer.(map[string]interface{})
			if ok {
				findOfferChanges(&changed, &line)
			} else {
				//
			}
		}

func findOfferChanges(changed *(map[string]interface{}), line *[]byte) {
	if len(*changed) < 0 {
		return
	}

	//Сравниваем
	ind := -1
	if len(Offers) > 0 {
		ind = slices.IndexFunc(Offers, func(o Offer) bool {
			return o.ID == (*changed)["id"]
		})
	}
	if ind < 0 {
		return
	}

	//for k, v := range *changed {
	//	if offers[ind].Get(){}
	//}
}*/
