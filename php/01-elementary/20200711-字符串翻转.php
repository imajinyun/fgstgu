<?php

$a = 'Hello World!';

// string(12) "!dlroW olleH"
var_dump(strrev($a));

// string(12) "!dlroW olleH"
var_dump(str_reverse($a));

// string(12) "!dlroW olleH"
var_dump(str_reverse2($a));

// string(12) "!dlroW olleH"
var_dump(str_reverse3($a));


function str_reverse(string $str): string
{
    $ans = '';
    $i = strlen($str) - 1;
    while ($i >= 0) {
        $ans .= $str[$i--];
    }

    return $ans;
}

function str_reverse2(string $str): string
{
    $length = strlen($str) - 1;
    $middle = floor($length / 2);
    for ($i = 0, $j = $length; $i <= $middle && $j > $middle; $i++, $j--) {
        [$str[$i], $str[$j]] = [$str[$j], $str[$i]];
    }

    return $str;
}

function str_reverse3(string $str): string
{
    $arr = [];
    $length = strlen($str);
    for ($i = 0; $i < $length; $i++) {
        $arr[$i] = $str[$i];
    }
    $chars = [];
    for ($i = 0; $i < $length; $i++) {
        $chars[$i] = array_pop($arr);
    }

    return implode('', $chars);
}
