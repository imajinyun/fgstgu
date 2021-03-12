#include "ds/list_sort.h"
#include "miniunit.h"
#include <assert.h>
#include <stdbool.h>
#include <string.h>

char *values[] = {"XXXX", "1234", "abcd", "xjvf", "DNSS"};

#define NUM_VALUES 5

List *createWords() {
  List *words = listToCreate();
  for (int i = 0; i < NUM_VALUES; i++) { listToPush(words, values[i]); }
  return words;
}

int isSortedList(List *list) {
  LIST_FOREACH(list, first, next, curr) {
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
  bool isSorted = isSortedList(words);

  TestAssert(n == 0, "bubble sort failed");
  TestAssert(isSorted, "words are not sorted after bubble sort");
  n = listBubbleSort(words, (listSortCompare) strcmp);
  TestAssert(n == 0, "bubble sort of already sorted failed");

  isSorted = isSortedList(words);
  TestAssert(isSorted, "words should be sort if already double sorted");
  listToClearNode(words);

  words = listToCreate(words);
  n = listBubbleSort(words, (listSortCompare) strcmp);
  TestAssert(n == 0, "bubble sort failed on empty list");
  isSorted = isSortedList(words);
  TestAssert(isSorted, "words should be sorted if empty");
  listToClearNode(words);
  return NULL;
}

char *testMergeSort() {
  List *words = createWords();

  List *l1 = listMergeSort(words, (listSortCompare) strcmp);
  bool isSorted = isSortedList(l1);
  TestAssert(isSorted, "words are not sorted after merge sort");

  List *l2 = listMergeSort(l1, (listSortCompare) strcmp);
  isSorted = isSortedList(l2);
  TestAssert(isSorted, "words are not sorted after merge sort");

  listToClearNode(l2);
  listToClearNode(l1);
  listToClearNode(words);
  return NULL;
}

char *allListSortTests() {
  TestStart();
  TestCase(testBubbleSort);
  TestCase(testMergeSort);
  return NULL;
}

RunTestMain(allListSortTests);
