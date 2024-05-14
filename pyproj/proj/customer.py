from proj.shop import Product, Shop
from yefpy import gotypes, yef


class Customer(yef.YefClass):
    @gotypes.gotype({"shop": "*goclass.Class"}, "")
    def __init__(self, shop: Shop):
        self.shop = shop

    @gotypes.gotype({"product": "*goclass.Class"}, "")
    def by_product(self, product: Product):
        self.shop.delete_products(product)
