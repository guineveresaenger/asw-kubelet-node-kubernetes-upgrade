package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
)

type Data struct {
	Reservations []struct {
		Instances []struct {
			InstanceId string `json:"InstanceId"`
		}
	}
}

func main() {

	nodeName := flag.String("nodename", "", "a string")
	flag.Parse()
	//
	// nodeName := *wordPtr

	filterString := "Name=private-dns-name, Values=" + *nodeName

	var data Data

	cmd := exec.Command("aws", "ec2", "describe-instances", "--filters", filterString)

	content, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print("Error: ", err)
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Print("Unmarshal-error: ", err)
	}
	// get the instance id and log to stdout
	fmt.Println(data.Reservations[0].Instances[0].InstanceId)
}
