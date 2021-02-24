#ifndef LINKEDLIST_H
#define LINKEDLIST_H

#include <stdbool.h>

#define TSIZE 45

struct film {
  char title[TSIZE];
  int rating;
};

typedef struct film Item;
typedef struct node {
  Item item;
  struct node *next;
} Node;
typedef Node *List;

void InitializeList(List *list);
bool IsEmptyList(const List *list);
bool IsFullList(const List *list);
unsigned int ListItemCount(const List *list);
bool ListAddItem(Item item, List *list);
void ListTraverse(const List *list, void (*fn)(Item item));
void ListRelease(List *list);

#endif
