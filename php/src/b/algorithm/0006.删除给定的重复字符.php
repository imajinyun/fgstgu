<?php

$data = [
    ['baibaidubaidudu', 'baidu', ''],
    ['818816168166', '816', ''],
    ['81818166', '816', '81']
];
foreach ($data as $item) {
    [$s, $t, $expect] = $item;
    $actual = removeRepeatString($s, $t);
    if (removeRepeatString($s, $t) !== $expect) {
        printf("\nActual: %s\nExpect: %s\n", $actual, $expect);
        break;
    }
}

function removeRepeatString(string $s, string $t): string
{
    $ans = str_replace($t, '', $s);
    if ($ans === $s) {
        return $ans;
    }
    return removeRepeatString($ans, $t);
}
