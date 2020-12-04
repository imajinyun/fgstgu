<?php

class Foo
{
    private string $name = 'FooClas';
    protected int $version = 20200721;
    public float $score = 78.12;
}

/**
 * 使用反射
 */
$class = new \ReflectionClass(Foo::class);
$properties = $class->getProperties();
foreach ($properties as $key => $property) {
    $property->setAccessible(true);
}

$foo = new Foo();
$keys = [];
foreach ($properties as $key => $property) {
    // 只能通过 ReflectionProperty 实例的 getValue 方法访问，而不能 $foo->name 直接访问
    $keys[$property->getName()] = $property->getValue($foo);
}

/*
array(3) {
  ["name"]=>
  string(7) "FooClas"
  ["version"]=>
  int(20200721)
  ["score"]=>
  float(78.12)
}
*/
var_dump($keys);


/**
 * 使用数组
 */
$foo = new Foo();
$array = (array)$foo;
$keys = [
    'name' => "\0" . Foo::class . "\0" . 'name',
    'version' => "\0*\0" . 'version',
    'score' => 'score',
];

/*
string(7) "FooClas"
int(20200721)
float(78.12)
*/
var_dump($array[$keys['name']], $array[$keys['version']], $array[$keys['score']]);


/**
 * 使用闭包
 */
$foo = new Foo();

/*
string(7) "FooClas"
int(20200721)
float(78.12)
NULL
*/
var_dump(
    inspect($foo, 'name'),
    inspect($foo, 'version'),
    inspect($foo, 'score'),
    inspect($foo, 'description'),
);

function inspect($object, string $property)
{
    return (function () use ($property) {
        if (isset($this->{$property})) {
            return $this->{$property};
        }
        return null;
    })->call($object);
}
