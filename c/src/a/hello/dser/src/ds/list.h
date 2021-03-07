#ifndef DS_LIST_H
#define DS_LIST_H

#include <stdlib.h>

struct ListNode;

typedef struct ListNode {
  struct ListNode *prev;
  struct ListNode *next;
  void *value;
} ListNode;

typedef struct List {
  int count;
  ListNode *first;
  ListNode *last;
} List;

List *listToCreate();
void listToClearNode(List *list);
void listToClearData(List *list);
void listToDestroy(List *list);

#define LIST_COUNT(L) ((L)->count)
#define LIST_FIRST(L) ((L)->first != NULL ? (L)->first->value : NULL)
#define LIST_LAST(L) ((L)->last != NULL ? (L)->last->value : NULL)

void listToPush(List *list, void *value);
void *listToPop(List *list);
void listToUnshift(List *list, void *value);
void *listToShift(List *list);
void *listToRemove(List *list, ListNode *node);

#define LIST_FOREACH(L, S, M, V)                                               \
  ListNode *node = NULL;                                                       \
  ListNode *V = NULL;                                                          \
  for (V = node = L->S; node != NULL; V = node = node->M)

#endif
