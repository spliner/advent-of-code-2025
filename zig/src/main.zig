const std = @import("std");
const day01 = @import("day01");
const day02 = @import("day02");
const day03 = @import("day03");
const day04 = @import("day04");
const day05 = @import("day05");
const day06 = @import("day06");
const day07 = @import("day07");

const ArgParseError = error{ MissingArgs, InvalidDay, InvalidPart };

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    const args = try std.process.argsAlloc(allocator);
    if (args.len < 4) {
        std.debug.print("Usage: {s} <day> <part> <input_file>\n", .{args[0]});
        return ArgParseError.MissingArgs;
    }

    const day = try std.fmt.parseInt(u32, args[1], 10);
    const part = try std.fmt.parseInt(u32, args[2], 10);
    const filepath = args[3];

    if (part != 1 and part != 2) {
        std.debug.print("Part must be 1 or 2\n", .{});
        return ArgParseError.InvalidPart;
    }

    const file = try std.fs.cwd().openFile(filepath, .{ .mode = .read_only });
    defer file.close();

    var buffer: [1024]u8 = undefined;
    var file_reader = file.reader(&buffer);
    const reader = &file_reader.interface;

    const answer = switch (day) {
        1 => if (part == 1) try day01.part1(allocator, reader) else try day01.part2(allocator, reader),
        2 => if (part == 1) try day02.part1(allocator, reader) else try day02.part2(allocator, reader),
        3 => if (part == 1) try day03.part1(allocator, reader) else try day03.part2(allocator, reader),
        4 => if (part == 1) try day04.part1(allocator, reader) else try day04.part2(allocator, reader),
        5 => if (part == 1) try day05.part1(allocator, reader) else try day05.part2(allocator, reader),
        6 => if (part == 1) try day06.part1(allocator, reader) else try day06.part2(allocator, reader),
        7 => if (part == 1) try day07.part1(allocator, reader) else try day07.part2(allocator, reader),
        else => {
            std.debug.print("Day {d} not implemented yet\n", .{day});
            return ArgParseError.InvalidDay;
        },
    };

    std.debug.print("Day {d} Part {d}: {f}\n", .{ day, part, answer });
}

test "simple test" {
    const gpa = std.testing.allocator;
    var list: std.ArrayList(i32) = .empty;
    defer list.deinit(gpa); // Try commenting this out and see if zig detects the memory leak!
    try list.append(gpa, 42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}

test "fuzz example" {
    const Context = struct {
        fn testOne(context: @This(), input: []const u8) anyerror!void {
            _ = context;
            // Try passing `--fuzz` to `zig build test` and see if it manages to fail this test case!
            try std.testing.expect(!std.mem.eql(u8, "canyoufindme", input));
        }
    };
    try std.testing.fuzz(Context{}, Context.testOne, .{});
}
