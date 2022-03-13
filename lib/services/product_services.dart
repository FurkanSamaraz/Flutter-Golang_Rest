import 'package:flutter_application_1/models/product.dart';
import 'package:http/http.dart' as http;

class ProductService {
  Future<List<ProductModel>> getAllProducts() async {
    var url = Uri.parse('http://localhost:3000/root');
    var response = await http.get(url);
    var res = productModelFromJson(response.body);
    return res;
  }
}
