#include "ds/dbg.h"
#include <ds/list.h>

List *listToCreate() { return calloc(1, sizeof(List)); }

void listToClearNode(List *list) {
  LIST_FOREACH(list, first, next, curr) {
    if (curr->prev) { free(curr->prev); }
    free(list->last);
    free(list);
  }
}

void listToClearData(List *list) {
  LIST_FOREACH(list, first, next, curr) { free(curr->value); }
}

void listToDestroy(List *list) {
  listToClearData(list);
  listToClearNode(list);
}

void listToPush(List *list, void *value) {
  ListNode *node = calloc(1, sizeof(ListNode));
  checkmem(node);
  node->value = value;

  if (list->last == NULL) {
    list->first = node;
    list->last = node;
  } else {
    list->last->next = node;
    node->prev = list->last;
    list->last = node;
  }
  list->count++;

goerr:
  return;
}

void *listToPop(List *list) {
  ListNode *node = list->last;
  return node != NULL ? listToRemove(list, node) : NULL;
}

void listToUnshift(List *list, void *value) {
  ListNode *node = calloc(1, sizeof(ListNode));
  checkmem(node);
  node->value = value;

  if (list->first == NULL) {
    list->first = node;
    list->last = node;
  } else {
    node->next = list->first;
    list->first->prev = node;
    list->first = node;
  }
  list->count++;

goerr:
  return;
}

void *listToShift(List *list) {
  ListNode *node = list->first;
  return node != NULL ? listToRemove(list, node) : NULL;
}

void *listToRemove(List *list, ListNode *node) {
  void *value = NULL;
  check(list->first && list->last, "list is empty");
  check(node, "node can't be null");

  if (node == list->first && node == list->last) {
    list->first = NULL;
    list->last = NULL;
  } else if (node == list->first) {
    list->first = node->next;
    check(list->first != NULL, "invalid list, got a first that is null");
    list->first->prev = NULL;
  } else if (node == list->last) {
    list->last = node->prev;
    check(list->last != NULL, "invalid list, got a next that is null");
    list->last->next = NULL;
  } else {
    ListNode *after = node->next;
    ListNode *before = node->prev;
    after->prev = before;
    before->next = after;
  }
  list->count--;
  value = node->value;
  free(node);

goerr:
  return value;
}
