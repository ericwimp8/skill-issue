# Request

The desktop picker now needs to hide commands whose required capabilities are unavailable in the active runtime. The bug was reported in `CommandPicker`, and the proposed fix adds a hardcoded command-name switch there. Inspect the connected source files and decide the smallest complete placement. Preserve the current manifest-driven design.
