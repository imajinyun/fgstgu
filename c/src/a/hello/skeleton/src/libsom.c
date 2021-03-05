#include "dbg.h"
#include <dlfcn.h>
#include <stdio.h>

typedef int (*lib_function)(const char *data);

/**
 * -> make
 * -> cd src
 * -> cc -c libsoc.c -o ../build/libsoc.o
 * -> cc -shared -o ../build/libsoc.so ../build/libsoc.o
 * -> cc -Wall -g -DNDEBUG libsom.c -ldl -o ../build/libsom
 * -> ../build/libsom ../build/libsoc.so message "Hello World"
 */
int main(int argc, char *argv[]) {
  int n = 0;
  check(argc == 4, "Usage: libsom libsom.so function data");

  char *file = argv[1];
  char *func = argv[2];
  char *data = argv[3];

  void *lib = dlopen(file, RTLD_NOW);
  check(lib != NULL, "failed to open the library %s: %s", file, dlerror());

  lib_function fn = dlsym(lib, func);
  check(fn != NULL, "did not find %s func in the library %s: %s", func, file,
        dlerror());

  n = fn(data);
  check(n == 0, "function %s return %d for data: %s", func, n, data);
  n = dlclose(lib);
  check(n == 0, "failed to close %s", file);
  return 0;

goerr:
  return -1;
}
