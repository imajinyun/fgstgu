#include "queue.h"
#include <stdio.h>

void InitializeQueue(Queue *queue) {
  queue->head = queue->tail = NULL;
  queue->count = 0;
}

bool IsEmptyQueue(Queue *queue) { return queue->count == 0; }
bool IsFullQueue(const Queue *queue) { return queue->count == MAXQUEUE; }
int QueueItemCount(Queue *queue) {
  return queue->count;
}
bool EnqueueItem(Item item, Queue *queue);
bool DequeueItem(Item *item, Queue *queue);
void QueueRelease(Queue *queue);
