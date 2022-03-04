# go-lerning
Capturing my Golang learning programs 
I followed Nigel Poulton on pluralsight for my learning.
https://www.pluralsight.com/courses/go-fundamentals

Consurrency
- Executing multiple processes independently
- Lots of tasks "on the go", but only one "active"

Concurrency: Dealting with lost of things and doing only one tasks or frequent context switching on processor
Parallelism: Doing lots of things at onece or running multiple tasks on multiple processors in parall

- Go does not use direct OS threads for consurrency instead it usages goroutine (os green threads) to maintain concurrency. These are lighter than OS threads. They can grow and shrink as needed. Go manages goroutines, hence fewer context switching. 
Go routines have faster startup time, communicate via channel.
Goroute can switch lots of work on a single OS thread

Go's concurrency model
- Actor
- Communication Sequential Processes (CSP)


'go' keyward makes function call as go routine
However main is also a go routine, if code comes out of main then whatever is running in other routines will get terminated and program exits. Hence we need to tie our code with ,ain method by using waitGrp

 