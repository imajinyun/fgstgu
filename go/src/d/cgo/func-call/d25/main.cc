#include "person.h"
#include <stdio.h>

int main(int argc, char *argv[]) {
  auto p = new Person("gopher", 10);
  char buf[64];
  char *name = p->GetName(buf, sizeof(buf) - 1);
  int age = p->GetAge();
  printf("%s, %d years old.\n", name, age);
  delete p;
  return 0;
}
