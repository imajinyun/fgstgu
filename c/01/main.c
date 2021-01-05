/**
 * 学生成绩管理系统。
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * 学员成绩信息结构体。
 */
struct student_score {
  int number;                 // 学员学号
  char name[30];              // 学员姓名
  float chinese;              // 语文成绩
  float mathematics;          // 数学成绩
  float english;              // 英语成绩
  struct student_score *next; // 下一个学员成绩信息结构体
};

typedef struct student_score score_t;

#define LENGTH sizeof(struct student_score)

// 是否调试
#define DEBUG 1

// 控制每行 80 个字符
#define LINE_LENGTH 80

int n; // n 统计输入的学员人数

void swap_number(score_t *a, score_t *b) {
  int number; // 学号
  number = a->number;
  a->number = b->number;
  b->number = number;
}

void swap_name(score_t *a, score_t *b) {
  char name[10];
  strcpy(name, a->name);
  strcpy(a->name, b->name);
  strcpy(b->name, name);
}

void swap_score(score_t *a, score_t *b) {
  float score; // 成绩
  score = a->chinese;
  a->chinese = b->chinese;
  b->chinese = score;

  score = a->mathematics;
  a->mathematics = b->mathematics;
  b->mathematics = score;

  score = a->english;
  a->english = b->english;
  b->english = score;
}

void print_border_symbol() {
  for (int i = 1; i < LINE_LENGTH; ++i) {
    if (LINE_LENGTH == i) {
      printf("*\n");
    } else {
      printf("*");
    }
  }
  printf("\n");
}

void print_table_head() {
  printf("--------------------------------------------------------------------------------\n");
  printf("|学号\t|姓名\t|语文\t|数学\t|英语\t|\n");
  printf("--------------------------------------------------------------------------------\n");
}

/**
 * 显示系统菜单信息。
 *
 * @param int num 用户选择的菜单编号
 * @return int
 */
int menu(int num) {
  printf("🎉 学生成绩管理系统");
  printf("\n");

  print_border_symbol();
  printf("1-创建学生成绩表\t\t2-显示学生的成绩\n");
  printf("3-查询学生的成绩\t\t4-添加学生的成绩\n");
  printf("5-删除学生的成绩\t\t6-排序学生的成绩\n");
  printf("7-保存学生的成绩\t\t8-统计学生的成绩\n");
  printf("9-加载学生的成绩\t\t0-退出学生成绩管理系统\n");
  print_border_symbol();

  printf("请选择您要操作的序号（选择 0 退出）：");
  scanf("%d", &num);
  getchar();

  return num;
}

/**
 * 创建学生成绩信息链表。
 *
 * @return
 */
