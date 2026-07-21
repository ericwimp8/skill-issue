def calculate_total(lines):
    return round(sum(line["price"] * line["quantity"] for line in lines), 2)


def store_invoice(lines, ledger):
    total = calculate_total(lines)
    ledger.append({"lines": lines, "total": total})
    return ledger[-1]


def api_invoice(invoice):
    return {**invoice, "total": round(invoice["total"], 2)}
