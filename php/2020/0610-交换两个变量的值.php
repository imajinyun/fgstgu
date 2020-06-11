<?php

$a = 1;
$b = 2;
$c = 3;
$d = 4;
$e = 5;
$f = 6;

swap1($a, $b);
swap2($c, $d);
swap3($e, $f);

/*
int(2)
int(1)
int(4)
int(3)
int(6)
int(5)
*/
var_dump($a, $b, $c, $d, $e, $f);


function swap1(int &$a, int &$b): void
{
    if ($a !== $b) {
        $t = $a;
        $a = $b;
        $b = $t;
    }
}

function swap2(int &$a, int &$b): void
{
    if ($a !== $b) {
        [$b, $a] = [$a, $b];
    }
}

function swap3(int &$a, int &$b): void
{
    if ($a !== $b) {
        $a ^= $b;
        $b ^= $a;
        $a ^= $b;
    }
}
