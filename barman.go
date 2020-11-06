package main

import (
	. "fmt"
	"math/rand"
	"strconv"
	"time"
)

const barmanName = "\"Bob (the barman)\""

var collectingTime int

func barman() {
	for {
		// Блокируем стол
		emptyTable.Lock()

		smoker := rand.Intn(3) + 1
		if smoker == 1 {				// Курить будет первый курильщик
			Println(barmanName + " is now collecting tobacco and matches (for the first smoker)...")
			collectionOfComponents()
			Println(barmanName + " laid out the tobacco and matches (" + strconv.Itoa(collectingTime) + " ms)")

			// На столе табак и спички
			tobaccoAndMatches.Unlock()
		} else if smoker == 2 {			// Курить будет второй курильщик
			Println(barmanName + " is now collecting paper and matches (for the second smoker)...")
			collectionOfComponents()
			Println(barmanName + " laid out the paper and matches (" + strconv.Itoa(collectingTime) + " ms)")

			// На столе бумага и спички
			paperAndMatches.Unlock()
		} else if smoker == 3 {			// Курить будет третий курильщик
			Println(barmanName + " is now collecting tobacco and paper (for the third smoker)...")
			collectionOfComponents()
			Println(barmanName + " laid out the tobacco and paper (" + strconv.Itoa(collectingTime) + " ms)")

			// На столе бумага и табак
			tobaccoAndPaper.Unlock()
		}
	}

	// Поток выполнил работу
	wg.Done()
}

// Ожидание сбора компонентов
func collectionOfComponents() {
	collectingTime = rand.Intn(maxCollectingTime-minCollectingTime) + minCollectingTime
	time.Sleep(time.Duration(collectingTime) * time.Millisecond)
}
