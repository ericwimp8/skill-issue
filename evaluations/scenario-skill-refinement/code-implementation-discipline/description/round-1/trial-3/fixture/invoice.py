def calculate_total(lines):
    return sum(line["price"] * line["quantity"] for line in lines)


def store_invoice(lines, ledger):
    total = calculate_total(lines)
    ledger.append({"lines": lines, "total": total})
    return ledger[-1]


def api_invoice(invoice):
    return {**invoice, "total": round(invoice["total"], 2)}

