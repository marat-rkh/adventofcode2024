const std = @import("std");
const fmt = std.fmt;
const fs = std.fs;
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
        const xRes: u16 = @intCast(try math.mod(i32, p.x + v.x * 100, width));
        const yRes: u16 = @intCast(try math.mod(i32, p.y + v.y * 100, height));
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

fn parseCoords(line: []const u8) !struct { x: i32, y: i32 } {
    const commaIndex = mem.indexOf(u8, line, ",").?;
    const xStr = line[0..commaIndex];
    const yStr = line[commaIndex + 1 ..];
    const x = try fmt.parseInt(i32, xStr, 10);
    const y = try fmt.parseInt(i32, yStr, 10);
    return .{ .x = x, .y = y };
}

test "part 1 input 0" {
    try testing.expectEqual(12, try solve1("src/day14/in0.txt", 11, 7));
}

test "part 1 input 1" {
    try testing.expectEqual(228690000, try solve1("src/day14/in1.txt", 101, 103));
}
