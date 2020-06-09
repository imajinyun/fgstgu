<?php

ob_start(function (string $str) {
    static $index = 1;

    $result = '';
    if ($str) {
        $result = $index++ . ': ' . $str . "\n";
    }

    return $result;
}, 10);

ob_start('handle', 3);

echo 'fo';

sleep(1);

echo 'o';

sleep(1);

echo 'bar';

sleep(1);

// 1: FooBarBazz
echo 'bazz';

sleep(1);

echo 'hello';

sleep(1);

// 2: HelloWorld
echo 'world';


function handle(string $str): string
{
    return ucfirst($str);
}
