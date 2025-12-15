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

pub fn part1(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

pub fn part2(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

test "part1" {
    const input =
        \\..@@.@@@@.
        \\@@@.@.@.@@
        \\@@@@@.@.@@
        \\@.@@@@..@.
        \\@@.@@@@.@@
        \\.@@@@@@@.@
        \\.@.@.@.@@@
        \\@.@@@.@@@@
        \\.@@@@@@@@.
        \\@.@.@@@.@.
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 13 }, result);
}

test "part2" {
    const input =
        \\..@@.@@@@.
        \\@@@.@.@.@@
        \\@@@@@.@.@@
        \\@.@@@@..@.
        \\@@.@@@@.@@
        \\.@@@@@@@.@
        \\.@.@.@.@@@
        \\@.@@@.@@@@
        \\.@@@@@@@@.
        \\@.@.@@@.@.
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 8310 }, result);
}
