package main

import (
	"sync"
)

const (
	// Границы времени сбора компонентов барменом
	minCollectingTime = 100
	maxCollectingTime = 300

	// Границы времени скручивания сигареты курильщиком
	minMakingTime = 2000
	maxMakingTime = 2500

	// Границы времени курения сигареты
	minSmokingTime = 3000
	maxSmokingTime = 4000
)

var (
	// Мьютекс, отвечающий за заполненность стола бармена
	emptyTable        = &sync.Mutex{}
	// Мьютекс, указывающий на то, что кто-то сейчас курит
	isSmoking        = &sync.Mutex{}
	// Мьютексы, указывающие на выложенные компоненты
	tobaccoAndMatches = &sync.Mutex{}
	paperAndMatches   = &sync.Mutex{}
	tobaccoAndPaper   = &sync.Mutex{}
	// Используется для определения группы горутин
	wg                = &sync.WaitGroup{}
)

func main() {
	// Изначально блокируем мьютексы компонентов
	tobaccoAndMatches.Lock()
	paperAndMatches.Lock()
	tobaccoAndPaper.Lock()

	// Указываем количество потоков в группе
	wg.Add(4)

	// Запускаем потоки
	go barman()
	go smoker(tobaccoAndMatches, "\"Jake (the first smoker)\"")
	go smoker(paperAndMatches, "\"Andrew (the second smoker)\"")
	go smoker(tobaccoAndPaper, "\"Mike (the third smoker)\"")

	// Ожидаем завершения работы всех потоков
	wg.Wait()
}
