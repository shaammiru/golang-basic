param (
    [string]$moduleName
)

New-Item -ItemType Directory -Path $moduleName
Set-Location $moduleName
go mod init $moduleName
New-Item main.go

Add-Content main.go "package main"
Add-Content main.go ""
Add-Content main.go "func main() {"
Add-Content main.go ""
Add-Content main.go "}"

Set-Location ..
go work use $moduleName

Write-Output "Module $moduleName created"
