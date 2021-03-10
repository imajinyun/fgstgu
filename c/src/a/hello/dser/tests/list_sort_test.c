#include "ds/list_sort.h"
#include "miniunit.h"
#include <assert.h>
#include <string.h>

char *values[] = {"XXXX", "1234", "abcd", "xjvf", "DNSS"};

#define NUM_VALUES 5

List *createWords() {
  List *words = listToCreate();
  for (int i = 0; i < NUM_VALUES; i++) { listToPush(words, values[i]); }
  return words;
}

int isSortedList(List *words) {
  LIST_FOREACH(words, first, next, curr) {
    if (curr->next && strcmp(curr->value, curr->next->value) > 0) {
      debug("%s %s", (char *) curr->value, (char *) curr->next->value);
      return 0;
    }
  }
  return 1;
}

char *testBubbleSort() {
  List *words = createWords();
  int n = listBubbleSort(words, (listSortCompare) strcmp);
  TestAssert(n == 0, "bubble sort failed");
  TestAssert(isSortedList(words), "words are not sorted after bubble sort");
  n = listBubbleSort(words, (listSortCompare) strcmp);
  TestAssert(n == 0, "bubble sort of already sorted failed");
  return NULL;
}

char *allListSortTests() {
  TestStart();
  TestCase(testBubbleSort);
  return NULL;
}

RunTestMain(allListSortTests);
