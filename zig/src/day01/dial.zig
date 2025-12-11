const std = @import("std");

pub const Dial = struct {
    position: i32,
    stoppedAtZeroCount: i32,
    passedThroughZeroCount: i32,

    pub fn init(position: i32) Dial {
        return Dial{
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
        self.position += times;
    }
};

test "turnLeft" {
    var dial = Dial.init(50);
    dial.turnLeft(1);
    std.testing.expectEqual(49, dial.position);
}
