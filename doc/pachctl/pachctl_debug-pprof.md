## ./pachctl debug-pprof

Analyze a profile of pachd in pprof.

### Synopsis


Analyze a profile of pachd in pprof.

```
./pachctl debug-pprof profile
```

### Options

```
      --binary-file string    File to write the binary to. (default "binary")
  -d, --duration duration     Duration to run a CPU profile for. (default 1m0s)
      --profile-file string   File to write the profile to. (default "profile")
```

### Options inherited from parent commands

```
      --no-metrics           Don't report user metrics for this command
      --no-port-forwarding   Disable implicit port forwarding
  -v, --verbose              Output verbose logs
```

