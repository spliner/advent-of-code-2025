const std = @import("std");
const day = @import("day.zig");

const Range = struct {
    start: usize,
    end: usize,
};

const RangeIterator = struct {
    iterator: std.mem.TokenIterator(u8, .scalar),

    fn init(line: []const u8) RangeIterator {
        return .{
            .iterator = std.mem.tokenizeScalar(u8, line, ','),
        };
    }

    fn next(self: *RangeIterator) ?Range {
        if (self.iterator.next()) |range| {
            const trimmed_range = std.mem.trim(u8, range, " \n\r\t");
            var it = std.mem.tokenizeScalar(u8, trimmed_range, '-');
            const start = std.fmt.parseUnsigned(usize, it.next().?, 10) catch unreachable;
            const end = std.fmt.parseUnsigned(usize, it.next().?, 10) catch unreachable;
            return Range{
                .start = start,
                .end = end,
            };
        }

        return null;
    }
};

fn countDigits(n: usize) usize {
    if (n == 0) {
        return 1;
    }

    var num_digits: usize = 0;
    var current = n;
    while (current > 0) : (current /= 10) {
        num_digits += 1;
    }
    return num_digits;
}

fn isRepeatingPattern(n: usize) bool {
    const num_digits = countDigits(n);
    if (num_digits == 1) {
        return false;
    }

    var pattern_len: usize = 1;
    while (pattern_len < num_digits) : (pattern_len += 1) {
        if (num_digits % pattern_len != 0) {
            continue;
        }

        const divisor = std.math.pow(usize, 10, num_digits - pattern_len);
        const pattern = n / divisor;

        const repetitions = num_digits / pattern_len;
        var reconstructed: usize = 0;
        const multiplier = std.math.pow(usize, 10, pattern_len);
        var rep_idx: usize = 0;
        while (rep_idx < repetitions) : (rep_idx += 1) {
            reconstructed = reconstructed * multiplier + pattern;
        }

        if (reconstructed == n) {
            return true;
        }
    }

    return false;
}

pub fn part1(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var sum: usize = 0;
    const line = try reader.takeDelimiterInclusive('\n');
    var range_iterator = RangeIterator.init(line);
    while (range_iterator.next()) |range| {
        var n = range.start;
        while (n <= range.end) : (n += 1) {
            const num_digits = countDigits(n);
            if (num_digits % 2 != 0) {
                continue;
            }

            const half = num_digits / 2;
            const divisor = std.math.pow(usize, 10, half);
            const left = n / divisor;
            const right = n % divisor;
            if (left == right) {
                sum += n;
            }
        }
    }

    return day.Answer{ .int = @intCast(sum) };
}

pub fn part2(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var sum: usize = 0;
    const line = try reader.takeDelimiterInclusive('\n');
    var range_iterator = RangeIterator.init(line);
    while (range_iterator.next()) |range| {
        var n = range.start;
        while (n <= range.end) : (n += 1) {
            if (isRepeatingPattern(n)) {
                sum += n;
            }
        }
    }

    return day.Answer{ .int = @intCast(sum) };
}

test "part1" {
    const input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-56565,824824821-824824827,2121212118-2121212124\n";
    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 1227775554 }, result);
}

test "part2" {
    const input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124\n";
    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 4174379265 }, result);
}

test "countDigits" {
    const cases = [_]struct {
        n: usize,
        expected: usize,
    }{
        .{ .n = 0, .expected = 1 },
        .{ .n = 1, .expected = 1 },
        .{ .n = 10, .expected = 2 },
        .{ .n = 9999, .expected = 4 },
    };

    for (cases) |case| {
        const digits = countDigits(case.n);
        try std.testing.expectEqual(case.expected, digits);
    }
}
