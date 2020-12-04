<?php

/**
 * 浅拷贝：创建一个新对象，这个对象有着原始对象属性值的一份精确拷贝。如果属性是基本类型，拷贝的就是基本类型的值，如果属性是引用类型，拷贝的就是内存地址 ，所以如果其中一个对象改变了这个地址，就会影响到另一个对象。
 * 深拷贝：将一个对象从内存中完整的拷贝一份出来,从堆内存中开辟一个新的区域存放新对象,且修改新对象不会影响原对象。
 */

class Student
{
    private string $name;

    public function getName(): string
    {
        return $this->name;
    }

    public function setName(string $name = ''): self
    {
        $this->name = $name;

        return $this;
    }
}

class Administrator
{
    private ?Student $classLeader;

    public function __clone()
    {
        foreach ($this as $key => $val) {
            if (is_object($val) || is_array($val)) {
                $this->{$key} = unserialize(serialize($val));
            }
        }
    }

    /**
     * @return Student
     */
    public function getClassLeader(): Student
    {
        return $this->classLeader;
    }

    public function setClassLeader(Student $student): self
    {
        $this->classLeader = $student;

        return $this;
    }
}

$student = new Student();
$student->setName('Jack');

$admin1 = new Administrator();
$admin1->setClassLeader($student);

/** @var Administrator */
$admin2 = unserialize(serialize($admin1));

// string(4) "Jack"
var_dump($admin2->getClassLeader()->getName());

$admin2->getClassLeader()->setName('Tom');

/*
string(4) "Jack"
string(3) "Tom"
*/
var_dump($admin1->getClassLeader()->getName(), $admin2->getClassLeader()->getName());

$admin2 = clone $admin1;
$admin2->getClassLeader()->setName('Lucy');

/*
string(4) "Jack"
string(4) "Lucy"
*/
var_dump($admin1->getClassLeader()->getName(), $admin2->getClassLeader()->getName());
