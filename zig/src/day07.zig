const std = @import("std");
const day = @import("day.zig");

const LineIterator = struct {
    reader: *std.Io.Reader,
    line_writer: std.Io.Writer.Allocating,

    fn init(allocator: std.mem.Allocator, reader: *std.Io.Reader) LineIterator {
        return .{
            .reader = reader,
            .line_writer = std.Io.Writer.Allocating.init(allocator),
        };
    }

    fn deinit(self: *LineIterator) void {
        self.line_writer.deinit();
    }

    fn next(self: *LineIterator) !?[]const u8 {
        self.line_writer.clearRetainingCapacity();

        const found_delimiter = blk: {
            _ = self.reader.streamDelimiter(&self.line_writer.writer, '\n') catch |err| switch (err) {
                error.EndOfStream => break :blk false,
                else => return err,
            };
            break :blk true;
        };

        if (found_delimiter) {
            self.reader.toss(1); // Skip delimiter
        }

        const line = self.line_writer.written();

        // Only return null at end of stream with no data
        if (!found_delimiter and line.len == 0) {
            return null;
        }

        return line;
    }
};

pub fn part1(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var it = LineIterator.init(allocator, reader);
    defer it.deinit();

    while (try it.next()) |l| {
        std.debug.print("{s}\n", .{l});
    }
    return day.Answer{ .int = 0 };
}

pub fn part2(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
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

    try std.testing.expectEqual(day.Answer{ .int = 0 }, result);
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

    try std.testing.expectEqual(day.Answer{ .int = 0 }, result);
}
