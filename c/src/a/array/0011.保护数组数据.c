#include <stdio.h>

#define SIZE 5

void showArray(const double arr[], int n);
void multArray(double arr[], int n, double mult);

int main(int argc, char *argv[]) {
  double data[SIZE] = {20.0, 17.66, 8.33, 22.22, 15.69};
  printf("The original data array:\n");
  showArray(data, SIZE);
  multArray(data, SIZE, 2.5);
  printf("The data array after calling `multArray()`:\n");
  showArray(data, SIZE);
  return 0;
}

void showArray(const double arr[], int n) {
  for (int i = 0; i < n; i++) { printf("%8.3f", arr[i]); }
  putchar('\n');
}

void multArray(double arr[], int n, double mult) {
  for (int i = 0; i < n; i++) { arr[i] *= mult; }
}
