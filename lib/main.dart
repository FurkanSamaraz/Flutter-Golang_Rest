import 'package:flutter/material.dart';
import 'package:flutter_application_1/models/product.dart';
import 'package:flutter_application_1/services/product_services.dart';

void main() {
  runApp(AppDemo());
}

class AppDemo extends StatefulWidget {
  AppDemo({Key key}) : super(key: key);

  @override
  _AppDemoState createState() => _AppDemoState();
}

class _AppDemoState extends State<AppDemo> {
  var productService = ProductService();
  var listem = <ProductModel>[];
  @override
  void initState() {
    super.initState();
    getir();
  }

  getir() async {
    listem = await productService.getAllProducts();
    setState(() {});
    print(listem.first.name);
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: Padding(
          padding: const EdgeInsets.all(8.0),
          child: ListView.builder(
            itemCount: listem != null ? listem.length : 0,
            itemBuilder: (BuildContext context, int index) {
              return ListTile(
                title: Text(listem[index].name),
                subtitle: Text(listem[index].id.toString()),
                leading: FlutterLogo(size: 50.0),
                onTap: () => print(listem[index].name),
                contentPadding: EdgeInsets.all(10),
                shape: RoundedRectangleBorder(
                  side: BorderSide(color: Colors.black, width: 1),
                  borderRadius: BorderRadius.circular(5),
                ),
              );
            },
          ),
        ),
      ),
    );
  }
}
