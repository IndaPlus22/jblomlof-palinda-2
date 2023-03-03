## Task 2 - many senders, many recievers
* What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?
    * It's going to close the channel before all producers have sent thier information. Thus making the program panic since producers tries to send on closed channel.

* What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?
    * The first producer to finish will close the channel, thus making another producer panic when it tries sending. Only works when you have 1 producer.

* What happens if you remove the statement `close(ch)` completely?
    * Honestly no big difference, the only difference is that some reciever goroutines will not exit and take up computing resources. But not a lot, since the program terminates after.

* What happens if you increase the number of consumers from 2 to 4?
    * The amount of consumers are doubled... (What is this question)? But on a more serious level, the program is faster since the bottle-neck was the consumers (producers was waiting for consumers to recieve)

* Can you be sure that all strings are printed before the program
  stops?
    * No you can't. When the last producer is finished. The message is consumed whilst it returns to WaitSync that it's done. So if you're unlucky and the printstatement is slow AF the waitsync, close, timeprint and program exit might be faster than the stringprinter. But rarley happens. Although if you would swap order in consumer (so the "processing time" happens before print) it will most often not print the string.

    * Bonus, I had a case where it printed a string after the timeprint. So take that Stefan Nilsson