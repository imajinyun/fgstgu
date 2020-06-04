<?php

$a = [1, 2, 3, 4, 5, 6, 7, 8, 9];
$b = 1;

$c = array_reduce($a, function ($x, $y) {
    echo $x . '@' . $y . PHP_EOL;
    return $x + $y;
}, $b);

/*
1@1
2@2
4@3
7@4
11@5
16@6
22@7
29@8
37@9
int(46)
*/
var_dump($c);

$d = array_reduce($a, function ($x, $y) {
    echo $x . '@' . $y . PHP_EOL;
    return $x * $y;
}, $b);

/*
1@1
1@2
2@3
6@4
24@5
120@6
720@7
5040@8
40320@9
int(362880)
*/
var_dump($d);

// int(45)
var_dump(array_reduce($a, 'sum'));

// int(46)
var_dump(array_reduce($a, 'sum', 1));

// int(0)
var_dump(array_reduce($a, 'fact'));

// int(362880)
var_dump(array_reduce($a, 'fact', 1));

function sum($carry, $item): int
{
    return $carry += $item;
}

function fact($carry, $item)
{
    return $carry *= $item;
}
