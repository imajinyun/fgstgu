#include <stdbool.h>
#include <stdio.h>
#include <string.h>

#define LENGTH 30

char *str_gets(char *str, int n);
enum spectrum { red, orange, yellow, green, blue, violet };
const char *colors[] = {"red", "orange", "yellow", "green", "blue"};

int main(int argc, char *argv[]) {
  char choice[LENGTH];
  enum spectrum color;
  bool color_is_found = false;
  puts("Enter a color (empty line to quit):");
  while (str_gets(choice, LENGTH) != NULL && choice[0] != '\0') {
    for (color = red; color <= violet; color++) {
      if (strcmp(choice, colors[color]) == 0) {
        color_is_found = true;
        break;
      }
    }
    if (color_is_found) {
      switch (color) {
        case red:
          puts("Roses are red");
          break;
        case orange:
          puts("Poppies are orange");
          break;
        case yellow:
          puts("Sunflowers are yellow");
          break;
        case green:
          puts("Grass is green");
          break;
        case blue:
          puts("Bluebells are blue");
          break;
        case violet:
          puts("Violets are violet");
          break;
      }
    } else {
      printf("I don't know about the color %s\n", choice);
    }
    color_is_found = false;
    puts("Next color, please (empty line to quit):");
  }
  puts("Goodbye!");
  return 0;
}

char *str_gets(char *str, int n) {
  char *res, *find;
  res = fgets(str, n, stdin);
  if (res) {
    if ((find = strchr(str, '\n')) != NULL) {
      *find = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
