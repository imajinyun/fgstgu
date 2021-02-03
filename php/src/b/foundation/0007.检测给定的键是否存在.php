<?php

$search = ['a', 'b', 'c'];
$data = ['a' => 1, 'b' => 2, 'c' => 3, 'd' => 4];

foreach ($data as $key => $val) {
    if (! array_key_exists($key, array_flip($search))) {
        unset($data[$key]);
    }
}

print_r($data);
