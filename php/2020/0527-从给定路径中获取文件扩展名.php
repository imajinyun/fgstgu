<?php

$path = '/Users/user/Codes/laravel/laravel.test/package.json';

/*
Array
(
    [0] => json
    [1] => json
    [2] => json
    [3] => json
    [4] => json
    [5] => json
    [6] => json
)
*/
print_r([
    ext_name1($path),
    ext_name2($path),
    ext_name3($path),
    ext_name4($path),
    ext_name5($path),
    ext_name6($path),
    ext_name7($path),
]);

function ext_name1(string $path): string
{
    return pathinfo($path, PATHINFO_EXTENSION);
}

function ext_name2(string $path): string
{
    return pathinfo($path)['extension'];
}

function ext_name3(string $path): string
{
    $array = explode('.', $path);

    // return array_pop($array)
    return end($array);
}

function ext_name4(string $path): string
{
    $array = preg_split('/[\\.]/', $path);

    return $array[count($array) - 1];
}

function ext_name5(string $path): string
{
    // return substr(strrchr($path, '.'), 1);
    return substr($path, strrpos($path, '.') + 1);
}

function ext_name6(string $path): string
{
    return preg_replace('/^.*\.([^.]+)$/D', '$1', $path);
}

function ext_name7(string $path): string
{
    return (new \SplFileInfo($path))->getExtension();
}
