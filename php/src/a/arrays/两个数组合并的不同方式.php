<?php

require dirname(dirname(__DIR__)).'/vendor/autoload.php';

$a = [0, 1, 3, 'x' => true];
$b = [1, 2, 3, 'y' => false];
$c = ['x' => false, 'y' => true, 'z', 6];

/*
^ array:8 [
  0 => 0
  1 => 1
  2 => 3
  "x" => true
  3 => 1
  4 => 2
  5 => 3
  "y" => false
]
^ array:5 [
  0 => 0
  1 => 1
  2 => 3
  "x" => true
  "y" => false
]
*/
dd(array_merge($a, $b), $a + $b);

/*
^ array:7 [
  0 => 0
  1 => 1
  2 => 3
  "x" => false
  "y" => true
  3 => "z"
  4 => 6
]
^ array:5 [
  0 => 0
  1 => 1
  2 => 3
  "x" => true
  "y" => true
]
*/
dd(array_merge($a, $c), $a + $c);
