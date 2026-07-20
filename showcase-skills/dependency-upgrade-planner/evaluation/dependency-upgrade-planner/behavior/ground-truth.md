# Behavior Ground Truth

## Case 1 — React 18.2 to React 19

- Production evidence includes aligned direct `react` and `react-dom` dependencies, aligned type packages, modern JSX transform, removed `ReactDOM.render`, removed `React.createFactory`, and function-component `defaultProps` usage.
- The plan should recognize React's official 18.3 warning bridge as a possible prerequisite stage, align React DOM and types, and distinguish direct/runtime/type/build effects.
- Compatibility with every third-party dependency cannot be claimed from the sparse lockfile; the scheduler edge is transitive and must be handled through lockfile resolution rather than edited directly.
- Implementation remains out of scope. Validation should include typecheck, build, focused runtime behavior, error-reporting behavior where applicable, and relevant application tests only after production impacts are mapped.

## Case 2 — Vite 5.4 to Vite 6

- Vite is a direct development/build-tool dependency; Rollup is transitive; Terser is a direct build dependency below Vite 6's documented minimum for terser minification.
- The custom `resolve.conditions` list needs comparison with Vite 6 defaults; Sass explicitly requests legacy API; library CSS filename behavior can affect the consumer export; Node 21 is a declared and deployed platform risk because Vite 6 dropped Node 21 support.
- The plan should separate build-tool, transitive, platform, configuration, and produced-artifact effects and order the Node decision before claiming a viable Vite 6 build environment.
- Validation should include clean lock resolution after authorization, configured builds, artifact filename/export inspection, condition-resolution behavior, Sass output, and supported deployment runtime; rollback should cover manifest/lock/config/artifact state.

## Case 3 — Pydantic 1.10 to Pydantic 2

- Pydantic is direct and runtime; FastAPI is direct and has a locked peer constraint excluding Pydantic 2, so FastAPI compatibility or upgrade is a prerequisite edge.
- Production usage includes V1 `Config.orm_mode`, `from_orm`, and `parse_obj`; Pydantic's migration guidance maps these concepts toward V2 configuration and `model_validate` behavior.
- The plan should not treat deprecation shims as proof of complete compatibility and should distinguish Python/runtime, framework/peer, API usage, and validation-semantics effects.
- Validation should cover dependency resolution, model parsing and ORM-backed validation semantics, FastAPI request/response behavior, generated OpenAPI where relevant, and representative runtime paths. Rollback must restore dependency and lock state plus any migrated model/config source.

## Shared Required Characteristics

- Cite repository-relative project evidence and authoritative upstream URLs near claims.
- Label observed project facts, authoritative upstream facts, reasoned implications, and unknowns.
- Sequence prerequisites before dependent work and end with the smallest safe first stage plus evidence required before implementation.
- Preserve planning-only authority and avoid compatibility, build, test, or runtime success claims unsupported by executed evidence.
