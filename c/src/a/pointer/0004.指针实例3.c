#include <stdio.h>

int main(int argc, char *argv[]) {
  int ages[] = {19, 39, 12, 48, 62};
  char *names[] = {"Alan", "Mary", "Jack", "Lisa", "Jhon"};
  int count = sizeof(ages) / sizeof(int);
  int i = 0;

  // first way using indexing
  for (int i = 0; i < count; i++) { printf("%s has %d years old\n", names[i], ages[i]); }
  printf("\n");

  int *age = ages;
  char **name = names;
  // second way using pointers
  for (i = 0; i < count; i++) { printf("%s has %d years old\n", *(name + i), *(age + i)); }
  printf("\n");

  // third way, pointers are just arrays
  for (i = 0; i < count; i++) { printf("%s has %d years old\n", name[i], age[i]); }
  printf("\n");

  // fourth way with pointers in a stupid complex way
  for (name = names, age = ages; (age - ages) < count; name++, age++) { printf("%s has %d years old\n", *name, *age); }
  return 0;
}
