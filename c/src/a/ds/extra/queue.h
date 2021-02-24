#ifndef QUEUE_H
#define QUEUE_H

#include <stdbool.h>

typedef int Item;

#define MAXQUEUE 10

typedef struct node {
  Item item;
  struct node *next;
} Node;

typedef struct queue {
  Node *head;
  Node *tail;
  int count;
} Queue;

void InitializeQueue(Queue *queue);
bool IsEmptyQueue(Queue *queue);
bool IsFullQueue(const Queue *queue);
int QueueItemCount(Queue *queue);
bool EnqueueItem(Item item, Queue *queue);
bool DequeueItem(Item *item, Queue *queue);
void QueueRelease(Queue *queue);

#endif
