#include <stdio.h>
#include <stdlib.h>

struct student {
  char name[20];
  int age;
  float score;
};

void inputStudent(struct student *stu, int n);
void outputStudent(const struct student *stu, int n);
void sortStudent(struct student *stu, int n);

int main(int argc, char *argv[]) {
  int n, i, j;
  struct student *ps;


  printf("请输入要录入的学员数量：\n");
  printf("n = ");
  scanf("%d", &n);
  ps = (struct student *) malloc(n * sizeof(struct student));
  inputStudent(ps, n);
  sortStudent(ps, n);
  outputStudent(ps, n);
  return 0;
}

void inputStudent(struct student *stu, int n) {
  for (int i = 0; i < n; i++) {
    printf("输入第 %d 个学员的信息：\n", i + 1);
    printf("name = ");
    scanf("%s", stu[i].name);
    printf("age = ");
    scanf("%d", &stu[i].age);
    printf("score = ");
    scanf("%f", &stu[i].score);
  }
}

void sortStudent(struct student *stu, int n) {
  struct student tmp;
  for (int i = 0; i < n - 1; i++) {
    for (int j = 0; j < j - 1 - i; j++) {
      if (stu[j].score > stu[j + 1].score) {
        tmp = stu[j];
        stu[j] = stu[j + 1];
        stu[j + 1] = tmp;
      }
    }
  }
}

void outputStudent(const struct student *stu, int n) {
  printf("\n");
  for (int i = 0; i < n; i++) {
    printf("第 %d 个学员的信息为：\n", i + 1);
    printf("name = %s, age = %d, score %.2f\n", stu[i].name, stu[i].age, stu[i].score);
  }
}
