#include "../hello/debug/dbg.h"
#include <stdio.h>

#define MAX_DATA 100

typedef enum EyeColor {
  BLUE_EYES,
  GREEN_EYES,
  BROWN_EYES,
  BLACK_EYES,
  OTHER_EYES
} EyeColor;

const char *EYE_COLOR_NAMES[] = {"Blue", "Green", "Brown", "Black", "Other"};

typedef struct Person {
  int age;
  char fname[MAX_DATA];
  char lname[MAX_DATA];
  EyeColor ecolor;
  float income;
} Person;

int main(int argc, char *argv[]) {
  Person p = {.age = 0};
  int i;
  char *in = NULL;

  // your first name
  printf("What's your first name? ");
  in = fgets(p.fname, MAX_DATA - 1, stdin);
  check(in != NULL, "failed to read first name");

  // your last name
  printf("What's your last name? ");
  in = fgets(p.lname, MAX_DATA - 1, stdin);
  check(in != NULL, "failed to read last name");

  // your age
  printf("How old are you? ");
  int num = fscanf(stdin, "%d", &p.age);
  check(num > 0, "You have to enter a number");

  // your eye color
  printf("What color are your eyes:\n");
  for (i = 0; i <= OTHER_EYES; i++) {
    printf("%d) %s\n", i + 1, EYE_COLOR_NAMES[i]);
  }
  printf("> ");

  int eyes = -1;
  num = fscanf(stdin, "%d", &eyes);
  check(num > 0, "You have to enter a number");
  p.ecolor = eyes - 1;
  check(p.ecolor <= OTHER_EYES && p.ecolor >= 0,
        "Do it right, that's not an option");

  printf("How much do you make an hour? ");
  num = fscanf(stdin, "%f", &p.income);
  check(num > 0, "Enter a floating point number");

  printf("\nYou basic info:\n");
  printf("  First name: %s", p.fname);
  printf("  Last name: %s", p.lname);
  printf("  Age: %d\n", p.age);
  printf("  EyeColor: %s\n", EYE_COLOR_NAMES[p.ecolor]);
  printf("  Income: %f\n", p.income);
  return 0;

goerr:
  return -1;
}
