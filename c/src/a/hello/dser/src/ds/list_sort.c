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

  while (LIST_COUNT(left) > 0 || LIST_COUNT(right) > 0) {
    if (LIST_COUNT(left) > 0 && LIST_COUNT(right) > 0) {
      if (cmp(LIST_FIRST(left), LIST_FIRST(right)) <= 0) {
        value = listToShift(left);
      } else {
        value = listToShift(right);
      }
      listToPush(list, value);
    } else if (LIST_COUNT(left) > 0) {
      value = listToShift(left);
      listToPush(list, value);
    } else if (LIST_COUNT(right) > 0) {
      value = listToShift(right);
      listToPush(list, value);
    }
  }
  return list;
}

List *listMergeSort(List *list, listSortCompare cmp) {
  if (LIST_COUNT(list) <= 1) { return list; }
  List *left = listToCreate();
  List *right = listToCreate();
  int middle = LIST_COUNT(list) / 2;

  LIST_FOREACH(list, first, next, cur) {
    if (middle > 0) {
      listToPush(left, cur->value);
    } else {
      listToPush(right, cur->value);
    }
    middle--;
  }

  List *sleft = listMergeSort(left, cmp);
  List *sright = listMergeSort(right, cmp);
  if (sleft != left) { listToClearNode(left); }
  if (sright != right) { listToClearNode(right); }
  return listToMerge(sleft, sright, cmp);
}
