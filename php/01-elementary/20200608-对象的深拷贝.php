<?php

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
