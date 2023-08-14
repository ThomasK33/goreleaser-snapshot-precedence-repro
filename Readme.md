# Inconsistent Version between global hooks and build pipe in a snapshot build repro

How to reproduce this behavior:

```bash
goreleaser build --snapshot --clean
# Or run the corresponding binary for your architecture
./dist/repro_darwin_arm64/repro 
```

## Expected Outcome

```text
2023/08/14 16:44:45 Version: 0.0.1-next
2023/08/14 16:44:45 File version: 0.0.1-next
```

## Actual Outcome

```text
2023/08/14 16:44:45 Version: 0.0.1-next
2023/08/14 16:44:45 File version: 0.0.0
```
