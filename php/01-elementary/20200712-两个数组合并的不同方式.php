<?php

$a = [0, 1, 3, 'x' => true];
$b = [1, 2, 3, 'y' => false];
$c = ['x' => false, 'y' => true, 'z', 6];

/*
array(8) {
  [0]=>
  int(0)
  [1]=>
  int(1)
  [2]=>
  int(3)
  ["x"]=>
  bool(true)
  [3]=>
  int(1)
  [4]=>
  int(2)
  [5]=>
  int(3)
  ["y"]=>
  bool(false)
}
array(5) {
  [0]=>
  int(0)
  [1]=>
  int(1)
  [2]=>
  int(3)
  ["x"]=>
  bool(true)
  ["y"]=>
  bool(false)
}
*/
var_dump(array_merge($a, $b), $a + $b);

/*
array(7) {
  [0]=>
  int(0)
  [1]=>
  int(1)
  [2]=>
  int(3)
  ["x"]=>
  bool(false)
  ["y"]=>
  bool(true)
  [3]=>
  string(1) "z"
  [4]=>
  int(6)
}
array(5) {
  [0]=>
  int(0)
  [1]=>
  int(1)
  [2]=>
  int(3)
  ["x"]=>
  bool(true)
  ["y"]=>
  bool(true)
}
*/
var_dump(array_merge($a, $c), $a + $c);
