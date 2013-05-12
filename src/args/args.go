
package args

import "flag"
import "runtime"
import "strconv"


type Config struct {
	Num_messages int
	Message_size int
	Num_cores int
	Buffer_size int
	Num_goroutines int
	Num_goroutine_work int
}


/**
* Our main entry point.
* @return{struct} Our config structure populated with configuration info.
*/
func Main() Config {

	config := Config{}
	flag.IntVar(&config.Num_messages, "num-messages", 1, 
		"How many messages to send?")
	flag.IntVar(&config.Message_size, "message-size", 1, 
		"How many bytes in each message?")
	flag.IntVar(&config.Num_cores, "num-cores", 1, 
		"How many cores to use? (Max:" + strconv.Itoa(runtime.NumCPU()) + ")")
	flag.IntVar(&config.Buffer_size, "buffer-size", 0, 
		"Buffer size for our channel?")
	flag.IntVar(&config.Num_goroutines, "num-goroutines", 1, 
		"How many goroutines to run?")
	flag.IntVar(&config.Num_goroutine_work, "num-goroutine-work", 1, 
		"How much work to do in each goroutine?")

	var h = flag.Bool("h", false, "Print this help")
	var help = flag.Bool("help", false, "Print this help")

	flag.Parse()

	if (*h || *help) {
		flag.PrintDefaults()
		panic("Bailing out due to asking for help")

	} else if (config.Num_goroutine_work < 1) {
		flag.PrintDefaults()
		panic("Goroutine work must be >= 1")

	}

	return(config)

} // End of Main()


