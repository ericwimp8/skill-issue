# Body Pre-Correction Source Snapshots

## Case 1 `cacheKey`

```js
/** Query property insertion order does not affect cache identity. */
export function cacheKey(path, query) {
  return `${path}?${JSON.stringify(query)}`;
}
```

## Case 3 `DiskProfileRepository.load`

```js
/** Repository implementations return a profile object. */
export class DiskProfileRepository {
  constructor(path) {
    this.path = path;
  }

  async load() {
    return readFile(this.path, 'utf8');
  }
}
```

The full unchanged surrounding fixture source remains in the original fixture specification and retained reports. These snapshots preserve the causal expressions replaced during successful body execution.
