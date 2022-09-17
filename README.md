# Command Line

## Install
```
go install github.com/longht021189/ops-cl@latest
```

## Git
### Remove Submodule
```
ops-cl git remove-submodule --path path/to/submodule

# or

ops-cl git remove-submodule
# to list all submodules and pick to remove
```

## Docker
### Copy Binary to Scratch
```
ops-cl docker build-scratch -b <path-to-binary> -o <output-file> --builder=<name-of-docker-image>
```