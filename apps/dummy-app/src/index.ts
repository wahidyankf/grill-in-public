import { createGreeting } from "@swe-grilling/dummy-lib";

/** Produces greetings for the dummy app. */
export type GreetingCreator = (name: string) => string;

/** Creates the dummy app with its greeting dependency. */
export function createApp(createGreetingForName: GreetingCreator): {
  run: () => string;
} {
  return {
    run: () => createGreetingForName("Wahidyan"),
  };
}

/** Produces the dummy app's output using the workspace library. */
export function run(): string {
  return createApp(createGreeting).run();
}
