<?php

$search = ['a', 'b', 'c'];
$array = ['a' => 1, 'b' => 2, 'c' => 43, 'd' => 4];

/*
Array
(
    [a] => 1
    [b] => 2
    [c] => 43
)
*/
print_r(filter_key_array($array, $search));


function filter_key_array(array $array, array $keys = []): array
{
    foreach ($array as $key => $val) {
        if (! array_key_exists($key, array_flip($keys))) {
            unset($array[$key]);
        }
    }

    return $array;
}
