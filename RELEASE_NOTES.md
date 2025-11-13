# Release Notes

## v1.0.0 - 2025-11-13

Summary
- Initial release of the SDK Demos repository, containing minimal, self-contained examples for Python, Node.js, and Go, plus helper files to create and package the repository.

New Items
- examples/python
  - demo.py â€” simple POST demo using requests
  - requirements.txt â€” dependency list
  - README.md â€” run instructions
- examples/js
  - demo.js â€” simple POST demo using axios
  - package.json â€” Node metadata and dependency
  - README.md â€” run instructions
- examples/go
  - main.go â€” simple POST demo using net/http
  - go.mod â€” module file
  - README.md â€” run instructions
- Root-level files
  - README.md â€” repository overview and quick start
  - LICENSE â€” MIT (placeholder text included; replace with full canonical MIT before publishing)
  - .gitignore â€” common ignores for Python, Node, Go
  - create_and_zip.ps1 â€” script to create repository layout and produce sdk-demos.zip
  - RELEASE_NOTES.md â€” this file

Improvements
- Consistent API key handling: every example reads API_KEY from environment for uniformity.
- Cross-language minimal pattern: each example shows a minimal initialization + POST request and basic error handling so you can adapt to your SDK quickly.
- Packaging convenience: this PowerShell script produces a reproducible zip archive (sdk-demos.zip) for easy distribution.

Fixed Issues
- None in this initial release.

Removed Item
- None in this initial release.

Known Issues
- Placeholder endpoint: Each demo uses the placeholder endpoint at https://api.example.com/v1/demo â€” replace this with your real API endpoint or SDK call.
- LICENSE placeholder: The LICENSE file contains a placeholder notice. Replace it with the full canonical MIT text before publishing.
- No CI or automated tests: There is no GitHub Actions workflow included yet to validate examples.
- No SDK-specific bindings: Examples show generic HTTP requests; replace placeholders with real SDK client initialization and calls where appropriate.
- Windows / PowerShell: This script targets Windows PowerShell / PowerShell Core. If you need a POSIX shell variant, use the other script.

Contributors
- Sanjay-automater (initial author)

Upgrade suggestions (for v1.1.0)
- Add examples for additional languages (Java, C#, Ruby).
- Add a GitHub Actions workflow to run linting and basic smoke tests for each example.
- Replace placeholder HTTP calls with real SDK client usage for each language (if SDKs exist).
- Provide a POSIX shell script variant for non-Windows users.
