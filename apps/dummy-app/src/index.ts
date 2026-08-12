import { createGreeting } from "@swe-grilling/dummy-lib";

/** Produces the dummy app's output using the workspace library. */
export function run(): string {
  return createGreeting("Wahidyan");
}
