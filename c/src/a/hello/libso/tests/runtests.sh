#!/usr/bin/env bash

echo ''
echo '✔️ Running unit tests...'

for i in tests/*_test; do
  if test -f $i; then
    if $VALGRIND ./$i 2>>tests/tests.log; then
      echo ''
      echo "=== END:    $i"
    else
      echo "--- FAIL    fatal error in test $i: here's tests/tests.log"
      echo ''
      tail -n 10 tests/tests.log
      exit 1
    fi
  fi
done

echo ''
echo '✔️ All tests is finished'
