#include <stdio.h>

int main(int argc, char *argv[]) {
  struct names {
    char first[20];
    char last[20];
  };
  struct person {
    int id;
    struct names name;
  };
  struct person stu = {1004, {"Jack", "Ma"}};
  puts(stu.name.first);
  puts(stu.name.last);

  struct human {
    int id;
    struct {
      char first[20];
      char last[20];
    }; // 匿名结构
  };
  struct human tec = {8888, {"Jet", "Li"}};
  printf("The first name: %s\n", tec.first);
  printf("The last name: %s\n", tec.last);
  return 0;
}
