<?php

$array = [
    'Clothing and Accessories🧥 👚 👕 👖 👔',
    'Emojis 👶🏻 👦🏻 👧🏻 👨🏻 👩🏻 👱🏻‍♀️',
    '🤳🏻 💅🏻 👂🏻 👃🏻Cream White Emojis',
    'Moderate Brown Emojis 👶🏽 👦🏽 👧🏽 👨🏽 👩🏽 👱🏽‍♀️',
    '👩🏿🎊🏞🎉💂🏽‍♂️👩🏿🖐🏿Hello World #1',
    'Hello 🇸🇮 World #2',
];
$pattern1 = '/^(\u00a9|\u00ae|[\u2000-\u3300]|\ud83c[\ud000-\udfff]|\ud83d[\ud000-\udfff]|\ud83e[\ud000-\udfff])+\w+.*/i';
$filter = [];
foreach ($array as $item) {
    preg_match($pattern1, $item, $matches);
    if (empty($matches)) {
        $filter[] = $item;
        var_dump($item);
    }
}

if (! $filter) {
    die('no match any item');
}

$result = [];
$pattern2 = '/^(\u00a9|\u00ae|[\u2000-\u3300]|\ud83c[\ud000-\udfff]|\ud83d[\ud000-\udfff]|\ud83e[\ud000-\udfff])|(\w+.*)/i';
foreach ($filter as $item) {
    preg_match($pattern2, $item, $matches);
    if (isset($matches[0]) && ! empty($matches[0])) {
        $result[] = $matches[0] ?? '';
    }
}
print_r($result);
