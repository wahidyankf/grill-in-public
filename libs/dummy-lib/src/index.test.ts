import assert from "node:assert/strict";
import test from "node:test";

import { createGreeting } from "./index.js";

test("creates a greeting for the supplied name", () => {
  assert.equal(createGreeting("Wahidyan"), "Hello, Wahidyan!");
});
