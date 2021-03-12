#ifndef DYNAMIC_ARRAY_H
#define DYNAMIC_ARRAY_H

#include "ds/dbg.h"
#include <assert.h>
#include <stdlib.h>

typedef struct DynamicArray {
  int end;
  int max;
  size_t sizt;
  size_t rate;
  void **items;
} DynamicArray;

DynamicArray *dynamicArrayToCreate(size_t size, size_t max);
void dynamicArrayToDestroy(DynamicArray *array);
void dynamicArrayToClear(DynamicArray *array);
int dynamicArrayToExpand(DynamicArray *array);
int dynamicArrayToContract(DynamicArray *array);
int dynamicArrayToPush(DynamicArray *array, void *elem);
void *dynamicArrayToPop(DynamicArray *array);
void dynamicArrayToClearDestroy(DynamicArray *array);

#define DA_LAST(A) ((A)->items[(A)->end - 1])
#define DA_FIRST(A) ((A)->items[0])
#define DA_END(A) ((A)->end)
#define DA_MAX(A) ((A)->max)
#define DA_COUNT(A) dynamicArrayToExpand(A)
#define DA_FREE(E) free((E))

#define DEFAULT_EXPAND_RATE 300

static inline void dynamicArrayToSet(DynamicArray *array, int i, void *elem) {
  check(i < array->max, "dynamic array attempt to set past max");
  if (i > array->end) { array->end = i; }
  array->items[i] = elem;

goerr:
  return;
}

static inline void *dynamicArrayToGet(DynamicArray *array, int i) {
  check(i < array->max, "dynamic array attempt to get past max");
  return array->items[i];

goerr:
  return NULL;
}

static inline void *dynamicArrayRemove(DynamicArray *array, int i) {
  void *elem = array->items[i];
  array->items[i] = NULL;
  return elem;
}

static inline void *dynamicArrayToNew(DynamicArray *array) {
  check(array->elementSize > 0, "can't use dynamicArrayToNew on 0 size dynamic "
                                "array") return calloc(1, array->elementSize);

goerr:
  return NULL;
}

#endif
