<?php

class RangeOverlap
{
    /**
     * 判断原始范围是否包含在目标范围中。
     *
     * @param array $source 原始范围
     * @param array $target 目标范围
     * @param bool $overlap 是否包含端点，默认不包含端点
     *
     * @return bool
     */
    public static function overlapWith(array $source, array $target, $overlap = false): bool
    {
        sort($source);
        sort($target);

        if ($overlap) {
            return self::includeOverlap($source, $target);
        }

        return self::exclusiveOverlap($source, $target);
    }

    private static function includeOverlap(array $source, array $target): bool
    {
        [$begin, $end] = $source;
        if ($target[1] < $begin) {
            return false;
        }

        if ($target[0] > $end) {
            return false;
        }

        return true;
    }

    private static function exclusiveOverlap(array $source, array $target): bool
    {
        [$begin, $end] = $source;
        if ($target[1] <= $begin) {
            return false;
        }

        if ($target[0] >= $end) {
            return false;
        }

        return true;
    }
}

$source = [1, 4];
$target = [4, 6];

// bool(false)
var_dump(RangeOverlap::overlapWith($source, $target));

// bool(true)
var_dump(RangeOverlap::overlapWith($source, $target, true));
