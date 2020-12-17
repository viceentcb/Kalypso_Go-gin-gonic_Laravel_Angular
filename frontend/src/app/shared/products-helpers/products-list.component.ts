import { Component, OnInit, Input } from '@angular/core';
import { Products, ProductsService } from '../../core';

@Component({
  selector: 'app-products',
  templateUrl: './products-list.component.html',
  styleUrls: ['./products-list.component.css']
})
export class ProductslistComponent implements OnInit {
  constructor(
    private productsService: ProductsService) { }


    products : Products[];

  ngOnInit() {
     this.products = [];
    this.productsService.getAll().subscribe(data => {
      this.products = data;
      console.log(this.products,'products laravel');
    })
  }

}


/**
 * NG ON INIT TESTING DATA GO
 */

// ngOnInit() {
//   this.buyProducts = [];
//  this.buysProducts.getAll().subscribe(data => {
//    this.buyProducts = data;
//    console.log(this.buyProducts,'products laravel');
//  })
// }