from products import create_product


def create_product_route(request, products):
    return create_product(request["code"].strip(), products)

