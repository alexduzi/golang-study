package main

import (
	"fmt"
	"sync"
	"time"
)

const PHILOSOPHERS_SIZE int = 5
const MEAL int = 3

type Philosopher struct {
	name          string
	leftFork      *sync.Mutex
	rightFork     *sync.Mutex
	tablePosition int
	mealsEaten    int
}

func (p *Philosopher) Thinking() {
	fmt.Printf("%s is thinking\n", p.name)
	time.Sleep(time.Millisecond * 500)
}

func (p *Philosopher) Eating(wg *sync.WaitGroup) {
	defer wg.Done()

	for p.mealsEaten < MEAL {
		p.Thinking()

		// This is the part that causes DEADLOCK!
		// All philosophers try to grab the left fork first
		fmt.Printf("%s is trying to grab the LEFT fork (position %d)\n", p.name, p.tablePosition)
		p.leftFork.Lock()
		fmt.Printf("%s GRABBED the left fork\n", p.name)

		// Small delay to increase the chance of deadlock
		time.Sleep(time.Millisecond * 100)

		fmt.Printf("%s is trying to grab the RIGHT fork (position %d)\n", p.name, (p.tablePosition+1)%PHILOSOPHERS_SIZE)
		p.rightFork.Lock()
		fmt.Printf("%s GRABBED the right fork\n", p.name)

		// Eating
		p.mealsEaten++
		fmt.Printf("%s is EATING! (meal %d/%d)\n", p.name, p.mealsEaten, MEAL)
		time.Sleep(time.Millisecond * 500)

		// Releasing the forks
		p.rightFork.Unlock()
		fmt.Printf("%s released the right fork\n", p.name)

		p.leftFork.Unlock()
		fmt.Printf("%s released the left fork\n", p.name)

		fmt.Printf("%s finished meal %d/%d\n\n", p.name, p.mealsEaten, MEAL)
	}
}

func NewPhilosopher(name string, tablePosition int, leftFork, rightFork *sync.Mutex) *Philosopher {
	return &Philosopher{
		name:          name,
		tablePosition: tablePosition,
		leftFork:      leftFork,
		rightFork:     rightFork,
		mealsEaten:    0,
	}
}

func CreatePhilosophers() []*Philosopher {
	// Create the forks (mutexes) - one between each pair of philosophers
	forks := make([]*sync.Mutex, PHILOSOPHERS_SIZE)
	for i := 0; i < PHILOSOPHERS_SIZE; i++ {
		forks[i] = &sync.Mutex{}
	}

	philosophers := []*Philosopher{}

	// Circular arrangement: each philosopher has a fork on the left and another on the right
	// ALL grab the left fork first - this CAUSES DEADLOCK!
	philosophers = append(philosophers, NewPhilosopher("Aristotle", 0, forks[0], forks[1]))
	philosophers = append(philosophers, NewPhilosopher("Socrates", 1, forks[1], forks[2]))
	philosophers = append(philosophers, NewPhilosopher("Plato", 2, forks[2], forks[3]))
	philosophers = append(philosophers, NewPhilosopher("Voltaire", 3, forks[3], forks[4]))
	philosophers = append(philosophers, NewPhilosopher("Thomas Hobbes", 4, forks[4], forks[0]))

	return philosophers
}

func main() {
	fmt.Println("   DINING PHILOSOPHERS PROBLEM DEMONSTRATION")
	fmt.Println("   (Dijkstra's Classic Concurrency Problem)")

	wg := sync.WaitGroup{}
	philosophers := CreatePhilosophers()

	fmt.Println("The philosophers are sitting at the table...")
	time.Sleep(time.Second)

	wg.Add(PHILOSOPHERS_SIZE)
	for _, philosopher := range philosophers {
		go philosopher.Eating(&wg)
	}

	// Timeout to detect deadlock
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("All philosophers finished all their meals!")
	case <-time.After(10 * time.Second):
		fmt.Println("DEADLOCK DETECTED!")
		fmt.Println("The philosophers are stuck waiting for the forks.")
	}
}
