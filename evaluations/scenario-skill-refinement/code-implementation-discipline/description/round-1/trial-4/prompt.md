# Trial 4 Prompt

Work only in the supplied `fixture/` directory. A caller currently prepends the tenant identifier before writing cache keys, but background refresh calls the cache directly and collides across tenants. Make tenant isolation intrinsic to cache writes without duplicating caller work. Trace the current call path, implement the smallest complete change, run the focused checks, and summarize any caller cleanup.

