#include <stdio.h>

#define MAX_TITLE 41
#define MAX_AUTHOR 81

struct book {
  char title[MAX_TITLE];
  char author[MAX_AUTHOR];
  float price;
};

/**
 * 复合字面量和结构。
 *
 * -> 复合字面量在所有函数的外部，具有静态存储期；
 *    如果复合字面量在块中，则具有自动存储期。
 *    复合字面量和普通初始化列表的语法规则相同。这意味着，可以在复合字面量中使用指定初始化器。
 */
int main(int argc, char *argv[]) {
  struct book mybook;
  int score;
  printf("Enter test score: ");
  scanf("%d", &score);
  if (score >= 70) {
    mybook = (struct book){"基督山伯爵", "大仲马", 74.50};
  } else {
    mybook = (struct book){"茶花女", "小仲马", 79.99};
  }
  printf("Your assigned reading: ");
  printf("%s by %s: ￥%.2f\n", mybook.title, mybook.author, mybook.price);
  return 0;
}
