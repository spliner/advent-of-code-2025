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

pub const LineIterator = struct {
    reader: *std.Io.Reader,
    line_writer: std.Io.Writer.Allocating,

    pub fn init(allocator: std.mem.Allocator, reader: *std.Io.Reader) LineIterator {
        return .{
            .reader = reader,
            .line_writer = std.Io.Writer.Allocating.init(allocator),
        };
    }

    pub fn deinit(self: *LineIterator) void {
        self.line_writer.deinit();
    }

    pub fn next(self: *LineIterator) !?[]const u8 {
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
