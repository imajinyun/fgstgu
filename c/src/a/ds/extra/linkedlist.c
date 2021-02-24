#include "linkedlist.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

static void CopyToNode(Item item, Node *node);

void InitializeList(List *list) { list = NULL; }

bool IsEmptyList(const List *list) { return (list == NULL) ? true : false; }

bool IsFullList(const List *list) {
  Node *node;
  bool full;
  node = (Node *) malloc(sizeof(Node));
  full = node == NULL ? true : false;
  free(node);
  return full;
}

unsigned int ListItemCount(const List *list) {
  unsigned int count = 0;
  Node *node = *list;
  while (node != NULL) {
    ++count;
    node = node->next;
  }
  return count;
}

bool ListAddItem(Item item, List *list) {
  Node *node;
  Node *curr = *list;
  node = (Node *) malloc(sizeof(Node));
  if (node == NULL) { return false; }
  CopyToNode(item, node);
  node->next = NULL;
  if (curr == NULL) {
    *list = node;
  } else {
    while (curr->next != NULL) { curr = curr->next; }
    curr->next = node;
  }
  return true;
}

void ListTraverse(const List *list, void (*fn)(Item item)) {
  Node *node = *list;
  while (node != NULL) {
    (*fn)(node->item);
    node = node->next;
  }
}

void ListRelease(List *list) {
  Node *node;
  while (*list != NULL) {
    node = (*list)->next;
    free(*list);
    *list = node;
  }
}

static void CopyToNode(Item item, Node *node) { node->item = item; }
