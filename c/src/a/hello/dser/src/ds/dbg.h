#ifndef __DBG_H
#define __DBG_H

#include <errno.h>
#include <stdio.h>
#include <string.h>

#ifdef NDEBUG
#define debug(M, ...)
#else
#define debug(M, ...) fprintf(stderr, "[🟣] (%s:%d) " M "\n", __FILE__, __LINE__, ##__VA_ARGS__)
#endif

#define clean() (errno == 0 ? "None" : strerror(errno))
#define error(M, ...) fprintf(stderr, "[🔴] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__, clean(), ##__VA_ARGS__)
#define warn(M, ...) fprintf(stderr, "[🟡] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__, clean(), ##__VA_ARGS__)
#define info(M, ...) fprintf(stderr, "[🔵] (%s:%d: errno: %s) " M "\n", __FILE__, __LINE__, clean(), ##__VA_ARGS__)

#define check(A, M, ...)                                                                                               \
  if (!(A)) {                                                                                                          \
    error(M, ##__VA_ARGS__);                                                                                           \
    errno = 0;                                                                                                         \
    goto goerr;                                                                                                        \
  }
#define fatal(M, ...)                                                                                                  \
  {                                                                                                                    \
    error(M, ##__VA_ARGS__);                                                                                           \
    errno = 0;                                                                                                         \
    goto goerr;                                                                                                        \
  }
#define checkmem(A) check((A), "Out of memory")
#define checkdbg(A, M, ...)                                                                                            \
  if (!(A)) {                                                                                                          \
    debug(M, ##__VA_ARGS__);                                                                                           \
    errno = 0;                                                                                                         \
    goto goerr;                                                                                                        \
  }

#endif
