const std = @import("std");
const testing = std.testing;

fn solve1(_: []const u8) i32 {
    return 0;
}

test "basic add functionality" {
    try testing.expect(solve1("in0.txt") == 0);
}
