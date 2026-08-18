import { createGreeting } from "@grind-in-public/dummy-lib";

/** Produces greetings for the dummy app. */
export type GreetingCreator = (name: string) => string;

/** Creates the dummy app with its greeting dependency. */
export function createApp(createGreetingForName: GreetingCreator): {
  run: () => string;
} {
  return {
    // Keep the app's fixed drill input at the composition boundary so
    // a unit test can replace only the collaboration, not the behavior under test.
    run: () => createGreetingForName("Wahidyan"),
  };
}

/** Produces the dummy app's output using the workspace library. */
export function run(): string {
  // Production composition supplies the real workspace dependency; unit tests
  // exercise createApp directly with a deterministic test double instead.
  return createApp(createGreeting).run();
}
