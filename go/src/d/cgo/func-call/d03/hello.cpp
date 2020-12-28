#include <iostream>
extern "C" {
	#include "hello.h"
}

void Hello(const char* s) {
	std::cout << s;
}
