const std = @import("std");
const day = @import("day.zig");

fn countDigits(n: u512) u512 {
    if (n == 0) {
        return 1;
    }

    var num_digits: u512 = 0;
    var current = n;
    while (current > 0) : (current /= 10) {
        num_digits += 1;
    }

    return num_digits;
}

// I know I should not be using integers for this, but I have to see how Zig handles it.
// Zig allows us to have arbitrary bit-width integers, like u9 :o
pub fn part1(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var sum: u512 = 0;
    while (try reader.takeDelimiter('\n')) |line| {
        const trimmed = std.mem.trim(u8, line, " \r\n\t");
        const bank = try std.fmt.parseUnsigned(u512, trimmed, 10);
        const num_digits = countDigits(bank);
        var i: u512 = 0;
        var max_joltage: u512 = 0;
        while (i < num_digits) : (i += 1) {
            const left_divisor = std.math.pow(u512, 10, num_digits - i - 1);
            const left = (bank / left_divisor) % 10;
            var j = i + 1;
            while (j < num_digits) : (j += 1) {
                const right_divisor = std.math.pow(u512, 10, num_digits - j - 1);
                const right = (bank / right_divisor) % 10;
                const joltage = left * 10 + right;
                if (joltage > max_joltage) {
                    max_joltage = joltage;
                }
            }
        }
        sum += max_joltage;
    }

    return day.Answer{ .int = @intCast(sum) };
}

const InvalidNumberOfDigitsError = error{};

pub fn part2(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var sum: usize = 0;
    while (try reader.takeDelimiter('\n')) |line| {
        const trimmed = std.mem.trim(u8, line, " \r\n\t");
        const max_joltage = maxJoltage(trimmed, 12);
        sum += max_joltage;
    }

    return day.Answer{ .int = @intCast(sum) };
}

fn maxJoltage(input: []const u8, n: usize) usize {
    var result: usize = 0;

    var index: usize = 0;
    var i: usize = 0;
    while (i < n) : (i += 1) {
        var j = index;
        while (j <= input.len - (n - i)) : (j += 1) {
            if (input[index] < input[j]) {
                index = j;
            }
        }

        const digit = input[index] - '0';
        result += digit * std.math.pow(usize, 10, n - i - 1);

        index += 1;
    }

    return result;
}

test "part1" {
    const input =
        \\987654321111111
        \\811111111111119
        \\234234234234278
        \\818181911112111
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 357 }, result);
}

test "part2" {
    const input =
        \\987654321111111
        \\811111111111119
        \\234234234234278
        \\818181911112111
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 3121910778619 }, result);
}
