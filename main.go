package main

import (
	"fmt"
	"log"
	queue "test-queue/queueService"
	tools "test-queue/toolsService"
)

func Cards(cards int, q queue.Queue[int], name string) ([]int, int) {
	totalLoops := tools.PerformanceParam{
		Name:  "total loops",
		Value: 0,
	}
	defer tools.Performance(name, &totalLoops)()

	removedCards := addCardsToQueue(cards, q, &totalLoops)
	removeRemainingCards(q, &removedCards, &totalLoops)
	remainingCard := q.Dequeue()

	return removedCards, remainingCard
}

func addCardsToQueue(cards int, q queue.Queue[int], totalLoops *tools.PerformanceParam) []int {
	removedCards := []int{}

	for i := 1; i <= cards; i++ {
		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}
		removedCards = append(removedCards, i)

		if i <= cards {
			err := q.Enqueue(i)
			if err != nil {
				log.Fatal(err)
			}

		}

	}
	return removedCards
}

func removeRemainingCards(q queue.Queue[int], removedCards *[]int, totalLoops *tools.PerformanceParam) {
	for q.Len() >= 1 {

		if totalLoopsInt, ok := totalLoops.Value.(int); ok {
			totalLoops.Value = totalLoopsInt + 1
		}

		_ = q.Dequeue()
	}
}

func main() {
	n1, n2 := Cards(10, queue.NewSliceQueue[int](), "test")
	fmt.Println(n1)
	fmt.Println(n2)
}
