Go Routines:

- Go language provides a special feature known as a Goroutines.
- A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program.
- Goroutine is a light weighted thread, The cost of creating Goroutines is very small as compared to the thread.
- Every program contains at least a single Goroutine and that Goroutine is known as the main Goroutine.
- All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated.
- Goroutine always works in the background.
- parallel execution of the function thus over comes serial execution.
- By default, Go tries to use One core.
- Scheduler runs one routine for one logical CPU core until it finishes or makes a blocking call (like an HTTP rqst).

- we can only use 'go' keyword in front of our function calls.

- Concurrency -> We can have multiple threads executing code, If one thread blocks another one is picked up & worked on.
- Parallelism -> Multiple threads executed at the same exact time, requires multiple CPU's.

- main routine  -> main routine created when we launched our program.
- child routine -> child routines created by the 'go' keyword.

- By default, when used go keyword, at some point of time main routine will come to an end but child routines will continue to run in the bg.
- this is a bug by default, to overcome this we need to implement channels.




Channel:

- A Channel is a technique which allows to let one goroutine to send data to another goroutine. 
- By default, sends and receives block until the other side is ready, This allows goroutines to synchronize without explicit locks or condition variables.