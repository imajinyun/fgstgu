<?php

/**
 * 在实例化对象时：
 *
 * - static：会根据运行时调用的类来决定实例化对象；
 * - self：会根据所在位置的类来决定实例化对象；
 *
 * 注意：只想实例化子类，并且不希望后续在对子类的使用中由于父类的变化对子类产生影响时，后期静态绑定就能发挥它的作用了。
 */

class A
{
    public static function who()
    {
        echo __CLASS__, PHP_EOL;
    }

    public static function test()
    {
        self::who();
    }
}

class B extends A
{
    public static function who()
    {
        echo __CLASS__, PHP_EOL;
    }
}

// A
B::test();
