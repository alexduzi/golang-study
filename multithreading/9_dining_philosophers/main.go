package main

import (
	"fmt"
	"sync"
	"time"
)

const PHILOSOPHERS_SIZE int = 5

const MEAL int = 1

type Philosopher struct {
	name          string
	leftFork      *sync.Mutex
	rightFork     *sync.Mutex
	tablePosition int
}

func (p *Philosopher) Thinking() {
	fmt.Printf("%s is thinking", p.name)
	time.Sleep(time.Second * time.Duration(2))
}

func (p *Philosopher) Eating(wg *sync.WaitGroup) {
	p.Thinking()

	fmt.Printf("%s is trying to get the left fork", p.name)
	p.leftFork.Lock()

	fmt.Printf("%s is trying to get the right fork", p.name)
	p.rightFork.Lock()

	fmt.Printf("%s is eating!", p.name)
	time.Sleep(time.Second * time.Duration(3))

	p.leftFork.Unlock()
	p.rightFork.Unlock()
	wg.Done()
}

func NewPhilosopher(name string, tablePosition int, leftFork, rightFork *sync.Mutex) *Philosopher {
	return &Philosopher{
		name:          name,
		tablePosition: tablePosition,
		leftFork:      leftFork,
		rightFork:     rightFork,
	}
}

func CreatePhilosophers() []*Philosopher {
	philosopers := []*Philosopher{}

	philosopers = append(philosopers, NewPhilosopher("Aristotle", 1))
	philosopers = append(philosopers, NewPhilosopher("Socrates", 2))
	philosopers = append(philosopers, NewPhilosopher("Plato", 3))
	philosopers = append(philosopers, NewPhilosopher("Voltaire", 4))
	philosopers = append(philosopers, NewPhilosopher("Thomas Hobbes", 5))

	return philosopers
}

func main() {
	fmt.Println("philosophers are sitting the table")
	wg := sync.WaitGroup{}

	philosopers := CreatePhilosophers()

	wg.Add(5)
	for _, philosopher := range philosopers {

		go func(philosopher *Philosopher) {
			philosopher.Eating()
		}(philosopher)
	}

	wg.Wait()
	fmt.Println("all philosophers eaten!")
}
