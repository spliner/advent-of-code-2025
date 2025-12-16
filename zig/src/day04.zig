const std = @import("std");
const day = @import("day.zig");

const Point = struct {
    x: isize,
    y: isize,

    fn adjacent(self: *Point) [8]Point {
        const points: [8]Point = .{
            .{ .x = self.x - 1, .y = self.y - 1 }, // Toself.left
            .{ .x = self.x, .y = self.y - 1 }, // Top
            .{ .x = self.x + 1, .y = self.y - 1 }, // Toself.right
            .{ .x = self.x - 1, .y = self.y }, // Left
            .{ .x = self.x + 1, .y = self.y }, // Right
            .{ .x = self.x + 1, .y = self.y + 1 }, // Bottom right
            .{ .x = self.x, .y = self.y + 1 }, // Bottom
            .{ .x = self.x - 1, .y = self.y + 1 }, // Bottom left
        };
        return points;
    }
};

pub fn part1(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var points = std.AutoHashMap(Point, void).init(allocator);
    defer points.deinit();

    var y: isize = 0;
    while (try reader.takeDelimiter('\n')) |line| {
        const trimmed = std.mem.trim(u8, line, " \r\t");
        for (trimmed, 0..) |c, x| {
            if (c == '@') {
                const p = Point{ .x = @as(isize, @intCast(x)), .y = y };
                try points.put(p, {});
            }
        }
        y += 1;
    }

    var total: isize = 0;
    var it = points.keyIterator();
    while (it.next()) |p| {
        const adj = p.adjacent();
        var count: usize = 0;
        for (adj) |a| {
            if (points.contains(a)) {
                count += 1;
            }
        }
        if (count < 4) {
            total += 1;
        }
    }

    return day.Answer{ .int = total };
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
