<?php

// 文件 A 相对于文件 B 的路径，即从文件 B 访问文件 A 的相对路径
$a = '/a/b/c/d/e/f/x.php';
$b = '/a/b/c/m/n/y.php';
$c = '/o/p/q/r/z.php';

/*
string(20) "../../../d/e/f/x.php"
string(17) ""
*/
var_dump(relative_path($a, $b), relative_path($a, $c));

/*
string(20) "../../../d/e/f/x.php"
string(17) ""
*/
var_dump(relative_path2($a, $b), relative_path2($a, $c));


function relative_path(string $a, string $b): string
{
    $ans = '';
    if (empty($a) || empty($b)) {
        return $ans;
    }

    if (strpos($a, '/') !== false) {
        $a = ltrim($a, '/');
    }
    if (strpos($b, '/') !== false) {
        $b = ltrim($b, '/');
    }
    $aArr = explode('/', $a);
    $bArr = explode('/', $b);
    if (isset($aArr[0], $bArr[0]) && $aArr[0] !== $bArr[0]) {
        return $ans;
    }
    $i = 0;
    while ($aArr[$i] === $bArr[$i]) {
        unset($aArr[$i]);
        array_unshift($aArr, '..');
        $i++;
    }
    $ans = implode('/', $aArr);

    return $ans;
}

function relative_path2(string $a, string $b): string
{
    $ans = '';
    if (empty($a) || empty($b)) {
        return $ans;
    }
    if (strpos($a, '/') !== false) {
        $a = ltrim($a, '/');
    }
    if (strpos($b, '/') !== false) {
        $b = ltrim($b, '/');
    }

    [$aArr, $bArr] = [explode('/', $a), explode('/', $b)];
    if (isset($aArr[0], $bArr[0]) && $aArr[0] !== $bArr[0]) {
        return $ans;
    }
    $same = array_intersect($aArr, $bArr);
    foreach ($aArr as $key => $value) {
        if (in_array($value, $same)) {
            unset($aArr[$key]);
            array_unshift($aArr, '..');
        }
    }
    $ans = implode('/', $aArr);

    return $ans;
}