score_t *create_linked_list() {
  printf("正在进行的操作是【创建学生成绩信息链表】...\n");
  score_t *head;
  score_t *p1, *p2, *p3, *max;
  int i, j;
  float score;
  char t[10];
  n = 0; // 统计输入的学员人数

  p1 = p2 = p3 = (score_t *) malloc(LENGTH);
  head = p3;

number_label1:
  printf("请输入学生学号（学号应大于 0，输入 0 退出）：");
  scanf("%d", &p1->number);
  while (p1->number < 0) {
    getchar();
    printf("输入错误，请重新输入学生学号：");
    scanf("%d", &p1->number);
  }
  if (p1->number == 0) {
    goto finished;
  } else {
    p3 = head;
    if (n > 0) {
      for (i = 0; i < n; ++i) {
        if (p1->number != p3->number) {
          p3 = p3->next;
        } else {
          printf("⚠️学号重复，请重新输入！\n");
          goto number_label1;
        }
      }
    }
  }

  printf("请输入学生姓名：");
  scanf("%s", p1->name);

  printf("请输入语文成绩（0~100）：");
  scanf("%f", &p1->chinese);
  while (p1->chinese < 0 || p1->chinese > 100) {
    getchar();
    printf("输入错误，请重新输入语文成绩：");
    scanf("%f", &p1->chinese);
  }

  printf("请输入数学成绩（0~100）：");
  scanf("%f", &p1->mathematics);
  while (p1->mathematics < 0 || p1->mathematics > 100) {
    getchar();
    printf("输入错误，请重新输入数学成绩：");
    scanf("%f", &p1->mathematics);
  }

  printf("请输入英文成绩（0~100）：");
  scanf("%f", &p1->english);
  while (p1->english < 0 || p1->english > 100) {
    getchar();
    printf("输入错误，请重新输入英文成绩：");
    scanf("%f", &p1->english);
  }

  head = NULL;
  while (p1->number != 0) {
    n += 1;
    if (n == 1) {
      head = p1;
    } else {
      p2->next = p1;
    }
    p2 = p1;
    p1 = (score_t *) malloc(LENGTH);

  number_label2:
    printf("请输入学生学号（学号应大于 0，输入 0 退出）：");
    scanf("%d", &p1->number);
    while (p1->number < 0) {
      getchar();
      printf("输入错误，请重新输入学生学号：");
      scanf("%d", &p1->number);
    }
    if (p1->number == 0) {
      goto finished;
    } else {
      p3 = head;
      if (n > 0) {
        for (i = 0; i < n; ++i) {
          if (p1->number != p3->number) {
            p3 = p3->next;
          } else {
            printf("⚠️学号重复，请重新输入！\n");
            goto number_label2;
          }
        }
      }
    }

    printf("请输入学生姓名：");
    scanf("%s", p1->name);

    printf("请输入语文成绩（0~100）：");
    scanf("%f", &p1->chinese);
    while (p1->chinese < 0 || p1->chinese > 100) {
      getchar();
      printf("输入错误，请重新输入语文成绩：");
      scanf("%f", &p1->chinese);
    }

    printf("请输入数学成绩（0~100）：");
    scanf("%f", &p1->mathematics);
    while (p1->mathematics < 0 || p1->mathematics > 100) {
      getchar();
      printf("输入错误，请重新输入数学成绩：");
      scanf("%f", &p1->mathematics);
    }

    printf("请输入英文成绩（0~100）：");
    scanf("%f", &p1->english);
    while (p1->english < 0 || p1->english > 100) {
      getchar();
      printf("输入错误，请重新输入英文成绩：");
      scanf("%f", &p1->english);
    }
  }

finished:
  p1 = head;
  for (i = 1; i < n; i++) {
    for (j = i + 1; j <= n; j++) {
      max = p1;
      p1 = p1->next;
      if (max->number > p1->number) {
        swap_number(max, p1);
        swap_name(max, p1);
        swap_score(max, p1);
      }
    }
    p1 = head;
  }
  p2->next = NULL;
  printf("输入的学生数为 %d 个！\n", n);

  return (head);
}

/**
 * 打印学生成绩信息。
 *
 * @param head
 */
void dump(score_t *head) {
  printf("\n正在进行的操作是【打印学生成绩信息】...\n");
  score_t *p;
  if (head == NULL) {
    printf("\n没有任何学生成绩信息！\n");
  } else {
    print_table_head();

    p = head;
    do {
      printf("|%d\t|%s\t|%.1f\t|%.1f\t|%.1f\t|\n", p->number, p->name, p->chinese, p->mathematics, p->english);
      printf("--------------------------------------------------------------------------------\n");
      p = p->next;
    } while (p != NULL);
  }
}

/**
 * 添加学生成绩信息。
 *
 * @param head
 * @param stu
 * @return
 */
score_t *add(score_t *head, score_t *stu) {
  printf("\n正在进行的操作是【添加学生成绩信息】...\n");
  score_t *p0, *p1, *p2, *p3, *max;
  int i;
  float score;
  char t[10];
  stu = (score_t *) malloc(LENGTH);

add_label:
  printf("请输入学生学号（学号应大于 0）：");
  scanf("%d", &stu->number);
  while (stu->number < 0) {
    getchar();
    printf("输入错误，请重新输入学生学号：");
    scanf("%d", &stu->number);
  }

  if (stu->number == 0) {
    goto end_label;
  } else {
    p3 = head;
    if (n > 0) {
      for (i = 0; i < n; i++) {
        if (stu->number != p3->number) {
          p3 = p3->next;
        } else {
          printf("⚠️学号重复，请重新输入！\n");
          goto add_label;
        }
      }
    }
  }

  printf("请输入学生姓名：");
  scanf("%s", stu->name);

  printf("请输入语文成绩（0~100）：");
  scanf("%f", &stu->chinese);
  while (stu->chinese < 0 || stu->chinese > 100) {
    getchar();
    printf("输入错误，请重新输入语文成绩：");
    scanf("%f", &stu->chinese);
  }

  printf("请输入数学成绩（0~100）：");
  scanf("%f", &stu->mathematics);
  while (stu->mathematics < 0 || stu->mathematics > 100) {
    getchar();
    printf("输入错误，请重新输入数学成绩：");
    scanf("%f", &stu->mathematics);
  }

  printf("请输入英文成绩（0~100）：");
  scanf("%f", &stu->english);
  while (stu->english < 0 || stu->english > 100) {
    getchar();
    printf("输入错误，请重新输入英文成绩：");
    scanf("%f", &stu->english);
  }

  p1 = head;
  p0 = stu;
  if (head == NULL) {
    head = p0;
    p0->next = NULL;
  } else {
    if (p1->next == NULL) {
      p1->next = p0;
      p0->next = NULL;
    } else {
      while (p1->next != NULL) {
        p2 = p1;
        p1 = p1->next;
      }
      p1->next = p0;
      p0->next = NULL;
    }
  }
  n += 1;
  p1 = head;
  p0 = stu;
  for (i = 0; i < n; ++i) {
    max = p1;
    p1 = p1->next;
    if (max->number > p1->number) {
      swap_number(max, p1);
      swap_name(max, p1);
      swap_score(max, p1);
    }
    max = head;
    p1 = head;
  }

end_label:
  printf("♻️目前输入的学生人数为 %d 个\n", n);

  return (head);
}

