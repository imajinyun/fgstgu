<?php

// 使用 call_user_func 函数的方式
call_user_func(static function () {
    echo 'Hello World', PHP_EOL;
}); // Hello World

// 使用闭包的方式
(static function () {
    echo 'Hello World', PHP_EOL;
})(); // Hello World

// 使用闭包传参的方式
(function () {
    $args = func_get_args();
    echo $args[0], PHP_EOL;
})('Hello PHP'); // Hello PHP

// 使用匿名类的方式
var_dump((new class() {
    public static function run()
    {
        return 'Hello World';
    }
})::run()); // string(11) "Hello World"
