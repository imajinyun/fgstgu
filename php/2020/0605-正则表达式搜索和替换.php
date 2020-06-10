<?php

$subject = ['1', 'a', '2', 'b', '3', 'A', 'B', '4'];
$pattern = ['/\d/', '/[a-z]/', '/[1a]/'];
$replace = ['A:$0', 'B:$0', 'C:$0'];

/*
Array
(
    [0] => A:C:1
    [1] => B:C:a
    [2] => A:2
    [3] => B:b
    [4] => A:3
    [7] => A:4
)
*/
print_r(preg_filter($pattern, $replace, $subject));

/*
Array
(
    [0] => A:C:1
    [1] => B:C:a
    [2] => A:2
    [3] => B:b
    [4] => A:3
    [5] => A
    [6] => B
    [7] => A:4
)
*/
print_r(preg_replace($pattern, $replace, $subject));
