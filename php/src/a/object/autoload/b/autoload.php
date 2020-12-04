<?php

if (! function_exists('classAutoloader')) {
    function classAutoloader($class)
    {
        $class = strtolower($class);
        $filename = './include/'.$class.'.php';
        if (is_file($filename) && !class_exists($class)) {
            include $filename;
        }
    }
}

spl_autoload_register('classAutoloader');
