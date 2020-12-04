<?php

/*
int(1)
int(1)
int(2)
int(6)
int(24)
int(120)
*/
var_dump(
    factorial(0),
    factorial(1),
    factorial(2),
    factorial(3),
    factorial(4),
    factorial(5),
);


function factorial(int $x): int
{
    if ($x == 0) {
        return 1;
    }

    return $x * factorial($x - 1);
}
