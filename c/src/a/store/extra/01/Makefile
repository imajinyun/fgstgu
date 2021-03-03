# Flag for implicit rules
CC = gcc

# Flag for implicit rules. Turn on debug info
CFLAGS = -Wall -g

objects = main.o diceroll.o
main: $(objects)
main.o: diceroll.h
diceroll.o: diceroll.h

.PHONY: clean
clean:
	-rm main $(objects)
