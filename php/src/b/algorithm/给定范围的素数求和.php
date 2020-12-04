<?php

// 参考：https://www.wolframalpha.com/input/?i=sum+of+primes+under+10%5E2

// int(1060)
var_dump(prime_sum(0, 100));

function prime_sum(int $start, int $end)
{
    if ($start <= 1) {
        $start = 2;
    }
    $sum = 0;
    if ($start <= $end) {
        for ($i = $start; $i <= $end; $i++) {
            if (is_prime($i)) {
                $sum += $i;
            }
        }
    }

    return $sum;
}

function is_prime(int $num): bool
{
    if ($num < 0) {
        return false;
    }
    for ($i = 2; $i < $num; $i++) {
        if ($num % $i === 0) {
            return false;
        }
    }

    return true;
}
