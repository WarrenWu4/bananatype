# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Dates are in the form: **mm-dd-yyyy**

---

## [Unreleased]

### Added
- Default 10MB file size limit for logger with automatic rotation

### Changes 
- Settings page renamed to "settings"
- Expanded word bank to 4000+ words 
- Changed word randomizer from uniform distribution to normal distribution for more realistic word length variety
- Embedded word data to remove external file dependencies and for easier distribution

### Fixed
- Prevented duplicate consecutive words from being generated
- Cleaned up word bank data (removed random periods)


---

## [0.0.4] - 01-29-2026

### Fixed
- Temporary fix for resource path in different environments
- Removed random "." in word bank
- Fixed duplicate consecutive words during line generation

---

## [0.0.3] - 01-29-2026

### Refactored
- Updated logging mechanism

---

## [0.0.2] - 01-28-2026

### Added
- Basic logger

### Fixed
- Resource path on production
- Moved pkgbuild to separate directory

---

## [0.0.1] - 01-26-2026

### Added
- Basic typer typing view
- Settings view with time adjustment
- Analytics screen with WPM & accuracy metrics
- Build files, readmes, and changelog

---

[0.0.4]: https://github.com/WarrenWu4/bananas/releases/tag/v0.0.4
[0.0.3]: https://github.com/WarrenWu4/bananas/releases/tag/v0.0.3
[0.0.2]: https://github.com/WarrenWu4/bananas/releases/tag/v0.0.2
[0.0.1]: https://github.com/WarrenWu4/bananas/releases/tag/v0.0.1


