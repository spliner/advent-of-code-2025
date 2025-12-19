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
        const has_more = blk: {
            _ = reader.streamDelimiter(&line_writer.writer, '\n') catch |err| switch (err) {
                error.EndOfStream => break :blk false,
                else => return err,
            };
            break :blk true;
        };
        const line = line_writer.written();
        if (line.len == 0) {
            break;
        }

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

        if (!has_more) {
            break;
        }

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

pub fn part2(allocator: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var lines: std.ArrayList([]u8) = .empty;
    defer lines.deinit(allocator);
    defer for (lines.items) |l| {
        allocator.free(l);
    };

    var line_writer = std.io.Writer.Allocating.init(allocator);
    defer line_writer.deinit();

    while (true) {
        const has_more = blk: {
            _ = reader.streamDelimiter(&line_writer.writer, '\n') catch |err| switch (err) {
                error.EndOfStream => break :blk false,
                else => return err,
            };
            break :blk true;
        };
        const line = line_writer.written();
        if (line.len == 0) {
            break;
        }

        const copied = try allocator.alloc(u8, line.len);
        std.mem.copyForwards(u8, copied, line);
        try lines.append(allocator, copied);

        line_writer.clearRetainingCapacity();
        if (!has_more) {
            break;
        }
        reader.toss(1); // Skip delimiter
    }

    const last_line = lines.items[lines.items.len - 1];
    var start: usize = 0;
    var limit = findLimit(last_line, start);
    var sum: usize = 0;
    while (start <= last_line.len) {
        var numbers: std.ArrayList(usize) = .empty;
        defer numbers.deinit(allocator);

        var col: usize = limit;
        while (col >= start) {
            var number: usize = 0;
            var row: usize = 0;
            while (row < lines.items.len - 1) : (row += 1) {
                const n = lines.items[row][col];
                if (n != ' ') {
                    number = number * 10 + (n - '0');
                }
            }

            try numbers.append(allocator, number);

            if (col == 0) {
                break;
            }
            col -= 1;
        }

        const operator = last_line[start];
        const func: *const fn (acc: usize, curr: usize) usize, var acc: usize =
            if (operator == '+') .{ sumNumbers, 0 } else .{ multiplyNumbers, 1 };
        for (numbers.items) |n| {
            acc = func(acc, n);
        }
        sum += acc;

        start = limit + 2;
        limit = findLimit(last_line, start);
    }

    return day.Answer{ .int = @intCast(sum) };
}

fn sumNumbers(x: usize, y: usize) usize {
    return x + y;
}

fn multiplyNumbers(x: usize, y: usize) usize {
    return x * y;
}

fn findLimit(line: []u8, start: usize) usize {
    var i = start + 1;
    while (i < line.len) : (i += 1) {
        if (line[i] == '*' or line[i] == '+') {
            // Accounting for operator index + whitespace
            return i - 2;
        }
    }

    return line.len - 1;
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

    try std.testing.expectEqual(day.Answer{ .int = 3263827 }, result);
}
