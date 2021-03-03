# Flag for implicit rules
CC = gcc

# Flag for implicit rules. Turn on debug info
CFLAGS = -Wall -g

objects = main.o scope.o
main: $(objects)
main.o: scope.h ../../../hello/debug/dbg.h
scope.o: scope.h ../../../hello/debug/dbg.h

.PHONY: clean
clean:
	-rm main $(objects)
