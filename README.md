# clearingway

Clearingway listens to a Discord channel for the message `/clears <world> <first-name> <last-name>`.

When it hears this message, it tries to find the relevant character in the Lodestone. If it finds them, it then parses their
fflogs and tries to assign them a few roles:

1. A role for the highest current parse they have in any relevant encounter (Gold, Orange, Purple, Blue, Green, Grey)
2. A role for every relevant encounter they've cleared ("P1S-Cleared," "P2S-Cleared," "P3S-Cleared," etc.)
3. A combo legend role purely for flexing purposes for every ultimate they've cleared (The Legend, The Double Legend, The Triple Legend, The Tetra Legend)
4. The role "NA's Comfiest" if they have any relevant encounter clears with a parse between 0 and 0.9
5. The role "The Nice Legend" if they have any ultimate clears with a parse between 69.0 and 69.9
6. The role "Nice" if they have any relevant clears with a parse between 69.0 and 69.9
7. The role "The Comfy Legend" if they have any ultimate clears with a parse between 0 and 0.9

It can be configured with the `config.yaml` file found in this repository.

## Running

Clearingway requires the following environment variables to start:

* **DISCORD_TOKEN**: You have to create a [Discord bot for Clearingway](https://discord.com/developers/applications). Once you've done so, you can add the bot token here.
* **FFLOGS_CLIENT_ID**: The client ID from [fflogs](https://www.fflogs.com/api/clients/).
* **FFLOGS_CLIENT_SECRET**: The client secret from [fflogs](https://www.fflogs.com/api/clients/).

## Directory Structure

.
├── cmd/
│   └── app/
│         ├── main.go            // Entrypoint, Bot initialization and dependency injection
│         ├── config.go          // Configuration handling
│         └── routes.go          // Discord command registry
├── config/                      // Configuration files
│   ├── examples/
│   │   ├── encounter-example.json
│   │   └── menu-example.json
│   ├── extreme/
│   ├── savage/
│   ├── ultimates/
│   └── menus/
├── internal/
│   ├── application/             // Use cases and orchestration
│   │   ├── commands/            // Command handlers
│   │   ├── queries/             // Query handlers
│   │   └── services/            // Application services
│   ├── domain/                  // Pure business logic (no external dependencies)
│   │   ├── entities/            // Domain entities
│   │   ├── repositories/        // Data access interfaces
│   │   └── services/            // Application services
│   ├── infrastructure/          // External integration
│   │   ├── discord/             // Discord API service
│   │   ├── fflogs/              // FFLogs API service
│   │   └── lodestone/           // Lodestone API service
│   ├── logging/                 // External Logging library implementations
│   │   ├── Interfaces/          // Logging Abstraction
│   │   └── zap/                 // Zap Logging implementation
│   ├── ui/                      // UI and interaction
│   │   ├── commands/            // Discord command definitions
│   │   ├── handlers/            // Discord interaction handlers
│   │   └── menus/               // Menu system
│   └── shared/                  // Shared utilities
│        ├── ffxiv/              // FFXIV-specific utilities
│        └── util/               // Common utilities
│  . . .

## About Each Package

### `cmd/`
Command is the layer that holds the entrypoint files (aka main) for the bot as well as any other tools we may need to add, i.e. migrations or adding an API. Currently by the end of the refactor, we should only have the app package in the command folder.

### `cmd/app/`
The App package holds the three files that actually run and wire the whole bot together. First, main.go should be solely used for the initialization of the bot and injecting dependencies throughout the bot. Second, config.go handles the ENV, the config files, and validates and defines the shape of the config data. Finally routes.go will handle the wiring of the commands and events, the registry so we can register the bot flow throughout the codebase.

### `config/`
The Config folder will hold all the config files we need for the codebase, the goal here is to make the config files readable and easy to add or swap them out as we need. There will be some consideration required as we move from the monolithic config.yaml to the more “micro-service” like scheme proposed here. First, I think we should have examples so when we need to add new fights to config, we can simply copy and paste and fill in the fields, the examples should give an explanation on where to get the information required as well when it may not be so clear, such as fight ID.

`menus/` is currently a tentative, possibly temporary folder to hold the menu configs until we decide how the redesign of the menus will look like.

### `internal/application/`
The Application layer holds the use cases and orchestration. We use the Application layer to connect the pure business logic in the Domain layer to the other layers.

**Testing:** We should test the Application layer with integration testing using external dependency mocks to ensure services are called correctly.

#### `commands/`
Here we keep the write-side use-cases, in other words, actions the system performs when a user asks for something to change. A command takes a request, coordinates domain rules, calls external systems through ports and returns a result.

#### `queries/`
Queries hold the read-side use-cases, each building a read model by coordinating repositories and read-only ports.

#### `services/`
Services are the interfaces that describe what the application needs done by external systems, along with some tiny shared types used by the interfaces.

### `internal/domain/`
The Domain layer is the layer we put pure business logic, structs, and methods. The Domain layer should depend on nothing, no external dependencies, no HTTP, no APIs, no database. Imports within the Domain layer are however allowed.

**Testing:** The Domain layer should test with pure unit tests, due to the requirement of no dependencies we wont need to mock anything and everything in the Domain layer should all work in isolation, so any methods or rules added should be tested.

#### `entities/`
The Entities package holds the definitions of the business nouns and the rules they must follow. Every definition must be stable so the rest of the codebase can depend on it. 

Every file within the Entities package contains a single business type with its state and the invariants that state must obey. The invariants should be enforced with factories and if pure methods are needed they would belong here as well.

#### `repositories/`
This is where interfaces are defined that describes how a domain asks for data to be loaded, saved or queried. Every interface should accept a `context.Context` and take and return domain types.

#### `services/`
Rules and calculations that span multiple domain nouns but still require no I/O. The place where business policies live when they don’t fit in a single entity. There should be one service per policy, a small, cohesive unit that performs a calculation or decision.

Domain nouns are classified as a service over an entity when the rule combines data from multiple nouns, applies policies/config that aren’t owned by a single noun or needs time windows, comparisons or catalogs that sit outside one entity and is still pure.

### `internal/infrastructure/`
The Infrastructure layer is for any external integrations, such as the discord API, pulling info from FFLogs or the lodestone, or any database integrations.

**Testing:** Since the Infrastructure layer is connecting all to external services, we should test how the logic integrates with the external services with the real service, ensuring we get the responses and structures we expect.

#### `discord/`
The discord package contains all logic that interacts with the Discord API.

#### `fflogs/`
The fflogs package contains all logic that interacts with the fflogs API.

#### `lodestone/`
The lodestone package contains all logic that interacts with the lodestone.

### `logging/`
The Logging layer is where all our logging will run through, keeping it easy to adjust and keeping code elsewhere readable.

**Testing:** Our biggest concerns in this layer will be to ensure the interfaces are mockable for testing other layers and ensuring the implementation integrates well with the parts of the codebase that is using logging.

#### `interfaces/`
The interfaces package is where we’ll keep all the abstraction for logging, a place to declare how logging functions work, but not where they’re implemented.

#### `zap/`
The zap package is where we implement actual logging with the zap library (https://pkg.go.dev/go.uber.org/zap)

**Note:** We are starting with Zap for logging, but if we encounter an edge case where Zap simply won't work or the implementation is not worth the time, we can add additional packages for other frameworks.

### `internal/ui/`
The UI layer is the UI of the bot, any logic that relates to what a user interacts with should go here.

**Testing:** In the UI layer, we should focus on tests that affect user interaction, ensuring we have proper request and response handling based on how users would interact with the bot, using mocks as needed.

#### `commands/`
A declarative catalog of all slash commands the bot exposes: their names, descriptions, options, permissions, and any metadata the registrar needs. We don’t handle events or logic here, we just define the specs.

#### `handlers/`
The Handlers package holds the code that reacts to Discord interactions and translates them into application layer requests. They then turn the results into user-facing replies. 

#### `menus/`
The Menus package contains the pure presentation builders, taking plain data and returning Discord ready payloads; embeds, components, content strings.

### `internal/shared/`
The Shared layer is for, as its name suggests, any utilities that are shared between the other layers (excluding Domain).

**Testing:** Testing in the Shared layer will depend a lot on what utilities will need to be added here, it is likely that these tests will consist largely of unit tests.

#### `ffxiv/`
Tiny pure toolkit for FFXIV domain basics. Here we would put enums/constants for jobs, roles, worlds, datacenters and tiers, lookup tables, and strict parser/normalizers.

#### `util/`
General-purpose, side-effect free helpers used everywhere else. String utilities, collections, generic paging/cursor structs, light error helpers, anything else that could be used by everything, it goes here. It should be small and testable.
