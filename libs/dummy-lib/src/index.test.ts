import assert from "node:assert/strict";
import test from "node:test";

import { createGreeting } from "./index.js";

test("creates a greeting for the supplied name", () => {
  // This literal output is the library contract consumed by Dummy App's
  // integration test, so it guards the cross-project example as well.
  assert.equal(createGreeting("Wahidyan"), "Hello, Wahidyan!");
});
