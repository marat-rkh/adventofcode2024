const std = @import("std");
const fmt = std.fmt;
const fs = std.fs;
const io = std.io;
const heap = std.heap;
const math = std.math;
const mem = std.mem;
const testing = std.testing;

fn solve1(path: []const u8, width: u16, height: u16) !u32 {
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    const file = try fs.cwd().openFile(path, .{});
    defer file.close();
    const input = try file.readToEndAlloc(allocator, math.maxInt(u32));
    defer allocator.free(input);
    var lines = mem.tokenizeSequence(u8, input, "\n");
    var q1: u32 = 0;
    var q2: u32 = 0;
    var q3: u32 = 0;
    var q4: u32 = 0;
    while (lines.next()) |line| {
        const spaceIndex = mem.indexOf(u8, line, " ").?;
        const pStr = line[2..spaceIndex];
        const vStr = line[spaceIndex + 3 ..];
        const p = try parseCoords(pStr);
        const v = try parseCoords(vStr);
        const xRes: u16 = @intCast(try math.mod(i32, p[0] + v[0] * 100, width));
        const yRes: u16 = @intCast(try math.mod(i32, p[1] + v[1] * 100, height));
        if (xRes < width / 2) {
            if (yRes < height / 2) {
                q1 += 1;
            } else if (yRes > height / 2) {
                q4 += 1;
            }
        } else if (xRes > width / 2) {
            if (yRes < height / 2) {
                q2 += 1;
            } else if (yRes > height / 2) {
                q3 += 1;
            }
        }
    }
    return q1 * q2 * q3 * q4;
}

fn parseCoords(line: []const u8) ![2]i32 {
    const commaIndex = mem.indexOf(u8, line, ",").?;
    const xStr = line[0..commaIndex];
    const yStr = line[commaIndex + 1 ..];
    const x = try fmt.parseInt(i32, xStr, 10);
    const y = try fmt.parseInt(i32, yStr, 10);
    return .{ x, y };
}

fn solve2(path: []const u8, width: comptime_int, height: comptime_int) !void {
    const stdout = io.getStdOut().writer();
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    const file = try fs.cwd().openFile(path, .{});
    defer file.close();
    const input = try file.readToEndAlloc(allocator, math.maxInt(u32));
    defer allocator.free(input);
    var lines = mem.tokenizeSequence(u8, input, "\n");
    var points = std.ArrayList([2]u16).init(allocator);
    defer points.deinit();
    var velocities = std.ArrayList([2]i32).init(allocator);
    defer velocities.deinit();
    while (lines.next()) |line| {
        const spaceIndex = mem.indexOf(u8, line, " ").?;
        const pStr = line[2..spaceIndex];
        const vStr = line[spaceIndex + 3 ..];
        const p = try parseCoords(pStr);
        const v = try parseCoords(vStr);
        try points.append(.{ @intCast(p[0]), @intCast(p[1]) });
        try velocities.append(v);
    }
    try stdout.print("Initial:\n", .{});
    try printField(points, width, height);
    // Answer is 7093
    for (0..10000) |step| {
        for (points.items, velocities.items) |*p, v| {
            p[0] = @intCast(try math.mod(i32, p[0] + v[0], width));
            p[1] = @intCast(try math.mod(i32, p[1] + v[1], height));
        }
        try stdout.print("Step {d}:\n", .{step + 1});
        try printField(points, width, height);
    }
}

fn printField(points: std.ArrayList([2]u16), width: comptime_int, height: comptime_int) !void {
    const stdout = io.getStdOut().writer();
    var field = mem.zeroes([width][height]bool);
    for (points.items) |p| {
        field[p[0]][p[1]] = true;
    }
    for (0..height) |y| {
        for (0..width) |x| {
            if (field[x][y]) {
                try stdout.print("#", .{});
            } else {
                try stdout.print(".", .{});
            }
        }
        try stdout.print("\n", .{});
    }
    try stdout.print("\n", .{});
}

test "part 1 input 0" {
    try testing.expectEqual(12, try solve1("src/day14/in0.txt", 11, 7));
}

test "part 1 input 1" {
    try testing.expectEqual(228690000, try solve1("src/day14/in1.txt", 101, 103));
}

pub fn main() !void {
    try solve2("src/day14/in1.txt", 101, 103);
}
