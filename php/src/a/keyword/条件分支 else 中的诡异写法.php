<?php

// else 后面完全可以跟 for/while 等循环
/*
string(11) "Hello World"
bool(true)
bool(true)
*/
$a = true;
if ($a) {
    if ('1e3' == 1000) {
        var_dump('Hello World');
    }
    var_dump($a);
} else label: {
    var_dump($a);
}
