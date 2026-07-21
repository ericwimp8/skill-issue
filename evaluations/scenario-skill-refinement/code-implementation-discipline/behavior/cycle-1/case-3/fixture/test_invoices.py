import unittest

from invoices import email_invoice, format_invoice_amount, pdf_invoice


class InvoicePresentationTest(unittest.TestCase):
    def test_negative_sign_precedes_currency_symbol(self):
        invoice = {"amount": -12.5}
        self.assertEqual(format_invoice_amount(invoice["amount"]), "-$12.50")
        self.assertEqual(email_invoice(invoice), "Invoice total: -$12.50")
        self.assertEqual(pdf_invoice(invoice)["totalLabel"], "-$12.50")

    def test_positive_format_is_preserved(self):
        self.assertEqual(format_invoice_amount(12.5), "$12.50")


if __name__ == "__main__":
    unittest.main()

