#include "queue.h"
#include <stdbool.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
  Queue queue;
  Item item;
  char ch;
  InitializeQueue(&queue);
  puts("Testing the Queue interface. Type a to add a value, "
       "type d to delete a value, and type q to quit:");
  while ((ch = getchar()) != 'q') {
    if (ch != 'a' && ch != 'd') { continue; }
    if (ch == 'a') {
      printf("Please enter a integer to add: ");
      scanf("%d", &item);
      if (!IsFullQueue(&queue)) {
        printf("Putting %d into queue\n", item);
        EnqueueItem(item, &queue);
      } else {
        puts("Queue is full!");
      }
    } else {
      if (IsEmptyQueue(&queue)) {
        puts("Nothing to delete!");
      } else {
        DequeueItem(&item, &queue);
        printf("Removing %d from queue\n", item);
      }
    }
    printf("%d items in queue\n", QueueItemCount(&queue));
    puts("Type a to add, d to delete, q to quit:");
  }
  QueueRelease(&queue);
  puts("Done!");
  return 0;
}
