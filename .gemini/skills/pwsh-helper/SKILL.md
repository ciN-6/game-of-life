---
name: pwsh-helper
description: Provides expert guidance on PowerShell 7 (pwsh) syntax, command-line usage, and scripting. Use when executing shell commands on Windows, writing pwsh scripts, or needing idiomatic PowerShell alternatives to bash/cmd.
---

# PowerShell Helper (pwsh)

This skill provides expert guidance for using PowerShell 7 (`pwsh`) effectively. It prioritizes modern, idiomatic syntax and cross-platform compatibility where applicable.

## Core Mandates

1. **Prefer PowerShell 7 (pwsh)**: Always assume the user has access to `pwsh`. Avoid legacy Windows PowerShell (5.1) quirks unless specified.
2. **Idiomatic Pipelining**: Use objects and filters (`Where-Object`, `Select-Object`) instead of complex string parsing (`awk`, `sed`).
3. **Silent Tools**: When writing scripts for Gemini CLI, use silent or quiet flags (e.g., `npm install --silent`) and ensure commands terminate.
4. **Command Chaining**: Avoid `&&` and `||` operators when targeting `powershell.exe` (Windows PowerShell 5.1). Use `;` for sequential execution or `if ($?) { ... }` for conditional execution.

## Quick Reference

- **Command Chaining (5.1 & 7)**: `cmd1; cmd2` (sequential) or `cmd1; if ($?) { cmd2 }` (conditional success).
- **File Searching**: `Get-ChildItem -Recurse | Select-String "pattern"`
- **JSON Handling**: `ConvertFrom-Json` and `ConvertTo-Json`
- **Error Handling**: Use `try { ... } catch { ... }` blocks for robust scripting.
- **Variables**: `$variable = "value"`
- **Execution Policy**: If scripts fail to run, check `Get-ExecutionPolicy`.

## Advanced Usage

For a comprehensive list of commands and their equivalents, see [references/common-commands.md](references/common-commands.md).

### JSON Manipulation Example

```powershell
# Read and update a JSON field
$pkg = Get-Content package.json | ConvertFrom-Json
$pkg.version = "1.0.1"
$pkg | ConvertTo-Json -Depth 10 | Set-Content package.json
```

### Filtering and Selecting

```powershell
# Get all .go files larger than 10KB
Get-ChildItem -Filter *.go -Recurse | Where-Object { $_.Length -gt 10kb } | Select-Object FullName, Length
```
