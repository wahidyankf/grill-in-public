import assert from "node:assert/strict";
import test from "node:test";

import { createApp } from "./index.js";

test("the dummy app delegates to its injected greeting creator", () => {
  const receivedNames: string[] = [];
  const app = createApp((name) => {
    receivedNames.push(name);
    return "Mock greeting";
  });

  assert.equal(app.run(), "Mock greeting");
  assert.deepEqual(receivedNames, ["Wahidyan"]);
});
