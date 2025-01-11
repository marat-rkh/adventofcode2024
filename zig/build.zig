const std = @import("std");

// Although this function looks imperative, note that its job is to
// declaratively construct a build graph that will be executed by an external
// runner.
pub fn build(b: *std.Build) void {
    // Standard target options allows the person running `zig build` to choose
    // what target to build for. Here we do not override the defaults, which
    // means any target is allowed, and the default is native. Other options
    // for restricting supported target set are available.
    const target = b.standardTargetOptions(.{});

    // Standard optimization options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall. Here we do not
    // set a preferred release mode, allowing the user to decide how to optimize.
    const optimize = b.standardOptimizeOption(.{});

    const run_test_day13 = b.addRunArtifact(b.addTest(.{
        .root_source_file = b.path("src/day13/solution.zig"),
        .target = target,
        .optimize = optimize,
    }));
    b.step("day13", "Run tests for day 13").dependOn(&run_test_day13.step);

    const run_test_day14 = b.addRunArtifact(b.addTest(.{
        .root_source_file = b.path("src/day14/solution.zig"),
        .target = target,
        .optimize = optimize,
    }));
    b.step("day14", "Run tests for day 14").dependOn(&run_test_day14.step);

    const run_day14_part2 = b.addRunArtifact(b.addExecutable(.{
        .name = "day14part2",
        .root_source_file = b.path("src/day14/solution.zig"),
        .target = target,
        .optimize = optimize,
    }));

    b.step("day14part2", "Run day 14 part 2").dependOn(&run_day14_part2.step);
}
