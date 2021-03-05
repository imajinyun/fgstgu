#undef NDEBUG
#include "../hello/debug/dbg.h"
#include <assert.h>
#include <stdio.h>

int safecopy(char *from, char *to, int fr_len, int to_len) {
  assert(from != NULL && to != NULL && "from and to can't be NULL");
  int i;
  int n = fr_len > to_len - 1 ? to_len - 1 : fr_len;

  if (fr_len < 0 || to_len <= 0) { return -1; }
  for (i = 0; i < n; i++) { to[i] = from[i]; }
  to[to_len - 1] = '\0';
  return i;
}

int main(int argc, char *argv[]) {
  char from[] = "0123456789";
  int fr_len = sizeof(from);
  char to[] = "0123456";
  int to_len = sizeof(to);

  debug("Copying '%s': %d to '%s': %d", from, fr_len, to, to_len);

  int n = safecopy(from, to, fr_len, to_len);
  check(n > 0, "failed to safe copy");
  check(to[to_len - 1] == '\0', "string not terminated");

  debug("Result is: '%s': %d", to, to_len);

  n = safecopy(from, to, fr_len, 0);
  check(n == -1, "safe copy should fail #2");
  check(to[to_len - 1] == '\0', "string not terminated");
  return 0;

goerr:
  return 1;
}
