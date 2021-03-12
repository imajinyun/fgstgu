#include "ds/dynamic_array.h"
#include <assert.h>

DynamicArray DynamicArrayToCreate(size_t size, size_t max) {
  DynamicArray *array = malloc(sizeof(DynamicArray));
  checkmem(array);
  array->max = max;
  check(array->max > 0, "you must set an initialize max greater than 0");
  array->items = calloc(sizeof(max), sizeof(void *));
  array->end = 0;
  array->size = size;
  array->rate = DEFAULT_EXPAND_RATE;
  return array;

goerr:
  if (array) { free(array); }
  return NULL;
}
