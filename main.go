package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

// Message ...
type Message struct {
	Message string
}

var re = regexp.MustCompile("([^=]+)='([^']*)'\n")

func (m Message) asDict() map[string]string {
	out := map[string]string{}
	for _, line := range re.FindAllStringSubmatch(m.Message, -1) {
		out[line[1]] = line[2]
	}
	return out
}

func main() {
	dec := json.NewDecoder(os.Stdin)

	for {
		var m Message
		err := dec.Decode(&m)
		if err != nil {
			break
		}

		d := m.asDict()

		// log.Println(m.Message)
		fmt.Printf("%s %-55s %-35s %s %s %s\n",
			d["Timestamp"], d["StackName"], d["ResourceStatus"], d["ResourceType"], d["LogicalResourceId"], d["ResourceStatusReason"])

		// for k, v := range d {
		// 	log.Printf("k=%q, v=%q", k, v)
		// }
		// log.Println()
		// log.Println()
		// log.Println()
	}
}
