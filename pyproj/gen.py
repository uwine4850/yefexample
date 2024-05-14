from yefpy.gen_start_sh import generate_start_sh
from yefpy.mod_data import get_modules_info, Module
from yefpy.generate import generate_yaml, get_yaml_modules_data

if __name__ == '__main__':
    generate_start_sh("/home/fhx/GolandProjects/yefgotest/pyproj/.venv", ["/home/fhx/GolandProjects/yefgotest/pyproj"],
        "/home/fhx/GolandProjects/yefgotest/pygen")
    shop = Module("proj.shop", "shop", "shop.go", [])
    customer = Module("proj.customer", "customer", "customer.go", [])
    modules_info = get_modules_info([shop, customer])
    yaml_modules_data = get_yaml_modules_data(modules_info)
    generate_yaml(yaml_modules_data, "pygen.yaml", "/home/fhx/GolandProjects/yefgotest/pygen")
