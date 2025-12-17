const std = @import("std");
const day = @import("day.zig");

const Range = struct {
    start: usize,
    end: usize,

    fn inRange(self: Range, item: usize) bool {
        return self.start <= item and item <= self.end;
    }
};

pub fn part1(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var ranges = try parseRanges(allocator, reader);
    defer ranges.deinit(allocator);

    var ids = try parseIds(allocator, reader);
    defer ids.deinit(allocator);
    var count: usize = 0;
    outer: for (ids.items) |i| {
        for (ranges.items) |r| {
            if (r.inRange(i)) {
                count += 1;
                continue :outer;
            }
        }
    }
    return day.Answer{ .int = @intCast(count) };
}

pub fn part2(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

fn parseRanges(allocator: std.mem.Allocator, reader: *std.Io.Reader) !std.ArrayList(Range) {
    var ranges: std.ArrayList(Range) = .empty;
    while (try reader.takeDelimiter('\n')) |line| {
        const trimmed = std.mem.trim(u8, line, " \r\t");
        if (std.mem.eql(u8, trimmed, "")) {
            return ranges;
        }

        var split = std.mem.splitScalar(u8, trimmed, '-');
        const first = split.next() orelse unreachable;
        const second = split.next() orelse unreachable;

        const range = Range{
            .start = try std.fmt.parseUnsigned(usize, first, 10),
            .end = try std.fmt.parseUnsigned(usize, second, 10),
        };
        try ranges.append(allocator, range);
    }

    return ranges;
}

fn parseIds(allocator: std.mem.Allocator, reader: *std.io.Reader) !std.ArrayList(usize) {
    var ids: std.ArrayList(usize) = .empty;
    while (try reader.takeDelimiter('\n')) |line| {
        const trimmed = std.mem.trim(u8, line, " \r\t");
        const id = try std.fmt.parseUnsigned(usize, trimmed, 10);
        try ids.append(allocator, id);
    }
    return ids;
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
