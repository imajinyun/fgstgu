<?php

$array = [
    '中国人民大学新校区',
    '中国人民大学（新  校区）',
    '中国人民大  学【新校区】 ',
    'Apple Inc. Limited ',
    ' Apple    Inc. Limited ',
    "Apple \n\tInc.\t Limited\n",
    "\n\n\tApple Inc. Limited",
];

$result = toNormalString($array);
print_r($result);

function toNormalString(array $array): array
{
    foreach ($array as $key => $item) {
        $item = trim($item);
        $item = preg_replace('/\s+/', ' ', $item);
        $array[$key] = $item;
    }

    return $array;
}
