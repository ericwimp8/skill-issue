class Cache:
    def __init__(self):
        self.values = {}

    def put(self, key, value):
        self.values[key] = value

    def get(self, key):
        return self.values.get(key)


def save_from_request(cache, tenant_id, key, value):
    cache.put(f"{tenant_id}:{key}", value)


def refresh_in_background(cache, tenant_id, key, value):
    cache.put(key, value)

