#include "number.h"
#include <stdio.h>

// -> gcc -o main.o main.c number.so
// -> ./main.o
int main(int argc, char *argv[]) {
	int a, b, x;
	a = 100;
	b = 128;
	x = average(a, b);
	printf("(%d+%d)/2=%d\n", a, b, x);
	return 0;
}
