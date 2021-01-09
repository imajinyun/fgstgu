package com.a.shopping;

import com.a.shopping.customer.Customer;
import com.a.shopping.supermarket.Merchandise;
import com.a.shopping.supermarket.Supermarket;

import java.util.Scanner;

public class Main {

    private static boolean isOpen = true;

    public static void main(String[] args) {
        Supermarket supermarket = Main.initializeSupermarket();
        Scanner scanner = new Scanner(System.in);
        while (isOpen) {
            Main.dumpSupermarket(supermarket);
            Customer customer = Main.initializeCustomer(supermarket);

            int sum = 0;
            while (true) {
                System.out.println("本店提供 " + supermarket.merchandises.length + " 种商品，请输入您要购买的商品编号：");
                int merchandiseId = scanner.nextInt();

                if (merchandiseId < 0) {
                    System.out.println("你本次购买了 " + sum + " 元的商品,欢迎下次光临");
                    break;
                }

                if (merchandiseId >= supermarket.merchandises.length) {
                    System.out.println("你选择的商品本店暂时没有，欢迎继续选购其它商品");
                    continue;
                }

                Merchandise buyMerchandise = supermarket.merchandises[merchandiseId];
                System.out.println("[" + buyMerchandise.name + "]单价 " + buyMerchandise.sellPrice + "，请问要购买几个？");

                int numToBuy = scanner.nextInt();
                if (numToBuy <= 0) {
                    System.out.println("不买就多看看，欢迎继续选购");
                    continue;
                }

                if (numToBuy > buyMerchandise.count) {
                    System.out.println();
                    continue;
                }

                if (numToBuy * buyMerchandise.sellPrice + sum > customer.money) {
                    System.out.println("你的余额不够，请理智消费");
                    continue;
                }

                sum += numToBuy * buyMerchandise.sellPrice;
                buyMerchandise.count -= numToBuy;
                supermarket.statistics[merchandiseId] += numToBuy;
            }

            customer.money -= sum;
            supermarket.totalIncome += sum;
            System.out.println("[" + customer.name + "] 总共消费 " + sum + " 元，欢迎下次光临");
            System.out.println("[" + customer.name + "] 身上带的现金剩余：" + customer.money + " 元");

            System.out.println("♻️ 您是否要继续购买（请输入 true or false）");
            isOpen = scanner.nextBoolean();

            if (customer.isDriverCar) {
                supermarket.totalParkingSpace++;
            }
        }
        Main.dumpStatistics(supermarket);
    }

    private static Supermarket initializeSupermarket() {
        Supermarket supermarket = new Supermarket();
        supermarket.name = "阳光小超市";
        supermarket.address = "坂雪岗大道万科里购物中心";
        supermarket.totalParkingSpace = 100;
        supermarket.merchandises = new Merchandise[300];
        supermarket.statistics = new int[supermarket.merchandises.length];

        // Initializes all items of the specified category.
        for (int i = 0; i < supermarket.merchandises.length; i++) {
            Merchandise m = new Merchandise();
            m.id = "#" + i;
            m.name = "商品" + i;
            m.count = 100;
            m.buyPrice = Math.random() * 100;
            m.sellPrice = m.buyPrice * (1 + Math.random());
            supermarket.merchandises[i] = m;
        }
        System.out.println("Supermarket in business...");
        return supermarket;
    }

    private static Customer initializeCustomer(Supermarket supermarket) {
        Customer customer = new Customer();
        customer.name = "顾客编号" + (int) (Math.random() * 100000);
        customer.money = (1 + Math.random()) * 1000;
        System.out.println("😄 [" + customer.name + "]，今天身上带了 " + customer.money + " 元人民币，我去逛逛超市");

        if (customer.isDriverCar) {
            if (supermarket.totalParkingSpace > 0) {
                System.out.println("欢迎 [" + customer.name + "]" + " 驾车而来，本店已经为您安排了免费的停车位");
            } else {
                System.out.println("不好意思,本店车位已满，欢迎下次i光临。");
            }
        } else {
            System.out.println("欢迎 [" + customer.name + "] 光临本店。");
        }
        return customer;
    }

    private static void dumpSupermarket(Supermarket supermarket) {
        String info = "";
        info += "本店名称：" + supermarket.name + "\n";
        info += "本店地址：" + supermarket.address + "\n";
        info += "本店总停车位：" + supermarket.totalParkingSpace + " 个\n";
        info += "本店今日营业额：" + supermarket.totalIncome + " 元\n";
        info += "本店总共商品种类：" + supermarket.merchandises.length + " 种";
        System.out.println(info);
    }

    private static void dumpStatistics(Supermarket supermarket) {
        System.out.println("😭 超市打烊了...");
        System.out.println("今日销售额为 " + supermarket.totalIncome + "，营业统计如下：");
        for (int i = 0; i < supermarket.merchandises.length; i++) {
            int total = supermarket.statistics[i];
            if (total > 0) {
                Merchandise m = supermarket.merchandises[i];
                double profit = total * (m.sellPrice - m.buyPrice);
                double income = total * m.sellPrice;
                String str = "";
                str += "[" + m.name + "] 进货价为 " + m.buyPrice + " 元，销售价为 " + m.sellPrice + "元";
                str += "售出 " + total + "，利润 " + profit + " 元，毛利润为 " + income + " 元";
                System.out.println(str);
            }
        }
    }
}
