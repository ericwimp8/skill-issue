class Cache:
    def __init__(self):
        self.values = {}

    def put(self, tenant_id, key, value):
        self.values[f"{tenant_id}:{key}"] = value

    def get(self, tenant_id, key):
        return self.values.get(f"{tenant_id}:{key}")


def save_from_request(cache, tenant_id, key, value):
    cache.put(tenant_id, key, value)


def refresh_in_background(cache, tenant_id, key, value):
    cache.put(tenant_id, key, value)
