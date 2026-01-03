const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});

    const exe = b.addExecutable(.{
        .name = "loocup",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/loocup.zig"),
            .target = target,
        }),
    });

    b.installArtifact(exe);

    const run = b.addRunArtifact(exe);
    const run_step = b.step("run", "Directly run loocup");
    run_step.dependOn(&run.step);
}
