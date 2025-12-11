const std = @import("std");

pub const Answer = union(enum) {
    int: i64,
    string: []const u8,

    pub fn format(self: Answer, writer: anytype) !void {
        switch (self) {
            .int => |val| try writer.print("{d}", .{val}),
            .string => |val| try writer.print("{s}", .{val}),
        }
    }
};

pub const SolutionFn = *const fn (allocator: std.mem.Allocator, reader: *std.Io.Reader) anyerror!Answer;
