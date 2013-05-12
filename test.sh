#!/bin/bash
#
# This script does runs our CPU testing script with different settings
#

#
# Errors are fatal
#
set -e
#set -x # Debugging

NUM_MESSAGES=10000000
BUFFER_SIZE=100
NUM_CORES=1
NUM_GOROUTINES=100
MESSAGE_SIZE=1
NUM_GOROUTINE_WORK=1


#
# Run the script in a single test
#
function run() {
	go run ./main.go \
		--num-messages ${NUM_MESSAGES} \
		--buffer-size ${BUFFER_SIZE} \
		--num-cores ${NUM_CORES} \
		--num-goroutines ${NUM_GOROUTINES} \
		--message-size ${MESSAGE_SIZE} \
		--num-goroutine-work ${NUM_GOROUTINE_WORK}

} # End of run()

#
# Run a series of tests with different cores
#
function run_tests() {

	NUM_CORES=1
	run

	NUM_CORES=2
	run
	
	NUM_CORES=4
	run

} # End of run_tests()


#
# Create a list of powers of 2 for the work we'll do per goroutine
#
NUM_GOROUTINE_LIST=""
NUM=1
for I in `seq 1 17`
do
	NUM_GOROUTINE_LIST="${NUM_GOROUTINE_LIST} ${NUM}"
	NUM=`expr $NUM \* 2`
done

#
# Now run our script for different numbers of goroutines
#
for I in `echo $NUM_GOROUTINE_LIST`
do
	NUM_GOROUTINE_WORK=$I
	run_tests

done



