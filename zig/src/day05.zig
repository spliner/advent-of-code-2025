const std = @import("std");
const day = @import("day.zig");

pub fn part1(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

pub fn part2(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

test "part1" {
    const input =
        \\3-5
        \\10-14
        \\16-20
        \\12-18
        \\
        \\1
        \\5
        \\8
        \\11
        \\17
        \\32
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 3 }, result);
}

test "part2" {
    const input =
        \\3-5
        \\10-14
        \\16-20
        \\12-18
        \\
        \\1
        \\5
        \\8
        \\11
        \\17
        \\32
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 14 }, result);
}