/**
 * 查询学生成绩信息。
 *
 * @param head
 * @return
 */
score_t *search(score_t *head) {
  printf("\n正在进行的操作是【查询学生成绩信息】...\n");
  int number;
  score_t *p1, *p2;
  printf("输入要查询的学生学号（输入 0 输出）：");
  scanf("%d", &number);

  while (number != 0) {
    if (head == NULL) {
      printf("\n没有任何学生资料！\n");
      return (head);
    }

    print_table_head();
    p1 = head;
    while (number != p1->number && p1->next != NULL) {
      p2 = p1;
      p1 = p1->next;
    }

    if (number == p1->number) {
      printf("|%d\t|%s\t|%.1f\t|%.1f\t|%.1f\t|\n", p1->number, p1->name, p1->chinese, p1->mathematics, p1->english);
      printf("--------------------------------------------------------------------------------\n");
    } else {
      printf("⚠️学号为 %d 的学生不存在！\n", number);
    }
    printf("输入要查询的学生学号（输入 0 退出）：");
    scanf("%d", &number);
  }
  printf("已经退出了查询学生学号功能！\n");

  return (head);
}

/**
 * 删除学生成绩信息。
 *
 * @param head
 * @return
 */
score_t *del(score_t *head) {
  printf("\n正在进行的操作是【删除学生成绩信息】...\n");
  score_t *p1, *p2;
  int number;

  printf("输入要删除的学生学号（输入 0 退出）：");
  scanf("%d", &number);
  getchar();
  while (number != 0) {
    if (head == NULL) {
      printf("\n没有任何学生资料！\n");
      return (head);
    }

    p1 = head;
    while (number != p1->number && p1->next != NULL) { // 找到要删除的学号的学生记录
      p2 = p1;
      p1 = p1->next;
    }

    if (number == p1->number) {
      if (p1 == head) {
        head = p1->next;
      } else {
        p2->next = p1->next;
      }
      printf("✅删除学号为 %d 的学生\n", number);
      n -= 1;
    } else {
      printf("⚠️学号为 %d 的学生不存在\n", number);
    }

    printf("输入要删除的学号学号（输入 0 退出）：");
    scanf("%d", &number);
    getchar();
  }

#ifdef DEBUG
  printf("已经退出了删除删除学生功能！\n");
#endif
  printf("目前的学生人数为 %d 个！\n", n);

  return (head);
}

/**
 * 排序学生成绩信息。
 *
 * @param head
 * @return
 */
