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
	fmt.Printf("%s está pensando\n", p.name)
	time.Sleep(time.Millisecond * 500)
}

func (p *Philosopher) Eating(wg *sync.WaitGroup) {
	defer wg.Done()

	for p.mealsEaten < MEAL {
		p.Thinking()

		// Esta é a parte que causa DEADLOCK!
		// Todos os filósofos tentam pegar o garfo da esquerda primeiro
		fmt.Printf("%s está tentando pegar o garfo da ESQUERDA (posição %d)\n", p.name, p.tablePosition)
		p.leftFork.Lock()
		fmt.Printf("%s PEGOU o garfo da esquerda\n", p.name)

		// Pequeno delay para aumentar a chance de deadlock
		time.Sleep(time.Millisecond * 100)

		fmt.Printf("%s está tentando pegar o garfo da DIREITA (posição %d)\n", p.name, (p.tablePosition+1)%PHILOSOPHERS_SIZE)
		p.rightFork.Lock()
		fmt.Printf("%s PEGOU o garfo da direita\n", p.name)

		// Comendo
		p.mealsEaten++
		fmt.Printf("%s está COMENDO! (refeição %d/%d)\n", p.name, p.mealsEaten, MEAL)
		time.Sleep(time.Millisecond * 500)

		// Soltando os garfos
		p.rightFork.Unlock()
		fmt.Printf("%s soltou o garfo da direita\n", p.name)

		p.leftFork.Unlock()
		fmt.Printf("%s soltou o garfo da esquerda\n", p.name)

		fmt.Printf("%s terminou a refeição %d/%d\n\n", p.name, p.mealsEaten, MEAL)
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
	// Criar os garfos (mutexes) - um entre cada par de filósofos
	forks := make([]*sync.Mutex, PHILOSOPHERS_SIZE)
	for i := 0; i < PHILOSOPHERS_SIZE; i++ {
		forks[i] = &sync.Mutex{}
	}

	philosophers := []*Philosopher{}

	// Disposição circular: cada filósofo tem um garfo à esquerda e outro à direita
	// TODOS pegam o garfo da esquerda primeiro - isso CAUSA DEADLOCK!
	philosophers = append(philosophers, NewPhilosopher("Aristotle", 0, forks[0], forks[1]))
	philosophers = append(philosophers, NewPhilosopher("Socrates", 1, forks[1], forks[2]))
	philosophers = append(philosophers, NewPhilosopher("Plato", 2, forks[2], forks[3]))
	philosophers = append(philosophers, NewPhilosopher("Voltaire", 3, forks[3], forks[4]))
	philosophers = append(philosophers, NewPhilosopher("Thomas Hobbes", 4, forks[4], forks[0]))

	return philosophers
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println("   DEMONSTRAÇÃO DO PROBLEMA DOS FILÓSOFOS JANTANDO")
	fmt.Println("   (Dining Philosophers Problem - Dijkstra)")
	fmt.Println("═══════════════════════════════════════════════════════")

	wg := sync.WaitGroup{}
	philosophers := CreatePhilosophers()

	fmt.Println("Os filósofos estão sentados à mesa...")
	time.Sleep(time.Second)

	wg.Add(PHILOSOPHERS_SIZE)
	for _, philosopher := range philosophers {
		go philosopher.Eating(&wg)
	}

	// Timeout para detectar deadlock
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("═══════════════════════════════════════════════════════")
		fmt.Println("Todos os filósofos comeram todas as refeições!")
		fmt.Println("═══════════════════════════════════════════════════════")
	case <-time.After(10 * time.Second):
		fmt.Println("═══════════════════════════════════════════════════════")
		fmt.Println("DEADLOCK DETECTADO!")
		fmt.Println("Os filósofos ficaram travados esperando pelos garfos.")
		fmt.Println("═══════════════════════════════════════════════════════")
	}
}
