<?php

$array = [
    'Clothing and Accessories🧥 👚 👕 👖 👔',
    'Emojis 👶🏻 👦🏻 👧🏻 👨🏻 👩🏻 👱🏻‍♀️',
    '🤳🏻 💅🏻 👂🏻 👃🏻Cream White Emojis',
    'Moderate Brown Emojis 👶🏽 👦🏽 👧🏽 👨🏽 👩🏽 👱🏽‍♀️',
    '👩🏿🎊🏞🎉💂🏽‍♂️👩🏿🖐🏿Hello World #1',
    'Hello 🇸🇮 World #2',
    'VSCode for Mac',
    '👍👊🏿👈🏼🤞🏻🤛',
    '🤝 This is a test #1💄🧠👩🏿‍🦰',
    '🧶👠This is a test #2💄🧠👩🏿‍🦰 sort 🤝🤟🏽',
];
$pattern1 = '/^((\u00a9|\u00ae|[\u2000-\u3300]|\ud83c[\ud000-\udfff]|\ud83d[\ud000-\udfff]|\ud83e[\ud000-\udfff])|(\w\s\d))+[\w\s\d].*$/i';
$filter = [];
foreach ($array as $item) {
    preg_match($pattern1, $item, $matches);
    if (empty($matches)) {
        var_dump($item);
        $filter[] = $item;
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
