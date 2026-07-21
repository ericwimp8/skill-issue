def format_invoice_amount(amount):
    sign = "-" if amount < 0 else ""
    return f"{sign}${abs(amount):.2f}"


def email_invoice(invoice):
    return f"Invoice total: {format_invoice_amount(invoice['amount'])}"


def pdf_invoice(invoice):
    return {"totalLabel": format_invoice_amount(invoice["amount"])}