score_t *sort(score_t *head) {
  printf("\n正在进行的操作是【排序学生成绩信息】...\n");
  score_t *p, *max;
  int i, j, x;

  if (head == NULL) {
    printf("没有任何学生资料，请先建立链表！\n");
    return (head);
  }

  max = p = head;
  for (i = 0; i < LINE_LENGTH; ++i) { printf("*"); }
  printf("\n1-按学生学号排序\t2-按学生姓名排序\n");
  printf("3-按语文成绩排序\t4-按数学成绩排序\n");
  printf("5-按英文成绩排序\n");
  for (i = 0; i < LINE_LENGTH; ++i) { printf("*"); }

  printf("\n请选择操作：");
  scanf("%d", &x);
  getchar();
  switch (x) {
    case 1: // 按学生学号排序
      for (i = 1; i < n; i++) {
        for (j = i + 1; j <= n; j++) {
          max = p;
          p = p->next;
          if (max->number > p->number) {
            swap_number(max, p);
            swap_name(max, p);
            swap_score(max, p);
          }
        }
        max = head;
        p = head;
      }
      dump(head);
      break;
    case 2: // 按学生姓名排序
      for (i = 1; i < n; i++) {
        for (j = i + 1; j <= n; j++) {
          max = p;
          p = p->next;
          if (strcmp(max->name, p->name) > 0) {
            swap_number(max, p);
            swap_name(max, p);
            swap_score(max, p);
          }
        }
        max = head;
        p = head;
      }
      dump(head);
      break;
    case 3: // 按语文成绩排序
      for (i = 1; i < n; i++) {
        for (j = i + 1; j <= n; j++) {
          max = p;
          p = p->next;
          if (max->chinese > p->chinese) {
            swap_number(max, p);
            swap_name(max, p);
            swap_score(max, p);
          }
        }
        max = head;
        p = head;
      }
      dump(head);
      break;
    case 4: // 按数学成绩排序
      for (i = 1; i < n; i++) {
        for (j = i + 1; j <= n; j++) {
          max = p;
          p = p->next;
          if (max->mathematics > p->mathematics) {
            swap_number(max, p);
            swap_name(max, p);
            swap_score(max, p);
          }
        }
        max = head;
        p = head;
      }
      dump(head);
      break;
    case 5: // 按英文成绩排序
      for (i = 1; i < n; i++) {
        for (j = i + 1; j <= n; j++) {
          max = p;
          p = p->next;
          if (max->english > p->english) {
            swap_number(max, p);
            swap_name(max, p);
            swap_score(max, p);
          }
        }
        max = head;
        p = head;
      }
      dump(head);
      break;
    default:
      printf("⚠️你的输入有误，请检查后重试！\n");
      return 0;
      break;
  }

  return (head);
}

/**
 * 保存学生成绩信息。
 *
 * @param info
 * @return
 */
int *save(score_t *p) {
  printf("\n正在进行的操作是【保存学生成绩信息】...\n");
  FILE *fp;
  char filename[20];

  printf("请输入文件路径及文件名称：");
  scanf("%s", filename);
  if ((fp = fopen(filename, "w+")) == NULL) {
    printf("不能打开文件！\n");
    return 0;
  }

  fprintf(fp, "学生成绩管理系统\n");
  fprintf(fp, "--------------------------------------------------------------------------------\n");
  fprintf(fp, "|学号\t|姓名\t|语文\t|数学\t|英文\t|\n");
  fprintf(fp, "--------------------------------------------------------------------------------\n");

  while (p != NULL) {
    fprintf(fp, "%d\t%s\t%.1f\t%.1f\t%.1f\t\n", p->number, p->name, p->chinese, p->mathematics, p->english);
    p = p->next;
  }
  fclose(fp);
  printf("文件已经保存！\n");

  return 0;
}

/**
 * 读取学生成绩信息。
 *
 * @param head
 * @return
 */
score_t *load(score_t *head) {
  printf("\n正在进行的操作是【加载学生成绩信息】...\n");
  score_t *p, *q;
  int m = 0;
  char filename[20];
  FILE *fp;

  printf("请输入文件路径及文件名称：");
  scanf("%s", filename);
  if ((fp = fopen(filename, "r+")) == NULL) {
    printf("不能打开文件！\n");
    return 0;
  }
  fscanf(fp, "学生成绩管理系统\n");
  fscanf(fp, "--------------------------------------------------------------------------------\n");
  fscanf(fp, "|学号\t|姓名\t|语文\t|数学\t|英文\t|\n");
  fscanf(fp, "--------------------------------------------------------------------------------\n");

  printf("学生成绩管理系统\n");
  print_table_head();

  m += 1;
  if (m == 1) {
    p = (score_t *) malloc(LENGTH);
    fscanf(fp, "%d%s%f%f%f", &p->number, p->name, &p->chinese, &p->mathematics, &p->english);
    printf("|%d\t|%s\t|%.1f\t|%.1f\t|%.1f\t|\n", p->number, p->name, p->chinese, p->mathematics, p->english);
    head = NULL;

    do {
      n += 1;
      if (n == 1) {
        head = p;
      } else {
        q->next = p;
      }
      q = p;
      p = (score_t *) malloc(LENGTH);
      fscanf(fp, "%d%s%f%f%f\n", &p->number, p->name, &p->chinese, &p->mathematics, &p->english);
      printf("|%d\t|%s\t|%.1f\t|%.1f\t|%.1f\t|\n", p->number, p->name, p->chinese, p->mathematics, p->english);
    } while (!feof(fp));

    q->next = p;
    p->next = NULL;
    n += 1;
  }
  printf("--------------------------------------------------------------------------------\n");
  fclose(fp);

  return (head);
}

