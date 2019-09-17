package internet_cafe

import "fmt"

type Randomizer interface {
	CalculateTime() int
}

type InternetCafe struct {
	computers  int
	randomizer Randomizer
}

func NewInternetCafe(computers int, randomizer Randomizer) InternetCafe {
	return InternetCafe{computers, randomizer}
}

func (i InternetCafe) Start(tourists int) []string {
	input := make(chan int, tourists)
	outputEvents := make(chan string)

	for n := 0; n < i.computers; n++ {
		go (&Computer{input, outputEvents, i.randomizer}).Start()
	}

	events := []string{}

	for tourist := 1; tourist <= tourists; tourist++ {
		input <- tourist
	}
	close(input)

	for event := range outputEvents {
		events = append(events, event)
	}

	events = append(events, "The place is empty, let's close up and go to the beach!")

	return events
}

type Computer struct {
	input        chan int
	outputEvents chan string
	randomizer   Randomizer
}

func (c *Computer) Start() {
	for tourist := range c.input {
		timeSpent := c.randomizer.CalculateTime()

		c.outputEvents <- fmt.Sprintf("Tourist %d is online", tourist)
		c.outputEvents <- fmt.Sprintf("Tourist %d is done, having spent %d minutes online.", tourist, timeSpent)
	}
	close(c.outputEvents)
}
