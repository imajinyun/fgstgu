#include <stdio.h>

struct student {
  double score;
};

struct teacher {
  double salary;
};

struct person {
  char name[20];
  int age;
  int status; // 1:学生 2:老师
  union {
    struct student stu;
    struct teacher tec;
  };
};

int main(int argc, char *argv[]) {
  struct person p;
  printf("sizeof(p) = %lu, sizeof(struct person) = %lu\n", sizeof(p), sizeof(struct person));
  return 0;
}
