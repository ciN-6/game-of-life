# PowerShell 7 (pwsh) Common Commands & Idioms

## File & Directory Operations

| Task | Command (Idiomatic pwsh) | Notes |
| :--- | :--- | :--- |
| List files | `Get-ChildItem` (alias: `ls`, `dir`, `gci`) | Use `-Recurse` for recursive listing. |
| Find string in files | `Get-ChildItem -Recurse | Select-String "pattern"` | Equivalent to `grep -r`. |
| Remove file/folder | `Remove-Item path -Force -Recurse` | Alias: `rm`, `del`. |
| Copy file | `Copy-Item source destination` | Alias: `cp`, `copy`. |
| Move file | `Move-Item source destination` | Alias: `mv`, `move`. |

## Pipeline & Redirection

- **Command Chaining**:
  - **Sequential**: `cmd1; cmd2` (Works in 5.1 and 7).
  - **Conditional (AND)**: `cmd1; if ($?) { cmd2 }` (PowerShell 5.1 equivalent of `&&`).
  - **Conditional (OR)**: `cmd1; if (-not $?) { cmd2 }` (PowerShell 5.1 equivalent of `||`).
  - **Note**: `&&` and `||` are only available in **PowerShell 7 (pwsh)**.
- **No BOM**: In PowerShell 7, `>` and `Out-File` default to UTF-8 without BOM. In 5.1, use `-Encoding UTF8`.
- **Filtering**: `... | Where-Object { $_.Property -eq "Value" }` (alias: `?`).
- **Selecting**: `... | Select-Object Property1, Property2`.
- **Sorting**: `... | Sort-Object PropertyName`.

## JSON Processing

- **From JSON**: `$data = Get-Content file.json | ConvertFrom-Json`
- **To JSON**: `$data | ConvertTo-Json -Depth 10 | Set-Content file.json`

## Environment Variables

- **Get**: `$env:VARIABLE_NAME`
- **Set (Session)**: `$env:VARIABLE_NAME = "value"`

## Scripting Idioms

- **Foreach**: `1..10 | ForEach-Object { Write-Host $_ }` (alias: `%`).
- **Splatting**: 
  ```powershell
  $params = @{
      Path = "test.txt"
      Value = "Hello"
      Encoding = "utf8"
  }
  Set-Content @params
  ```
