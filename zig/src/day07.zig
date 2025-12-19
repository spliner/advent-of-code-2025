const std = @import("std");
const day = @import("day.zig");

pub fn part1(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var it = day.LineIterator.init(allocator, reader);
    defer it.deinit();

    const first_line = try it.next() orelse unreachable;

    var beams = try std.ArrayList(usize).initCapacity(allocator, first_line.len);
    defer beams.deinit(allocator);

    for (first_line) |c| {
        const val: usize = if (c == 'S') 1 else 0;
        try beams.append(allocator, val);
    }

    var splits: usize = 0;
    while (try it.next()) |line| {
        for (beams.items, 0..) |val, i| {
            if (val != 0 and line[i] == '^') {
                splits += 1;
                beams.items[i - 1] += val;
                beams.items[i + 1] += val;
                beams.items[i] = 0;
            }
        }
    }

    return day.Answer{ .int = @intCast(splits) };
}

pub fn part2(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var it = day.LineIterator.init(allocator, reader);
    defer it.deinit();

    const first_line = try it.next() orelse unreachable;

    var beams = try std.ArrayList(usize).initCapacity(allocator, first_line.len);
    defer beams.deinit(allocator);

    for (first_line) |c| {
        const val: usize = if (c == 'S') 1 else 0;
        try beams.append(allocator, val);
    }

    var timelines: usize = 1;
    while (try it.next()) |line| {
        for (beams.items, 0..) |val, i| {
            if (val != 0 and line[i] == '^') {
                timelines += val;
                beams.items[i - 1] += val;
                beams.items[i + 1] += val;
                beams.items[i] = 0;
            }
        }
    }

    return day.Answer{ .int = @intCast(timelines) };
}

test "part1" {
    const input =
        \\.......S.......
        \\...............
        \\.......^.......
        \\...............
        \\......^.^......
        \\...............
        \\.....^.^.^.....
        \\...............
        \\....^.^...^....
        \\...............
        \\...^.^...^.^...
        \\...............
        \\..^...^.....^..
        \\...............
        \\.^.^.^.^.^...^.
        \\...............
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 21 }, result);
}

test "part2" {
    const input =
        \\.......S.......
        \\...............
        \\.......^.......
        \\...............
        \\......^.^......
        \\...............
        \\.....^.^.^.....
        \\...............
        \\....^.^...^....
        \\...............
        \\...^.^...^.^...
        \\...............
        \\..^...^.....^..
        \\...............
        \\.^.^.^.^.^...^.
        \\...............
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 40 }, result);
}
