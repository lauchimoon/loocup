const std = @import("std");

pub fn FuncArg() type {
    return struct {
       const Self = @This(); 

        typ: []const u8,
        name: []const u8,

        pub fn init(typ: []const u8, name: []const u8) Self {
            return Self{
                .typ = typ,
                .name = name,
            };
        }
    };
}

const ArgList = std.ArrayList(FuncArg());

pub fn Function() type {
    return struct {
        const Self = @This();

        ret_type: []const u8,
        name: []const u8,
        args: ArgList,

        pub fn init(ret_type: []const u8, name: []const u8, args: ArgList) Self {
            return Self{
                .ret_type = ret_type,
                .name = name,
                .args = args,
            };
        }

        pub fn print(f: Self) void {
            std.debug.print("{s} {s}(", .{f.ret_type, f.name});
            const n_args = f.args.items.len;
            for (f.args.items, 0 .. n_args) |arg, i| {
                if (i + 1 >= n_args) {
                    std.debug.print("{s} {s}", .{arg.typ, arg.name});
                } else {
                    std.debug.print("{s} {s}, ", .{arg.typ, arg.name});
                }
            }
            std.debug.print(");\n", .{});
        }
    };
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const capacity: usize = 256;
    var args: ArgList = try ArgList.initCapacity(allocator, capacity);
    defer args.deinit(allocator);

    try args.append(allocator, FuncArg().init("int", "a"));
    try args.append(allocator, FuncArg().init("int", "b"));

    const f = Function().init("int", "add", args);
    f.print();
}
