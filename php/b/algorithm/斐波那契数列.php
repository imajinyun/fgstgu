<?php
/*
int(0)
int(1)
int(1)
int(2)
int(3)
int(5)
int(8)
int(13)
*/
var_dump(
    fibonacci(0),
    fibonacci(1),
    fibonacci(2),
    fibonacci(3),
    fibonacci(4),
    fibonacci(5),
    fibonacci(6),
    fibonacci(7),
);

function fibonacci(int $x)
{
    if ($x <= 0) {
        return 0;
    }

    if ($x == 1 || $x == 2) {
        return 1;
    }

    return fibonacci($x - 1) + fibonacci($x - 2);
}
