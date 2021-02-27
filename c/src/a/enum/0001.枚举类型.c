#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  enum weeks { Mon = 1, Tue, Wed, Thu, Fri, Sat, Sun } day;
  printf("Please enter a number between 1 and 7: ");
  scanf("%d", &day);
  switch (day) {
    case Mon:
      puts("Monday");
      break;
    case Tue:
      puts("Tuesday");
      break;
    case Wed:
      puts("Wednesday");
      break;
    case Thu:
      puts("Thursday");
      break;
    case Fri:
      puts("Friday");
      break;
    case Sat:
      puts("Saturday");
      break;
    case Sun:
      puts("Sunday");
      break;
    default:
      puts("Error");
      exit(EXIT_FAILURE);
  }
  return 0;
}
