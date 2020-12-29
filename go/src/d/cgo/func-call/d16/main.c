#include <stdio.h>
#include "hello.h"

// -> go build -buildmode=c-shared -o hello.so hello.go
// -> gcc -o main main.c hello.so
// -> ./main
int main(int argc, const char* argv[]) {
	int res = sum(10, 20);
	printf("result: %d\n", res);
	return 0;
}
