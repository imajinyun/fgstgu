#include "miniunit.h"
#include <assert.h>
#include <ds/list.h>

char *testListCreate() {
  List *list = listToCreate();
  TestAssert(list != NULL, "failed to create list");
  return NULL;
}

char *testListDestroy() {
  List *list = listToCreate();
  listToDestroy(list);
  TestAssert(list != NULL, "failed to destroy list");
  return NULL;
}

char *testListPush() {
  char *t1 = "#1 test data";
  char *t2 = "#2 test data";
  char *t3 = "#3 test data";

  List *list = listToCreate();
  listToPush(list, t1);
  TestAssert(LIST_LAST(list) == t1, "wrong last value");

  listToPush(list, t2);
  TestAssert(LIST_LAST(list) == t2, "wrong last value");

  listToPush(list, t3);
  TestAssert(LIST_LAST(list) == t3, "wrong last value");

  TestAssert(LIST_COUNT(list) == 3, "wrong count on push");
  return NULL;
}

char *testListPop() {
  char *t1 = "#1 test data";
  char *t2 = "#2 test data";
  char *t3 = "#3 test data";

  List *list = listToCreate();
  listToPush(list, t1);
  listToPush(list, t2);
  listToPush(list, t3);
  TestAssert(LIST_COUNT(list) == 3, "wrong count on push");

  char *value = listToPop(list);
  TestAssert(value == t3, "wrong last value");
  TestAssert(LIST_COUNT(list) == 2, "wrong count on push");

  value = listToPop(list);
  // TestAssert(value == t2, "wrong last value");
  TestAssert(LIST_COUNT(list) == 1, "wrong count on push");

  value = listToPop(list);
  // TestAssert(value == t1, "wrong last value");
  TestAssert(LIST_COUNT(list) == 0, "wrong count on push");
  return NULL;
}

char *testListUnshift() {
  char *t1 = "#1 test data";
  char *t2 = "#2 test data";
  char *t3 = "#3 test data";
  List *list = listToCreate();
  TestAssert(LIST_COUNT(list) == 0, "wrong count on push");

  listToUnshift(list, t1);
  TestAssert(LIST_FIRST(list) == t1, "wrong first value");

  listToUnshift(list, t2);
  TestAssert(LIST_FIRST(list) == t2, "wrong first value");

  listToUnshift(list, t3);
  TestAssert(LIST_FIRST(list) == t3, "wrong first value");
  TestAssert(LIST_COUNT(list) == 3, "wrong count on push");
  return NULL;
}

char *testListShift() {
  char *t1 = "#1 test data";
  char *t2 = "#2 test data";
  char *t3 = "#3 test data";
  List *list = listToCreate();
  listToPush(list, t1);
  listToPush(list, t2);
  listToPush(list, t3);
  TestAssert(LIST_COUNT(list) == 3, "wrong count on push");
  TestAssert(LIST_COUNT(list) != 0, "wrong count before shift");

  char *value = listToShift(list);
  TestAssert(value == t1, "wrong value on shift");

  value = listToShift(list);
  TestAssert(value == t2, "wrong value on shift");

  value = listToShift(list);
  TestAssert(value == t3, "wrong value on shift");
  TestAssert(LIST_COUNT(list) == 0, "wrong count on shift");
  return NULL;
}

char *testListRemove() {
  char *t1 = "#1 test data";
  char *t2 = "#2 test data";
  char *t3 = "#3 test data";
  List *list = listToCreate();
  listToPush(list, t1);
  listToPush(list, t2);
  listToPush(list, t3);
  char *value = listToRemove(list, list->first->next);
  TestAssert(value == t2, "wrong removed element");
  return NULL;
}

char *allListTests() {
  TestStart();
  TestCase(testListCreate);
  TestCase(testListDestroy);
  TestCase(testListPush);
  TestCase(testListPop);
  TestCase(testListUnshift);
  TestCase(testListShift);
  TestCase(testListRemove);
  return NULL;
}

RunTestMain(allListTests);
