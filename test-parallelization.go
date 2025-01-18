package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Scenario: Fetching Wather Info from Different Cities
// Imagine we want to fetch weather info from several different cities at the same time and combine the results.
// Each goroutine performs a “weather fetch” simulation, and then sends the result to a channel. The main function collects those results.
// fetchWeather simulates fetching weather data for a given city
// - city: the name of the city
// - wg: pointer to the WaitGroup tracking concurrent tasks
// - results: a "send-only" channel (chan<- string) where we send (not receive) weather info
func fetchWeather(city string, wg *sync.WaitGroup, results chan<- string) {
	// defer wg.Done() means: when this function returns (ends),
	// call wg.Done() to signal that this goroutine is finished.
	defer wg.Done()

	// Simulate a random delay (like network latency or processing time).
	delay := time.Duration(rand.Intn(2000)) * time.Millisecond
	time.Sleep(delay)

	// Build a fake weather info string for the city.
	weatherInfo := fmt.Sprintf("Weather in %s: %d°C, %d%% humidity",
		city, rand.Intn(35), rand.Intn(100))

	// Send (push) the weather info string into the results channel.
	// results <- weatherInfo means "send weatherInfo to the channel".
	results <- weatherInfo
}

func mainAdditional() {
	// Seed the random number generator using the current time in nanoseconds.
	// This ensures random values vary each time you run the program.
	rand.Seed(time.Now().UnixNano())

	// List of cities we want to fetch weather data for.
	cities := []string{"New York", "London", "Tokyo", "Berlin", "Sydney"}

	// Create a channel that carries strings, with buffer capacity = len(cities).
	// This means the channel can hold up to len(cities) items before it blocks.
	results := make(chan string, len(cities))

	// We create a WaitGroup to manage concurrency.
	// We'll add 1 for each goroutine we start, and call wg.Done() when each finishes.
	var wg sync.WaitGroup

	// Launch a goroutine for each city to fetch weather data in parallel.
	for _, city := range cities {
		wg.Add(1) // Increment the WaitGroup counter by 1.

		// The 'go' keyword starts a new goroutine (lightweight thread) that will run fetchWeather concurrently.
		go fetchWeather(city, &wg, results)
	}

	// wg.Wait() blocks until the WaitGroup counter goes back down to 0,
	// i.e., until all goroutines have called wg.Done() and finished.
	wg.Wait()

	// Close the channel to signal that no more data will be sent into 'results'.
	// This allows the 'range results' loop below to end.
	close(results)

	// Read from the 'results' channel until it's closed. Each iteration receives one string from the channel.
	for weatherData := range results {
		fmt.Println(weatherData)
	}

	// Execution reaches here after we've read all results from the channel.
	fmt.Println("All weather information collected.")
}
