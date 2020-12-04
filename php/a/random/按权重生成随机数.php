<?php

/**
 * @see https://www.rubyguides.com/2016/05/weighted-random-numbers/
 * @see https://softwareengineering.stackexchange.com/questions/150616/get-weighted-random-item
 * @see https://gist.github.com/rickydunlop/3720869
 */

$weights = [
    'a' => 800,
    'b' => 500,
    'c' => 400,
    'd' => 200,
    'e' => 100,
    'f' => 50,
];
var_dump(random_by_weight($weights, 1));
var_dump(random_by_weight($weights, 2));
var_dump(random_by_weight($weights, 3));

function random_by_weight(array $array, int $num): array
{
    $ans = [];
    if (count($array) <= 0 || $num <= 0) {
        return $ans;
    }

    $sum = array_sum($array);
    do {
        $random = random_int(1, $sum);
        foreach ($array as $i => $value) {
            if ($random < $value && ! in_array((int)$value, $ans, true)) {
                $ans[$i] = $value;
                break;
            }
            $random -= $value;
        }
    } while (count($ans) !== $num);

    return $ans;
}
