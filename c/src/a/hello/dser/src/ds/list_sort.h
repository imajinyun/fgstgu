#ifndef DS_LIST_SORT_H
#define DS_LIST_SORT_H

#include "ds/list.h"

typedef int (*listSortCompare)(const void *a, const void *b);
int listBubbleSort(List *list, listSortCompare cmp);
List *listMergeSort(List *list, listSortCompare cmp);

#endif
