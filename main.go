package main

import "fmt"
import "runtime"
import "time"

import "./src/args"

func main() {

	config := args.Main()

	max_procs := runtime.NumCPU()
	if config.Num_cores > max_procs {
		panic(fmt.Sprintf("We only have %d cores, but you asked for %d!", max_procs, config.Num_cores))
	}
	runtime.GOMAXPROCS(config.Num_cores)

	c := make(chan string, config.Buffer_size)

	//
	// Our goroutine for processing messages.
	// Make as many as were requested and do fake work in each goroutine
	// to spin the core.
	//
	for i := 0; i < config.Num_goroutines; i++ {
		go func(c chan string) {
			for {
				<-c
				for i := 0; i < config.Num_goroutine_work; i++ {
				}
			}
		}(c)
	}

	//
	// Create our message to send
	//
	var message string
	for i := 0; i < config.Message_size; i++ {
		message += "x"
	}

	//
	// Now actually send our messages!
	//
	start_time := time.Now()

	for i := 0; i < config.Num_messages; i++ {
		c <- message
	}

	//
	// All done! Calculate our metrics and print them up in
	// tab-delimited format so I can dump this into a spreadsheet.
	//
	duration := time.Since(start_time)
	seconds := duration.Seconds()
	messages_per_sec := float64(config.Num_messages) / seconds
	work_per_sec := messages_per_sec * float64(config.Num_goroutine_work)

	fmt.Println("Num_messages\tMessage_size\tNum_cores\tBuffer_size\tNum_goroutines\tNum_goroutine_work\tTime elapsed\tMessages/sec\tWork/sec")
	fmt.Printf("%-10d\t%-10d\t%-10d\t%-10d\t%-10d\t%-18d\t%-10.3f\t%-10.0f\t%-10.0f\n",
		config.Num_messages, config.Message_size, config.Num_cores,
		config.Buffer_size, config.Num_goroutines,
		config.Num_goroutine_work,
		seconds,
		messages_per_sec,
		work_per_sec,
	)

} // End of main()
