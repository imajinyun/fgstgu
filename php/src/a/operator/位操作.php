<?php

/**
 * 位操作。
 */

/*
bool(true)
bool(true)
bool(true)
bool(true)
*/
$x = 4; // 100
var_dump(($x ^ 0) === $x);
var_dump(($x ^ 1) === 5);
var_dump(($x ^ (~$x)) === -1);
var_dump(($x ^ $x) === 0);

$a = 1;
$b = 2;
$a ^= $b;
$b ^= $a;
$a ^= $b;
/*
int(2)
int(1)
*/
var_dump($a, $b);

$c = 3;
/*
int(0)
bool(true)
*/
var_dump(($a ^ ($b ^ $c)), ($a ^ ($b ^ $c)) === (($a ^ $b) ^ $c));

$c = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
$d = $e = [];
foreach ($c as $k => $v) {
    // 奇数
    if (($v & 1) === 1) {
        $d[] = $v;
    }

    // 偶数
    if (($v & 1) === 0) {
        $e[] = $v;
    }
}

/*
[1,3,5,7,9]
[2,4,6,8,10]
*/
echo '[', implode(',', $d), ']', PHP_EOL, '[', implode(',', $e), ']', PHP_EOL;

// 清零最低位的 1
$m = 8;
var_dump(($m & ($m - 1))); // int(0)

// 得到最低位的 1
$n = 8;
var_dump(($n & (-$n))); // int(8)
