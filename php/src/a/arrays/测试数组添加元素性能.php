<?php

/**
 * 结论：`$array[] = 1;` 这种形式向数组中添加元素的性能要好于 `array_push($array, 1)`。
 */

$starttime = get_microtime();
$array = [];
for ($i = 0; $i < 1000000; $i++) {
    $array[] = $i;
}
$endtime = get_microtime();
printf("Time: %f ms\r\n", ($endtime - $starttime) * 1000);

$starttime = get_microtime();
$array = [];
for ($i = 0; $i < 1000000; $i++) {
    array_push($array, $i);
}
$endtime = get_microtime();
printf("Time: %f ms\r\n", ($endtime - $starttime) * 1000);


function get_microtime()
{
    list($usec, $sec) = explode(' ', microtime());

    return (float) $usec + (float) $sec;
}
