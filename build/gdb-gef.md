```
echo source "~/.gdbinit-gef.py" >> ~/.gdbinit
./stream-mng  -conf  ./stream-mng.toml
gdb ./stream-mng
set args -conf ./stream-mng.toml
```