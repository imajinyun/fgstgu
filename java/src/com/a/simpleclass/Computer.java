package com.a.simpleclass;

public class Computer {

    String name;
    String brand;
    int count;
    double price;

    @Override
    public String toString() {
        return "Computer{" +
                "name='" + name + '\'' +
                ", brand='" + brand + '\'' +
                ", count=" + count +
                ", price=" + price +
                '}';
    }
}
