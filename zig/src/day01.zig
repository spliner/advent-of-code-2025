const std = @import("std");

const Dial = struct {
    position: i32,
    stoppedAtZeroCount: i32,
    passedThroughZeroCount: i32,

    pub fn init(position: i32) Dial {
        return .{
            .position = position,
            .stoppedAtZeroCount = 0,
            .passedThroughZeroCount = 0,
        };
    }

    pub fn turnLeft(self: *Dial, times: i32) void {
        const fullTurns = @divTrunc(times, 100);
        self.passedThroughZeroCount += fullTurns;

        const effectiveTimes = @mod(times, 100);
        var newPosition = self.position - effectiveTimes;
        var passedThroughZero = false;
        if (newPosition < 0) {
            newPosition = 100 + newPosition;
            passedThroughZero = self.position != 0;
        }

        if (newPosition == 0) {
            self.stoppedAtZeroCount += 1;
        } else if (passedThroughZero) {
            self.passedThroughZeroCount += 1;
        }

        self.position = newPosition;
    }

    pub fn turnRight(self: *Dial, times: i32) void {
        const fullTurns = @divTrunc(times, 100);
        self.passedThroughZeroCount += fullTurns;

        const effectiveTimes = @mod(times, 100);
        var newPosition = self.position + effectiveTimes;
        var passedThroughZero = false;
        if (newPosition > 99) {
            newPosition -= 100;
            passedThroughZero = true;
        }

        if (newPosition == 0) {
            self.stoppedAtZeroCount += 1;
        } else if (passedThroughZero) {
            self.passedThroughZeroCount += 1;
        }

        self.position = newPosition;
    }
};

pub fn part1(_: std.mem.Allocator, reader: *std.Io.Reader) !i32 {
    var d = Dial.init(50);

    while (try reader.takeDelimiter('\n')) |line| {
        const trim = std.mem.trim(u8, line, " \r\t");
        const direction = trim[0];
        const rawTimes = trim[1..];
        const times = try std.fmt.parseInt(i32, rawTimes, 10);
        if (direction == 'L') {
            d.turnLeft(times);
        } else if (direction == 'R') {
            d.turnRight(times);
        }
    }

    return d.stoppedAtZeroCount;
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
    const allocator = std.testing.allocator;

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
    const result = try part1(allocator, &reader.interface);
    try std.testing.expectEqual(3, result);
}
