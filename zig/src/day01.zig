const std = @import("std");
const day = @import("day.zig");

const Dial = struct {
    position: i32,
    stopped_zero_count: i32,
    passed_zero_count: i32,

    pub fn init(position: i32) Dial {
        return .{
            .position = position,
            .stopped_zero_count = 0,
            .passed_zero_count = 0,
        };
    }

    pub fn turnLeft(self: *Dial, times: i32) void {
        const full_turns = @divTrunc(times, 100);
        self.passed_zero_count += full_turns;

        const effective_times = @mod(times, 100);
        var new_position = self.position - effective_times;
        var passed_through_zero = false;
        if (new_position < 0) {
            new_position = 100 + new_position;
            passed_through_zero = self.position != 0;
        }

        if (new_position == 0) {
            self.stopped_zero_count += 1;
            self.passed_zero_count += 1;
        } else if (passed_through_zero) {
            self.passed_zero_count += 1;
        }

        self.position = new_position;
    }

    pub fn turnRight(self: *Dial, times: i32) void {
        const full_turns = @divTrunc(times, 100);
        self.passed_zero_count += full_turns;

        const effective_times = @mod(times, 100);
        var new_position = self.position + effective_times;
        var passed_through_zero = false;
        if (new_position > 99) {
            new_position -= 100;
            passed_through_zero = true;
        }

        if (new_position == 0) {
            self.stopped_zero_count += 1;
            self.passed_zero_count += 1;
        } else if (passed_through_zero) {
            self.passed_zero_count += 1;
        }

        self.position = new_position;
    }
};

pub fn part1(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var d = Dial.init(50);

    while (try reader.takeDelimiter('\n')) |line| {
        const trim = std.mem.trim(u8, line, " \r\t");
        const direction = trim[0];
        const times = try std.fmt.parseInt(i32, trim[1..], 10);
        if (direction == 'L') {
            d.turnLeft(times);
        } else if (direction == 'R') {
            d.turnRight(times);
        }
    }

    return day.Answer{ .int = d.stopped_zero_count };
}

pub fn part2(_: std.mem.Allocator, reader: *std.Io.Reader) !day.Answer {
    var d = Dial.init(50);

    while (try reader.takeDelimiter('\n')) |line| {
        const trim = std.mem.trim(u8, line, " \r\t");
        const direction = trim[0];
        const times = try std.fmt.parseInt(i32, trim[1..], 10);
        if (direction == 'L') {
            d.turnLeft(times);
        } else if (direction == 'R') {
            d.turnRight(times);
        }
    }

    return day.Answer{ .int = d.passed_zero_count };
}

test {
    std.testing.refAllDecls(@This());
}

test "turnLeft" {
    const cases = [_]struct {
        startPos: i32,
        times: i32,
        expectedPos: i32,
    }{
        .{ .startPos = 50, .times = 1, .expectedPos = 49 },
        .{ .startPos = 50, .times = 100, .expectedPos = 50 },
        .{ .startPos = 0, .times = 1, .expectedPos = 99 },
    };

    for (cases) |case| {
        var dial = Dial.init(case.startPos);
        dial.turnLeft(case.times);
        try std.testing.expectEqual(case.expectedPos, dial.position);
    }
}

test "turnRight" {
    const cases = [_]struct {
        startPos: i32,
        times: i32,
        expectedPos: i32,
    }{
        .{ .startPos = 50, .times = 1, .expectedPos = 51 },
        .{ .startPos = 50, .times = 100, .expectedPos = 50 },
        .{ .startPos = 99, .times = 1, .expectedPos = 0 },
    };

    for (cases) |case| {
        var dial = Dial.init(case.startPos);
        dial.turnRight(case.times);
        try std.testing.expectEqual(case.expectedPos, dial.position);
    }
}

test "part1" {
    const input =
        \\L68
        \\L30
        \\R48
        \\L5
        \\R60
        \\L55
        \\L1
        \\L99
        \\R14
        \\L82
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
        \\L68
        \\L30
        \\R48
        \\L5
        \\R60
        \\L55
        \\L1
        \\L99
        \\R14
        \\L82
    ;
    var reader_buf: [1024]u8 = undefined;
    var reader = std.testing.Reader.init(&reader_buf, &.{
        .{ .buffer = input },
    });

    const result = try part2(std.testing.allocator, &reader.interface);

    try std.testing.expectEqual(day.Answer{ .int = 6 }, result);
}
