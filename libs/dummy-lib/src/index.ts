/** Returns a message for applications that consume this library. */
export function createGreeting(name: string): string {
  // Keep this tiny library pure: callers can exercise it without setup, and the
  // app can demonstrate workspace consumption separately from side effects.
  return `Hello, ${name}!`;
}
