<?php

function __autoload($class)
{
    $filename = './'.$class.'.php';
    include $filename;
}
