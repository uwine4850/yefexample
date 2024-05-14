from yefpy import gotypes, yef


class Product(yef.YefClass):
    @gotypes.gotype({"name": "string"}, "")
    def __init__(self, name: str):
        self.name = name


class Shop(yef.YefClass):
    def __init__(self):
        self.products: list[Product] = []

    @gotypes.gotype({}, "[]Product")
    def get_products(self) -> list[Product]:
        return self.products

    @gotypes.gotype({"product": "*goclass.Class"}, "")
    def add_product(self, product: Product):
        self.products.append(product)

    @gotypes.gotype({"product": "*goclass.Class"}, "")
    def delete_products(self, product: Product):
        self.products.remove(product)
