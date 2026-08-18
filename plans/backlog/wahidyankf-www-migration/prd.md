# Product Requirements

## Overview

The product is the owner's personal website, unchanged in behavior, served from this repository instead of `ose-public`. The requirements below describe the migrated system's observable behavior, not the migration steps.

## Personas

**The owner** maintains the site and needs to change it without leaving this repository.

**A visitor** reads the site and must not be able to tell that anything moved.

## User Stories

**US-1** — As the owner, I want to build and run the site from this repository, so that I can change it without switching repositories.

**US-2** — As the owner, I want the site's tests to run under this workspace's Nx targets, so that its checks participate in the same gates as everything else here.

**US-3** — As the owner, I want the site's Gherkin scenarios in this repository's `specs/` tree, so that its behavior is described where the specs policy says behavior lives.

**US-4** — As a visitor, I want the site to serve the same pages as before, so that the move is invisible to me.

**US-5** — As the owner, I want the `ose-public` copy removed only after the migrated copy is verified, so that no window exists where neither copy is trustworthy.

## Acceptance Criteria

```gherkin
Feature: Migrated site builds and runs

  Scenario: The site builds from this repository
    Given the workspace has installed its dependencies
    When "npx nx run wahidyankf-www:build" runs
    Then the command exits 0
    And a Next.js build output exists under "apps/wahidyankf-www/.next"

  Scenario: The development server serves the home page
    Given the site has been built
    When "npx nx run wahidyankf-www:dev" runs and port 3201 is requested
    Then the home page responds with status 200

  Scenario: The site's quick checks pass under this workspace
    Given the site project defines a "test:quick" target
    When "npx nx run wahidyankf-www:test:quick" runs
    Then the command exits 0

  Scenario: Migrated scenarios live in the specs tree
    Given the site's behavior is described by Gherkin feature files
    When the specs tree is inspected at "specs/apps/wahidyankf"
    Then every feature file from the source repository is present
    And each scenario names the test that binds it

  Scenario: The rendered pages match the source repository
    Given both copies of the site have been built
    When the migrated home page and the source home page are compared
    Then their rendered text content is identical

  Scenario: The source copy is removed only after verification
    Given the migrated site passes every gate in this plan
    When the removal step runs in the source repository
    Then the migrated copy is already serving and verified
```

## Product Scope

**In scope:** the site's pages, its component library usage, its end-to-end suite, its Gherkin specs, and its local development setup.

**Out of scope:** content changes, visual changes, new pages, analytics changes, and SEO work. Each is a separate plan, and each becomes possible once the site lives here.

## Product Risks

A visitor-facing regression is the only risk with an audience. It is contained by keeping the source copy live until the migrated copy is verified, and by comparing rendered output rather than trusting that a green build means an unchanged site.
