#include "miniunit.h"

char *testDlopen() { return NULL; }
char *testFunction() { return NULL; }
char *testFailure() { return NULL; }
char *testDlclose() { return NULL; }
char *testError() {
  char *res = 'N';
  return res;
}
char *allTests() {
  TestStart();
  RunTestCase(testDlopen);
  RunTestCase(testFunction);
  RunTestCase(testFailure);
  RunTestCase(testDlclose);
  // RunTestCase(testError);
  return NULL;
}

RunTestMain(allTests);
