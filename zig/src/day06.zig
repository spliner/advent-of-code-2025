const std = @import("std");
const day = @import("day.zig");

pub fn part1(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var numbers: std.ArrayList(std.ArrayList(isize)) = .empty;
    defer numbers.deinit(allocator);
    defer for (numbers.items) |*n| {
        n.deinit(allocator);
    };

    var operators: std.ArrayList(u8) = .empty;
    defer operators.deinit(allocator);

    var line_writer = std.io.Writer.Allocating.init(allocator);
    defer line_writer.deinit();

    while (true) {
        _ = reader.streamDelimiter(&line_writer.writer, '\n') catch |err| blk: switch (err) {
            error.EndOfStream => {
                break :blk;
            },
            else => {
                return err;
            },
        };

        const line = line_writer.written();

        var token_iterator = std.mem.tokenizeScalar(u8, line, ' ');
        if (line[0] == '*' or line[0] == '+') {
            while (token_iterator.next()) |op| {
                try operators.append(allocator, op[0]);
            }

            break;
        }

        var row: std.ArrayList(isize) = .empty;
        while (token_iterator.next()) |str| {
            const n = try std.fmt.parseInt(isize, str, 10);
            try row.append(allocator, n);
        }

        try numbers.append(allocator, row);

        line_writer.clearRetainingCapacity();

        // Skip delimiter
        reader.toss(1);
    }

    const cols = numbers.items[0].items.len;
    const rows = numbers.items.len;

    var sum: isize = 0;
    var c: usize = 0;
    while (c < cols) {
        const op = operators.items[c];
        var r: usize = 0;
        var result: isize = if (op == '*') 1 else 0;
        while (r < rows) {
            const n = numbers.items[r].items[c];
            result = if (op == '*') result * n else result + n;
            r += 1;
        }

        sum += result;
        c += 1;
    }
    return day.Answer{ .int = sum };
}

pub fn part2(_: std.mem.Allocator, _: *std.Io.Reader) !day.Answer {
    return day.Answer{ .int = 0 };
}

test "part1" {
    const input =
        \\123 328  51 64 
        \\ 45 64  387 23 
        \\  6 98  215 314
        \\*   +   *   +  
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part1(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 4277556 }, result);
}

test "part2" {
    const input =
        \\123 328  51 64 
        \\ 45 64  387 23 
        \\  6 98  215 314
        \\*   +   *   +  
    ;

    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 0 }, result);
}
