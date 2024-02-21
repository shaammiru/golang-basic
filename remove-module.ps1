param (
    [string]$moduleName
)

if (Select-String -Path "go.work" -Pattern $moduleName) {
    (Get-Content "go.work") -notmatch $moduleName | Set-Content "go.work"
    Remove-Item $moduleName -Recurse -Force
    Write-Output "Module $moduleName removed"
} else {
    Write-Output "Module $moduleName not found"
}
