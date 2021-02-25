#include "queue.h"
#include <stdio.h>
#include <stdlib.h>

static void CopyToNode(Item item, Node *node);
static void CopyToItem(Node *node, Item *item);

void InitializeQueue(Queue *queue) {
  queue->head = queue->tail = NULL;
  queue->count = 0;
}

bool IsEmptyQueue(Queue *queue) { return queue->count == 0; }

bool IsFullQueue(const Queue *queue) { return queue->count == MAXQUEUE; }

int QueueItemCount(Queue *queue) { return queue->count; }

bool EnqueueItem(Item item, Queue *queue) {
  Node *node;
  if (IsFullQueue(queue)) { return false; }
  node = (Node *) malloc(sizeof(Node));
  if (node == NULL) {
    fprintf(stderr, "Unable to allocate memory!\n");
    exit(EXIT_FAILURE);
  }
  CopyToNode(item, node);
  node->next = NULL;
  if (IsEmptyQueue(queue)) {
    queue->head = node;
  } else {
    queue->tail->next = node;
  }
  queue->tail = node;
  queue->count++;
  return true;
}

bool DequeueItem(Item *item, Queue *queue) {
  Node *node;
  if (IsEmptyQueue(queue)) { return false; }
  CopyToItem(queue->head, item);
  node = queue->head;
  queue->head = queue->head->next;
  free(node);
  queue->count--;
  if (queue->count == 0) { queue->tail = NULL; }
  return true;
}

void QueueRelease(Queue *queue) {
  Item item;
  while (!IsEmptyQueue(queue)) { DequeueItem(&item, queue); }
}

static void CopyToNode(Item item, Node *node) { node->item = item; }
static void CopyToItem(Node *node, Item *item) { *item = node->item; }
