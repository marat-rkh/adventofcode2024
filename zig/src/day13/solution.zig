const std = @import("std");
const fmt = std.fmt;
const fs = std.fs;
const heap = std.heap;
const math = std.math;
const mem = std.mem;
const testing = std.testing;

// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
//
// X: a * 94 + b * 22 = 8400
// Y: a * 34 + b * 67 = 5400
//
// D = 94 * 67 - 34 * 22 = 5550
// Da = 8400 * 67 - 5400 * 22 = 444000
// Db = 94 * 5400 - 34 * 8400 = 222000
// a = Da / D = 80
// b = Db / D = 40
fn solve1(path: []const u8) !i64 {
    return try solve(path, 0);
}

fn solve2(path: []const u8) !i64 {
    return try solve(path, 10000000000000);
}

fn solve(path: []const u8, shift: i64) !i64 {
    var gpa = heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    const file = try fs.cwd().openFile(path, .{});
    defer file.close();
    const input = try file.readToEndAlloc(allocator, math.maxInt(u32));
    defer allocator.free(input);
    var blocks = mem.tokenizeSequence(u8, input, "\n\n");
    var total: i64 = 0;
    while (blocks.next()) |block| {
        var lines = mem.tokenizeSequence(u8, block, "\n");
        const lineA = lines.next().?;
        const coordsA = try parseCoords(lineA, "X+", "Y+");
        const lineB = lines.next().?;
        const coordsB = try parseCoords(lineB, "X+", "Y+");
        const linePrize = lines.next().?;
        var coordsPrize = try parseCoords(linePrize, "X=", "Y=");
        coordsPrize.x += shift;
        coordsPrize.y += shift;
        const D = coordsA.x * coordsB.y - coordsA.y * coordsB.x;
        const Da = coordsPrize.x * coordsB.y - coordsPrize.y * coordsB.x;
        const Db = coordsA.x * coordsPrize.y - coordsA.y * coordsPrize.x;
        const a = math.divExact(i64, Da, D) catch continue;
        const b = math.divExact(i64, Db, D) catch continue;
        if (a >= 0 or b >= 0) {
            total += (a * 3 + b);
        }
    }
    return total;
}

fn parseCoords(line: []const u8, xPrefix: []const u8, yPrefix: []const u8) !struct { x: i64, y: i64 } {
    const xStr = line[mem.indexOf(u8, line, xPrefix).? + xPrefix.len .. mem.indexOf(u8, line, ",").?];
    const yStr = line[mem.indexOf(u8, line, yPrefix).? + yPrefix.len ..];
    const x = try fmt.parseInt(i64, xStr, 10);
    const y = try fmt.parseInt(i64, yStr, 10);
    return .{ .x = x, .y = y };
}

test "part 1 input 0" {
    try testing.expectEqual(480, try solve1("src/day13/in0.txt"));
}

test "part 1 input 1" {
    try testing.expectEqual(40369, try solve1("src/day13/in1.txt"));
}

test "part 2 input 0" {
    try testing.expectEqual(875318608908, try solve2("src/day13/in0.txt"));
}

test "part 2 input 1" {
    try testing.expectEqual(72587986598368, try solve2("src/day13/in1.txt"));
}
