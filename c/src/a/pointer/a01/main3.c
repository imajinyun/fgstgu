#include <stdio.h>

#define SIZE 10

// 以下四种原型声明都是等价的
int sum1(const int *arr, int n);
int sum2(const int *, int);
int sum3(const int arr[], int n);
int sum4(const int [], int);

int main(int argc, char *argv[]) {
  int marbles[SIZE] = {20, 10, 5, 39, 4, 16, 19, 26, 31, 20,};
  printf("The total number of marbles is %ld\n", (long) sum1(marbles, SIZE));
  printf("The total number of marbles is %d\n", sum3(marbles, SIZE));
  printf("The size of marbles is %zd bytes\n", sizeof marbles);
  return 0;
}

int sum1(const int *arr, int n) {
  int sum = 0;
  for (int i = 0; i < n; i++) {
    sum += arr[i];
  }
  printf("The size of arr is %zd bytes\n", sizeof arr);
  return sum;
}

int sum3(const int arr[], int n) {
  int sum = 0;
  for (int i = 0; i < n; i++) {
    sum += *(arr + i);
  }
  printf("The size of arr is %zd bytes\n", sizeof arr);
  return sum;
}
