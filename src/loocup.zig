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

        pub fn eql(self: Self, other: Self) bool {
            return std.mem.eql(u8, self.typ, other.typ);
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
    var argsf: ArgList = try ArgList.initCapacity(allocator, capacity);
    defer argsf.deinit(allocator);

    var argsg: ArgList = try ArgList.initCapacity(allocator, capacity);
    defer argsg.deinit(allocator);

    try argsf.append(allocator, FuncArg().init("int", "a"));
    try argsf.append(allocator, FuncArg().init("int", "b"));

    try argsg.append(allocator, FuncArg().init("int", "a"));
    try argsg.append(allocator, FuncArg().init("int", "b"));

    const f = Function().init("int", "f", argsf);
    const g = Function().init("bool", "g", argsg);
    f.print();
    g.print();

    std.debug.print("{}\n", .{functionsMatch(f, g)});
}

// Function names doesn't matter for now, we care about return type and arguments' type
pub fn functionsMatch(f: Function(), g: Function()) bool {
    const f_args_len = f.args.items.len;
    if (g.args.items.len != f.args.items.len) return false;

    for (0..f_args_len) |i| {
        const f_arg = f.args.items[i];
        const g_arg = g.args.items[i];
        if (!f_arg.eql(g_arg)) return false;
    }

    return std.mem.eql(u8, f.ret_type, g.ret_type);
}
