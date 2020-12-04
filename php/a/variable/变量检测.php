<?php

$array = ['', ' ', false, true, [], null, '0', 0, 0.0, "\0"];

/*
'':
bool(true)
bool(true)
bool(false)

' ':
bool(true)
bool(false)
bool(false)

false:
bool(true)
bool(true)
bool(false)

true:
bool(true)
bool(false)
bool(false)

[]:
bool(true)
bool(true)
bool(false)

null:
bool(false)
bool(true)
bool(true)

'0':
bool(true)
bool(true)
bool(false)

0:
bool(true)
bool(true)
bool(false)

0.0:
bool(true)
bool(true)
bool(false)

"\0"
bool(true)
bool(false)
bool(false)
*/
foreach ($array as $value) {
    var_dump(isset($value), empty($value), null === $value);
    echo PHP_EOL;
}
