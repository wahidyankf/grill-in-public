import assert from "node:assert/strict";
import test from "node:test";

import { run } from "./index.js";

test("the dummy app consumes the workspace library", () => {
  // Unlike the quick unit test, this intentionally crosses the app/library
  // boundary and therefore belongs in the explicit integration suite.
  assert.equal(run(), "Hello, Wahidyan!");
});
