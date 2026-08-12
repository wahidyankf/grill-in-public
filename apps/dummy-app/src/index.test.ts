import assert from "node:assert/strict";
import test from "node:test";

import { run } from "./index.js";

test("the dummy app consumes the workspace library", () => {
  assert.equal(run(), "Hello, Wahidyan!");
});
