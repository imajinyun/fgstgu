<?php

$a = [1, 2, -1];
$b = [[1, 2, 3], [1, 2, -1]];

// bool(true)
var_dump(array_contains($a, $b));


function array_contains(array $needle, array $haystack): bool
{
    foreach ($haystack as $item) {
        sort($item);
        sort($needle);
        if ($item === $needle) {
            return true;
        }
    }

    return false;
}
