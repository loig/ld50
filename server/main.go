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
	"log"
	"net/http"
	"strconv"
	"sync"
)

type score struct {
	name  string
	level int
}

var theScores []score
var mut sync.Mutex

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	name := r.PostFormValue("uname")
	level := r.PostFormValue("level")

	nlevel, err := strconv.Atoi(level)

	if err == nil && name != "" {

		mut.Lock()

		newScore := score{name: name, level: nlevel}

		position := 0

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

		log.Print("Added a score: ", name, ",", level, ",", newPos, " from request ", *r)

		fmt.Fprint(w, newPos, ":", theScores[0], ":", theScores[1], ":", theScores[2], ":", len(theScores))

		return
	}

	log.Print("Wrong request: ", *r, name, level)
	fmt.Fprintf(w, "error")
}

func main() {

	theScores = []score{{"AAA", 0}, {"AAA", 0}, {"AAA", 0}}

	http.HandleFunc("/", HandleRequest)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
