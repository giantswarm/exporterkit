# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [3.0.0] - 2021-12-16

### Changed

- Replace SIGKILL with SIGTERM signal
- Upgrade to Go 1.17
- Upgrade github.com/giantswarm/microendpoint v0.2.0 to v1.0.0
- Upgrade github.com/giantswarm/microerror v0.3.0 to v0.4.0
- Upgrade github.com/giantswarm/microkit v0.2.2 to v1.0.0
- Upgrade github.com/giantswarm/micrologger v0.5.0 to v0.6.0
- Upgrade github.com/prometheus/client_golang v1.9.0 to v1.11.0
- Upgrade golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e to v0.0.0-20210220032951-036812b2e83c
- Upgrade github.com/giantswarm/architect-orb v0.7.0 to v4.8.1

## [0.2.1] 2021-02-08

### Added

- Add `Set.Stop` method to stop collector after is has been booted.

### Fixed

- Reduce noisy debug logging when metrics are collected.

## [0.2.0] 2020-03-23

### Changed

- Switch from dep to Go modules.
- Use architect orb.



## [0.1.0] 2020-03-18

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/exporterkit/compare/v3.0.0...HEAD
[3.0.0]: https://github.com/giantswarm/exporterkit/compare/v0.2.1...v3.0.0
[0.2.1]: https://github.com/giantswarm/exporterkit/compare/v0.1.0...v0.2.0
[0.2.0]: https://github.com/giantswarm/exporterkit/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/exporterkit/releases/tag/v0.1.0
