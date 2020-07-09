<?php

$a = [1, 2, 3];
print_r(duplicate($a, 3));

$b = ['a', 'b'];
print_r(duplicate($b, 3));


function duplicate(array $array, int $num): array
{
    if ($num <= 0) {
        return $array;
    }

    return array_filter(explode(',', str_repeat(implode(',', $array) . ',', $num)));
}
