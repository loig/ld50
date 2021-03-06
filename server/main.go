/*LD50, a game for Ludum Dare 50
  Copyright (C) 2022  Loïg Jezequel

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see https://www.gnu.org/licenses/.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type score struct {
	name  string
	level int
}

const logfile string = "server.log"

var theScores []score
var mut sync.Mutex

func insertScore(name string, nlevel int) int {

	newScore := score{name: name, level: nlevel}

	position := 0

	mut.Lock()

	for i := 0; i < len(theScores); i++ {
		position = i
		if newScore.level > theScores[i].level {
			newScore, theScores[i] = theScores[i], newScore
			break
		}
	}

	newPos := position
	oldScore := score{name: name, level: nlevel}
	if newScore == oldScore {
		newPos++
	}

	position++

	for ; position < len(theScores); position++ {
		newScore, theScores[position] = theScores[position], newScore
	}

	if newScore.level > 0 {
		theScores = append(theScores, newScore)
	}

	mut.Unlock()

	newPos++

	return newPos
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://v6p9d9t4.ssl.hwcdn.net")

	name := r.PostFormValue("uname")
	level := r.PostFormValue("level")

	nlevel, err := strconv.Atoi(level)

	if err == nil && name != "" {

		newPos := insertScore(name, nlevel)

		fmt.Print(name, ":", level, ";")

		log.Print("Added a score: ", name, ",", level, ",", newPos, " from request ", *r)

		fmt.Fprint(w, newPos, ":", theScores[0], ":", theScores[1], ":", theScores[2], ":", len(theScores))

		return
	}

	log.Print("Wrong request: ", *r, name, level)
	fmt.Fprintf(w, "error")
}

func main() {

	content, err := ioutil.ReadFile(logfile)
	if err != nil {
		log.Print(err)
	}

	previous := strings.Split(string(content), ";")

	theScores = []score{{"AAA", 0}, {"AAA", 0}, {"AAA", 0}}

	for _, ascore := range previous {
		info := strings.Split(ascore, ":")
		if len(info) == 2 {
			nlevel, err := strconv.Atoi(info[1])
			if err == nil {
				insertScore(info[0], nlevel)
			}
		}
	}

	http.HandleFunc("/", HandleRequest)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
