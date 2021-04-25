<?php

$array = [
    'Clothing and Accessories🧥 👚 👕 👖 👔',
    'Emojis 👶🏻 👦🏻 👧🏻 👨🏻 👩🏻 👱🏻‍♀️',
    '🤳🏻 💅🏻 👂🏻 👃🏻Cream White Emojis',
    'Moderate Brown Emojis 👶🏽 👦🏽 👧🏽 👨🏽 👩🏽 👱🏽‍♀️',
    '👩🏿🎊🏞🎉💂🏽‍♂️Hello World',
];
$pattern = '/^(\u00a9|\u00ae|[\u2000-\u3300]|\ud83c[\ud000-\udfff]|\ud83d[\ud000-\udfff]|\ud83e[\ud000-\udfff])+(\/\w+)*/';

foreach ($array as $item) {
    preg_match($pattern, $item, $matches);
    echo ($matches[0] ?? 'no match') . PHP_EOL;
}
