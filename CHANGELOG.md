# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

When a new release is proposed:

1. Create a new branch `bump/x.x.x` (this isn't a long-lived branch!!!);
2. The Unreleased section on `CHANGELOG.md` gets a version number and date;
3. Open a Pull Request with the bump version changes targeting the `main` branch;
4. When the Pull Request is merged a new git tag must be created using [GitLab environment](https://gitlab.com/intruderlabs/toolbox/intrusearch/-/tags).

Releases to productive environments should run from a tagged version.
Exceptions are acceptable depending on the circumstances (critical bug fixes that can be cherry-picked, etc.).

## [Unreleased]

### Added

- added `README.md` about the library to the community - [4157141881](https://intruderlabs.monday.com/boards/3790337872/pulses/4157141881)
- added `CONTRIBUTING.md` and `CODE_OF_CONDUCT.md` to the community - [4157141881](https://intruderlabs.monday.com/boards/3790337872/pulses/4157141881)
- added feature to search by index _id
- added the field `index` in the search_request
- added source in the OS response body - [PDT-10](https://intruderlabs.atlassian.net/browse/PDT-10)

### Changed

- corrected the error deserialization when it's not found response
- corrected the attempt to read the `io.ReadCloser` more than once in the method `doRequest` - [RD-175](https://intruderlabs.atlassian.net/browse/RD-175)
- created method for search with `hidra` - [RD-151](https://intruderlabs.atlassian.net/browse/RD-151)
- created file `client_interface` and add methods inside - [RD-43.19](https://intruderlabs.monday.com/boards/3797906866/pulses/3984323973)
- code sample has been improved in `README.md` - [4162823734](https://intruderlabs.monday.com/boards/3790337872/pulses/4162823734)

### Removed

-
