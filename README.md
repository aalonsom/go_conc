# go_conc
Examples of concurrency programs using Go

This set of simple exercises are intended to show how to use the functionalities
in the go programming for developing concurrent programs. These examples 
illustrate the goroutines and the use of channels for illustrating the 
synchronization and shared memory for a number of goroutines. 
The included codes are:

- Mine: is a model of a mine where there is a motor for moving out water when
it gets a given level. However, this motor cannot be activated if the methane
is to high, as it may explode. There are a couple of goroutines. One generates
a random value of the level water and another one for the methane level. If
the methane is greater than a threshold. If it is below the level, the motor can be
activated if the water is greater than a threshold. The motor is stopped if the 
water level below is the lowest threshold. (This is example is inspired from
the book "", by A. Burns and A. Wellings)

- Single buffer: This is the first out of three examples for the classical
producer/consumer buffer. In this case, the buffer can only store a data. Then,
a producer cannot write a new data, if the buffer is full and it has to block 
until a hole is available. The consumer cannot get a data, if the buffer is 
empty. Then, it is blocked until a new data is stored. The synchronization is 
only made by a couple of channels for letting the producers to put data and
consumers to get data. The buffer is passive. 

- Buffer Multi: This case supports a buffer for storing a limited set of data. 
As below, consumers must block if tha buffer is empty, while producers must 
block if the buffer is full. The solution is simple and elegant that relies on 
channels with buffer

- Buffer chan: This is a similar example that is a bit artificial for showing
the select functionality within channels. This allows a goroutine buffer that 
can randomly handle invoking goroutines  if neither the buffer is full nor empty.

 
