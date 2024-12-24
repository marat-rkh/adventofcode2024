const std = @import("std");
const testing = std.testing;

fn solve1(_: []const u8) i32 {
    return 0;
}

test "input 0" {
    // X: a * 94 + b * 22 = 8400
    // Y: a * 34 + b * 67 = 5400
    // X: 94a = 8400 - 22b
    // X: a = (8400 - 22 * b) / 94
    // Y: (8400 - 22 * b) / 94 * 34 + b * 67 = 5400
    // Basically, we need to find out if there are if there exist any a and b from N.
    try testing.expect(solve1("in0.txt") == 0);
}
