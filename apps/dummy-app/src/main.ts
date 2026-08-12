import { run } from "./index.js";

// Keep process concerns in this thin entry point; index.ts stays reusable by
// tests and other callers without having to capture console output.
console.log(run());
