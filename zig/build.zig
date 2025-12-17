const std = @import("std");

// Although this function looks imperative, it does not perform the build
// directly and instead it mutates the build graph (`b`) that will be then
// executed by an external runner. The functions in `std.Build` implement a DSL
// for defining build steps and express dependencies between them, allowing the
// build runner to parallelize the build automatically (and the cache system to
// know when a step doesn't need to be re-run).
pub fn build(b: *std.Build) void {
    // Standard target options allow the person running `zig build` to choose
    // what target to build for. Here we do not override the defaults, which
    // means any target is allowed, and the default is native. Other options
    // for restricting supported target set are available.
    const target = b.standardTargetOptions(.{});
    // Standard optimization options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall. Here we do not
    // set a preferred release mode, allowing the user to decide how to optimize.
    const optimize = b.standardOptimizeOption(.{});
    // It's also possible to define more custom flags to toggle optional features
    // of this build script using `b.option()`. All defined flags (including
    // target and optimize options) will be listed when running `zig build --help`
    // in this directory.

    // Shared day module for common types
    const day = b.addModule("day", .{
        .root_source_file = b.path("src/day.zig"),
        .target = target,
    });

    const day01 = b.addModule("day01", .{
        .root_source_file = b.path("src/day01.zig"),
        .target = target,
        .imports = &.{
            .{ .name = "day.zig", .module = day },
        },
    });
    const day02 = b.addModule("day02", .{
        .root_source_file = b.path("src/day02.zig"),
        .target = target,
        .imports = &.{
            .{ .name = "day.zig", .module = day },
        },
    });
    const day03 = b.addModule("day03", .{
        .root_source_file = b.path("src/day03.zig"),
        .target = target,
        .imports = &.{
            .{ .name = "day.zig", .module = day },
        },
    });
    const day04 = b.addModule("day04", .{
        .root_source_file = b.path("src/day04.zig"),
        .target = target,
        .imports = &.{
            .{ .name = "day.zig", .module = day },
        },
    });
    const day05 = b.addModule("day05", .{
        .root_source_file = b.path("src/day05.zig"),
        .target = target,
        .imports = &.{
            .{ .name = "day.zig", .module = day },
        },
    });

    // Main executable for running Advent of Code solutions
    const exe = b.addExecutable(.{
        .name = "aoc2025",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/main.zig"),
            .target = target,
            .optimize = optimize,
            // Import day modules - add new days here
            .imports = &.{
                .{ .name = "day01", .module = day01 },
                .{ .name = "day02", .module = day02 },
                .{ .name = "day03", .module = day03 },
                .{ .name = "day04", .module = day04 },
            },
        }),
    });

    // This declares intent for the executable to be installed into the
    // install prefix when running `zig build` (i.e. when executing the default
    // step). By default the install prefix is `zig-out/` but can be overridden
    // by passing `--prefix` or `-p`.
    b.installArtifact(exe);

    // This creates a top level step. Top level steps have a name and can be
    // invoked by name when running `zig build` (e.g. `zig build run`).
    // This will evaluate the `run` step rather than the default step.
    // For a top level step to actually do something, it must depend on other
    // steps (e.g. a Run step, as we will see in a moment).
    const run_step = b.step("run", "Run the app");

    // This creates a RunArtifact step in the build graph. A RunArtifact step
    // invokes an executable compiled by Zig. Steps will only be executed by the
    // runner if invoked directly by the user (in the case of top level steps)
    // or if another step depends on it, so it's up to you to define when and
    // how this Run step will be executed. In our case we want to run it when
    // the user runs `zig build run`, so we create a dependency link.
    const run_cmd = b.addRunArtifact(exe);
    run_step.dependOn(&run_cmd.step);

    // By making the run step depend on the default step, it will be run from the
    // installation directory rather than directly from within the cache directory.
    run_cmd.step.dependOn(b.getInstallStep());

    // This allows the user to pass arguments to the application in the build
    // command itself, like this: `zig build run -- arg1 arg2 etc`
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    // Setup tests for the executable and each day module
    const exe_tests = b.addTest(.{
        .root_module = exe.root_module,
    });
    const run_exe_tests = b.addRunArtifact(exe_tests);

    const day01_tests = b.addTest(.{
        .root_module = day01,
    });
    const run_day01_tests = b.addRunArtifact(day01_tests);
    const day02_tests = b.addTest(.{
        .root_module = day02,
    });
    const run_day02_tests = b.addRunArtifact(day02_tests);
    const day03_tests = b.addTest(.{
        .root_module = day03,
    });
    const run_day03_tests = b.addRunArtifact(day03_tests);
    const day04_tests = b.addTest(.{
        .root_module = day04,
    });
    const run_day04_tests = b.addRunArtifact(day04_tests);
    const day05_tests = b.addTest(.{
        .root_module = day05,
    });
    const run_day05_tests = b.addRunArtifact(day05_tests);

    const test_step = b.step("test", "Run all tests");
    test_step.dependOn(&run_exe_tests.step);
    test_step.dependOn(&run_day01_tests.step);
    test_step.dependOn(&run_day02_tests.step);
    test_step.dependOn(&run_day03_tests.step);
    test_step.dependOn(&run_day04_tests.step);
    test_step.dependOn(&run_day05_tests.step);

    // Just like flags, top level steps are also listed in the `--help` menu.
    //
    // The Zig build system is entirely implemented in userland, which means
    // that it cannot hook into private compiler APIs. All compilation work
    // orchestrated by the build system will result in other Zig compiler
    // subcommands being invoked with the right flags defined. You can observe
    // these invocations when one fails (or you pass a flag to increase
    // verbosity) to validate assumptions and diagnose problems.
    //
    // Lastly, the Zig build system is relatively simple and self-contained,
    // and reading its source code will allow you to master it.
}
