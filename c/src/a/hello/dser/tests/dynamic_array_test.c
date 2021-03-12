#include "ds/dynamic_array.h"
#include "miniunit.h"

static DynamicArray *array = NULL;
// static int *v1 = NULL;
// static int *v2 = NULL;

char *testDynamicArrayCreate() {
  array = dynamicArrayToCreate(sizeof(int), 100);
  TestAssert(array != NULL, "create dynamic array failed");
  TestAssert(array->items != NULL, "items are wrong in dynamic array");
  TestAssert(array->end == 0, "end isn't at the right spot");
  return NULL;
}

char *allDynamicArrayTests() {
  TestStart();
  TestCase(testDynamicArrayCreate);
  return NULL;
}

RunTestMain(allDynamicArrayTests);
