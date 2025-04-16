#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const child_process_1 = require("child_process");
function getExePath() {
    const arch = process.arch;
    let os = process.platform;
    let extension = '';
    if (['win32', 'cygwin'].includes(process.platform)) {
        os = 'windows';
        extension = '.exe';
    }
    try {
        // Since the bin will be located inside `node_modules`, we can simply call require.resolve
        return require.resolve(`dofy-${os}-${arch}/bin/dofy${extension}`);
    }
    catch (e) {
        throw new Error(`Couldn't find dofy binary inside node_modules for ${os}-${arch}`);
    }
}
function run() {
    var _a;
    const args = process.argv.slice(2);
    const processResult = (0, child_process_1.spawnSync)(getExePath(), args, { stdio: "inherit" });
    process.exit((_a = processResult.status) !== null && _a !== void 0 ? _a : 0);
}
run();
