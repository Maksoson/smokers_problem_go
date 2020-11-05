package main

import (
	. "fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	makingTime int
	smokeTime  int
)

func smoker(componentsNeeded *sync.Mutex, name string) {
	for {
		// Блокируем мьютекс по выложенным компонентам
		componentsNeeded.Lock()
		// Процесс курения начинается
		isSmoking.Lock()

		Println(name + " is making new cigarette...")
		// Скручивание сигареты
		makeCigarette()
		Println(name + " make his cigarette (" + strconv.Itoa(makingTime) + " ms)")

		// После скручивания сигареты, стол освобождается
		emptyTable.Unlock()

		Println(name + " is smoking cigarette...")
		// Курение сигареты
		smokeCigarette()
		Println(name + " smoked his cigarette (" + strconv.Itoa(smokeTime) + " ms)")

		// Курильщик докурил сигарету
		isSmoking.Unlock()
	}

	// Поток выполнил работу
	wg.Done()
}

// Ожидание скручивания сигареты
func makeCigarette() {
	makingTime = rand.Intn(maxMakingTime-minMakingTime) + minMakingTime
	time.Sleep(time.Duration(makingTime) * time.Millisecond)
}

// Ожидание скуривания сигареты
func smokeCigarette() {
	smokeTime = rand.Intn(maxSmokingTime-minSmokingTime) + minSmokingTime
	time.Sleep(time.Duration(smokeTime) * time.Millisecond)
}
