<?php

var_dump(find(10, 1, 2));

function find(int $n, int $x, int $y)
{
    while (true) {
        while ($n > 0) {
            if ($x === $y) {
                $found = true;
                break;
            }
            $n--;
        }
    }
}
