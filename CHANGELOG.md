## v0.3.0 (2016-06-14)

* lib: Fix `Events()` to return all chaos events.
* lib: Move library code. You need to import `github.com/mlafeldt/chaosmonkey/lib` now.
* lib: Allow to configure custom User Agent.
* cli: Send custom User Agent `chaosmonkey Go client <version>`.
* cli: Allow to wipe state of Chaos Monkey via `-wipe-state`.
* cli: Add `--version` to show program version.

## v0.2.0 (2016-05-24)

* lib: Introduce `Strategy` type.
* lib: Add `Strategies` variable -- a list of chaos strategies.
* lib: Rename `ChaosEvent` to `Event`.
* cli: Allow to list chaos strategies via `-list-strategies`.
* cli: Allow to list auto scaling groups via `-list-groups`.

## v0.1.0 (2016-05-15)

* Initial version.
