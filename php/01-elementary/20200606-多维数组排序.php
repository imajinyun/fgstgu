<?php

$array = [
    ['id' => 3, 'name' => 'B', 'created_at' => '2018-10-10 12:33:00'],
    ['id' => 4, 'name' => 'D', 'created_at' => '2019-11-09 14:47:33'],
    ['id' => 1, 'name' => 'A', 'created_at' => '2021-12-31 02:18:53'],
    ['id' => 2, 'name' => 'C', 'created_at' => '2016-12-03 18:23:23'],
];

/*
Array
(
    [0] => Array
        (
            [id] => 1
            [name] => A
            [created_at] => 2021-12-31 02:18:53
        )

    [1] => Array
        (
            [id] => 4
            [name] => D
            [created_at] => 2019-11-09 14:47:33
        )

    [2] => Array
        (
            [id] => 3
            [name] => B
            [created_at] => 2018-10-10 12:33:00
        )

    [3] => Array
        (
            [id] => 2
            [name] => C
            [created_at] => 2016-12-03 18:23:23
        )

)
*/
print_r(multi_array_sort($array, 'created_at', SORT_DESC));


function multi_array_sort()
{
    $args = func_get_args();
    $data = array_shift($args);
    foreach ($args as $index => $field) {
        if (is_string($field)) {
            $tmp = [];
            foreach ($data as $key => $val) {
                $tmp[$key] = $val[$field];
            }
            $args[$index] = $tmp;
        }
    }
    $args[] = &$data;
    // call_user_func_array('array_multisort', $args);
    array_multisort(...$args);

    return array_pop($args);
}
