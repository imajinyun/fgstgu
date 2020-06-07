package org.b;

public class Main {

    public static void main(String[] args) {
        String name1 = "XPS 15(9500) 15.6英寸旗舰创作本";
        String brand1 = "Dell";
        int count1 = 10;
        double price1 = 10989.00;
        int num1 = 1;

        System.out.println("购买电脑：");
        System.out.println("电脑名称：" + name1);
        System.out.println("电脑品牌：" + brand1);
        System.out.println("电脑数量：" + num1);
        System.out.println("电脑价格：" + price1);
        System.out.println("剩余库存：" + (count1 - num1));

        System.out.println();

        String name2 = "HUAWEI MateBook 13 2020款";
        String brand2 = "华为";
        int count2 = 15;
        double price2 = 5900.00;
        int num2 = 2;

        System.out.println("购买信息：");
        System.out.println("电脑名称：" + name2);
        System.out.println("电脑品牌" + brand2);
        System.out.println("电脑数量：" + num2);
        System.out.println("电脑价格：" + price2);
        System.out.println("剩余库存：" + (count2 - num2));
    }

}
