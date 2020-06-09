<?php

$one = ['orange', 'banana', 'apple'];
$two = ['fruit' => ['orange', 'banana', 'apple']];
$thr = ['' => ['fruit' => ['orange', 'banana', 'apple']]];
$fou = [0 => [1 => ['fruit' => ['orange', 'banana', 'apple']]]];

var_dump(get_array_dimension($one));
var_dump(get_array_dimension($two));
var_dump(get_array_dimension($thr));
var_dump(get_array_dimension($fou));


function get_array_dimension($array = []): int
{
    if (! is_array($array)) {
        return 0;
    }

    $dimension = 0;
    foreach ($array as $item) {
        $value = get_array_dimension($item);
        if ($value > $dimension) {
            $dimension = $value;
        }
    }

    return $dimension + 1;
}
