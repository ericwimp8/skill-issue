import unittest

from invoice import api_invoice, calculate_total, store_invoice


class InvoiceTest(unittest.TestCase):
    def test_total_is_canonical_for_every_consumer(self):
        lines = [{"price": 0.1, "quantity": 1}, {"price": 0.2, "quantity": 1}]
        ledger = []
        invoice = store_invoice(lines, ledger)
        self.assertEqual(calculate_total(lines), 0.3)
        self.assertEqual(invoice["total"], 0.3)
        self.assertEqual(api_invoice(invoice)["total"], 0.3)


if __name__ == "__main__":
    unittest.main()

