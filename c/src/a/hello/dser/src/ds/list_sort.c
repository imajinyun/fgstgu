#include "ds/list_sort.h"
#include "ds/dbg.h"
#include <stdbool.h>

inline void listNodeSwap(ListNode *a, ListNode *b) {
  void *t = a->value;
  a->value = b->value;
  b->value = t;
}

int listBubbleSort(List *list, listSortCompare cmp) {
  bool sorted = true;
  if (LIST_COUNT(list) <= 1) { return 0; }

  do {
    sorted = true;
    LIST_FOREACH(list, first, next, curr) {
      if (curr->next) {
        if (cmp(curr->value, curr->next->value) > 0) {
          listNodeSwap(curr, curr->next);
          sorted = false;
        }
      }
    }
  } while (!sorted);
  return 0;
}

List *listToMerge(List *left, List *right, listSortCompare cmp) {
  List *list = listToCreate();
  void *value = NULL;
  int m = LIST_COUNT(left), n = LIST_COUNT(right);

  while (m > 0 || n > 0) {
    if (m > 0 && n > 0) {
      if (cmp(LIST_FIRST(left), LIST_FIRST(right)) <= 0) {
        value = listToShift(left);
      } else {
        value = listToShift(right);
      }
      listToPush(list, value);
    } else if (m > 0) {
      value = listToShift(left);
      listToPush(list, value);
    } else if (n > 0) {
      value = listToShift(right);
      listToPush(list, value);
    }
  }
  return list;
}

List *listMergeSort(List *list, listSortCompare cmp) {
  int n = LIST_COUNT(list);
  if (n <= 1) { return list; }
  List *left = listToCreate();
  List *right = listToCreate();
  int middle = n / 2;

  LIST_FOREACH(list, first, next, curr) {
    if (middle > 0) {
      listToPush(left, curr->value);
    } else {
      listToPush(right, curr->value);
    }
    middle--;
  }

  List *sleft = listMergeSort(left, cmp);
  List *sright = listMergeSort(right, cmp);
  if (sleft != left) { listToClearNode(left); }
  if (sright != right) { listToClearNode(right); }
  return listToMerge(sleft, sright, cmp);
}
