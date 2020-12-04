<?php

$a = 1;
$b = $a;

// 先 unset $a，后打印变量
// unset($a);
// var_dump($a); // PHP Notice:  Undefined variable: a
// var_dump($b); // int(1)

// 先 unset $b，后打印变量
// unset($b);
// var_dump($a); // int(1)
// var_dump($b); // PHP Notice:  Undefined variable: b

$a = 'hello world';
$b = &$a;

// 先 unset $a，后打印变量
// unset($a);
// var_dump($a); // PHP Notice:  Undefined variable: a
// var_dump($b); // string(11) "hello world"

// 先 unset $b，后打印变量
// unset($b);
// var_dump($a); // string(11) "hello world"
// var_dump($b); // PHP Notice:  Undefined variable: b

// 全局变量在局部区域使用
// $a = 100;
// var_dump(call_user_func(static function () {
//     $b = 10;
//     return $a * $b;
// })); // PHP Notice:  Undefined variable: a

// 全局变量在局部区域使用
// $a = 100;
// var_dump(call_user_func(static function () {
//     global $a;
//     $b = 8;

//     return $a * $b;
// })); // int(800)

// 全局变量在局部区域中修改
// $a = 100;
// var_dump(call_user_func(static function () {
//     global $a;
//     $a = 200;
//     $b = 4;
//     return $a * $b;
// })); // int(800)
// var_dump($a); // int(200)

// 循环执行完后的变量的访问
// for ($i = 0; $i < 10; $i++) {
// }
// var_dump($i); // int(10)

// 循环执行完后变量的访问
foreach ([1, 2, 3, 4, 5] as $key => $val) {
}
var_dump($key); // int(4)
var_dump($val); // int(5)
