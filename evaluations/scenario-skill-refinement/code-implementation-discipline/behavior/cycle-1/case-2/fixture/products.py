def create_product(code, products):
    product = {"code": code}
    products.append(product)
    return product


def import_product(row, products):
    return create_product(row["code"], products)

