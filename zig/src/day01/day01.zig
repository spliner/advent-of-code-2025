const std = @import("std");

const Dial = @import("dial.zig").Dial;

pub fn part1(_: std.mem.Allocator, reader: *std.Io.Reader) !void {
    var d = Dial.init(50);
    d.turnLeft(10);

    while (try reader.takeDelimiter('\n')) |line| {
        const trim = std.mem.trim(u8, line, " \r\t");
        const direction = trim[0];
        std.debug.print("direction: {c}\n", .{direction});
        const rawTimes = trim[1..];
        std.debug.print("times: {s}\n", .{rawTimes});
        const times = try std.fmt.parseInt(i32, rawTimes, 10);
        if (direction == 'L') {
            d.turnLeft(times);
        } else if (direction == 'R') {
            d.turnRight(times);
        }
        std.debug.print("{s}\n", .{line});
    }
}
