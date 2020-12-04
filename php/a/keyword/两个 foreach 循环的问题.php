<?php

/**
 * @see https://www.php.net/manual/en/control-structures.foreach.php
 */

$array = [1, 2, 3];
foreach ($array as $key => &$val);

/*
Array
(
[0] => 1
[1] => 2
[2] => 3
)
*/
print_r($array);

foreach ($array as $key => $val);
/*
Array
(
[0] => 1
[1] => 2
[2] => 2
)
*/
print_r($array);

// 解决方式在两个 foreach 循环之间加入 unset($val); 语句
