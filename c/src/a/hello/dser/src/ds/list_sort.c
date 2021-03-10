#include "ds/dbg.h"
#include "ds/list_sort.h"
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
      if (curr->value) {
        if (cmp(curr->value, curr->next->value) > 0) {
          listNodeSwap(curr, curr->next);
          sorted = false;
        }
      }
    }
  } while (!sorted);
  return 0;
}
