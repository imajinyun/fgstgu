<?php

/**
 * [深入解析 composer 的自动加载原理](https://segmentfault.com/a/1190000014948542?utm_source=tag-newest)
 */

require './autoload.php';

var_dump((new A())->getName());
var_dump((new B())->getName());
