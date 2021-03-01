#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Person {
  char *name;
  int age;
  int height;
  int weight;
};

struct Person *create(char *name, int age, int height, int weight) {
  struct Person *who = malloc(sizeof(struct Person));
  assert(who != NULL);
  who->name = strdup(name);
  who->age = age;
  who->height = height;
  who->weight = weight;
  return who;
}

void destroy(struct Person *who) {
  assert(who != NULL);
  free(who->name);
  free(who);
}

void show(struct Person *who) {
  printf("Name: %s\n", who->name);
  printf("  Age: %d\n", who->age);
  printf("  Height: %d\n", who->height);
  printf("  Weight: %d\n", who->weight);
  printf("\n");
}

int main(int argc, char *argv[]) {
  struct Person *jack = create("Jack", 64, 175, 160);
  struct Person *mary = create("Mary", 38, 178, 170);

  printf("Jack is at memory location: %p\n", jack);
  show(jack);
  printf("Mary is at memory location: %p\n", mary);
  show(mary);

  jack->age += 10;
  jack->height -= 2;
  jack->weight += 10;
  show(jack);

  mary->age += 10;
  mary->height += 2;
  mary->weight += 10;
  show(mary);

  destroy(jack);
  destroy(mary);
  return 0;
}