/**
 * 统计学生成绩信息。
 * @param head
 * @return
 */
score_t *statistics(score_t *head) {
  printf("\n正在进行的操作是【统计学生成绩信息】...\n");
  score_t *p;
  float sum, avg, sum1, sum2, sum3, avg1, avg2, avg3, min, max;
  int x, y = 0, i = 0;
  char minname[10], maxname[10];
  sum = avg = sum1 = sum2 = sum3 = avg1 = avg2 = avg3 = min = max = 0;
  p = head;

  for (int j = 0; j < LINE_LENGTH; j++) { printf("*"); }
  printf("\n1-个人总分和平均分\t2-单科平均分\t3-总分最高分和最低分\n");
  for (int j = 0; j < LINE_LENGTH; j++) { printf("*"); }
  printf("\n请输入要操作的序号：");
  scanf("%d", &x);
  getchar();

  switch (x) {
    case 1:
      if (head == NULL) {
        printf("\n没有任何学生资料！\n");
        return (head);
      } else {
        printf("--------------------------------------------------------------------------------\n");
        printf("|学号\t|姓名\t|语文\t|数学\t|英文\t|总分\t|平均分\t|\n");
        printf("--------------------------------------------------------------------------------\n");
        char format[] = "|%d\t|%s\t|%.1f\t|%.1f\t|%.1f\t|%.1f\t|%1.f\t|\n";
        while (p != NULL) {
          sum = p->chinese + p->mathematics + p->english;
          avg = sum / 3;
          printf(format, p->number, p->name, p->chinese, p->mathematics, p->english, sum, avg);
          printf("--------------------------------------------------------------------------------\n");
          p = p->next;
        }
        printf("\n");
      }
      return (head);
      break;
    case 2:
      if (head == NULL) {
        printf("\n没有任何学生资料！\n");
        return (head);
      }
      while (p != NULL) {
        sum1 += p->chinese;
        sum2 += p->mathematics;
        sum3 += p->english;

        y += 1;
        avg1 = sum1 / y;
        avg2 = sum2 / y;
        avg3 = sum3 / y;
        p = p->next;
      }
      printf("语文成绩平均分为：%0.1f\n", avg1);
      printf("数学成绩平均分为：%0.1f\n", avg2);
      printf("英文成绩平均分为：%0.1f\n", avg3);
      printf("\n");
      return (head);
      break;
    case 3:
      if (head == NULL) {
        printf("\n没有任何学生资料！\n");
        return (head);
      }
      min = max = p->chinese + p->mathematics + p->english;
      while (i < n) {
        i += 1;
        sum = p->chinese + p->mathematics + p->english;
        if (max < sum) {
          max = sum;
          strcpy(maxname, p->name);
        }

        if (min > sum) {
          min = sum;
          strcpy(minname, p->name);
        }

        p = p->next;
      }
      printf("总分最高分：%.1f，姓名：%s\n", max, maxname);
      printf("总分最低分：%.1f，姓名：%s\n", min, minname);
      printf("\n");
      return (head);
      break;
    default:
      printf("⚠️你的输入有误，请检查后重试！\n");
      break;
  }

  return (head);
}

/**
 * 主入口。
 *
 * @param argc
 * @param argv
 * @return
 */
int main(int argc, char *argv[]) {
  int code = 0;
  score_t *head = NULL, *stu = NULL;

  while (1) {
    code = menu(code);

    switch (code) {
      case 1: // 创建学生成绩信息链表
        head = create_linked_list();
        break;
      case 2: // 打印学生的成绩
        dump(head);
        break;
      case 3: // 查询学生的成绩
        head = search(head);
        break;
      case 4: // 添加学生的成绩
        head = add(head, stu);
        break;
      case 5: // 删除学生的成绩
        head = del(head);
        break;
      case 6: // 排序学生的成绩
        head = sort(head);
        break;
      case 7: // 保存学生的成绩
        save(head);
        break;
      case 8: // 统计学生的成绩
        head = statistics(head);
        break;
      case 9: // 加载学生的成绩
        head = load(head);
        break;
      case 0: // 退出学生成绩管理系统
        exit(0);
        break;
      default:
        printf("🙏 你的输入有误，请检查后重试！\n");
        break;
    }
  }

  return 0;
}
