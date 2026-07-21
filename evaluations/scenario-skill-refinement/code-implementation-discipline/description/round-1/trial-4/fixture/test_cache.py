import unittest

from cache import Cache, refresh_in_background, save_from_request


class CacheTest(unittest.TestCase):
    def test_all_writes_are_tenant_scoped(self):
        cache = Cache()
        save_from_request(cache, "tenant-a", "profile", "request")
        refresh_in_background(cache, "tenant-b", "profile", "refresh")
        self.assertEqual(cache.get("tenant-a", "profile"), "request")
        self.assertEqual(cache.get("tenant-b", "profile"), "refresh")


if __name__ == "__main__":
    unittest.main()

